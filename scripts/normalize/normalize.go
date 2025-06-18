package main

import (
	"fmt"
	"strings"

	"github.com/firehydrant/terraform-provider-firehydrant/scripts/common"
)

func normalizeSpec(spec map[string]interface{}) NormalizationReport {
	report := NormalizationReport{
		ConflictDetails: make([]ConflictDetail, 0),
	}

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No components found in spec")
		return report
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No schemas found in components")
		return report
	}

	// Add spec reference to each schema for context during normalization
	for _, schema := range schemas {
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			schemaMap["__spec"] = spec
		}
	}

	paths, pathsOk := spec["paths"].(map[string]interface{})
	if !pathsOk {
		fmt.Println("Warning: No paths found in spec")
	}

	entityMap := buildEntityRelationships(schemas, paths)
	requiredFieldsMap := make(map[string]map[string]bool)

	for entityName, related := range entityMap {
		requiredFields := make(map[string]bool)
		if related.CreateSchema != "" {
			if createSchema, ok := schemas[related.CreateSchema].(map[string]interface{}); ok {
				if required, ok := createSchema["required"].([]interface{}); ok {
					for _, field := range required {
						if fieldName, ok := field.(string); ok {
							requiredFields[fieldName] = true
						}
					}
				}
			}
		}
		requiredFieldsMap[entityName] = requiredFields
	}

	for entityName, related := range entityMap {
		entitySchema, ok := schemas[entityName].(map[string]interface{})
		if !ok {
			continue
		}

		requiredFields := requiredFieldsMap[entityName]

		if related.CreateSchema != "" {
			if createSchema, ok := schemas[related.CreateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, createSchema, requiredFields)
				report.ConflictDetails = append(report.ConflictDetails, conflicts...)
			}
		}

		if related.UpdateSchema != "" {
			if updateSchema, ok := schemas[related.UpdateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, updateSchema, requiredFields)
				report.ConflictDetails = append(report.ConflictDetails, conflicts...)
			}
		}
	}

	// Clean up spec reference
	for _, schema := range schemas {
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			delete(schemaMap, "__spec")
		}
	}

	globalFixes := applyGlobalNormalizations(schemas)
	report.ConflictDetails = append(report.ConflictDetails, globalFixes...)

	enumFixes := normalizeEnums(schemas)
	report.ConflictDetails = append(report.ConflictDetails, enumFixes...)

	if pathsOk {
		parameterFixes := normalizePathParameters(paths)
		report.ConflictDetails = append(report.ConflictDetails, parameterFixes...)

		pathEnumFixes := normalizePathEnums(paths)
		report.ConflictDetails = append(report.ConflictDetails, pathEnumFixes...)
	}

	report.TotalFixes = len(report.ConflictDetails)
	for _, detail := range report.ConflictDetails {
		if detail.ConflictType == "map-class" {
			report.MapClassFixes++
		} else {
			report.PropertyFixes++
		}
	}

	return report
}

func applyGlobalNormalizations(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("Applying global normalizations to %d schemas\n", len(schemas))

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		schemaConflicts := normalizeAdditionalProperties(schemaName, schemaMap, "")
		conflicts = append(conflicts, schemaConflicts...)
	}

	return conflicts
}

type EntityRelationship struct {
	EntityName   string
	CreateSchema string
	UpdateSchema string
}

