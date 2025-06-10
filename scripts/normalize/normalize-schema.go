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

	report := normalizeSpec(spec)

	printNormalizationReport(report)

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
	fmt.Printf("   Enum extractions: %d\n", report.EnumExtractions)
}

func printUsage() {
	fmt.Println("OpenAPI Schema Normalizer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-normalize <input.json> [output.json]")
}

type NormalizationReport struct {
	TotalFixes      int
	MapClassFixes   int
	PropertyFixes   int
	EnumExtractions int
	ConflictDetails []ConflictDetail
	ExtractedEnums  []EnumExtraction
}

type ConflictDetail struct {
	Schema       string
	Property     string
	ConflictType string
	Resolution   string
}

type EntityRelationship struct {
	EntityName   string
	CreateSchema string
	UpdateSchema string
}

// EnumInfo represents an enum found in the spec
type EnumInfo struct {
	Type        string   `json:"type"`
	Values      []string `json:"values"`
	Description string   `json:"description"`
	Signature   string   `json:"signature"` // For deduplication
}

// EnumLocation tracks where an enum is used
type EnumLocation struct {
	SchemaName   string   `json:"schema_name"`
	PropertyPath string   `json:"property_path"`
	EnumInfo     EnumInfo `json:"enum_info"`
}

// EnumExtraction represents the result of enum extraction
type EnumExtraction struct {
	ExtractedName string         `json:"extracted_name"`
	EnumInfo      EnumInfo       `json:"enum_info"`
	Locations     []EnumLocation `json:"locations"`
}

func normalizeSpec(spec map[string]interface{}) NormalizationReport {
	report := NormalizationReport{
		ConflictDetails: make([]ConflictDetail, 0),
		ExtractedEnums:  make([]EnumExtraction, 0),
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

	// Get paths for parameter normalization
	paths, pathsOk := spec["paths"].(map[string]interface{})
	if !pathsOk {
		fmt.Println("Warning: No paths found in spec")
	}

	// Step 1: Extract and normalize enums
	extractions := normalizeEnums(spec)
	report.ExtractedEnums = extractions
	report.EnumExtractions = len(extractions)

	// Step 2: Build entity relationships
	entityMap := buildEntityRelationships(schemas)

	// Step 3: Normalize each entity and its related schemas
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

	// Step 4: Apply global normalizations to schemas
	globalFixes := applyGlobalNormalizations(schemas)
	report.ConflictDetails = append(report.ConflictDetails, globalFixes...)

	// Step 5: Normalize path parameters to match entity IDs
	if pathsOk {
		parameterFixes := normalizePathParameters(paths, schemas)
		report.ConflictDetails = append(report.ConflictDetails, parameterFixes...)
	}

	// Calculate totals
	report.TotalFixes = len(report.ConflictDetails) + report.EnumExtractions
	for _, detail := range report.ConflictDetails {
		if detail.ConflictType == "map-class" {
			report.MapClassFixes++
		} else {
			report.PropertyFixes++
		}
	}

	return report
}

func normalizeEnums(spec map[string]interface{}) []EnumExtraction {
	fmt.Printf("\n=== Extracting and Normalizing Enums ===\n")

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No components found in spec")
		return []EnumExtraction{}
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No schemas found in components")
		return []EnumExtraction{}
	}

	// Step 1: Collect all enum usages
	enumLocations := collectAllEnums(schemas)
	fmt.Printf("Found %d enum usages across schemas\n", len(enumLocations))

	if len(enumLocations) == 0 {
		return []EnumExtraction{}
	}

	// Step 2: Group by signature for deduplication
	enumGroups := groupEnumsBySignature(enumLocations)
	fmt.Printf("Found %d unique enum signatures\n", len(enumGroups))

	// Step 3: Extract enums that appear in multiple locations or are complex
	extractions := []EnumExtraction{}

	for signature, locations := range enumGroups {
		if len(locations) > 1 {
			// Multiple usages - extract to shared enum
			extraction := extractSharedEnum(signature, locations, schemas)
			extractions = append(extractions, extraction)
			fmt.Printf("🔄 Extracted shared enum: %s (used in %d locations)\n",
				extraction.ExtractedName, len(locations))
		} else if shouldExtractSingleEnum(locations[0]) {
			// Single usage but complex enough to warrant extraction
			extraction := extractSingleEnum(locations[0], schemas)
			extractions = append(extractions, extraction)
			fmt.Printf("📤 Extracted single enum: %s\n", extraction.ExtractedName)
		}
	}

	return extractions
}

