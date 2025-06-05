package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	specPath := os.Args[1]

	fmt.Printf("=== Terraform Overlay Generator ===\n")
	fmt.Printf("Input: %s\n", specPath)

	specData, err := ioutil.ReadFile(specPath)
	if err != nil {
		fmt.Printf("Error reading spec file: %v\n", err)
		os.Exit(1)
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(specData, &spec); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d paths and %d schemas\n\n", len(spec.Paths), len(spec.Components.Schemas))

	// Analyze the spec to find resources
	resources := analyzeSpec(spec)

	// Generate overlay
	overlay := generateOverlay(resources, spec)

	// Write overlay file
	if err := writeOverlay(overlay); err != nil {
		fmt.Printf("Error writing overlay: %v\n", err)
		os.Exit(1)
	}

	// Print summary
	printOverlaySummary(resources, overlay)
}

func printUsage() {
	fmt.Println("OpenAPI Terraform Overlay Generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-overlay <input.json>")
}

// ============================================================================
// OVERLAY GENERATION LOGIC
// ============================================================================

type OpenAPISpec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]PathItem    `json:"paths"`
	Components Components             `json:"components"`
}

type Components struct {
	Schemas         map[string]Schema      `json:"schemas"`
	SecuritySchemes map[string]interface{} `json:"securitySchemes,omitempty"`
}

type Schema struct {
	Type       string                 `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
	AllOf      []interface{}          `json:"allOf,omitempty"`
	Nullable   bool                   `json:"nullable,omitempty"`
	Items      interface{}            `json:"items,omitempty"`
	Raw        map[string]interface{} `json:"-"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
}

type Operation struct {
	OperationID string                 `json:"operationId,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Parameters  []Parameter            `json:"parameters,omitempty"`
	RequestBody map[string]interface{} `json:"requestBody,omitempty"`
	Responses   map[string]interface{} `json:"responses,omitempty"`
}

type Parameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required,omitempty"`
	Schema   Schema `json:"schema,omitempty"`
}

type ResourceInfo struct {
	EntityName   string
	SchemaName   string
	ResourceName string
	Operations   map[string]OperationInfo
	CreateSchema string // Track the create request schema
	UpdateSchema string // Track the update request schema
}

type OperationInfo struct {
	OperationID   string
	Path          string
	Method        string
	RequestSchema string // Track request schema separately
}

type PropertyMismatch struct {
	PropertyName string
	MismatchType string
	Description  string
}

type OverlayAction struct {
	Target string                 `yaml:"target"`
	Update map[string]interface{} `yaml:"update,omitempty"`
}

type Overlay struct {
	Overlay string `yaml:"overlay"`
	Info    struct {
		Title       string `yaml:"title"`
		Version     string `yaml:"version"`
		Description string `yaml:"description"`
	} `yaml:"info"`
	Actions []OverlayAction `yaml:"actions"`
}

func analyzeSpec(spec OpenAPISpec) map[string]*ResourceInfo {
	resources := make(map[string]*ResourceInfo)

	// First pass: identify all entity schemas
	entitySchemas := identifyEntitySchemas(spec.Components.Schemas)
	fmt.Printf("Identified %d entity schemas\n", len(entitySchemas))

	// Second pass: match operations to entities
	for path, pathItem := range spec.Paths {
		analyzePathOperations(path, pathItem, entitySchemas, resources, spec)
	}

	// Third pass: validate and filter resources
	validResources := make(map[string]*ResourceInfo)
	for name, resource := range resources {
		if isValidTerraformResource(resource) {
			validResources[name] = resource
			opTypes := make([]string, 0)
			for crudType := range resource.Operations {
				entityOp := mapCrudToEntityOperation(crudType, resource.EntityName)
				opTypes = append(opTypes, fmt.Sprintf("%s->%s", crudType, entityOp))
			}
			fmt.Printf("Valid Terraform resource: %s with operations: %v\n",
				name, opTypes)
		}
	}

	return validResources
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
	hasSuffix := strings.HasSuffix(name, "Entity")

	// Be strict: require both conditions
	return hasID && hasSuffix
}

