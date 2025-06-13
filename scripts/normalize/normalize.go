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
		fmt.Printf("Analyzing entity: %s\n", entityName)

		entitySchema, ok := schemas[entityName].(map[string]interface{})
		if !ok {
			continue
		}

		// Check against create schema
		if related.CreateSchema != "" {
			if createSchema, ok := schemas[related.CreateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, related.CreateSchema, createSchema)
				report.ConflictDetails = append(report.ConflictDetails, conflicts...)
			}
		}

		// Check against update schema
		if related.UpdateSchema != "" {
			if updateSchema, ok := schemas[related.UpdateSchema].(map[string]interface{}); ok {
				conflicts := normalizeSchemas(entityName, entitySchema, related.UpdateSchema, updateSchema)
				report.ConflictDetails = append(report.ConflictDetails, conflicts...)
			}
		}
	}

	// Apply global normalizations to schemas
	globalFixes := applyGlobalNormalizations(schemas)
	report.ConflictDetails = append(report.ConflictDetails, globalFixes...)

	enumFixes := normalizeEnums(schemas)
	report.ConflictDetails = append(report.ConflictDetails, enumFixes...)

	// Normalize path parameters to match entity IDs
	if pathsOk {
		parameterFixes := normalizePathParameters(paths)
		report.ConflictDetails = append(report.ConflictDetails, parameterFixes...)
	}

	// Calculate totals
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

	fmt.Printf("\n=== Scanning for all additionalProperties instances ===\n")
	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		additionalPropsFound := findAllAdditionalProperties(schemaName, schemaMap, "")
		if len(additionalPropsFound) > 0 {
			fmt.Printf("Schema %s has additionalProperties at:\n", schemaName)
			for _, path := range additionalPropsFound {
				fmt.Printf("  - %s\n", path)
			}
		}
	}

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		fmt.Printf("  Normalizing schema: %s\n", schemaName)
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

func normalizeSchemas(entityName string, entitySchema map[string]interface{},
	requestName string, requestSchema map[string]interface{}) []ConflictDetail {

	conflicts := make([]ConflictDetail, 0)

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		return conflicts
	}

	fmt.Printf("  Comparing %s vs %s\n", entityName, requestName)
	fmt.Printf("    Entity properties: %d, Request properties: %d\n", len(entityProps), len(requestProps))

	// Check each property that exists in both schemas
	// Terraform requires exact matches for properties across requests and responses
	for propName, requestProp := range requestProps {
		if entityProp, exists := entityProps[propName]; exists {
			fmt.Printf("    Checking property: %s\n", propName)
			conflict := checkAndFixProperty(entityName, propName, entityProp, requestProp, entityProps, requestProps)
			if conflict != nil {
				fmt.Printf("    ✅ Fixed: %s - %s\n", propName, conflict.Resolution)
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

	// Check for map vs class conflict
	entityType, _ := entityPropMap["type"].(string)
	requestType, _ := requestPropMap["type"].(string)

	fmt.Printf("      Property types - Entity: %s, Request: %s\n", entityType, requestType)

	if entityType == "object" && requestType == "object" {
		_, entityHasProps := entityPropMap["properties"]
		_, entityHasAdditional := entityPropMap["additionalProperties"]
		_, requestHasProps := requestPropMap["properties"]
		_, requestHasAdditional := requestPropMap["additionalProperties"]

		fmt.Printf("      Entity - hasProps: %v, hasAdditional: %v\n", entityHasProps, entityHasAdditional)
		fmt.Printf("      Request - hasProps: %v, hasAdditional: %v\n", requestHasProps, requestHasAdditional)

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
