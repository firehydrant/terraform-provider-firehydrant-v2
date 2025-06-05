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

	resources := analyzeSpec(spec)

	overlay := generateOverlay(resources, spec)

	if err := writeOverlay(overlay); err != nil {
		fmt.Printf("Error writing overlay: %v\n", err)
		os.Exit(1)
	}

	printOverlaySummary(resources, overlay)
}

func printUsage() {
	fmt.Println("OpenAPI Terraform Overlay Generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-overlay <input.json>")
}

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
	CreateSchema string
	UpdateSchema string
}

type OperationInfo struct {
	OperationID   string
	Path          string
	Method        string
	RequestSchema string
}

type PropertyMismatch struct {
	PropertyName string
	MismatchType string
	Description  string
}

type CRUDInconsistency struct {
	PropertyName      string
	InconsistencyType string
	Description       string
	SchemasToIgnore   []string
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

// Check if a resource is viable for Terraform
func isTerraformViable(resource *ResourceInfo, spec OpenAPISpec) bool {
	// Must have at least create and read operations
	_, hasCreate := resource.Operations["create"]
	_, hasRead := resource.Operations["read"]

	if !hasCreate || !hasRead {
		return false
	}

	// Check for problematic CRUD patterns that can't be handled by property ignoring
	if resource.CreateSchema != "" && resource.UpdateSchema != "" {
		// Re-parse the spec to get raw schema data for analysis
		specData, err := json.Marshal(spec)
		if err != nil {
			return true // If we can't analyze, assume it's viable
		}

		var rawSpec map[string]interface{}
		if err := json.Unmarshal(specData, &rawSpec); err != nil {
			return true // If we can't analyze, assume it's viable
		}

		components, _ := rawSpec["components"].(map[string]interface{})
		schemas, _ := components["schemas"].(map[string]interface{})

		createProps := getSchemaProperties(schemas, resource.CreateSchema)
		updateProps := getSchemaProperties(schemas, resource.UpdateSchema)

		// Count manageable properties (non-system fields)
		createManageableProps := 0
		updateManageableProps := 0
		commonManageableProps := 0

		for prop := range createProps {
			if !isSystemProperty(prop) {
				createManageableProps++
			}
		}

		for prop := range updateProps {
			if !isSystemProperty(prop) {
				updateManageableProps++
				// Check if this property also exists in create
				if createProps[prop] != nil && !isSystemProperty(prop) {
					commonManageableProps++
				}
			}
		}

		// Reject resources with fundamentally incompatible CRUD patterns
		if createManageableProps <= 1 && updateManageableProps >= 3 && commonManageableProps == 0 {
			fmt.Printf("    Incompatible CRUD pattern: Create=%d manageable, Update=%d manageable, Common=%d\n",
				createManageableProps, updateManageableProps, commonManageableProps)
			return false
		}
	}

	return true
}

func getSchemaProperties(schemas map[string]interface{}, schemaName string) map[string]interface{} {
	if schemaName == "" {
		return map[string]interface{}{}
	}

	schema, exists := schemas[schemaName]
	if !exists {
		return map[string]interface{}{}
	}

	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	properties, ok := schemaMap["properties"].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}

	return properties
}

func isSystemProperty(propName string) bool {
	systemProps := []string{
		"id", "created_at", "updated_at", "created_by", "updated_by",
		"version", "etag", "revision", "last_modified",
	}

	lowerProp := strings.ToLower(propName)

	for _, sysProp := range systemProps {
		if lowerProp == sysProp || strings.HasSuffix(lowerProp, "_"+sysProp) {
			return true
		}
	}

	// Also consider ID fields as system properties
	if strings.HasSuffix(lowerProp, "_id") {
		return true
	}

	return false
}