// Collect all enum usages from schemas
func collectAllEnums(schemas map[string]interface{}) []EnumLocation {
	var locations []EnumLocation

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		// Recursively find enums in this schema
		enumsInSchema := findEnumsInSchema(schemaName, schemaMap, "")
		locations = append(locations, enumsInSchema...)
	}

	return locations
}

// Recursively find enums in a schema object
func findEnumsInSchema(schemaName string, obj interface{}, path string) []EnumLocation {
	var locations []EnumLocation

	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this object is an enum
		if enumValues, hasEnum := v["enum"]; hasEnum {
			if enumArray, ok := enumValues.([]interface{}); ok && len(enumArray) > 0 {
				// Convert enum values to strings
				var values []string
				for _, val := range enumArray {
					if str, ok := val.(string); ok {
						values = append(values, str)
					}
				}

				if len(values) > 0 {
					enumType := "string"
					if typeVal, hasType := v["type"].(string); hasType {
						enumType = typeVal
					}

					description := ""
					if desc, hasDesc := v["description"].(string); hasDesc {
						description = desc
					}

					location := EnumLocation{
						SchemaName:   schemaName,
						PropertyPath: path,
						EnumInfo: EnumInfo{
							Type:        enumType,
							Values:      values,
							Description: description,
							Signature:   createEnumSignature(enumType, values),
						},
					}
					locations = append(locations, location)
				}
			}
		}

		// Recursively check nested objects (but skip certain keys)
		for key, value := range v {
			// Skip certain keys that aren't property definitions
			if key == "enum" || key == "type" || key == "description" || key == "example" {
				continue
			}

			newPath := path
			if newPath != "" {
				newPath += "." + key
			} else {
				newPath = key
			}
			nested := findEnumsInSchema(schemaName, value, newPath)
			locations = append(locations, nested...)
		}

	case []interface{}:
		// Check array items
		for i, item := range v {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			nested := findEnumsInSchema(schemaName, item, newPath)
			locations = append(locations, nested...)
		}
	}

	return locations
}

// Create a signature for deduplication
func createEnumSignature(enumType string, values []string) string {
	// Sort values for consistent signature
	sortedValues := make([]string, len(values))
	copy(sortedValues, values)

	// Simple bubble sort
	for i := 0; i < len(sortedValues); i++ {
		for j := i + 1; j < len(sortedValues); j++ {
			if sortedValues[i] > sortedValues[j] {
				sortedValues[i], sortedValues[j] = sortedValues[j], sortedValues[i]
			}
		}
	}

	return fmt.Sprintf("%s:[%s]", enumType, strings.Join(sortedValues, ","))
}

// Group enum locations by signature
func groupEnumsBySignature(locations []EnumLocation) map[string][]EnumLocation {
	groups := make(map[string][]EnumLocation)

	for _, location := range locations {
		signature := location.EnumInfo.Signature
		groups[signature] = append(groups[signature], location)
	}

	return groups
}

// Check if a single enum should be extracted
func shouldExtractSingleEnum(location EnumLocation) bool {
	// Extract if:
	// - Has more than 3 values (complex enum)
	// - Values suggest it's a standard enum (encoding, status, etc.)
	if len(location.EnumInfo.Values) > 3 {
		return true
	}

	// Check for common enum patterns
	return isStandardEnumType(location.EnumInfo.Values)
}

// Check if enum values represent standard types
func isStandardEnumType(values []string) bool {
	return isEncodingEnum(values) || isStatusEnum(values) || isLevelEnum(values) || isTypeEnum(values)
}

// Extract a shared enum used in multiple locations
func extractSharedEnum(signature string, locations []EnumLocation, schemas map[string]interface{}) EnumExtraction {
	// Generate a meaningful name for the extracted enum
	extractedName := generateExtractedEnumName(locations)

	// Ensure the name is unique
	extractedName = ensureUniqueName(extractedName, schemas)

	// Use the first location's enum info as the canonical definition
	canonicalEnum := locations[0].EnumInfo

	// Create the extracted enum schema
	enumSchema := map[string]interface{}{
		"type": canonicalEnum.Type,
		"enum": convertStringsToInterfaces(canonicalEnum.Values),
	}

	if canonicalEnum.Description != "" {
		enumSchema["description"] = canonicalEnum.Description
	}

	// Add to schemas
	schemas[extractedName] = enumSchema

	// Replace all usages with $ref
	for _, location := range locations {
		replaceEnumWithRef(schemas, location, extractedName)
	}

	return EnumExtraction{
		ExtractedName: extractedName,
		EnumInfo:      canonicalEnum,
		Locations:     locations,
	}
}

