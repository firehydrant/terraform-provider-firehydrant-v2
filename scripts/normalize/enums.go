package main

import (
	"fmt"
	"strings"
	"unicode"
)

// EnumProperty represents an enum found during normalization
type EnumNormalizationInfo struct {
	SchemaName   string
	PropertyPath string
	PropertyName string
	EnumValues   []string
	Target       map[string]interface{}
}

// Main enum normalization function
func normalizeEnums(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Enum Properties ===\n")

	// Find all enum properties across all schemas
	allEnums := findAllEnumProperties(schemas)

	if len(allEnums) == 0 {
		fmt.Printf("No enum properties found to normalize\n")
		return conflicts
	}

	fmt.Printf("Found %d enum properties to normalize\n", len(allEnums))

	// Transform each enum property
	for _, enumInfo := range allEnums {
		conflict := transformEnumProperty(enumInfo)
		if conflict != nil {
			conflicts = append(conflicts, *conflict)
		}
	}

	fmt.Printf("Successfully normalized %d enum properties\n", len(conflicts))
	return conflicts
}

// Find all enum properties recursively across all schemas
func findAllEnumProperties(schemas map[string]interface{}) []EnumNormalizationInfo {
	var allEnums []EnumNormalizationInfo

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		schemaEnums := findEnumsInSchemaRecursive(schemaName, schemaMap, "", schemaMap)
		allEnums = append(allEnums, schemaEnums...)
	}

	return allEnums
}

// Recursively find enum properties within a single schema
func findEnumsInSchemaRecursive(schemaName string, currentObj map[string]interface{}, path string, rootSchema map[string]interface{}) []EnumNormalizationInfo {
	var enums []EnumNormalizationInfo

	// Check if current object has enum values
	if enumValues, hasEnum := currentObj["enum"]; hasEnum {
		if enumArray, ok := enumValues.([]interface{}); ok && len(enumArray) > 0 {
			// Convert enum values to strings
			var stringValues []string
			for _, val := range enumArray {
				if str, ok := val.(string); ok {
					stringValues = append(stringValues, str)
				}
			}

			if len(stringValues) > 0 {
				// Extract property name from path
				propertyName := extractPropertyNameFromNormalizationPath(path, schemaName)

				enumInfo := EnumNormalizationInfo{
					SchemaName:   schemaName,
					PropertyPath: path,
					PropertyName: propertyName,
					EnumValues:   stringValues,
					Target:       currentObj, // Reference to modify
				}

				enums = append(enums, enumInfo)
				fmt.Printf("📋 Found enum in %s.%s: %d values\n", schemaName, propertyName, len(stringValues))
			}
		}
	}

	// Recursively check properties
	if properties, hasProps := currentObj["properties"].(map[string]interface{}); hasProps {
		for propName, propValue := range properties {
			if propMap, ok := propValue.(map[string]interface{}); ok {
				newPath := propName
				if path != "" {
					newPath = fmt.Sprintf("%s.properties.%s", path, propName)
				} else {
					newPath = fmt.Sprintf("properties.%s", propName)
				}

				propEnums := findEnumsInSchemaRecursive(schemaName, propMap, newPath, rootSchema)
				enums = append(enums, propEnums...)
			}
		}
	}

	// Check array items
	if items, hasItems := currentObj["items"].(map[string]interface{}); hasItems {
		newPath := path
		if path != "" {
			newPath = fmt.Sprintf("%s.items", path)
		} else {
			newPath = "items"
		}

		itemEnums := findEnumsInSchemaRecursive(schemaName, items, newPath, rootSchema)
		enums = append(enums, itemEnums...)
	}

	return enums
}