func generateOverlay(resources map[string]*ResourceInfo, spec OpenAPISpec) *Overlay {
	overlay := &Overlay{
		Overlay: "1.0.0",
	}

	overlay.Info.Title = "Terraform Provider Overlay"
	overlay.Info.Version = "1.0.0"
	overlay.Info.Description = "Auto-generated overlay for Terraform resources"

	// Separate viable and non-viable resources
	viableResources := make(map[string]*ResourceInfo)
	skippedResources := make([]string, 0)

	for name, resource := range resources {
		if isTerraformViable(resource, spec) {
			viableResources[name] = resource
		} else {
			skippedResources = append(skippedResources, name)
		}
	}

	fmt.Printf("\n=== Overlay Generation Analysis ===\n")
	fmt.Printf("Total resources found: %d\n", len(resources))
	fmt.Printf("Viable for Terraform: %d\n", len(viableResources))
	fmt.Printf("Skipped (non-viable): %d\n", len(skippedResources))

	if len(skippedResources) > 0 {
		fmt.Printf("\nSkipped resources:\n")
		for _, skipped := range skippedResources {
			fmt.Printf("  - %s\n", skipped)
		}
	}

	// Update description with actual count
	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(viableResources))

	// Detect property mismatches for viable resources only
	resourceMismatches := detectPropertyMismatches(viableResources, spec)

	// Detect CRUD inconsistencies for viable resources only
	resourceCRUDInconsistencies := detectCRUDInconsistencies(viableResources, spec)

	// Track which properties already have ignore actions to avoid duplicates
	ignoreTracker := make(map[string]map[string]bool) // map[schemaName][propertyName]bool

	// Generate actions only for viable resources
	for _, resource := range viableResources {
		// Mark the response entity schema
		entityUpdate := map[string]interface{}{
			"x-speakeasy-entity": resource.EntityName,
		}

		overlay.Actions = append(overlay.Actions, OverlayAction{
			Target: fmt.Sprintf("$.components.schemas.%s", resource.SchemaName),
			Update: entityUpdate,
		})

		// Initialize ignore tracker for this resource's schemas
		if ignoreTracker[resource.EntityName] == nil {
			ignoreTracker[resource.EntityName] = make(map[string]bool)
		}
		if resource.CreateSchema != "" && ignoreTracker[resource.CreateSchema] == nil {
			ignoreTracker[resource.CreateSchema] = make(map[string]bool)
		}
		if resource.UpdateSchema != "" && ignoreTracker[resource.UpdateSchema] == nil {
			ignoreTracker[resource.UpdateSchema] = make(map[string]bool)
		}

		// Add speakeasy ignore for property mismatches
		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			addIgnoreActionsForMismatches(overlay, resource, mismatches, ignoreTracker)
		}

		// Add speakeasy ignore for CRUD inconsistencies
		if inconsistencies, exists := resourceCRUDInconsistencies[resource.EntityName]; exists {
			addIgnoreActionsForInconsistencies(overlay, resource, inconsistencies, ignoreTracker)
		}

		// Add entity operations
		for crudType, opInfo := range resource.Operations {
			entityOp := mapCrudToEntityOperation(crudType, resource.EntityName)

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

	fmt.Printf("\n=== Overlay Generation Complete ===\n")
	fmt.Printf("Generated %d actions for %d viable resources\n", len(overlay.Actions), len(viableResources))

	// Count ignore actions
	totalIgnores := 0
	for _, action := range overlay.Actions {
		if _, hasIgnore := action.Update["x-speakeasy-ignore"]; hasIgnore {
			totalIgnores++
		}
	}

	if totalIgnores > 0 {
		fmt.Printf("✅ %d speakeasy ignore actions added for property issues\n", totalIgnores)
	}

	return overlay
}

func addIgnoreActionsForMismatches(overlay *Overlay, resource *ResourceInfo, mismatches []PropertyMismatch, ignoreTracker map[string]map[string]bool) {
	// Add speakeasy ignore for create schema properties
	if resource.CreateSchema != "" {
		for _, mismatch := range mismatches {
			// Ignore in request schema
			if !ignoreTracker[resource.CreateSchema][mismatch.PropertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.CreateSchema, mismatch.PropertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[resource.CreateSchema][mismatch.PropertyName] = true
			}

			// Also ignore in response entity schema
			if !ignoreTracker[resource.EntityName][mismatch.PropertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, mismatch.PropertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[resource.EntityName][mismatch.PropertyName] = true
			}
		}
	}

	// Add speakeasy ignore for update schema properties
	if resource.UpdateSchema != "" {
		for _, mismatch := range mismatches {
			// Ignore in request schema
			if !ignoreTracker[resource.UpdateSchema][mismatch.PropertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.UpdateSchema, mismatch.PropertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[resource.UpdateSchema][mismatch.PropertyName] = true
			}

			// Also ignore in response entity schema (avoid duplicates)
			if !ignoreTracker[resource.EntityName][mismatch.PropertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, mismatch.PropertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[resource.EntityName][mismatch.PropertyName] = true
			}
		}
	}
}

