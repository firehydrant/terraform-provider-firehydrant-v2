package main

import "fmt"

// Recursively find all additionalProperties in a schema
func findAllAdditionalProperties(schemaName string, obj interface{}, path string) []string {
	var found []string

	switch v := obj.(type) {
	case map[string]interface{}:
		if _, hasAdditional := v["additionalProperties"]; hasAdditional {
			fullPath := schemaName
			if path != "" {
				fullPath += "." + path
			}
			found = append(found, fullPath)
		}

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
		if _, hasAdditional := v["additionalProperties"]; hasAdditional {
			objType, _ := v["type"].(string)
			_, hasProperties := v["properties"]

			if objType == "object" || hasProperties || (!hasProperties && hasAdditional) {
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
