package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := inputPath
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	fmt.Printf("=== OpenAPI Schema Normalizer ===\n")
	fmt.Printf("Input: %s\n", inputPath)
	fmt.Printf("Output: %s\n\n", outputPath)

	// Read the spec
	specData, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading spec: %v\n", err)
		os.Exit(1)
	}

	var spec map[string]interface{}
	if err := json.Unmarshal(specData, &spec); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Normalize the spec
	report := normalizeSpec(spec)

	// Print report
	printNormalizationReport(report)

	// Write normalized spec
	normalizedData, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling normalized spec: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outputPath, normalizedData, 0644); err != nil {
		fmt.Printf("Error writing normalized spec: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Successfully normalized OpenAPI spec\n")
	fmt.Printf("   Total fixes applied: %d\n", report.TotalFixes)
}

func printUsage() {
	fmt.Println("OpenAPI Schema Normalizer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-normalize <input.json> [output.json]")
}

// ============================================================================
// NORMALIZATION LOGIC
// ============================================================================

type NormalizationReport struct {
	TotalFixes      int
	MapClassFixes   int
	PropertyFixes   int
	ConflictDetails []ConflictDetail
}

type ConflictDetail struct {
	Schema       string
	Property     string
	ConflictType string
	Resolution   string
}

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

	// Build entity relationships
	entityMap := buildEntityRelationships(schemas)

	// Normalize each entity and its related schemas
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

	// Apply global normalizations
	globalFixes := applyGlobalNormalizations(schemas)
	report.ConflictDetails = append(report.ConflictDetails, globalFixes...)

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

			// Look for create schema
			createName := "create_" + baseName
			if _, exists := schemas[createName]; exists {
				rel.CreateSchema = createName
			}

			// Look for update schema
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

		// Normalize to consistent structure without adding additionalProperties
		// Option 1: Both have empty properties - make them consistent
		if entityHasProps && requestHasProps {
			entityPropsObj, _ := entityPropMap["properties"].(map[string]interface{})
			requestPropsObj, _ := requestPropMap["properties"].(map[string]interface{})

			if len(entityPropsObj) == 0 && len(requestPropsObj) == 0 {
				// Both have empty properties - this is already consistent
				fmt.Printf("      Both have empty properties - already consistent\n")
				return nil
			}
		}

		// Option 2: One has properties, one has additionalProperties - make them both use properties
		if entityHasAdditional && !requestHasAdditional && requestHasProps {
			// Entity uses additionalProperties, request uses properties
			// Convert entity to use empty properties like request
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
			// Request uses additionalProperties, entity uses properties
			// Convert request to use empty properties like entity
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

func applyGlobalNormalizations(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("Applying global normalizations to %d schemas\n", len(schemas))

	// Fix common patterns across all schemas
	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		// Apply global normalizations to ALL schemas (don't skip request schemas)
		props, ok := schemaMap["properties"].(map[string]interface{})
		if !ok {
			continue
		}

		fmt.Printf("  Checking schema: %s\n", schemaName)

		// Check for common problematic properties
		for propName, prop := range props {
			propMap, ok := prop.(map[string]interface{})
			if !ok {
				continue
			}

			// Fix empty properties objects - but don't add additionalProperties
			if propType, _ := propMap["type"].(string); propType == "object" {
				if propsObj, hasProps := propMap["properties"].(map[string]interface{}); hasProps && len(propsObj) == 0 {
					_, hasAdditional := propMap["additionalProperties"]

					// Debug output for labels property specifically
					if propName == "labels" {
						fmt.Printf("    Found labels property in %s:\n", schemaName)
						fmt.Printf("      Has empty properties: %v\n", hasProps)
						fmt.Printf("      Has additionalProperties: %v\n", hasAdditional)
					}

					// Only normalize if it has additionalProperties - convert to consistent empty properties
					if hasAdditional {
						delete(propMap, "additionalProperties")
						// Keep the empty properties object for consistency
						props[propName] = propMap

						conflicts = append(conflicts, ConflictDetail{
							Schema:       schemaName,
							Property:     propName,
							ConflictType: "map-class",
							Resolution:   "Converted additionalProperties to empty properties for consistency",
						})

						if propName == "labels" {
							fmt.Printf("      ✅ Converted labels from additionalProperties to empty properties in %s\n", schemaName)
						}
					}
				}
			}
		}
	}

	return conflicts
}

func printNormalizationReport(report NormalizationReport) {
	fmt.Println("\n=== Normalization Report ===")
	fmt.Printf("Total fixes applied: %d\n", report.TotalFixes)
	fmt.Printf("Map/Class fixes: %d\n", report.MapClassFixes)
	fmt.Printf("Other property fixes: %d\n", report.PropertyFixes)

	if len(report.ConflictDetails) > 0 {
		fmt.Println("\nDetailed fixes:")
		for _, detail := range report.ConflictDetails {
			fmt.Printf("  - %s.%s [%s]: %s\n",
				detail.Schema, detail.Property, detail.ConflictType, detail.Resolution)
		}
	}
}