func addIgnoreActionsForInconsistencies(overlay *Overlay, resource *ResourceInfo, inconsistencies []CRUDInconsistency, ignoreTracker map[string]map[string]bool) {
	for _, inconsistency := range inconsistencies {
		// Add ignore actions for each schema listed in SchemasToIgnore
		for _, schemaName := range inconsistency.SchemasToIgnore {
			if !ignoreTracker[schemaName][inconsistency.PropertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", schemaName, inconsistency.PropertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[schemaName][inconsistency.PropertyName] = true
			}
		}
	}
}

// Enhanced mismatch detection - same as before but only for viable resources
func detectPropertyMismatches(resources map[string]*ResourceInfo, spec OpenAPISpec) map[string][]PropertyMismatch {
	mismatches := make(map[string][]PropertyMismatch)

	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return mismatches
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return mismatches
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	for _, resource := range resources {
		var resourceMismatches []PropertyMismatch

		// Check create operation mismatches
		if resource.CreateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.CreateSchema].(map[string]interface{}); exists {
					createMismatches := findPropertyMismatches(entitySchema, requestSchema, "create")
					resourceMismatches = append(resourceMismatches, createMismatches...)
				}
			}
		}

		// Check update operation mismatches
		if resource.UpdateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.UpdateSchema].(map[string]interface{}); exists {
					updateMismatches := findPropertyMismatches(entitySchema, requestSchema, "update")
					resourceMismatches = append(resourceMismatches, updateMismatches...)
				}
			}
		}

		if len(resourceMismatches) > 0 {
			mismatches[resource.EntityName] = resourceMismatches
		}
	}

	return mismatches
}

func findPropertyMismatches(entitySchema, requestSchema map[string]interface{}, operation string) []PropertyMismatch {
	var mismatches []PropertyMismatch

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		return mismatches
	}

	for propName, entityProp := range entityProps {
		if requestProp, exists := requestProps[propName]; exists {
			if hasStructuralMismatch(propName, entityProp, requestProp) {
				mismatch := PropertyMismatch{
					PropertyName: propName,
					MismatchType: "structural-mismatch",
					Description:  describeStructuralDifference(entityProp, requestProp),
				}
				mismatches = append(mismatches, mismatch)
			}
		}
	}

	return mismatches
}

