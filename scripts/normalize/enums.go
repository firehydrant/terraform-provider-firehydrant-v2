package main

import (
	"fmt"
	"strings"
	"unicode"
)

type EnumNormalizationInfo struct {
	SchemaName   string
	PropertyPath string
	PropertyName string
	EnumValues   []string
	Target       map[string]interface{}
}

func normalizeEnums(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Enum Properties ===\n")

	allEnums := findAllEnumProperties(schemas)

	if len(allEnums) == 0 {
		fmt.Printf("No enum properties found to normalize\n")
		return conflicts
	}

	fmt.Printf("Found %d enum properties to normalize\n", len(allEnums))

	for _, enumInfo := range allEnums {
		conflict := transformEnumProperty(enumInfo)
		if conflict != nil {
			conflicts = append(conflicts, *conflict)
		}
	}

	fmt.Printf("Successfully normalized %d enum properties\n", len(conflicts))
	return conflicts
}

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

func findEnumsInSchemaRecursive(schemaName string, currentObj map[string]interface{}, path string, rootSchema map[string]interface{}) []EnumNormalizationInfo {
	var enums []EnumNormalizationInfo

	if enumValues, hasEnum := currentObj["enum"]; hasEnum {
		if enumArray, ok := enumValues.([]interface{}); ok && len(enumArray) > 0 {
			var stringValues []string
			for _, val := range enumArray {
				if str, ok := val.(string); ok {
					stringValues = append(stringValues, str)
				}
			}

			if len(stringValues) > 0 {
				propertyName := extractPropertyNameFromNormalizationPath(path, schemaName)

				enumInfo := EnumNormalizationInfo{
					SchemaName:   schemaName,
					PropertyPath: path,
					PropertyName: propertyName,
					EnumValues:   stringValues,
					Target:       currentObj,
				}

				enums = append(enums, enumInfo)
			}
		}
	}

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

func transformEnumProperty(enumInfo EnumNormalizationInfo) *ConflictDetail {
	var members []map[string]interface{}

	for _, value := range enumInfo.EnumValues {
		memberName := generateEnumMemberName(value, enumInfo.PropertyName, enumInfo.SchemaName)
		members = append(members, map[string]interface{}{
			"name":  memberName,
			"value": value,
		})
	}

	delete(enumInfo.Target, "enum")
	enumInfo.Target["x-speakeasy-enums"] = members

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

func generateEnumMemberName(value, fieldName, entityName string) string {
	if strings.TrimSpace(value) == "" {
		value = "Empty"
	}

	var prefix string

	switch {
	case entityName == "RequestParam":
		prefix = "Request"
	case entityName == "RequestBody":
		if strings.HasPrefix(fieldName, "Form") {
			prefix = "" // Form prefix is already in fieldName
		} else {
			prefix = "Request"
		}
	case entityName == "Response":
		prefix = "Response"
	case strings.HasSuffix(entityName, "Entity"):
		prefix = convertToEnumMemberName(strings.TrimSuffix(entityName, "Entity"))
	default:
		prefix = convertToEnumMemberName(entityName)
	}

	cleanFieldName := convertToEnumMemberName(fieldName)
	cleanValue := convertToEnumMemberName(value)

	memberName := prefix + cleanFieldName + cleanValue

	// Ensure it starts with a letter
	if len(memberName) > 0 && unicode.IsDigit(rune(memberName[0])) {
		memberName = "Value" + memberName
	}

	if memberName == "" {
		memberName = prefix + cleanFieldName + "Unknown"
	}

	return memberName
}

func extractPropertyNameFromNormalizationPath(path, schemaName string) string {
	if path == "" {
		return strings.TrimSuffix(schemaName, "Entity")
	}

	parts := strings.Split(path, ".")

	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]

		if part == "properties" || part == "items" {
			continue
		}

		return part
	}

	return strings.TrimSuffix(schemaName, "Entity")
}

func convertToEnumMemberName(value string) string {
	if strings.TrimSpace(value) == "" {
		return "Empty"
	}

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, value)

	cleaned = strings.Trim(cleaned, "_")
	for strings.Contains(cleaned, "__") {
		cleaned = strings.ReplaceAll(cleaned, "__", "_")
	}

	if cleaned == "" {
		return "Empty"
	}

	parts := strings.Split(cleaned, "_")
	var result strings.Builder

	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(string(part[0])))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	memberName := result.String()

	if len(memberName) > 0 && unicode.IsDigit(rune(memberName[0])) {
		memberName = "Value" + memberName
	}

	if memberName == "" {
		memberName = "Empty"
	}

	return memberName
}

