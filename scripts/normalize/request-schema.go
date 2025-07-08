package main

import (
	"fmt"
	"strings"
)

type RequestSchemaInfo struct {
	Path        string
	Method      string
	OperationId string
	ContentType string
	Schema      map[string]interface{}
	SchemaName  string
	IsInline    bool
}

func findInlineRequestSchemas(paths map[string]interface{}) []RequestSchemaInfo {
	var inlineSchemas []RequestSchemaInfo

	for pathName, pathItem := range paths {
		pathMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		// Check POST, PUT, PATCH methods for request bodies
		methods := []string{"post", "put", "patch"}
		for _, method := range methods {
			if operation, exists := pathMap[method]; exists {
				opMap, ok := operation.(map[string]interface{})
				if !ok {
					continue
				}

				operationId, _ := opMap["operationId"].(string)

				if requestBody, hasReqBody := opMap["requestBody"]; hasReqBody {
					reqBodyMap, ok := requestBody.(map[string]interface{})
					if !ok {
						continue
					}

					if content, hasContent := reqBodyMap["content"].(map[string]interface{}); hasContent {
						for contentType, contentSchema := range content {
							if contentMap, ok := contentSchema.(map[string]interface{}); ok {
								if schema, hasSchema := contentMap["schema"].(map[string]interface{}); hasSchema {
									// Check if it's an inline schema (no $ref)
									if _, hasRef := schema["$ref"]; !hasRef {
										// Generate schema name
										schemaName := generateRequestSchemaName(operationId, contentType)

										inlineSchemas = append(inlineSchemas, RequestSchemaInfo{
											Path:        pathName,
											Method:      method,
											OperationId: operationId,
											ContentType: contentType,
											Schema:      schema,
											SchemaName:  schemaName,
											IsInline:    true,
										})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return inlineSchemas
}

func generateRequestSchemaName(operationId, contentType string) string {
	// Tests in laddetruck required operationIds to be added to routes
	// so we will rely on these here
	schemaName := operationId

	// Only add suffix for non-JSON content types
	if contentType != "application/json" {
		schemaName = schemaName + convertContentTypeToSuffix(contentType)
	}

	return schemaName
}

func convertContentTypeToSuffix(contentType string) string {
	switch contentType {
	case "multipart/form-data":
		return "_form"
	case "application/x-www-form-urlencoded":
		return "_form_encoded"
	case "text/plain":
		return "_text"
	case "application/xml":
		return "_xml"
	default:
		suffix := strings.ReplaceAll(contentType, "/", "_")
		suffix = strings.ReplaceAll(suffix, "-", "_")
		suffix = strings.ReplaceAll(suffix, ".", "_")
		suffix = strings.ReplaceAll(suffix, "+", "_")
		return "_" + suffix
	}
}

func ensureUniqueSchemaName(baseName string, schemas map[string]interface{}) string {
	if _, exists := schemas[baseName]; !exists {
		return baseName
	}

	counter := 1
	for {
		candidate := fmt.Sprintf("%s_%d", baseName, counter)
		if _, exists := schemas[candidate]; !exists {
			return candidate
		}
		counter++
	}
}

// normalizeRequestSchemasWithPaths extracts inline request body schemas and moves them to components/schemas
func normalizeRequestSchemasWithPaths(paths map[string]interface{}, schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("\n=== Normalizing Request Schemas ===\n")

	inlineSchemas := findInlineRequestSchemas(paths)

	if len(inlineSchemas) == 0 {
		fmt.Printf("No inline request schemas found to normalize\n")
		return conflicts
	}

	fmt.Printf("Found %d inline request schemas to normalize\n", len(inlineSchemas))

	for _, schemaInfo := range inlineSchemas {
		finalSchemaName := ensureUniqueSchemaName(schemaInfo.SchemaName, schemas)

		schemaCopy := make(map[string]interface{})
		for k, v := range schemaInfo.Schema {
			schemaCopy[k] = v
		}

		schemas[finalSchemaName] = schemaCopy

		// Replace the original inline schema with $ref
		ref := fmt.Sprintf("#/components/schemas/%s", finalSchemaName)

		// Clear the original schema and replace with $ref
		for key := range schemaInfo.Schema {
			delete(schemaInfo.Schema, key)
		}
		schemaInfo.Schema["$ref"] = ref

		conflicts = append(conflicts, ConflictDetail{
			Schema:       finalSchemaName,
			Property:     fmt.Sprintf("%s.%s", schemaInfo.Method, schemaInfo.Path),
			ConflictType: "request-schema-extraction",
			Resolution:   fmt.Sprintf("Extracted inline request schema to %s", finalSchemaName),
		})
	}

	fmt.Printf("Successfully normalized %d request schemas\n", len(conflicts))
	return conflicts
}