// Check if two property structures are different
func hasStructuralMismatch(propName string, entityProp, requestProp interface{}) bool {
	entityStructure := getPropertyStructure(entityProp)
	requestStructure := getPropertyStructure(requestProp)
	return entityStructure != requestStructure
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
		_, hasAdditional := propMap["additionalProperties"]

		if hasProps {
			propsMap, _ := properties.(map[string]interface{})
			if len(propsMap) == 0 {
				return "object{empty}"
			}
			return "object{defined}"
		}

		if hasAdditional {
			return "object{additional}"
		}

		return "object{}"

	case "string", "integer", "number", "boolean":
		return propType

	default:
		if propType == "" {
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

// Detect CRUD inconsistencies - same as before but only for viable resources
func detectCRUDInconsistencies(resources map[string]*ResourceInfo, spec OpenAPISpec) map[string][]CRUDInconsistency {
	inconsistencies := make(map[string][]CRUDInconsistency)

	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return inconsistencies
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return inconsistencies
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	for _, resource := range resources {
		// Get properties from each schema
		entityProps := getSchemaProperties(schemas, resource.EntityName)
		createProps := map[string]interface{}{}
		updateProps := map[string]interface{}{}

		if resource.CreateSchema != "" {
			createProps = getSchemaProperties(schemas, resource.CreateSchema)
		}
		if resource.UpdateSchema != "" {
			updateProps = getSchemaProperties(schemas, resource.UpdateSchema)
		}

		// Collect all property names across CRUD operations
		allProps := make(map[string]bool)
		for prop := range entityProps {
			allProps[prop] = true
		}
		for prop := range createProps {
			allProps[prop] = true
		}
		for prop := range updateProps {
			allProps[prop] = true
		}

		var resourceInconsistencies []CRUDInconsistency

		// Check each property for consistency across CRUD operations
		for propName := range allProps {
			// Skip ID properties - they have separate handling logic
			if propName == "id" {
				continue
			}

			entityHas := entityProps[propName] != nil
			createHas := createProps[propName] != nil
			updateHas := updateProps[propName] != nil

			// Check for CRUD inconsistencies
			var schemasToIgnore []string
			var inconsistencyType string
			var description string
			hasInconsistency := false

			if resource.CreateSchema != "" && resource.UpdateSchema != "" {
				// Full CRUD resource - all three must be consistent
				if !(entityHas && createHas && updateHas) {
					hasInconsistency = true
					inconsistencyType = "crud-property-mismatch"
					description = fmt.Sprintf("Property not present in all CRUD operations (Entity:%v, Create:%v, Update:%v)", entityHas, createHas, updateHas)

					// Ignore in schemas where property exists but shouldn't for consistency
					if entityHas && (!createHas || !updateHas) {
						schemasToIgnore = append(schemasToIgnore, resource.EntityName)
					}
					if createHas && (!entityHas || !updateHas) {
						schemasToIgnore = append(schemasToIgnore, resource.CreateSchema)
					}
					if updateHas && (!entityHas || !createHas) {
						schemasToIgnore = append(schemasToIgnore, resource.UpdateSchema)
					}
				}
			} else if resource.CreateSchema != "" {
				// Create + Read resource - both must be consistent
				if !(entityHas && createHas) {
					hasInconsistency = true
					inconsistencyType = "create-read-mismatch"
					description = fmt.Sprintf("Property not present in both CREATE and READ (Entity:%v, Create:%v)", entityHas, createHas)

					if entityHas && !createHas {
						schemasToIgnore = append(schemasToIgnore, resource.EntityName)
					}
					if createHas && !entityHas {
						schemasToIgnore = append(schemasToIgnore, resource.CreateSchema)
					}
				}
			}

			if hasInconsistency {
				inconsistency := CRUDInconsistency{
					PropertyName:      propName,
					InconsistencyType: inconsistencyType,
					Description:       description,
					SchemasToIgnore:   schemasToIgnore,
				}
				resourceInconsistencies = append(resourceInconsistencies, inconsistency)
			}
		}

		if len(resourceInconsistencies) > 0 {
			inconsistencies[resource.EntityName] = resourceInconsistencies
		}
	}

	return inconsistencies
}

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

// Simplified pluralization logic
func pluralizeEntityName(entityName string) string {
	// Remove "Entity" suffix
	baseName := strings.TrimSuffix(entityName, "Entity")

	// Simple pluralization
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
			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
					path, method, paramName),
				Update: map[string]interface{}{
					"x-speakeasy-match": "id",
				},
			})
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
	viableCount := 0
	for _, resource := range resources {
		if isTerraformViable(resource, OpenAPISpec{}) {
			viableCount++
		}
	}

	fmt.Println("\n=== Summary ===")
	fmt.Printf("✅ Successfully generated overlay with %d actions\n", len(overlay.Actions))
	fmt.Printf("📊 Resources: %d total, %d viable for Terraform, %d skipped\n",
		len(resources), viableCount, len(resources)-viableCount)

	fmt.Println("\nOverlay approach:")
	fmt.Println("1. Skip annotations for non-viable resources")
	fmt.Println("2. Mark viable entity schemas with x-speakeasy-entity")
	fmt.Println("3. Tag operations with x-speakeasy-entity-operation")
	fmt.Println("4. Mark ID parameters with x-speakeasy-match")
	fmt.Println("5. Apply x-speakeasy-ignore: true to problematic properties")
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
