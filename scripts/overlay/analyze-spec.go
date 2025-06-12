package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type ResourceInfo struct {
	EntityName   string
	SchemaName   string
	ResourceName string
	Operations   map[string]OperationInfo
	CreateSchema string
	UpdateSchema string
	// Store the identified primary ID parameter
	// We need to track this so that we use the correct ID field where there are multiple IDs in path params
	PrimaryID string
}

type OperationInfo struct {
	OperationID   string
	Path          string
	Method        string
	RequestSchema string
}

func analyzeSpec(spec OpenAPISpec, manualMappings *ManualMappings) map[string]*ResourceInfo {
	resources := make(map[string]*ResourceInfo)

	// First pass: identify all entity schemas
	entitySchemas := identifyEntitySchemas(spec.Components.Schemas)
	fmt.Printf("Identified %d entity schemas\n", len(entitySchemas))

	// Second pass: match operations to entities
	for path, pathItem := range spec.Paths {
		analyzePathOperations(path, pathItem, entitySchemas, resources, spec, manualMappings)
	}

	// Third pass: validate resources but keep all for analysis
	fmt.Printf("\n=== Resource Validation ===\n")
	for name, resource := range resources {
		opTypes := make([]string, 0)
		for crudType := range resource.Operations {
			opTypes = append(opTypes, crudType)
		}
		fmt.Printf("Resource: %s with operations: %v\n", name, opTypes)

		if isTerraformViable(resource, spec) {
			fmt.Printf("  ✅ Viable for Terraform\n")
		} else {
			fmt.Printf("  ❌ Not viable for Terraform - will skip annotations\n")
		}
	}

	return resources
}

func identifyEntitySchemas(schemas map[string]Schema) map[string]bool {
	entities := make(map[string]bool)

	for name, schema := range schemas {
		if isEntitySchema(name, schema) {
			entities[name] = true
		}
	}

	return entities
}

func isEntitySchema(name string, schema Schema) bool {
	// Skip request/response wrappers
	lowerName := strings.ToLower(name)
	if strings.HasPrefix(lowerName, "create_") ||
		strings.HasPrefix(lowerName, "update_") ||
		strings.HasPrefix(lowerName, "delete_") ||
		strings.Contains(lowerName, "request") ||
		strings.Contains(lowerName, "response") ||
		strings.HasSuffix(name, "Paginated") {
		return false
	}

	// Skip nullable wrapper schemas
	if strings.HasPrefix(name, "Nullable") {
		return false
	}

	// Must be an object with properties
	if schema.Type != "object" || len(schema.Properties) == 0 {
		return false
	}

	// Entities should have an id property and end with "Entity"
	_, hasID := schema.Properties["id"]
	_, hasSlug := schema.Properties["slug"]
	hasSuffix := strings.HasSuffix(name, "Entity")

	hasIdentifier := hasID || hasSlug
	// Be strict: require both conditions
	return hasIdentifier && hasSuffix
}

func analyzePathOperations(path string, pathItem PathItem, entitySchemas map[string]bool,
	resources map[string]*ResourceInfo, spec OpenAPISpec, manualMappings *ManualMappings) {

	operations := []struct {
		method string
		op     *Operation
	}{
		{"get", pathItem.Get},
		{"post", pathItem.Post},
		{"put", pathItem.Put},
		{"patch", pathItem.Patch},
		{"delete", pathItem.Delete},
	}

	for _, item := range operations {
		if item.op == nil {
			continue
		}

		// Check if this operation should be manually ignored
		if shouldIgnoreOperation(path, item.method, manualMappings) {
			continue
		}

		resourceInfo := extractResourceInfo(path, item.method, item.op, entitySchemas, spec, manualMappings)
		if resourceInfo != nil {
			if existing, exists := resources[resourceInfo.ResourceName]; exists {
				// Merge operations
				for opType, opInfo := range resourceInfo.Operations {
					existing.Operations[opType] = opInfo
				}

				// Preserve create/update schema info - don't overwrite with empty values
				if resourceInfo.CreateSchema != "" {
					existing.CreateSchema = resourceInfo.CreateSchema
				}
				if resourceInfo.UpdateSchema != "" {
					existing.UpdateSchema = resourceInfo.UpdateSchema
				}
			} else {
				resources[resourceInfo.ResourceName] = resourceInfo
			}
		}
	}
}