func analyzePathOperations(path string, pathItem PathItem, entitySchemas map[string]bool,
	resources map[string]*ResourceInfo, spec OpenAPISpec) {

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

		resourceInfo := extractResourceInfo(path, item.method, item.op, entitySchemas, spec)
		if resourceInfo != nil {
			if existing, exists := resources[resourceInfo.ResourceName]; exists {
				fmt.Printf("    Merging operations for %s\n", resourceInfo.ResourceName)
				fmt.Printf("      Existing: CreateSchema='%s', UpdateSchema='%s'\n", existing.CreateSchema, existing.UpdateSchema)
				fmt.Printf("      New: CreateSchema='%s', UpdateSchema='%s'\n", resourceInfo.CreateSchema, resourceInfo.UpdateSchema)

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

				fmt.Printf("      After merge: CreateSchema='%s', UpdateSchema='%s'\n", existing.CreateSchema, existing.UpdateSchema)
			} else {
				resources[resourceInfo.ResourceName] = resourceInfo
				fmt.Printf("    New resource: %s with CreateSchema='%s', UpdateSchema='%s'\n",
					resourceInfo.ResourceName, resourceInfo.CreateSchema, resourceInfo.UpdateSchema)
			}
		}
	}
}

func extractResourceInfo(path, method string, op *Operation,
	entitySchemas map[string]bool, spec OpenAPISpec) *ResourceInfo {

	// Determine CRUD type
	crudType := determineCrudType(path, method, op.OperationID)
	if crudType == "" {
		return nil
	}

	// Find associated entity schema
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
		fmt.Printf("    Extracting request schema for %s %s operation\n", entityName, crudType)
		fmt.Printf("      Operation: %s %s %s\n", method, path, op.OperationID)

		if op.RequestBody != nil {
			fmt.Printf("      RequestBody exists\n")
			if content, ok := op.RequestBody["content"].(map[string]interface{}); ok {
				fmt.Printf("      Content exists\n")
				if jsonContent, ok := content["application/json"].(map[string]interface{}); ok {
					fmt.Printf("      JSON content exists\n")
					if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
						fmt.Printf("      Schema exists: %+v\n", schema)
						if ref, ok := schema["$ref"].(string); ok {
							requestSchemaName := extractSchemaName(ref)
							opInfo.RequestSchema = requestSchemaName
							fmt.Printf("      ✅ Found request schema: %s\n", requestSchemaName)

							if crudType == "create" {
								info.CreateSchema = requestSchemaName
							} else if crudType == "update" {
								info.UpdateSchema = requestSchemaName
							}
						} else {
							fmt.Printf("      No $ref found in schema\n")
						}
					} else {
						fmt.Printf("      No schema found in JSON content\n")
					}
				} else {
					fmt.Printf("      No application/json content found\n")
				}
			} else {
				fmt.Printf("      No content found in RequestBody\n")
			}
		} else {
			fmt.Printf("      No RequestBody found\n")
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
	// Remove Entity suffix
	resource := strings.TrimSuffix(entityName, "Entity")

	// Convert to snake_case
	resource = toSnakeCase(resource)

	// Handle special cases
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
		if i > 0 && isUpper(r) {
			if i == len(s)-1 || !isUpper(rune(s[i+1])) {
				result = append(result, '_')
			}
		}
		result = append(result, toLower(r))
	}
	return string(result)
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func isValidTerraformResource(resource *ResourceInfo) bool {
	// Must have at least create and read operations for a full Terraform resource
	_, hasCreate := resource.Operations["create"]
	_, hasRead := resource.Operations["read"]

	// Some resources might be read-only data sources
	if hasRead && !hasCreate {
		fmt.Printf("  Note: %s appears to be a read-only data source\n", resource.EntityName)
	}

	return hasCreate && hasRead
}

