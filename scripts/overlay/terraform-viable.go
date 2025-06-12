package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Check if a resource is viable for Terraform
func isTerraformViable(resource *ResourceInfo, spec OpenAPISpec) bool {
	// Must have at least create and read operations
	_, hasCreate := resource.Operations["create"]
	_, hasRead := resource.Operations["read"]

	if !hasCreate || !hasRead {
		fmt.Printf("    Missing required operations: Create=%v, Read=%v\n", hasCreate, hasRead)
		return false
	}

	// Must have a create schema to be manageable by Terraform
	if resource.CreateSchema == "" {
		fmt.Printf("    No create schema found\n")
		return false
	}

	// Identify the primary ID for this entity
	primaryID, validPrimaryID := identifyEntityPrimaryID(resource)
	if !validPrimaryID {
		fmt.Printf("    Cannot identify valid primary ID parameter\n")
		return false
	}

	// Validate all operations against the primary ID
	validOperations := validateOperationParameters(resource, primaryID, spec)

	// Must still have CREATE and READ after validation
	_, hasValidCreate := validOperations["create"]
	_, hasValidRead := validOperations["read"]

	if !hasValidCreate || !hasValidRead {
		fmt.Printf("    Lost required operations after parameter validation: Create=%v, Read=%v\n", hasValidCreate, hasValidRead)
		return false
	}

	// Update resource with only valid operations and primary ID
	resource.Operations = validOperations
	resource.PrimaryID = primaryID

	// Check for overlapping properties between create and entity schemas
	if !hasValidCreateReadConsistency(resource, spec) {
		fmt.Printf("    Create and Read operations have incompatible schemas\n")
		return false
	}

	// Check for problematic CRUD patterns that can't be handled by property ignoring
	if resource.CreateSchema != "" && resource.UpdateSchema != "" {
		// Re-parse the spec to get raw schema data for analysis
		specData, err := json.Marshal(spec)
		if err != nil {
			return true // If we can't analyze, assume it's viable
		}

		var rawSpec map[string]interface{}
		if err := json.Unmarshal(specData, &rawSpec); err != nil {
			return true // If we can't analyze, assume it's viable
		}

		components, _ := rawSpec["components"].(map[string]interface{})
		schemas, _ := components["schemas"].(map[string]interface{})

		createProps := getSchemaProperties(schemas, resource.CreateSchema)
		updateProps := getSchemaProperties(schemas, resource.UpdateSchema)

		// Count manageable properties (non-system fields)
		createManageableProps := 0
		updateManageableProps := 0
		commonManageableProps := 0

		for prop := range createProps {
			if !isSystemProperty(prop) {
				createManageableProps++
			}
		}

		for prop := range updateProps {
			if !isSystemProperty(prop) {
				updateManageableProps++
				// Check if this property also exists in create
				if createProps[prop] != nil && !isSystemProperty(prop) {
					commonManageableProps++
				}
			}
		}

		// Reject resources with fundamentally incompatible CRUD patterns
		if createManageableProps <= 1 && updateManageableProps >= 3 && commonManageableProps == 0 {
			fmt.Printf("    Incompatible CRUD pattern: Create=%d manageable, Update=%d manageable, Common=%d\n",
				createManageableProps, updateManageableProps, commonManageableProps)
			return false
		}
	}

	return true
}

// Validate operations against the identified primary ID
func validateOperationParameters(resource *ResourceInfo, primaryID string, spec OpenAPISpec) map[string]OperationInfo {
	validOperations := make(map[string]OperationInfo)

	// Get entity properties once for this resource
	entityProps := getEntityProperties(resource.EntityName, spec)

	for crudType, opInfo := range resource.Operations {
		pathParams := extractPathParameters(opInfo.Path)

		if crudType == "create" || crudType == "list" {
			// These operations should not have the entity's primary ID in path
			hasPrimaryID := false
			for _, param := range pathParams {
				if param == primaryID {
					hasPrimaryID = true
					break
				}
			}

			if hasPrimaryID {
				fmt.Printf("    Skipping %s operation %s: unexpectedly has primary ID %s in path\n",
					crudType, opInfo.Path, primaryID)
				continue
			}

			validOperations[crudType] = opInfo
			continue
		}

		// READ, UPDATE, DELETE should have exactly the primary ID
		hasPrimaryID := false
		hasConflictingEntityIDs := false

		for _, param := range pathParams {
			if param == primaryID {
				hasPrimaryID = true
			} else if isEntityID(param) {
				// This is another ID-like parameter
				// Check if it maps to a field in the entity (not the primary id field)
				if checkFieldExistsInEntity(param, entityProps) {
					// This parameter maps to a real entity field - it's valid
					fmt.Printf("    Parameter %s maps to entity field - keeping operation %s %s\n",
						param, crudType, opInfo.Path)
				} else {
					// This ID parameter doesn't map to any entity field
					if mapsToEntityID(param, resource.EntityName) {
						fmt.Printf("    Skipping %s operation %s: parameter %s would conflict with primary ID %s (both map to entity.id)\n",
							crudType, opInfo.Path, param, primaryID)
						hasConflictingEntityIDs = true
						break
					} else {
						// This is an unmappable ID parameter
						fmt.Printf("    Skipping %s operation %s: unmappable ID parameter %s (not in entity schema)\n",
							crudType, opInfo.Path, param)
						hasConflictingEntityIDs = true
						break
					}
				}
			}
			// Non-ID parameters are always OK
		}

		if !hasPrimaryID {
			fmt.Printf("    Skipping %s operation %s: missing primary ID %s\n",
				crudType, opInfo.Path, primaryID)
			continue
		}

		if hasConflictingEntityIDs {
			continue
		}

		validOperations[crudType] = opInfo
	}

	fmt.Printf("    Valid operations after parameter validation: %v\n", getOperationTypes(validOperations))
	return validOperations
}