func extractResourceInfo(path, method string, op *Operation,
	entitySchemas map[string]bool, spec OpenAPISpec, manualMappings *ManualMappings) *ResourceInfo {

	// Determine CRUD type
	crudType := determineCrudType(path, method, op.OperationID)
	if crudType == "" {
		return nil
	}

	// Check for manual entity mapping first
	if manualEntityName, hasManual := getManualEntityMapping(path, method, manualMappings); hasManual {
		// Use manual entity mapping
		entityName := manualEntityName
		resourceName := deriveResourceName(entityName, op.OperationID, path)

		info := &ResourceInfo{
			EntityName:   entityName,
			SchemaName:   entityName,
			ResourceName: resourceName,
			Operations:   make(map[string]OperationInfo),
		}

		opInfo := OperationInfo{
			OperationID: op.OperationID,
			Path:        path,
			Method:      method,
		}

		// Extract request schema for create/update operations
		if crudType == "create" || crudType == "update" {
			if op.RequestBody != nil {
				if content, ok := op.RequestBody["content"].(map[string]interface{}); ok {
					if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
						if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
							if ref, ok := schema["$ref"].(string); ok {
								requestSchemaName := extractSchemaName(ref)
								opInfo.RequestSchema = requestSchemaName

								if crudType == "create" {
									info.CreateSchema = requestSchemaName
								} else if crudType == "update" {
									info.UpdateSchema = requestSchemaName
								}
							}
						}
					}
				}
			}
		}

		info.Operations[crudType] = opInfo
		return info
	}

	// Find associated entity schema using automatic detection
	entityName := findEntityFromOperation(op, entitySchemas, spec)
	if entityName == "" {
		return nil
	}

	resourceName := deriveResourceName(entityName, op.OperationID, path)

	info := &ResourceInfo{
		EntityName:   entityName,
		SchemaName:   entityName,
		ResourceName: resourceName,
		Operations:   make(map[string]OperationInfo),
	}

	opInfo := OperationInfo{
		OperationID: op.OperationID,
		Path:        path,
		Method:      method,
	}

	// Extract request schema for create/update operations
	if crudType == "create" || crudType == "update" {
		if op.RequestBody != nil {
			if content, ok := op.RequestBody["content"].(map[string]interface{}); ok {
				if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
					if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
						if ref, ok := schema["$ref"].(string); ok {
							requestSchemaName := extractSchemaName(ref)
							opInfo.RequestSchema = requestSchemaName

							if crudType == "create" {
								info.CreateSchema = requestSchemaName
							} else if crudType == "update" {
								info.UpdateSchema = requestSchemaName
							}
						}
					}
				}
			}
		}
	}

	info.Operations[crudType] = opInfo
	return info
}

func determineCrudType(path, method, operationID string) string {
	lowerOp := strings.ToLower(operationID)

	// Check operation ID first
	if strings.Contains(lowerOp, "create") {
		return "create"
	}
	if strings.Contains(lowerOp, "update") || strings.Contains(lowerOp, "patch") {
		return "update"
	}
	if strings.Contains(lowerOp, "delete") {
		return "delete"
	}
	if strings.Contains(lowerOp, "list") {
		return "list"
	}
	if strings.Contains(lowerOp, "get") && strings.Contains(path, "{") {
		return "read"
	}

	// Fallback to method-based detection
	switch method {
	case "post":
		if !strings.Contains(path, "{") {
			return "create"
		}
	case "get":
		if strings.Contains(path, "{") {
			return "read"
		} else {
			return "list"
		}
	case "patch", "put":
		return "update"
	case "delete":
		return "delete"
	}

	return ""
}

func findEntityFromOperation(op *Operation, entitySchemas map[string]bool, spec OpenAPISpec) string {
	// Check response schemas first
	if op.Responses != nil {
		for _, response := range op.Responses {
			if respMap, ok := response.(map[string]interface{}); ok {
				if content, ok := respMap["content"].(map[string]interface{}); ok {
					if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
						if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
							entityName := findEntityInSchema(schema, entitySchemas)
							if entityName != "" {
								return entityName
							}
						}
					}
				}
			}
		}
	}

	// Check tags
	if len(op.Tags) > 0 {
		for _, tag := range op.Tags {
			possibleEntity := tag + "Entity"
			if entitySchemas[possibleEntity] {
				return possibleEntity
			}
		}
	}

	return ""
}

func findEntityInSchema(schema map[string]interface{}, entitySchemas map[string]bool) string {
	// Direct reference
	if ref, ok := schema["$ref"].(string); ok {
		schemaName := extractSchemaName(ref)
		if entitySchemas[schemaName] {
			return schemaName
		}
	}

	// Check in data array for paginated responses
	if props, ok := schema["properties"].(map[string]interface{}); ok {
		if data, ok := props["data"].(map[string]interface{}); ok {
			if dataType, ok := data["type"].(string); ok && dataType == "array" {
				if items, ok := data["items"].(map[string]interface{}); ok {
					if ref, ok := items["$ref"].(string); ok {
						schemaName := extractSchemaName(ref)
						if entitySchemas[schemaName] {
							return schemaName
						}
					}
				}
			}
		}
	}

	return ""
}

func extractSchemaName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func deriveResourceName(entityName, operationID, path string) string {
	resource := strings.TrimSuffix(entityName, "Entity")

	resource = toSnakeCase(resource)

	if strings.Contains(resource, "_") {
		parts := strings.Split(resource, "_")
		if len(parts) > 1 && parts[0] == parts[1] {
			// Remove duplicate prefix (e.g., incidents_incident -> incident)
			resource = parts[1]
		}
	}

	return resource
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			if i == len(s)-1 || !unicode.IsUpper(rune(s[i+1])) {
				result = append(result, '_')
			}
		}
		result = append(result, []rune(strings.ToLower(string(r)))...)
	}
	return string(result)
}

func extractPathParameters(path string) []string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(path, -1)

	var params []string
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}

	return params
}