func normalizePathEnums(paths map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Path/Operation Enum Properties ===\n")

	allPathEnums := findAllPathEnumProperties(paths)

	if len(allPathEnums) == 0 {
		fmt.Printf("No path enum properties found to normalize\n")
		return conflicts
	}

	fmt.Printf("Found %d path enum properties to normalize\n", len(allPathEnums))

	for _, enumInfo := range allPathEnums {
		conflict := transformEnumProperty(enumInfo)
		if conflict != nil {
			conflicts = append(conflicts, *conflict)
		}
	}

	fmt.Printf("Successfully normalized %d path enum properties\n", len(conflicts))
	return conflicts
}

func findAllPathEnumProperties(paths map[string]interface{}) []EnumNormalizationInfo {
	var allEnums []EnumNormalizationInfo

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
					if ok {
						for _, param := range paramsList {
							paramMap, ok := param.(map[string]interface{})
							if ok {
								paramName, _ := paramMap["name"].(string)
								schemaName := fmt.Sprintf("RequestParam_%s_%s_%s", method, pathName, paramName)

								if schema, hasSchema := paramMap["schema"].(map[string]interface{}); hasSchema {
									enumEnums := findEnumsInSchemaRecursive(schemaName, schema, "", schema)
									for j := range enumEnums {
										enumEnums[j].PropertyName = convertToEnumMemberName(paramName)
										enumEnums[j].SchemaName = "RequestParam"
									}
									allEnums = append(allEnums, enumEnums...)
								}
							}
						}
					}
				}

				if requestBody, hasReqBody := opMap["requestBody"]; hasReqBody {
					reqBodyMap, ok := requestBody.(map[string]interface{})
					if ok {
						if content, hasContent := reqBodyMap["content"].(map[string]interface{}); hasContent {
							for contentType, contentSchema := range content {
								if contentMap, ok := contentSchema.(map[string]interface{}); ok {
									if schema, hasSchema := contentMap["schema"].(map[string]interface{}); hasSchema {
										schemaName := fmt.Sprintf("RequestBody_%s_%s", method, sanitizePathForSchemaName(pathName))
										reqEnums := findEnumsInSchemaRecursive(schemaName, schema, "", schema)

										for j := range reqEnums {
											if contentType == "multipart/form-data" {
												reqEnums[j].PropertyName = "Form" + convertToEnumMemberName(reqEnums[j].PropertyName)
											} else {
												reqEnums[j].PropertyName = "Request" + convertToEnumMemberName(reqEnums[j].PropertyName)
											}
											reqEnums[j].SchemaName = "RequestBody"
										}
										allEnums = append(allEnums, reqEnums...)

										fmt.Printf("📋 Found request body enum in %s %s (%s)\n", method, pathName, contentType)
									}
								}
							}
						}
					}
				}

				// Checks for responses with inline enums enums, uncommon but included as a safeguard
				if responses, hasResponses := opMap["responses"]; hasResponses {
					respMap, ok := responses.(map[string]interface{})
					if ok {
						for statusCode, response := range respMap {
							if respBody, ok := response.(map[string]interface{}); ok {
								if content, hasContent := respBody["content"].(map[string]interface{}); hasContent {
									if jsonContent, hasJson := content["application/json"].(map[string]interface{}); hasJson {
										if schema, hasSchema := jsonContent["schema"].(map[string]interface{}); hasSchema {
											schemaName := fmt.Sprintf("Response_%s_%s_%s", method, pathName, statusCode)
											respEnums := findEnumsInSchemaRecursive(schemaName, schema, "", schema)
											allEnums = append(allEnums, respEnums...)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return allEnums
}

func sanitizePathForSchemaName(path string) string {
	cleaned := strings.ReplaceAll(path, "/", "_")
	cleaned = strings.ReplaceAll(cleaned, "{", "")
	cleaned = strings.ReplaceAll(cleaned, "}", "")
	cleaned = strings.Trim(cleaned, "_")
	return cleaned
}