func generateOverlay(resources map[string]*ResourceInfo, spec OpenAPISpec) *Overlay {
	overlay := &Overlay{
		Overlay: "1.0.0",
	}

	overlay.Info.Title = "Terraform Provider Overlay"
	overlay.Info.Version = "1.0.0"
	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d Terraform resources", len(resources))

	// Detect property mismatches between request and response schemas
	resourceMismatches := detectPropertyMismatches(resources, spec)

	// Report on request/response patterns
	fmt.Println("\n=== Request/Response Schema Analysis ===")
	for _, resource := range resources {
		if resource.CreateSchema != "" || resource.UpdateSchema != "" {
			fmt.Printf("%s:\n", resource.EntityName)
			if resource.CreateSchema != "" {
				fmt.Printf("  Create: %s (request) -> %s (response)\n", resource.CreateSchema, resource.EntityName)
			}
			if resource.UpdateSchema != "" {
				fmt.Printf("  Update: %s (request) -> %s (response)\n", resource.UpdateSchema, resource.EntityName)
			}
			if mismatches, hasMismatch := resourceMismatches[resource.EntityName]; hasMismatch {
				fmt.Printf("  ⚠️  WARNING: Has property mismatches:\n")
				for _, mismatch := range mismatches {
					fmt.Printf("    - %s: %s\n", mismatch.PropertyName, mismatch.Description)
				}
			} else {
				fmt.Printf("  ✅ No property mismatches detected\n")
			}
		}
	}

	// Generate actions
	for _, resource := range resources {
		// Mark the response entity schema
		entityUpdate := map[string]interface{}{
			"x-speakeasy-entity": resource.EntityName,
		}

		overlay.Actions = append(overlay.Actions, OverlayAction{
			Target: fmt.Sprintf("$.components.schemas.%s", resource.SchemaName),
			Update: entityUpdate,
		})

		// Add speakeasy ignore for mismatched properties in request AND response schemas
		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			// Add speakeasy ignore for create schema properties
			if resource.CreateSchema != "" {
				for _, mismatch := range mismatches {
					// Ignore in request schema
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.CreateSchema, mismatch.PropertyName),
						Update: map[string]interface{}{
							"x-speakeasy-ignore": true,
						},
					})

					// Also ignore in response entity schema
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, mismatch.PropertyName),
						Update: map[string]interface{}{
							"x-speakeasy-ignore": true,
						},
					})

					fmt.Printf("  ✅ Added speakeasy ignore for %s.%s in both request (%s) and response (%s) schemas\n",
						resource.EntityName, mismatch.PropertyName, resource.CreateSchema, resource.EntityName)
				}
			}

			// Add terraform ignore for update schema properties
			if resource.UpdateSchema != "" {
				for _, mismatch := range mismatches {
					// Ignore in request schema
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.UpdateSchema, mismatch.PropertyName),
						Update: map[string]interface{}{
							"x-speakeasy-ignore": true,
						},
					})

					// Also ignore in response entity schema (avoid duplicates)
					// Only add if we didn't already add it for create schema
					if resource.CreateSchema == "" {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, mismatch.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-ignore": true,
							},
						})
					}

					fmt.Printf("  ✅ Added speakeasy ignore for %s.%s in update schema\n", resource.EntityName, mismatch.PropertyName)
				}
			}
		}

		// Add entity operations
		for crudType, opInfo := range resource.Operations {
			entityOp := mapCrudToEntityOperation(crudType, resource.EntityName)

			if crudType == "list" {
				fmt.Printf("  List operation: %s -> %s\n", resource.EntityName, entityOp)
			}

			operationUpdate := map[string]interface{}{
				"x-speakeasy-entity-operation": entityOp,
			}

			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: fmt.Sprintf("$.paths[\"%s\"].%s", opInfo.Path, opInfo.Method),
				Update: operationUpdate,
			})

			// Add parameter matches for operations with path parameters
			if crudType == "read" || crudType == "update" || crudType == "delete" {
				addParameterMatches(overlay, opInfo.Path, opInfo.Method, resource.ResourceName)
			}
		}
	}

	fmt.Println("\n=== Overlay Generation Complete ===")
	fmt.Printf("Generated %d actions for %d resources\n", len(overlay.Actions), len(resources))

	// Count ignore actions
	ignoreCount := 0
	for _, action := range overlay.Actions {
		if _, hasIgnore := action.Update["x-speakeasy-ignore"]; hasIgnore {
			ignoreCount++
		}
	}

	if len(resourceMismatches) > 0 {
		fmt.Printf("✅ %d resources had property mismatches\n", len(resourceMismatches))
		fmt.Printf("✅ %d speakeasy ignore actions added\n", ignoreCount)
	}

	return overlay
}