// Extract a single complex enum
func extractSingleEnum(location EnumLocation, schemas map[string]interface{}) EnumExtraction {
	extractedName := generateSingleEnumName(location)
	extractedName = ensureUniqueName(extractedName, schemas)

	// Create the extracted enum schema
	enumSchema := map[string]interface{}{
		"type": location.EnumInfo.Type,
		"enum": convertStringsToInterfaces(location.EnumInfo.Values),
	}

	if location.EnumInfo.Description != "" {
		enumSchema["description"] = location.EnumInfo.Description
	}

	// Add to schemas
	schemas[extractedName] = enumSchema

	// Replace usage with $ref
	replaceEnumWithRef(schemas, location, extractedName)

	return EnumExtraction{
		ExtractedName: extractedName,
		EnumInfo:      location.EnumInfo,
		Locations:     []EnumLocation{location},
	}
}

// Generate name for shared enum
func generateExtractedEnumName(locations []EnumLocation) string {
	// Try to infer from enum values
	if len(locations) > 0 {
		values := locations[0].EnumInfo.Values
		if isEncodingEnum(values) {
			return "EncodingType"
		}
		if isStatusEnum(values) {
			return "StatusType"
		}
		if isLevelEnum(values) {
			return "LevelType"
		}
		if isTypeEnum(values) {
			return "ItemType"
		}
	}

	// Fallback to generic name
	return fmt.Sprintf("SharedEnum%d", len(locations[0].EnumInfo.Values))
}

// Generate name for single enum
func generateSingleEnumName(location EnumLocation) string {
	// Use property path to create meaningful name
	parts := strings.Split(location.PropertyPath, ".")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if lastPart != "properties" && lastPart != "" {
			return toPascalCase(lastPart) + "Type"
		}
	}

	// Try to infer from enum values
	values := location.EnumInfo.Values
	if isEncodingEnum(values) {
		return "EncodingType"
	}
	if isStatusEnum(values) {
		return "StatusType"
	}
	if isLevelEnum(values) {
		return "LevelType"
	}

	return "ExtractedEnumType"
}

// Ensure extracted enum name is unique
func ensureUniqueName(baseName string, schemas map[string]interface{}) string {
	if _, exists := schemas[baseName]; !exists {
		return baseName
	}

	// Add counter suffix if name exists
	counter := 1
	for {
		candidateName := fmt.Sprintf("%s%d", baseName, counter)
		if _, exists := schemas[candidateName]; !exists {
			return candidateName
		}
		counter++
	}
}

// Helper functions for enum type detection
func isEncodingEnum(values []string) bool {
	encodingPatterns := []string{"json", "yaml", "xml", "text", "application"}
	for _, value := range values {
		lowerValue := strings.ToLower(value)
		for _, pattern := range encodingPatterns {
			if strings.Contains(lowerValue, pattern) {
				return true
			}
		}
	}
	return false
}

func isStatusEnum(values []string) bool {
	statusPatterns := []string{"active", "inactive", "pending", "completed", "failed", "success", "error", "open", "closed"}
	for _, value := range values {
		lowerValue := strings.ToLower(value)
		for _, pattern := range statusPatterns {
			if lowerValue == pattern {
				return true
			}
		}
	}
	return false
}

func isLevelEnum(values []string) bool {
	levelPatterns := []string{"low", "medium", "high", "debug", "info", "warn", "error", "critical"}
	for _, value := range values {
		lowerValue := strings.ToLower(value)
		for _, pattern := range levelPatterns {
			if lowerValue == pattern {
				return true
			}
		}
	}
	return false
}

func isTypeEnum(values []string) bool {
	// Check if any value contains "type" or represents common type patterns
	for _, value := range values {
		lowerValue := strings.ToLower(value)
		if strings.Contains(lowerValue, "type") {
			return true
		}
	}

	// Check for common type patterns
	typePatterns := []string{"string", "number", "boolean", "object", "array", "null"}
	for _, value := range values {
		lowerValue := strings.ToLower(value)
		for _, pattern := range typePatterns {
			if lowerValue == pattern {
				return true
			}
		}
	}
	return false
}