func buildEntityRelationships(schemas map[string]interface{}, paths map[string]interface{}) map[string]EntityRelationship {
	relationships := make(map[string]EntityRelationship)

	// First: identify all entities
	for schemaName := range schemas {
		if strings.HasSuffix(schemaName, "Entity") && !strings.Contains(schemaName, "Nullable") && !strings.Contains(schemaName, "Paginated") {
			rel := EntityRelationship{
				EntityName: schemaName,
			}
			relationships[schemaName] = rel
		}
	}

	// Second: Try standard naming convention
	for entityName, rel := range relationships {
		baseName := strings.ToLower(strings.TrimSuffix(entityName, "Entity"))

		createName := "create_" + baseName
		if _, exists := schemas[createName]; exists {
			rel.CreateSchema = createName
			relationships[entityName] = rel
		}

		updateName := "update_" + baseName
		if _, exists := schemas[updateName]; exists {
			rel.UpdateSchema = updateName
			relationships[entityName] = rel
		}
	}

	// Third: Use paths to find non-standard mappings
	for pathName, pathItem := range paths {
		pathMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		// Check PUT, POST, PATCH operations
		for method, operation := range pathMap {
			if method != "put" && method != "post" && method != "patch" {
				continue
			}

			opMap, ok := operation.(map[string]interface{})
			if !ok {
				continue
			}

			requestSchema := getRequestSchema(opMap)
			responseEntity := getResponseEntity(opMap)

			if requestSchema != "" && responseEntity != "" {
				if rel, exists := relationships[responseEntity]; exists {
					operationID, _ := opMap["operationId"].(string)

					if strings.Contains(operationID, "create") ||
						(method == "put" && !strings.Contains(pathName, "{")) {
						if rel.CreateSchema == "" || rel.CreateSchema == requestSchema {
							rel.CreateSchema = requestSchema
							relationships[responseEntity] = rel
						}
					} else if strings.Contains(operationID, "update") ||
						(method == "patch") {
						if rel.UpdateSchema == "" || rel.UpdateSchema == requestSchema {
							rel.UpdateSchema = requestSchema
							relationships[responseEntity] = rel
						}
					}
				}
			}
		}
	}

	return relationships
}

func getRequestSchema(operation map[string]interface{}) string {
	if requestBody, ok := operation["requestBody"].(map[string]interface{}); ok {
		if content, ok := requestBody["content"].(map[string]interface{}); ok {
			if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
				if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
					if ref, ok := schema["$ref"].(string); ok {
						parts := strings.Split(ref, "/")
						if len(parts) > 0 {
							return parts[len(parts)-1]
						}
					}
				}
			}
		}
	}
	return ""
}