// Enhanced mismatch detection that covers all types of impedance mismatches
func detectPropertyMismatches(resources map[string]*ResourceInfo, spec OpenAPISpec) map[string][]PropertyMismatch {
	mismatches := make(map[string][]PropertyMismatch)

	// We need to work with the raw schemas as map[string]interface{} for proper detection
	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		fmt.Printf("Error marshaling spec for mismatch detection: %v\n", err)
		return mismatches
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		fmt.Printf("Error unmarshaling spec for mismatch detection: %v\n", err)
		return mismatches
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	fmt.Printf("\n=== Mismatch Detection Debug ===\n")
	fmt.Printf("Total resources to check: %d\n", len(resources))

	for _, resource := range resources {
		var resourceMismatches []PropertyMismatch

		fmt.Printf("Checking resource: %s\n", resource.EntityName)
		fmt.Printf("  CreateSchema: %s\n", resource.CreateSchema)
		fmt.Printf("  UpdateSchema: %s\n", resource.UpdateSchema)

		// Check create operation mismatches
		if resource.CreateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.CreateSchema].(map[string]interface{}); exists {
					fmt.Printf("  Found both entity and create schemas - checking for mismatches\n")
					createMismatches := findPropertyMismatches(entitySchema, requestSchema, "create")
					resourceMismatches = append(resourceMismatches, createMismatches...)
					fmt.Printf("  Create mismatches found: %d\n", len(createMismatches))
				} else {
					fmt.Printf("  Warning: Create schema %s not found\n", resource.CreateSchema)
				}
			} else {
				fmt.Printf("  Warning: Entity schema %s not found\n", resource.EntityName)
			}
		}

		// Check update operation mismatches
		if resource.UpdateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.UpdateSchema].(map[string]interface{}); exists {
					fmt.Printf("  Found both entity and update schemas - checking for mismatches\n")
					updateMismatches := findPropertyMismatches(entitySchema, requestSchema, "update")
					resourceMismatches = append(resourceMismatches, updateMismatches...)
					fmt.Printf("  Update mismatches found: %d\n", len(updateMismatches))
				} else {
					fmt.Printf("  Warning: Update schema %s not found\n", resource.UpdateSchema)
				}
			} else {
				fmt.Printf("  Warning: Entity schema %s not found\n", resource.EntityName)
			}
		}

		if len(resourceMismatches) > 0 {
			mismatches[resource.EntityName] = resourceMismatches
			fmt.Printf("  Total mismatches for %s: %d\n", resource.EntityName, len(resourceMismatches))
			for _, mismatch := range resourceMismatches {
				fmt.Printf("    - %s: %s\n", mismatch.PropertyName, mismatch.Description)
			}
		} else {
			fmt.Printf("  No mismatches found for %s\n", resource.EntityName)
		}
	}

	fmt.Printf("=== End Mismatch Detection Debug ===\n\n")
	return mismatches
}

func findPropertyMismatches(entitySchema, requestSchema map[string]interface{}, operation string) []PropertyMismatch {
	var mismatches []PropertyMismatch

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		fmt.Printf("    Warning: Could not access properties for %s operation\n", operation)
		return mismatches
	}

	fmt.Printf("    Checking %d entity properties vs %d request properties for %s\n",
		len(entityProps), len(requestProps), operation)

	for propName, entityProp := range entityProps {
		if requestProp, exists := requestProps[propName]; exists {
			if hasStructuralMismatch(propName, entityProp, requestProp) {
				mismatch := PropertyMismatch{
					PropertyName: propName,
					MismatchType: "structural-mismatch",
					Description:  describeStructuralDifference(entityProp, requestProp),
				}
				fmt.Printf("    ✅ Found structural mismatch: %s - %s\n", propName, mismatch.Description)
				mismatches = append(mismatches, mismatch)
			}
		}
	}

	return mismatches
}