// Convert string slice to interface slice (for JSON marshaling)
func convertStringsToInterfaces(strings []string) []interface{} {
	result := make([]interface{}, len(strings))
	for i, s := range strings {
		result[i] = s
	}
	return result
}

// Replace enum definition with $ref
func replaceEnumWithRef(schemas map[string]interface{}, location EnumLocation, refName string) {
	schema, ok := schemas[location.SchemaName].(map[string]interface{})
	if !ok {
		return
	}

	// Navigate to the enum location and replace it
	if location.PropertyPath == "" {
		// Enum is at schema root (shouldn't happen for properties, but handle it)
		return
	}

	// Navigate through the path to find the enum
	parts := strings.Split(location.PropertyPath, ".")
	current := schema

	// Navigate to parent of the enum
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if next, ok := current[part].(map[string]interface{}); ok {
			current = next
		} else {
			return // Path not found
		}
	}

	// Replace the enum with $ref
	finalPart := parts[len(parts)-1]
	current[finalPart] = map[string]interface{}{
		"$ref": fmt.Sprintf("#/components/schemas/%s", refName),
	}
}

// Convert string to PascalCase
func toPascalCase(s string) string {
	if s == "" {
		return s
	}

	// Handle snake_case and kebab-case
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(string(part[0])))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	return result.String()
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

// Normalize path parameters to match entity ID types
func normalizePathParameters(paths map[string]interface{}, schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Path Parameters ===\n")

	for pathName, pathItem := range paths {
		pathMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		// Check all HTTP methods in this path
		methods := []string{"get", "post", "put", "patch", "delete"}
		for _, method := range methods {
			if operation, exists := pathMap[method]; exists {
				opMap, ok := operation.(map[string]interface{})
				if !ok {
					continue
				}

				// Check parameters in this operation
				if parameters, hasParams := opMap["parameters"]; hasParams {
					paramsList, ok := parameters.([]interface{})
					if !ok {
						continue
					}

					for _, param := range paramsList {
						paramMap, ok := param.(map[string]interface{})
						if !ok {
							continue
						}

						// Check if this is a path parameter with integer type that should be string
						paramIn, _ := paramMap["in"].(string)
						paramName, _ := paramMap["name"].(string)

						if paramIn == "path" && (strings.Contains(paramName, "id") || strings.HasSuffix(paramName, "_id")) {
							schema, hasSchema := paramMap["schema"]
							if hasSchema {
								schemaMap, ok := schema.(map[string]interface{})
								if ok {
									paramType, _ := schemaMap["type"].(string)
									paramFormat, _ := schemaMap["format"].(string)

									// Convert integer ID parameters to string
									if paramType == "integer" {
										fmt.Printf("  Found integer ID parameter: %s %s.%s (type: %s, format: %s)\n",
											method, pathName, paramName, paramType, paramFormat)

										schemaMap["type"] = "string"
										delete(schemaMap, "format") // Remove int32/int64 format

										conflicts = append(conflicts, ConflictDetail{
											Schema:       fmt.Sprintf("path:%s", pathName),
											Property:     fmt.Sprintf("%s.%s", method, paramName),
											ConflictType: "parameter-type",
											Resolution:   fmt.Sprintf("Converted path parameter %s from integer to string", paramName),
										})

										fmt.Printf("    ✅ Converted %s parameter from integer to string\n", paramName)
									}
								}
							}
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
	fmt.Printf("Property fixes: %d\n", report.PropertyFixes)
	fmt.Printf("Enum extractions: %d\n", report.EnumExtractions)

	if len(report.ExtractedEnums) > 0 {
		fmt.Println("\nExtracted enums:")
		for _, extraction := range report.ExtractedEnums {
			fmt.Printf("  - %s: %v (used in %d locations)\n",
				extraction.ExtractedName, extraction.EnumInfo.Values, len(extraction.Locations))
			for _, location := range extraction.Locations {
				fmt.Printf("    └─ %s.%s\n", location.SchemaName, location.PropertyPath)
			}
		}
	}

	if len(report.ConflictDetails) > 0 {
		fmt.Println("\nOther fixes applied:")
		for _, detail := range report.ConflictDetails {
			fmt.Printf("  - %s.%s [%s]: %s\n",
				detail.Schema, detail.Property, detail.ConflictType, detail.Resolution)
		}
	}
}