// Transform a single enum property
func transformEnumProperty(enumInfo EnumNormalizationInfo) *ConflictDetail {
	// Generate x-speakeasy-enums members
	var members []map[string]interface{}

	for _, value := range enumInfo.EnumValues {
		memberName := generateEnumMemberName(value, enumInfo.PropertyName, enumInfo.SchemaName)
		members = append(members, map[string]interface{}{
			"name":  memberName,
			"value": value,
		})
	}

	// Remove original enum and add x-speakeasy-enums
	delete(enumInfo.Target, "enum")
	enumInfo.Target["x-speakeasy-enums"] = members

	// Create conflict detail for reporting
	targetPath := enumInfo.SchemaName
	if enumInfo.PropertyPath != "" {
		targetPath = fmt.Sprintf("%s.%s", enumInfo.SchemaName, enumInfo.PropertyPath)
	}

	return &ConflictDetail{
		Schema:       enumInfo.SchemaName,
		Property:     enumInfo.PropertyName,
		ConflictType: "enum-normalization",
		Resolution:   fmt.Sprintf("Replaced enum array with x-speakeasy-enums at %s (%d values)", targetPath, len(members)),
	}
}

// Generate unique enum member name using EntityName+FieldName+EnumValue
func generateEnumMemberName(value, fieldName, entityName string) string {
	// Handle empty/whitespace-only values
	originalValue := value
	if strings.TrimSpace(value) == "" {
		value = "Empty"
	}

	// Clean and convert parts
	cleanEntityName := convertToEnumMemberName(strings.TrimSuffix(entityName, "Entity"))
	cleanFieldName := convertToEnumMemberName(fieldName)
	cleanValue := convertToEnumMemberName(value)

	// Combine all three parts
	memberName := cleanEntityName + cleanFieldName + cleanValue

	// Ensure it starts with a letter (Go requirement)
	if len(memberName) > 0 && unicode.IsDigit(rune(memberName[0])) {
		memberName = "Value" + memberName
	}

	// Final fallback
	if memberName == "" {
		memberName = cleanEntityName + cleanFieldName + "Unknown"
	}

	// Debug logging for empty strings
	if originalValue == "" {
		fmt.Printf("🔍 Empty string enum: entity=%s, field=%s, value='%s' -> memberName='%s'\n",
			entityName, fieldName, originalValue, memberName)
	}

	return memberName
}

// Extract property name from normalization path
func extractPropertyNameFromNormalizationPath(path, schemaName string) string {
	if path == "" {
		// Schema-level enum, use schema name without Entity suffix
		return strings.TrimSuffix(schemaName, "Entity")
	}

	// Extract the last property name from the path
	// Examples:
	// "properties.status" -> "status"
	// "properties.integration.properties.type" -> "type"
	// "properties.alerts.items" -> "alerts"

	parts := strings.Split(path, ".")

	// Work backwards to find the actual property name
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]

		// Skip structural keywords
		if part == "properties" || part == "items" {
			continue
		}

		// This should be the actual property name
		return part
	}

	// Fallback to schema name if we can't determine property name
	return strings.TrimSuffix(schemaName, "Entity")
}

// Convert arbitrary string to Go-style enum member name component
func convertToEnumMemberName(value string) string {
	// Handle empty or whitespace-only strings
	if strings.TrimSpace(value) == "" {
		return "Empty"
	}

	// Replace special characters with underscores
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, value)

	// Remove leading/trailing underscores and collapse multiple underscores
	cleaned = strings.Trim(cleaned, "_")
	// Simple approach to collapse multiple underscores
	for strings.Contains(cleaned, "__") {
		cleaned = strings.ReplaceAll(cleaned, "__", "_")
	}

	// Split by underscores and convert to PascalCase
	if cleaned == "" {
		return "Empty"
	}

	parts := strings.Split(cleaned, "_")
	var result strings.Builder

	for _, part := range parts {
		if len(part) > 0 {
			// Capitalize first letter, lowercase the rest
			result.WriteString(strings.ToUpper(string(part[0])))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	memberName := result.String()

	// Ensure it starts with a letter (Go requirement)
	if len(memberName) > 0 && unicode.IsDigit(rune(memberName[0])) {
		memberName = "Value" + memberName
	}

	// Handle empty result after cleaning
	if memberName == "" {
		memberName = "Empty"
	}

	return memberName
}
