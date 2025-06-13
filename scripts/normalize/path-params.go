package main

import (
	"fmt"
	"strings"
)

// Normalize path parameters to match entity ID types
func normalizePathParameters(paths map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Path Parameters ===\n")

	for pathName, pathItem := range paths {
		pathMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		methods := []string{"get", "post", "put", "patch", "delete"}
		for _, method := range methods {
			if operation, exists := pathMap[method]; exists {
				opMap, ok := operation.(map[string]interface{})
				if !ok {
					continue
				}

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

						// normailze int and string parameters
						paramIn, _ := paramMap["in"].(string)
						paramName, _ := paramMap["name"].(string)

						if paramIn == "path" && (strings.Contains(paramName, "id") || strings.HasSuffix(paramName, "_id")) {
							schema, hasSchema := paramMap["schema"]
							if hasSchema {
								schemaMap, ok := schema.(map[string]interface{})
								if ok {
									paramType, _ := schemaMap["type"].(string)
									paramFormat, _ := schemaMap["format"].(string)

									if paramType == "integer" {
										fmt.Printf("  Found integer ID parameter: %s %s.%s (type: %s, format: %s)\n",
											method, pathName, paramName, paramType, paramFormat)

										schemaMap["type"] = "string"
										delete(schemaMap, "format")

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
