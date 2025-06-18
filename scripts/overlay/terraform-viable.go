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

	if (!hasCreate) || !hasRead {
		fmt.Printf("    Missing required operations in %v: Create=%v Read=%v\n", resource.ResourceName, hasCreate, hasRead)
		return false
	}

	// Must have a create schema to be manageable by Terraform
	if resource.CreateSchema == "" {
		fmt.Printf("    No create schema found for %v\n", resource.ResourceName)
		return false
	}

	primaryID, validPrimaryID := identifyEntityPrimaryID(resource, spec)
	if !validPrimaryID {
		fmt.Printf("    Cannot identify valid primary ID parameter for %v\n", resource.EntityName)
		return false
	}

	operationsCopy := make(map[string]OperationInfo)
	for k, v := range resource.Operations {
		operationsCopy[k] = v
	}

	// Validate all operations against the primary ID
	validOperations := validateOperationParameters(&ResourceInfo{
		EntityName:   resource.EntityName,
		SchemaName:   resource.SchemaName,
		ResourceName: resource.ResourceName,
		Operations:   operationsCopy,
		CreateSchema: resource.CreateSchema,
		UpdateSchema: resource.UpdateSchema,
		PrimaryID:    primaryID,
	}, primaryID, spec)

	// Must still have CREATE and READ after validation
	_, hasValidCreate := validOperations["create"]
	_, hasValidRead := validOperations["read"]

	if !hasValidCreate || !hasValidRead {
		fmt.Printf("    Lost required operations after parameter validation: Create=%v, Read=%v\n", hasValidCreate, hasValidRead)
		return false
	}

	// Only update the resource if this is the first validation
	if resource.PrimaryID == "" {
		resource.Operations = validOperations
		resource.PrimaryID = primaryID
	}

	// Check for overlapping properties between create and entity schemas
	if !hasValidCreateReadConsistency(resource, spec) {
		fmt.Printf("    %v Create and Read operations have incompatible schemas\n", resource.EntityName)
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
				if createProps[prop] != nil && !isSystemProperty(prop) {
					commonManageableProps++
				}
			}
		}

		if createManageableProps <= 1 && updateManageableProps >= 3 && commonManageableProps == 0 {
			fmt.Printf("    Incompatible CRUD pattern: Create=%d manageable, Update=%d manageable, Common=%d\n",
				createManageableProps, updateManageableProps, commonManageableProps)
			return false
		}
	}

	return true
}

func validateOperationParameters(resource *ResourceInfo, primaryID string, spec OpenAPISpec) map[string]OperationInfo {
	validOperations := make(map[string]OperationInfo)

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

		hasPrimaryID := false
		hasConflictingEntityIDs := false

		// We need to validate that we do not have conflicting ID path parameters
		// At time of this comment, we only map a single ID parameter to the corresponding entity id, using x-speakeasy-match

		// If this constraint prevents us from adding resources that customers need/want
		//   we'll need to figure out how to map additional IDs to right field and entities
		//   for the time, being, multiple ID params in path are not supported and corresponding operations are ignored,
		//   unless they are already exact match to the field name on the entity
		for _, param := range pathParams {
			if param == primaryID {
				hasPrimaryID = true
			} else if isEntityID(param) {
				// This is another ID-like parameter
				// Check if it maps to a field in the entity (not the primary id field)
				if checkFieldExistsInEntity(param, entityProps) {
					// This parameter maps to a real entity field - it's valid
					continue
				} else {
					// This ID parameter doesn't map to any entity field
					hasConflictingEntityIDs = true
					break
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
			fmt.Printf("    Skipping %s operation %s: has conflicting entity ID parameters\n",
				crudType, opInfo.Path)
			continue
		}

		validOperations[crudType] = opInfo
	}

	return validOperations
}

func getEntityProperties(entityName string, spec OpenAPISpec) map[string]interface{} {
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
// We need to ensure that the create and read operations are exactly the same, after accounting for ignored properties and normalization
func hasValidCreateReadConsistency(resource *ResourceInfo, spec OpenAPISpec) bool {
	if resource.CreateSchema == "" {
		return false
	}

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

	// If there is any overlap, try to use the schemas
	return true
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

	return strings.HasSuffix(lowerProp, "_id")
}

func checkFieldExistsInEntity(paramName string, entityProps map[string]interface{}) bool {
	if _, exists := entityProps[paramName]; exists {
		return true
	}

	variations := []string{
		paramName,
		strings.TrimSuffix(paramName, "_id"),
		strings.TrimSuffix(paramName, "Id"),
	}

	for _, variation := range variations {
		if _, exists := entityProps[variation]; exists {
			return true
		}
	}

	return false
}

func identifyEntityPrimaryID(resource *ResourceInfo, spec OpenAPISpec) (string, bool) {
	allParams := make(map[string]bool)

	for crudType, opInfo := range resource.Operations {
		if crudType == "create" || crudType == "list" {
			continue
		}

		pathParams := extractPathParameters(opInfo.Path)
		for _, param := range pathParams {
			allParams[param] = true
		}
	}

	if len(allParams) == 0 {
		return "", false
	}

	var entityPrimaryID string
	matchCount := 0

	for param := range allParams {
		if mapsToEntityID(param, resource.EntityName) {
			entityPrimaryID = param
			matchCount++
		}
	}

	if matchCount == 1 {
		return entityPrimaryID, true
	}

	entityProps := getEntityProperties(resource.EntityName, spec)
	_, hasID := entityProps["id"]
	_, hasSlug := entityProps["slug"]

	if hasSlug {
		for param := range allParams {
			if strings.HasSuffix(param, "_slug") {
				return param, true
			}
		}

		for param := range allParams {
			if strings.Contains(strings.ToLower(param), "slug") {
				return param, true
			}
		}
	}

	if hasID {
		entityBase := strings.ToLower(strings.TrimSuffix(resource.EntityName, "Entity"))
		for param := range allParams {
			if strings.HasSuffix(param, "_id") {
				paramBase := strings.TrimSuffix(param, "_id")
				if strings.Contains(paramBase, entityBase) || strings.Contains(entityBase, paramBase) {
					return param, true
				}
			}
		}
	}

	if allParams["id"] {
		return "id", true
	}

	if len(allParams) == 1 && (hasID || hasSlug) {
		for param := range allParams {
			return param, true
		}
	}

	return "", false
}

func mapsToEntityID(paramName, entityName string) bool {
	entityBase := strings.TrimSuffix(entityName, "Entity")

	expectedParam := toSnakeCase(entityBase) + "_id"

	return strings.EqualFold(paramName, expectedParam)
}

func isEntityID(paramName string) bool {
	return strings.HasSuffix(strings.ToLower(paramName), "_id") || strings.ToLower(paramName) == "id"
}
