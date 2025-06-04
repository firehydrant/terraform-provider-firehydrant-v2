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

	// First pass: Find and report all additionalProperties instances
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

	// Second pass: Fix all additionalProperties instances
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

// Recursively find all additionalProperties in a schema
func findAllAdditionalProperties(schemaName string, obj interface{}, path string) []string {
	var found []string

	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this object has additionalProperties
		if _, hasAdditional := v["additionalProperties"]; hasAdditional {
			fullPath := schemaName
			if path != "" {
				fullPath += "." + path
			}
			found = append(found, fullPath)
		}

		// Recursively check all nested objects
		for key, value := range v {
			newPath := path
			if newPath != "" {
				newPath += "." + key
			} else {
				newPath = key
			}
			nested := findAllAdditionalProperties(schemaName, value, newPath)
			found = append(found, nested...)
		}
	case []interface{}:
		// Check array items
		for i, item := range v {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			nested := findAllAdditionalProperties(schemaName, item, newPath)
			found = append(found, nested...)
		}
	}

	return found
}

// Recursively normalize all additionalProperties in a schema
func normalizeAdditionalProperties(schemaName string, obj interface{}, path string) []ConflictDetail {
	var conflicts []ConflictDetail

	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this object has additionalProperties
		if _, hasAdditional := v["additionalProperties"]; hasAdditional {
			objType, _ := v["type"].(string)
			_, hasProperties := v["properties"]

			// Remove additionalProperties if:
			// 1. It's explicitly type "object", OR
			// 2. It has "properties" (implicit object), OR
			// 3. It has additionalProperties but no other structure
			if objType == "object" || hasProperties || (!hasProperties && hasAdditional) {
				// Remove additionalProperties and ensure empty properties
				delete(v, "additionalProperties")
				if !hasProperties {
					v["properties"] = map[string]interface{}{}
				}

				fullPath := schemaName
				if path != "" {
					fullPath += "." + path
				}

				conflicts = append(conflicts, ConflictDetail{
					Schema:       schemaName,
					Property:     path,
					ConflictType: "map-class",
					Resolution:   fmt.Sprintf("Converted additionalProperties to empty properties at %s", fullPath),
				})

				fmt.Printf("      ✅ Converted additionalProperties to empty properties at %s\n", fullPath)
			}
		}

		// Recursively normalize all nested objects
		for key, value := range v {
			newPath := path
			if newPath != "" {
				newPath += "." + key
			} else {
				newPath = key
			}
			nested := normalizeAdditionalProperties(schemaName, value, newPath)
			conflicts = append(conflicts, nested...)
		}
	case []interface{}:
		// Normalize array items
		for i, item := range v {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			nested := normalizeAdditionalProperties(schemaName, item, newPath)
			conflicts = append(conflicts, nested...)
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
