package main

import (
	"fmt"
	"strings"
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

	paths, pathsOk := spec["paths"].(map[string]interface{})
	if !pathsOk {
		fmt.Println("Warning: No paths found in spec")
	}

	entityMap := buildEntityRelationships(schemas)
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

func buildEntityRelationships(schemas map[string]interface{}) map[string]EntityRelationship {
	relationships := make(map[string]EntityRelationship)

	for schemaName := range schemas {
		if strings.HasSuffix(schemaName, "Entity") && !strings.Contains(schemaName, "Nullable") && !strings.Contains(schemaName, "Paginated") {
			baseName := strings.ToLower(strings.TrimSuffix(schemaName, "Entity"))

			rel := EntityRelationship{
				EntityName: schemaName,
			}

			createName := "create_" + baseName
			if _, exists := schemas[createName]; exists {
				rel.CreateSchema = createName
			}

			updateName := "update_" + baseName
			if _, exists := schemas[updateName]; exists {
				rel.UpdateSchema = updateName
			}

			relationships[schemaName] = rel
		}
	}

	return relationships
}

func normalizeSchemas(entityName string, entitySchema map[string]interface{}, requestSchema map[string]interface{}, requiredFields map[string]bool) []ConflictDetail {

	conflicts := make([]ConflictDetail, 0)

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		return conflicts
	}

	if entityProps == nil {
		entityProps = make(map[string]interface{})
		entitySchema["properties"] = entityProps
	}

	if requestProps == nil {
		return conflicts
	}

	for propName, requestProp := range requestProps {
		if requiredFields[propName] {
			if _, exists := entityProps[propName]; !exists {
				// If the property is required in the request but missing in the entity, add it as a nullable field
				copiedProp := deepCopyProperty(requestProp)
				if propMap, ok := copiedProp.(map[string]interface{}); ok {
					propMap["nullable"] = true
				}
				entityProps[propName] = copiedProp

				conflicts = append(conflicts, ConflictDetail{
					Schema:       entityName,
					Property:     propName,
					ConflictType: "missing-required-field",
					Resolution:   fmt.Sprintf("Added missing required field '%s' as nullable in entity schema", propName),
				})
			}
		}
	}

	// Check each property that exists in both schemas
	// Terraform requires exact matches for properties across requests and responses
	for propName, requestProp := range requestProps {
		if entityProp, exists := entityProps[propName]; exists {
			isRequired := requiredFields[propName]
			conflict := checkAndFixProperty(entityName, propName, entityProp, requestProp, entityProps, requestProps, isRequired)
			if conflict != nil {
				conflicts = append(conflicts, *conflict)
			}
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

	// Check for map vs class conflict - event if these are the same shape, they need to be the same type or generation will fail
	entityType, _ := entityPropMap["type"].(string)
	requestType, _ := requestPropMap["type"].(string)

	if entityType == "object" && requestType == "object" {
		_, entityHasProps := entityPropMap["properties"]
		_, entityHasAdditional := entityPropMap["additionalProperties"]
		requestPropsValue, requestHasProps := requestPropMap["properties"]
		_, requestHasAdditional := requestPropMap["additionalProperties"]

		if isRequired && requestHasProps && !entityHasProps {
			// If the request has properties but the entity does not, we need to convert the entity to have the required properties
			entityPropMap["properties"] = deepCopyProperty(requestPropsValue)
			delete(entityPropMap, "additionalProperties")
			// Ensure entity properties are nullable since some of them will not actually be returned by the api
			// we just inject them to have valid terraform structures
			entityPropMap["nullable"] = true
			entityProps[propName] = entityPropMap

			return &ConflictDetail{
				Schema:       entityName,
				Property:     propName,
				ConflictType: "required-structure-alignment",
				Resolution:   fmt.Sprintf("Converted entity property '%s' to match request structure", propName),
			}
		}

		if entityHasProps && requestHasProps {
			entityPropsObj, _ := entityPropMap["properties"].(map[string]interface{})
			requestPropsObj, _ := requestPropMap["properties"].(map[string]interface{})

			if len(entityPropsObj) == 0 && len(requestPropsObj) == 0 {
				// already the same - both empty properties
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

func deepCopyProperty(prop interface{}) interface{} {
	switch v := prop.(type) {
	case map[string]interface{}:
		copied := make(map[string]interface{})
		for key, val := range v {
			copied[key] = deepCopyProperty(val)
		}
		return copied
	case []interface{}:
		copied := make([]interface{}, len(v))
		for i, val := range v {
			copied[i] = deepCopyProperty(val)
		}
		return copied
	default:
		return v
	}
}