// Check if two property structures are different
func hasStructuralMismatch(propName string, entityProp, requestProp interface{}) bool {
	// Convert both to normalized structure representations
	entityStructure := getPropertyStructure(entityProp)
	requestStructure := getPropertyStructure(requestProp)

	fmt.Printf("      Property '%s':\n", propName)
	fmt.Printf("        Request structure: %s\n", requestStructure)
	fmt.Printf("        Entity structure: %s\n", entityStructure)

	// If structures are different, we have a mismatch
	different := entityStructure != requestStructure
	if different {
		fmt.Printf("        ✅ Structures differ - will ignore\n")
	} else {
		fmt.Printf("        ✓ Structures match\n")
	}

	return different
}

// Get a normalized string representation of a property's structure
func getPropertyStructure(prop interface{}) string {
	propMap, ok := prop.(map[string]interface{})
	if !ok {
		return "unknown"
	}

	// Check for $ref
	if ref, hasRef := propMap["$ref"].(string); hasRef {
		return fmt.Sprintf("$ref:%s", ref)
	}

	propType, _ := propMap["type"].(string)

	switch propType {
	case "array":
		items, hasItems := propMap["items"]
		if hasItems {
			itemStructure := getPropertyStructure(items)
			return fmt.Sprintf("array[%s]", itemStructure)
		}
		return "array[unknown]"

	case "object":
		properties, hasProps := propMap["properties"]
		additionalProps, hasAdditional := propMap["additionalProperties"]

		if hasProps {
			propsMap, _ := properties.(map[string]interface{})
			if len(propsMap) == 0 {
				return "object{empty}"
			}

			// Get structure of nested properties
			var propStructures []string
			for key, value := range propsMap {
				propStructures = append(propStructures, fmt.Sprintf("%s:%s", key, getPropertyStructure(value)))
			}
			return fmt.Sprintf("object{%v}", propStructures)
		}

		if hasAdditional {
			additionalStructure := getPropertyStructure(additionalProps)
			return fmt.Sprintf("object{additional:%s}", additionalStructure)
		}

		return "object{}"

	case "string", "integer", "number", "boolean":
		return propType

	default:
		if propType == "" {
			// No explicit type - check what we have
			if _, hasProps := propMap["properties"]; hasProps {
				return "implicit-object"
			}
			if _, hasItems := propMap["items"]; hasItems {
				return "implicit-array"
			}
		}
		return fmt.Sprintf("type:%s", propType)
	}
}

// Describe the structural difference for reporting
func describeStructuralDifference(entityProp, requestProp interface{}) string {
	entityStructure := getPropertyStructure(entityProp)
	requestStructure := getPropertyStructure(requestProp)

	return fmt.Sprintf("request structure '%s' != response structure '%s'", requestStructure, entityStructure)
}

// Remove the old detectPropertyMismatch function - replaced with comprehensive structural comparison

func mapCrudToEntityOperation(crudType, entityName string) string {
	switch crudType {
	case "create":
		return entityName + "#create"
	case "read":
		return entityName + "#read"
	case "update":
		return entityName + "#update"
	case "delete":
		return entityName + "#delete"
	case "list":
		// For list operations, pluralize the entity name and use #read
		pluralEntityName := pluralizeEntityName(entityName)
		return pluralEntityName + "#read"
	default:
		return entityName + "#" + crudType
	}
}

