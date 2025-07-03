package main

import (
	"fmt"
	"strings"
)

// Terraform reserved keywords that should be replaced with empty objects
// Comprehensive list of Terraform reserved root attribute names
// Most recent information found at time of writing is from: https://medium.com/@d3vpasha/reserved-keywords-in-terraform-f37a4cbf3a81
// We removed data and resource from this list as we can successfully compile without ignoring
var terraformReservedKeywords = map[string]bool{
	"connection":    true,
	"count":         true,
	"depends_on":    true,
	"for_each":      true,
	"import":        true,
	"lifecycle":     true,
	"locals":        true,
	"module":        true,
	"output":        true,
	"postcondition": true,
	"precondition":  true,
	"provider":      true,
	"provisioner":   true,
	"removed":       true,
	"terraform":     true,
	"variable":      true,
	"moved":         true,
	"check":         true,
}

func normalizeTerraformKeywords(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("Normalizing Terraform reserved keywords in %d schemas\n", len(schemas))

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		schemaConflicts := replaceTerraformKeywordsInSchema(schemaName, schemaMap, "")
		conflicts = append(conflicts, schemaConflicts...)
	}

	return conflicts
}

func replaceTerraformKeywordsInSchema(schemaName string, obj interface{}, path string) []ConflictDetail {
	var conflicts []ConflictDetail

	switch v := obj.(type) {
	case map[string]interface{}:
		if properties, hasProperties := v["properties"].(map[string]interface{}); hasProperties {
			for propName, propValue := range properties {
				if isReservedKeyword(propName) {
					if shouldReplaceProperty(propValue) {
						properties[propName] = map[string]interface{}{
							"type":       "object",
							"properties": map[string]interface{}{},
						}

						fullPath := schemaName
						if path != "" {
							fullPath += "." + path
						}
						fullPath += ".properties." + propName

						conflicts = append(conflicts, ConflictDetail{
							Schema:       schemaName,
							Property:     fullPath,
							ConflictType: "terraform-keyword",
							Resolution:   fmt.Sprintf("Replaced reserved keyword '%s' with empty object", propName),
						})

						fmt.Printf("      ✅ Replaced Terraform keyword '%s' with empty object at %s\n", propName, fullPath)
					} else {
						fmt.Printf("      ⚠️  Skipped Terraform keyword '%s' at %s (not an object or ref)\n", propName, schemaName+".properties."+propName)
					}
				} else {
					// Recursively check nested properties
					newPath := path
					if newPath != "" {
						newPath += ".properties." + propName
					} else {
						newPath = "properties." + propName
					}
					nested := replaceTerraformKeywordsInSchema(schemaName, propValue, newPath)
					conflicts = append(conflicts, nested...)
				}
			}
		}

		// Also check other nested objects (like items, additionalProperties, etc.)
		for key, value := range v {
			if key != "properties" { // We already handled properties above
				newPath := path
				if newPath != "" {
					newPath += "." + key
				} else {
					newPath = key
				}
				nested := replaceTerraformKeywordsInSchema(schemaName, value, newPath)
				conflicts = append(conflicts, nested...)
			}
		}

	case []interface{}:
		for i, item := range v {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			nested := replaceTerraformKeywordsInSchema(schemaName, item, newPath)
			conflicts = append(conflicts, nested...)
		}
	}

	return conflicts
}

func isReservedKeyword(propertyName string) bool {
	if terraformReservedKeywords[propertyName] {
		return true
	}

	if terraformReservedKeywords[strings.ToLower(propertyName)] {
		return true
	}

	return false
}

// Only replace if the property is an object type or references another schema
func shouldReplaceProperty(propValue interface{}) bool {
	propMap, ok := propValue.(map[string]interface{})
	if !ok {
		return false
	}

	// Check if it's a direct object type
	if propType, hasType := propMap["type"].(string); hasType {
		if propType == "object" {
			return true
		}
	}

	// Check if it has a $ref (reference to another schema)
	if _, hasRef := propMap["$ref"]; hasRef {
		return true
	}

	// Check if it has properties (making it an object even without explicit type)
	if _, hasProperties := propMap["properties"]; hasProperties {
		return true
	}

	return false
}