func getResponseEntity(operation map[string]interface{}) string {
	if responses, ok := operation["responses"].(map[string]interface{}); ok {
		for _, status := range []string{"200", "201", "202"} {
			if response, ok := responses[status].(map[string]interface{}); ok {
				if content, ok := response["content"].(map[string]interface{}); ok {
					if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
						if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
							if ref, ok := schema["$ref"].(string); ok {
								parts := strings.Split(ref, "/")
								if len(parts) > 0 {
									schemaName := parts[len(parts)-1]
									if strings.HasSuffix(schemaName, "Entity") {
										return schemaName
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return ""
}

func normalizeSchemas(entityName string, entitySchema map[string]interface{},
	requestSchema map[string]interface{}, requiredFields map[string]bool) []ConflictDetail {

	conflicts := make([]ConflictDetail, 0)

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil {
		entityProps = make(map[string]interface{})
		entitySchema["properties"] = entityProps
	}

	if requestProps == nil {
		return conflicts
	}

	// First pass: Add ALL required fields that are missing (existing logic)
	for propName := range requiredFields {
		requestProp, existsInRequest := requestProps[propName]
		if !existsInRequest {
			// Required field not in request schema - this shouldn't happen
			fmt.Printf("WARNING: Required field '%s' not found in request schema\n", propName)
			continue
		}

		_, existsInEntity := entityProps[propName]
		if !existsInEntity {
			copiedProp := deepCopyProperties(requestProp)

			// Make the field nullable in entity since it might not be returned by API
			if propMap, ok := copiedProp.(map[string]interface{}); ok {
				propMap["nullable"] = true
			}

			entityProps[propName] = copiedProp

			conflicts = append(conflicts, ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "missing-required-field",
				Resolution:   fmt.Sprintf("Added missing required field '%s' from request to entity", propName),
			})
		}
	}

	// Second pass: Check and fix ALL properties with deep structure analysis
	// Get schemas for ref resolution
	spec, _ := entitySchema["__spec"].(map[string]interface{})
	components, _ := spec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	for propName, requestProp := range requestProps {
		entityProp, exists := entityProps[propName]

		if exists {
			// Property exists in both - check if structures match
			if !common.ComparePropertyStructures(entityProp, requestProp, schemas) {
				// Structures don't match - normalize the entity property
				normalized := normalizePropertyStructure(entityProp, requestProp, schemas, propName)
				entityProps[propName] = normalized

				conflicts = append(conflicts, ConflictDetail{
					Schema:       entityName,
					Property:     propName,
					ConflictType: "structure-normalized",
					Resolution:   fmt.Sprintf("Normalized structure of '%s' to match request schema", propName),
				})
			} else {
				// Structures match but check for the existing special cases
				isRequired := requiredFields[propName]
				conflict := checkAndFixProperty(entityName, propName, entityProp, requestProp,
					entityProps, requestProps, isRequired)
				if conflict != nil {
					conflicts = append(conflicts, *conflict)
				}
			}
		} else {
			// Property missing in entity - add it
			copiedProp := deepCopyProperties(requestProp)

			if propMap, ok := copiedProp.(map[string]interface{}); ok {
				propMap["nullable"] = true
			}

			entityProps[propName] = copiedProp

			conflicts = append(conflicts, ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "missing-field",
				Resolution:   fmt.Sprintf("Added missing field '%s' from request to entity", propName),
			})
		}
	}

	return conflicts
}

func checkAndFixProperty(entityName, propName string, entityProp, requestProp interface{},
	entityProps, requestProps map[string]interface{}, isRequired bool) *ConflictDetail {

	entityPropMap, _ := entityProp.(map[string]interface{})
	requestPropMap, _ := requestProp.(map[string]interface{})

	if entityPropMap == nil || requestPropMap == nil {
		return nil
	}

	entityType, _ := entityPropMap["type"].(string)
	requestType, _ := requestPropMap["type"].(string)

	if isRequired && entityType == "object" && requestType == "object" {
		entityPropsValue, entityHasProps := entityPropMap["properties"].(map[string]interface{})
		requestPropsValue, requestHasProps := requestPropMap["properties"].(map[string]interface{})

		needsAlignment := false

		if requestHasProps && !entityHasProps {
			needsAlignment = true
		} else if requestHasProps && entityHasProps {
			// Both have properties - check if they match
			requestPropCount := len(requestPropsValue)
			entityPropCount := len(entityPropsValue)

			// Entity has empty properties, request has nested structure
			if entityPropCount == 0 && requestPropCount > 0 {
				needsAlignment = true
			} else if requestPropCount != entityPropCount {
				needsAlignment = true
			} else {
				// Check if property names match
				for propKey := range requestPropsValue {
					if _, exists := entityPropsValue[propKey]; !exists {
						needsAlignment = true
						break
					}
				}
			}
		}

		if needsAlignment {
			alignedProp := make(map[string]interface{})
			alignedProp["type"] = "object"
			alignedProp["nullable"] = true

			if desc, ok := requestPropMap["description"]; ok {
				alignedProp["description"] = desc
			}

			if requestHasProps {
				alignedProp["properties"] = deepCopyProperties(requestPropsValue)
			} else {
				alignedProp["properties"] = make(map[string]interface{})
			}

			entityProps[propName] = alignedProp

			return &ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "required-structure-alignment",
				Resolution:   fmt.Sprintf("Aligned structure of required field '%s' to match request schema", propName),
			}
		}
	}

	if entityType == "object" && requestType == "object" {
		_, entityHasProps := entityPropMap["properties"]
		_, entityHasAdditional := entityPropMap["additionalProperties"]
		_, requestHasProps := requestPropMap["properties"]
		_, requestHasAdditional := requestPropMap["additionalProperties"]

		if entityHasProps && requestHasProps {
			entityPropsObj, _ := entityPropMap["properties"].(map[string]interface{})
			requestPropsObj, _ := requestPropMap["properties"].(map[string]interface{})

			if len(entityPropsObj) == 0 && len(requestPropsObj) == 0 {
				return nil
			}
		}

		if entityHasAdditional && !requestHasAdditional && requestHasProps {
			delete(entityPropMap, "additionalProperties")
			entityPropMap["properties"] = map[string]interface{}{}
			entityProps[propName] = entityPropMap

			return &ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "map-class",
				Resolution:   "Converted entity from additionalProperties to empty properties",
			}
		}

		if requestHasAdditional && !entityHasAdditional && entityHasProps {
			delete(requestPropMap, "additionalProperties")
			requestPropMap["properties"] = map[string]interface{}{}
			requestProps[propName] = requestPropMap

			return &ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "map-class",
				Resolution:   "Converted request from additionalProperties to empty properties",
			}
		}
	}

	return nil
}

// deepCopyProperties recursively copies properties to ensure they are nullable
// We need to make sure everything is nullable since we are generally copying from request schemas into resposne entities with this function
// This is necessary so that required fields in request schemas are not ignored in generation and make it into our terraform resources, as they would otherwise mismatch between request and response
func deepCopyProperties(props interface{}) interface{} {
	switch v := props.(type) {
	case map[string]interface{}:
		copied := make(map[string]interface{})
		for key, val := range v {
			if propMap, ok := val.(map[string]interface{}); ok {

				copiedProp := make(map[string]interface{})
				for pk, pv := range propMap {
					copiedProp[pk] = deepCopyProperties(pv)
				}

				// Ensure nested objects are also nullable
				if propType, ok := copiedProp["type"].(string); ok && propType == "object" {
					copiedProp["nullable"] = true
				}

				copied[key] = copiedProp
			} else {
				copied[key] = deepCopyProperties(val)
			}
		}
		return copied
	case []interface{}:
		copied := make([]interface{}, len(v))
		for i, val := range v {
			copied[i] = deepCopyProperties(val)
		}
		return copied
	default:
		return v
	}
}

// normalizePropertyStructure aligns entity property structure with request property structure
func normalizePropertyStructure(entityProp, requestProp interface{}, schemas map[string]interface{}, propName string) interface{} {
	requestResolved, requestType := common.GetResolvedPropertyType(requestProp, schemas)
	entityResolved, _ := common.GetResolvedPropertyType(entityProp, schemas)

	if requestResolved == nil {
		return entityProp
	}

	normalized := make(map[string]interface{})

	if propType, ok := requestResolved["type"].(string); ok {
		normalized["type"] = propType
	}
	if description, ok := requestResolved["description"].(string); ok {
		normalized["description"] = description
	}

	// Always make nullable in entity as it might not be returned by API
	normalized["nullable"] = true

	if propType, _ := requestResolved["type"].(string); propType == "object" {
		if requestProps, ok := requestResolved["properties"].(map[string]interface{}); ok {
			normalizedProps := make(map[string]interface{})

			for subPropName, requestSubProp := range requestProps {
				normalizedProps[subPropName] = deepCopyProperties(requestSubProp)
			}

			if entityResolved != nil {
				if entityProps, ok := entityResolved["properties"].(map[string]interface{}); ok {
					for subPropName, entitySubProp := range entityProps {
						if _, existsInRequest := requestProps[subPropName]; !existsInRequest {
							// This is a computed field, keep it as-is
							normalizedProps[subPropName] = entitySubProp
						}
					}
				}
			}

			normalized["properties"] = normalizedProps
		}

		if required, ok := requestResolved["required"].([]interface{}); ok {
			normalized["required"] = required
		}
	} else if propType, _ := requestResolved["type"].(string); propType == "array" {
		if requestItems, ok := requestResolved["items"].(map[string]interface{}); ok {
			if entityResolved != nil && entityResolved["items"] != nil {
				normalized["items"] = normalizePropertyStructure(entityResolved["items"], requestItems, schemas, propName+"[]")
			} else {
				normalized["items"] = deepCopyProperties(requestItems)
			}
		}
	}

	if entityMap, ok := entityProp.(map[string]interface{}); ok {
		if ref, hadRef := entityMap["$ref"].(string); hadRef && requestType == "inline" {
			fmt.Printf("      Normalized %s: inlined $ref %s to match request structure\n", propName, ref)
		}
	}

	return normalized
}