// Get entity properties for field existence checking
func getEntityProperties(entityName string, spec OpenAPISpec) map[string]interface{} {
	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return map[string]interface{}{}
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return map[string]interface{}{}
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	return getSchemaProperties(schemas, entityName)
}

// Check if create and read operations have compatible schemas
func hasValidCreateReadConsistency(resource *ResourceInfo, spec OpenAPISpec) bool {
	if resource.CreateSchema == "" {
		return false
	}

	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return false
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return false
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	entityProps := getSchemaProperties(schemas, resource.EntityName)
	createProps := getSchemaProperties(schemas, resource.CreateSchema)

	if len(entityProps) == 0 || len(createProps) == 0 {
		return false
	}

	// Count overlapping manageable properties
	commonManageableProps := 0
	createManageableProps := 0

	for prop := range createProps {
		if !isSystemProperty(prop) {
			createManageableProps++
			if entityProps[prop] != nil {
				commonManageableProps++
			}
		}
	}

	// Need at least some manageable properties
	if createManageableProps == 0 {
		return false
	}

	// Require at least 30% overlap of create properties to exist in entity
	// This is more lenient than the 50% I had before
	overlapRatio := float64(commonManageableProps) / float64(createManageableProps)
	return overlapRatio >= 0.3
}

func getSchemaProperties(schemas map[string]interface{}, schemaName string) map[string]interface{} {
	if schemaName == "" {
		return map[string]interface{}{}
	}

	schema, exists := schemas[schemaName]
	if !exists {
		return map[string]interface{}{}
	}

	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	properties, ok := schemaMap["properties"].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	return properties
}

func isSystemProperty(propName string) bool {
	systemProps := []string{
		"id", "created_at", "updated_at", "created_by", "updated_by",
		"version", "etag", "revision", "last_modified",
	}

	lowerProp := strings.ToLower(propName)

	for _, sysProp := range systemProps {
		if lowerProp == sysProp || strings.HasSuffix(lowerProp, "_"+sysProp) {
			return true
		}
	}

	// Also consider ID fields as system properties
	if strings.HasSuffix(lowerProp, "_id") {
		return true
	}

	return false
}

// Check if a field exists in the entity properties
func checkFieldExistsInEntity(paramName string, entityProps map[string]interface{}) bool {
	// Direct field name match
	if _, exists := entityProps[paramName]; exists {
		return true
	}

	// Check for common variations
	variations := []string{
		paramName,
		strings.TrimSuffix(paramName, "_id"), // Remove _id suffix
		strings.TrimSuffix(paramName, "Id"),  // Remove Id suffix
	}

	for _, variation := range variations {
		if _, exists := entityProps[variation]; exists {
			return true
		}
	}

	return false
}

// Identify the primary ID parameter that belongs to this specific entity
func identifyEntityPrimaryID(resource *ResourceInfo) (string, bool) {
	// Get all unique path parameters across operations
	allParams := make(map[string]bool)

	for crudType, opInfo := range resource.Operations {
		if crudType == "create" || crudType == "list" {
			continue // Skip operations that typically don't have entity-specific IDs
		}

		pathParams := extractPathParameters(opInfo.Path)
		for _, param := range pathParams {
			allParams[param] = true
		}
	}

	if len(allParams) == 0 {
		return "", false // No path parameters found
	}

	// Find the parameter that matches this entity
	var entityPrimaryID string
	matchCount := 0

	for param := range allParams {
		if mapsToEntityID(param, resource.EntityName) {
			entityPrimaryID = param
			matchCount++
		}
	}

	if matchCount == 0 {
		// No parameter maps to this entity - check for generic 'id' parameter
		if allParams["id"] {
			fmt.Printf("    Using generic 'id' parameter for entity %s\n", resource.EntityName)
			return "id", true
		}
		fmt.Printf("    No parameter maps to entity %s\n", resource.EntityName)
		return "", false
	}

	if matchCount > 1 {
		// Multiple parameters claim to map to this entity - ambiguous
		fmt.Printf("    Multiple parameters map to entity %s: ambiguous primary ID\n", resource.EntityName)
		return "", false
	}

	fmt.Printf("    Identified primary ID '%s' for entity %s\n", entityPrimaryID, resource.EntityName)
	return entityPrimaryID, true
}

// Check if a parameter name maps to a specific entity's ID field
func mapsToEntityID(paramName, entityName string) bool {
	// Extract base name from entity (e.g., "ChangeEvent" from "ChangeEventEntity")
	entityBase := strings.TrimSuffix(entityName, "Entity")

	// Convert to snake_case and add _id suffix
	expectedParam := toSnakeCase(entityBase) + "_id"

	return strings.ToLower(paramName) == strings.ToLower(expectedParam)
}

// Check if parameter looks like an entity ID
func isEntityID(paramName string) bool {
	return strings.HasSuffix(strings.ToLower(paramName), "_id") || strings.ToLower(paramName) == "id"
}

func getOperationTypes(operations map[string]OperationInfo) []string {
	var types []string
	for opType := range operations {
		types = append(types, opType)
	}
	return types
}
