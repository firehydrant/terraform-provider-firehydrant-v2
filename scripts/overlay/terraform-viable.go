package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Check if a resource is viable for Terraform
func isTerraformViable(resource *ResourceInfo, manualMappings *ManualMappings, schemas map[string]interface{}) bool {
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

	primaryID, validPrimaryID := identifyEntityPrimaryID(resource, schemas)
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
	}, primaryID, schemas, manualMappings)

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
	if !hasValidCreateReadConsistency(resource, schemas) {
		fmt.Printf("    %v Create and Read operations have incompatible schemas\n", resource.EntityName)
		return false
	}

	// Check for problematic CRUD patterns that can't be handled by property ignoring
	if resource.CreateSchema != "" && resource.UpdateSchema != "" {
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

func validateOperationParameters(resource *ResourceInfo, primaryID string, schemas map[string]interface{}, manualMappings *ManualMappings) map[string]OperationInfo {
	validOperations := make(map[string]OperationInfo)

	entityProps := getSchemaProperties(schemas, resource.EntityName)

	for crudType, opInfo := range resource.Operations {
		pathParams := extractPathParameters(opInfo.Path)

		if crudType == "create" || crudType == "list" {
			// These operations should not have the entity's primary ID in path
			hasPrimaryID := false
			for _, param := range pathParams {
				// Check if this parameter maps to the primary ID (either directly or via manual mapping)
				if param == primaryID {
					hasPrimaryID = true
					break
				}
				// Check if there's a manual mapping for this parameter to the primary ID
				if manualMatch, hasManual := getManualParameterMatch(opInfo.Path, opInfo.Method, param, manualMappings); hasManual {
					// Support nested properties in manual mappings
					if manualMatch == primaryID || strings.HasSuffix(manualMatch, "."+primaryID) {
						hasPrimaryID = true
						break
					}
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

		for _, param := range pathParams {
			if param == primaryID {
				hasPrimaryID = true
				continue
			}

			if manualMatch, hasManual := getManualParameterMatch(opInfo.Path, opInfo.Method, param, manualMappings); hasManual {
				// If the manual mapping points to the primary ID (could be nested), this satisfies our primary ID requirement
				if manualMatch == primaryID || strings.HasSuffix(manualMatch, "."+primaryID) {
					hasPrimaryID = true
					continue
				}

				// If the manual mapping points to a valid entity field (including nested), it's acceptable
				if checkFieldExistsInEntityWithRefResolution(manualMatch, entityProps, schemas) {

					continue
				} else {

					hasConflictingEntityIDs = true
					break
				}
			} else if isEntityID(param) {
				// This is another ID-like parameter without manual mapping
				// Only consider it problematic if it conflicts with the primary ID or claims to be this entity's ID
				if param != primaryID && mapsToEntityID(param, resource.EntityName) {
					// This parameter claims to be this entity's ID but isn't the primary ID - that's a conflict
					hasConflictingEntityIDs = true
					break
				}

				// For other ID parameters (like team_id, parent_id, etc.), they're just foreign keys
				// and don't need to exist in the entity schema - they're fine
				// we weren't able to map them but they don't conflict with the primary ID
			}
			// Non-ID parameters are always OK
		}

		if !hasPrimaryID {
			fmt.Printf("    Skipping %s operation %s: missing primary ID %s (either direct or via manual mapping)\n",
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

// Check if create and read operations have compatible schemas
// We need to ensure that the create and read operations are exactly the same, after accounting for ignored properties and normalization
func hasValidCreateReadConsistency(resource *ResourceInfo, schemas map[string]interface{}) bool {
	if resource.CreateSchema == "" {
		return false
	}

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

func checkFieldExistsInEntityWithRefResolution(fieldPath string, entityProps map[string]interface{}, schemas map[string]interface{}) bool {
	parts := strings.Split(fieldPath, ".")

	fmt.Printf("    Debug: Checking field path: %s (parts: %v)\n", fieldPath, parts)

	currentLevel := entityProps

	for i, part := range parts {
		fmt.Printf("    Debug: Looking for part '%s' at level %d\n", part, i)

		if prop, exists := currentLevel[part]; exists {
			fmt.Printf("    Debug: Found part '%s'\n", part)

			if i == len(parts)-1 {
				fmt.Printf("    Debug: Reached final part, field exists: %s\n", fieldPath)
				return true
			}

			if propMap, ok := prop.(map[string]interface{}); ok {
				// Check if it has direct properties
				if nestedProps, hasProps := propMap["properties"].(map[string]interface{}); hasProps {
					fmt.Printf("    Debug: Found direct properties for '%s'\n", part)
					currentLevel = nestedProps
					continue
				}

				// Check if it has a $ref that needs resolution
				if ref, hasRef := propMap["$ref"].(string); hasRef {
					fmt.Printf("    Debug: Found $ref '%s' for part '%s'\n", ref, part)

					resolvedSchema := resolveSchemaRef(ref, schemas)
					if resolvedSchema != nil {
						fmt.Printf("    Debug: Successfully resolved $ref '%s'\n", ref)

						if refProps, hasRefProps := resolvedSchema["properties"].(map[string]interface{}); hasRefProps {
							fmt.Printf("    Debug: Found properties in resolved schema\n")
							currentLevel = refProps
							continue
						} else {
							fmt.Printf("    Debug: Resolved schema has no properties\n")
						}
					} else {
						fmt.Printf("    Debug: Failed to resolve $ref '%s'\n", ref)
					}

					fmt.Printf("    Warning: Cannot resolve $ref for nested property validation: %s (ref: %s)\n", fieldPath, ref)
					return false
				}
			}

			// If we can't navigate deeper but haven't reached the end, the path is invalid
			fmt.Printf("    Debug: Cannot navigate deeper from part '%s' - not a valid object structure\n", part)
			fmt.Printf("    Cannot navigate to nested property: %s at part: %s\n", fieldPath, part)
			return false
		} else {
			fmt.Printf("    Debug: Part '%s' not found in current level\n", part)
			fmt.Printf("    Debug: Available keys in current level: %v\n", getKeys(currentLevel))
			fmt.Printf("    Field does not exist: %s at part: %s\n", fieldPath, part)
			return false
		}
	}

	return false
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func resolveSchemaRef(ref string, schemas map[string]interface{}) map[string]interface{} {
	if !strings.HasPrefix(ref, "#/components/schemas/") {
		return nil
	}

	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")

	if referencedSchema, exists := schemas[schemaName]; exists {
		specData, _ := json.Marshal(referencedSchema)
		var schemaMap map[string]interface{}
		json.Unmarshal(specData, &schemaMap)

		// At time of writing we have validation rules which prevent oneOfs in our swagger generation within laddertruck.
		// So we only have allOf patterns to work with, these are generally created by our NullableWrappers
		// NullableWrappers are used to allow nullable properties in the API (also created in laddertruck), so we need to resolve them
		if allOf, hasAllOf := schemaMap["allOf"].([]interface{}); hasAllOf {
			for _, item := range allOf {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if innerRef, hasInnerRef := itemMap["$ref"].(string); hasInnerRef {
						resolved := resolveSchemaRef(innerRef, schemas)
						if resolved != nil {
							return resolved
						}
					}
					if _, hasProps := itemMap["properties"].(map[string]interface{}); hasProps {
						return itemMap
					}
				}
			}
		}

		if _, hasProps := schemaMap["properties"].(map[string]interface{}); hasProps {
			return schemaMap
		}

		return schemaMap
	}

	return nil
}

func identifyEntityPrimaryID(resource *ResourceInfo, schemas map[string]interface{}) (string, bool) {
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

	entityProps := getSchemaProperties(schemas, resource.EntityName)
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

		if allParams["id"] {
			return "id", true
		}
	}

	if len(allParams) == 1 && (hasID || hasSlug) {
		for param := range allParams {
			if !strings.Contains(param, "team_id") && !strings.Contains(param, "parent_id") {
				return param, true
			}
		}
	}

	// For complex cases with multiple parameters, try to identify the most likely primary ID
	// Look for parameters that end with the entity name or are simple "id"
	entityBase := strings.ToLower(strings.TrimSuffix(resource.EntityName, "Entity"))
	for param := range allParams {
		lowerParam := strings.ToLower(param)
		if lowerParam == entityBase+"_id" || lowerParam == "id" {
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

func isViableDatasource(resource *ResourceInfo) bool {
	// A resource can be a data source if it has:
	// 1. A read operation (individual record lookup)
	// 2. OR a list operation (multiple records lookup)
	// 3. AND has a valid entity schema
	// 4. AND is not manually ignored

	_, hasRead := resource.Operations["read"]
	_, hasList := resource.Operations["list"]

	if !hasRead && !hasList {
		return false
	}

	// Must have an entity name
	if resource.EntityName == "" {
		return false
	}

	return true
}
