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

	for entityName, related := range entityMap {

		entitySchema, ok := schemas[entityName].(map[string]interface{})
		if !ok {
			continue
		}

		if related.CreateSchema != "" {
			if createSchema, ok := schemas[related.CreateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, createSchema)
				report.ConflictDetails = append(report.ConflictDetails, conflicts...)
			}
		}

		if related.UpdateSchema != "" {
			if updateSchema, ok := schemas[related.UpdateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, updateSchema)
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

func normalizeSchemas(entityName string, entitySchema map[string]interface{}, requestSchema map[string]interface{}) []ConflictDetail {

	conflicts := make([]ConflictDetail, 0)

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		return conflicts
	}

	// Check each property that exists in both schemas
	// Terraform requires exact matches for properties across requests and responses
	for propName, requestProp := range requestProps {
		if entityProp, exists := entityProps[propName]; exists {
			conflict := checkAndFixProperty(entityName, propName, entityProp, requestProp, entityProps, requestProps)
			if conflict != nil {
				conflicts = append(conflicts, *conflict)
			}
		}
	}

	return conflicts
}

func checkAndFixProperty(entityName, propName string, entityProp, requestProp interface{},
	entityProps, requestProps map[string]interface{}) *ConflictDetail {

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
		_, requestHasProps := requestPropMap["properties"]
		_, requestHasAdditional := requestPropMap["additionalProperties"]

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