// Simplified pluralization logic - keeping essential rules
func pluralizeEntityName(entityName string) string {
	// Remove "Entity" suffix
	baseName := strings.TrimSuffix(entityName, "Entity")

	// Essential special cases
	specialCases := map[string]string{
		"Person":   "People",
		"Child":    "Children",
		"Status":   "Statuses",
		"Process":  "Processes",
		"Policy":   "Policies",
		"Category": "Categories",
		"Entry":    "Entries",
		"Activity": "Activities",
		"Property": "Properties",
		"Entity":   "Entities",
		"Query":    "Queries",
		"Library":  "Libraries",
		"History":  "Histories",
		"Summary":  "Summaries",
		"Country":  "Countries",
		"City":     "Cities",
		"Company":  "Companies",
	}

	if plural, ok := specialCases[baseName]; ok {
		return plural + "Entities"
	}

	// Simple pluralization rules
	if strings.HasSuffix(baseName, "y") && len(baseName) > 1 && !isVowel(baseName[len(baseName)-2]) {
		baseName = baseName[:len(baseName)-1] + "ies"
	} else if strings.HasSuffix(baseName, "s") ||
		strings.HasSuffix(baseName, "ss") ||
		strings.HasSuffix(baseName, "sh") ||
		strings.HasSuffix(baseName, "ch") ||
		strings.HasSuffix(baseName, "x") ||
		strings.HasSuffix(baseName, "z") {
		baseName = baseName + "es"
	} else {
		baseName = baseName + "s"
	}

	return baseName + "Entities"
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func addParameterMatches(overlay *Overlay, path, method, resourceName string) {
	// Find all path parameters
	re := regexp.MustCompile(`\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(path, -1)

	for _, match := range matches {
		paramName := match[1]

		if paramName != "id" && (strings.Contains(paramName, "id") || paramName == resourceName) {
			fmt.Printf("    Adding x-speakeasy-match for parameter: %s (not exact 'id')\n", paramName)
			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
					path, method, paramName),
				Update: map[string]interface{}{
					"x-speakeasy-match": "id",
				},
			})
		} else if paramName == "id" {
			fmt.Printf("    Skipping x-speakeasy-match for parameter: %s (already exact 'id')\n", paramName)
		} else {
			fmt.Printf("    Skipping x-speakeasy-match for parameter: %s (not an ID parameter)\n", paramName)
		}
	}
}

func writeOverlay(overlay *Overlay) error {
	// Marshal to YAML
	data, err := yaml.Marshal(overlay)
	if err != nil {
		return fmt.Errorf("marshaling overlay: %w", err)
	}

	// Write file to current directory
	overlayPath := "terraform-overlay.yaml"
	if err := ioutil.WriteFile(overlayPath, data, 0644); err != nil {
		return fmt.Errorf("writing overlay file: %w", err)
	}

	fmt.Printf("Overlay written to: %s\n", overlayPath)
	return nil
}

func printOverlaySummary(resources map[string]*ResourceInfo, overlay *Overlay) {
	fmt.Println("\n=== Summary ===")
	fmt.Printf("✅ Successfully generated overlay with %d actions for %d resources\n",
		len(overlay.Actions), len(resources))

	fmt.Println("\nOverlay approach:")
	fmt.Println("1. Mark entity schemas with x-speakeasy-entity")
	fmt.Println("2. Tag operations with x-speakeasy-entity-operation")
	fmt.Println("3. Mark ID parameters with x-speakeasy-match")
	fmt.Println("4. Apply x-speakeasy-ignore: true to mismatched properties")

	fmt.Println("\nResources configured:")
	for name, resource := range resources {
		ops := make([]string, 0, len(resource.Operations))
		for op := range resource.Operations {
			ops = append(ops, op)
		}
		fmt.Printf("  - %s: [%s]\n", name, strings.Join(ops, ", "))
	}
}

// UnmarshalJSON custom unmarshaler for Schema to handle both structured and raw data
func (s *Schema) UnmarshalJSON(data []byte) error {
	type Alias Schema
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Also unmarshal into raw map
	if err := json.Unmarshal(data, &s.Raw); err != nil {
		return err
	}

	return nil
}
