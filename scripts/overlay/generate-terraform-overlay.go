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

// Manual mapping configuration
type ManualMapping struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Action string `yaml:"action"` // "ignore", "entity", "match"
	Value  string `yaml:"value,omitempty"`
}

type ManualMappings struct {
	Operations []ManualMapping `yaml:"operations"`
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
	PrimaryID    string // Store the identified primary ID parameter
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

// IDPattern represents the ID parameter pattern for an operation
type IDPattern struct {
	Path       string
	Method     string
	Operation  string
	Parameters []string
}

// PathParameterInconsistency represents inconsistencies in path parameters
type PathParameterInconsistency struct {
	InconsistencyType string
	Description       string
	Operations        []string
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	specPath := os.Args[1]
	var mappingsPath string
	if len(os.Args) > 2 {
		mappingsPath = os.Args[2]
	} else {
		mappingsPath = "manual-mappings.yaml"
	}

	fmt.Printf("=== Terraform Overlay Generator ===\n")
	fmt.Printf("Input: %s\n", specPath)

	// Load manual mappings
	manualMappings := loadManualMappings(mappingsPath)

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

	resources := analyzeSpec(spec, manualMappings)

	overlay := generateOverlay(resources, spec, manualMappings)

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

// Load manual mappings from file
func loadManualMappings(mappingsPath string) *ManualMappings {
	data, err := ioutil.ReadFile(mappingsPath)
	if err != nil {
		// File doesn't exist - return empty mappings
		fmt.Printf("No manual mappings file found at %s (this is optional)\n", mappingsPath)
		return &ManualMappings{}
	}

	var mappings ManualMappings
	if err := yaml.Unmarshal(data, &mappings); err != nil {
		fmt.Printf("Error parsing manual mappings file: %v\n", err)
		return &ManualMappings{}
	}

	fmt.Printf("Loaded %d manual mappings from %s\n", len(mappings.Operations), mappingsPath)
	return &mappings
}

// Check if an operation should be manually ignored
func shouldIgnoreOperation(path, method string, manualMappings *ManualMappings) bool {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.ToLower(mapping.Method) == strings.ToLower(method) && mapping.Action == "ignore" {
			fmt.Printf("    Manual mapping: Ignoring operation %s %s\n", method, path)
			return true
		}
	}
	return false
}

// Check if an operation has a manual entity mapping
func getManualEntityMapping(path, method string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.ToLower(mapping.Method) == strings.ToLower(method) && mapping.Action == "entity" {
			fmt.Printf("    Manual mapping: Operation %s %s -> Entity %s\n", method, path, mapping.Value)
			return mapping.Value, true
		}
	}
	return "", false
}

// Check if a parameter has a manual match mapping
func getManualParameterMatch(path, method, paramName string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.ToLower(mapping.Method) == strings.ToLower(method) && mapping.Action == "match" {
			// For match mappings, we expect the value to be in format "param_name:field_name"
			parts := strings.SplitN(mapping.Value, ":", 2)
			if len(parts) == 2 && parts[0] == paramName {
				fmt.Printf("    Manual mapping: Parameter %s in %s %s -> %s\n", paramName, method, path, parts[1])
				return parts[1], true
			}
		}
	}
	return "", false
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
	hasSuffix := strings.HasSuffix(name, "Entity")

	// Be strict: require both conditions
	return hasID && hasSuffix
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

// Extract path parameters in order from a path string
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

// Get entity properties for field existence checking
func getEntityProperties(entityName string, spec OpenAPISpec) map[string]interface{} {
	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return map[string]interface{}{}
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return map[string]interface{}{}
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	return getSchemaProperties(schemas, entityName)
}

// Check if a field exists in the entity properties
func checkFieldExistsInEntity(paramName string, entityProps map[string]interface{}) bool {
	// Direct field name match
	if _, exists := entityProps[paramName]; exists {
		return true
	}

	// Check for common variations
	variations := []string{
		paramName,
		strings.TrimSuffix(paramName, "_id"), // Remove _id suffix
		strings.TrimSuffix(paramName, "Id"),  // Remove Id suffix
	}

	for _, variation := range variations {
		if _, exists := entityProps[variation]; exists {
			return true
		}
	}

	return false
}

// Check if a resource is viable for Terraform
func isTerraformViable(resource *ResourceInfo, spec OpenAPISpec) bool {
	// Must have at least create and read operations
	_, hasCreate := resource.Operations["create"]
	_, hasRead := resource.Operations["read"]

	if !hasCreate || !hasRead {
		fmt.Printf("    Missing required operations: Create=%v, Read=%v\n", hasCreate, hasRead)
		return false
	}

	// Must have a create schema to be manageable by Terraform
	if resource.CreateSchema == "" {
		fmt.Printf("    No create schema found\n")
		return false
	}

	// Identify the primary ID for this entity
	primaryID, validPrimaryID := identifyEntityPrimaryID(resource)
	if !validPrimaryID {
		fmt.Printf("    Cannot identify valid primary ID parameter\n")
		return false
	}

	// Validate all operations against the primary ID
	validOperations := validateOperationParameters(resource, primaryID, spec)

	// Must still have CREATE and READ after validation
	_, hasValidCreate := validOperations["create"]
	_, hasValidRead := validOperations["read"]

	if !hasValidCreate || !hasValidRead {
		fmt.Printf("    Lost required operations after parameter validation: Create=%v, Read=%v\n", hasValidCreate, hasValidRead)
		return false
	}

	// Update resource with only valid operations and primary ID
	resource.Operations = validOperations
	resource.PrimaryID = primaryID

	// Check for overlapping properties between create and entity schemas
	if !hasValidCreateReadConsistency(resource, spec) {
		fmt.Printf("    Create and Read operations have incompatible schemas\n")
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

// Identify the primary ID parameter that belongs to this specific entity
func identifyEntityPrimaryID(resource *ResourceInfo) (string, bool) {
	// Get all unique path parameters across operations
	allParams := make(map[string]bool)

	for crudType, opInfo := range resource.Operations {
		if crudType == "create" || crudType == "list" {
			continue // Skip operations that typically don't have entity-specific IDs
		}

		pathParams := extractPathParameters(opInfo.Path)
		for _, param := range pathParams {
			allParams[param] = true
		}
	}

	if len(allParams) == 0 {
		return "", false // No path parameters found
	}

	// Find the parameter that matches this entity
	var entityPrimaryID string
	matchCount := 0

	for param := range allParams {
		if mapsToEntityID(param, resource.EntityName) {
			entityPrimaryID = param
			matchCount++
		}
	}

	if matchCount == 0 {
		// No parameter maps to this entity - check for generic 'id' parameter
		if allParams["id"] {
			fmt.Printf("    Using generic 'id' parameter for entity %s\n", resource.EntityName)
			return "id", true
		}
		fmt.Printf("    No parameter maps to entity %s\n", resource.EntityName)
		return "", false
	}

	if matchCount > 1 {
		// Multiple parameters claim to map to this entity - ambiguous
		fmt.Printf("    Multiple parameters map to entity %s: ambiguous primary ID\n", resource.EntityName)
		return "", false
	}

	fmt.Printf("    Identified primary ID '%s' for entity %s\n", entityPrimaryID, resource.EntityName)
	return entityPrimaryID, true
}

// Check if a parameter name maps to a specific entity's ID field
func mapsToEntityID(paramName, entityName string) bool {
	// Extract base name from entity (e.g., "ChangeEvent" from "ChangeEventEntity")
	entityBase := strings.TrimSuffix(entityName, "Entity")

	// Convert to snake_case and add _id suffix
	expectedParam := toSnakeCase(entityBase) + "_id"

	return strings.ToLower(paramName) == strings.ToLower(expectedParam)
}

// Check if parameter looks like an entity ID
func isEntityID(paramName string) bool {
	return strings.HasSuffix(strings.ToLower(paramName), "_id") || strings.ToLower(paramName) == "id"
}

// Validate operations against the identified primary ID
func validateOperationParameters(resource *ResourceInfo, primaryID string, spec OpenAPISpec) map[string]OperationInfo {
	validOperations := make(map[string]OperationInfo)

	// Get entity properties once for this resource
	entityProps := getEntityProperties(resource.EntityName, spec)

	for crudType, opInfo := range resource.Operations {
		pathParams := extractPathParameters(opInfo.Path)

		if crudType == "create" || crudType == "list" {
			// These operations should not have the entity's primary ID in path
			hasPrimaryID := false
			for _, param := range pathParams {
				if param == primaryID {
					hasPrimaryID = true
					break
				}
			}

			if hasPrimaryID {
				fmt.Printf("    Skipping %s operation %s: unexpectedly has primary ID %s in path\n",
					crudType, opInfo.Path, primaryID)
				continue
			}

			validOperations[crudType] = opInfo
			continue
		}

		// READ, UPDATE, DELETE should have exactly the primary ID
		hasPrimaryID := false
		hasConflictingEntityIDs := false

		for _, param := range pathParams {
			if param == primaryID {
				hasPrimaryID = true
			} else if isEntityID(param) {
				// This is another ID-like parameter
				// Check if it maps to a field in the entity (not the primary id field)
				if checkFieldExistsInEntity(param, entityProps) {
					// This parameter maps to a real entity field - it's valid
					fmt.Printf("    Parameter %s maps to entity field - keeping operation %s %s\n",
						param, crudType, opInfo.Path)
				} else {
					// This ID parameter doesn't map to any entity field
					if mapsToEntityID(param, resource.EntityName) {
						// This would also try to map to the primary ID - CONFLICT!
						fmt.Printf("    Skipping %s operation %s: parameter %s would conflict with primary ID %s (both map to entity.id)\n",
							crudType, opInfo.Path, param, primaryID)
						hasConflictingEntityIDs = true
						break
					} else {
						// This is an unmappable ID parameter
						fmt.Printf("    Skipping %s operation %s: unmappable ID parameter %s (not in entity schema)\n",
							crudType, opInfo.Path, param)
						hasConflictingEntityIDs = true
						break
					}
				}
			}
			// Non-ID parameters are always OK
		}

		if !hasPrimaryID {
			fmt.Printf("    Skipping %s operation %s: missing primary ID %s\n",
				crudType, opInfo.Path, primaryID)
			continue
		}

		if hasConflictingEntityIDs {
			continue // Already logged above
		}

		validOperations[crudType] = opInfo
	}

	fmt.Printf("    Valid operations after parameter validation: %v\n", getOperationTypes(validOperations))
	return validOperations
}

// Remove the helper function since we don't need it anymore

// Helper function to get operation types for logging
func getOperationTypes(operations map[string]OperationInfo) []string {
	var types []string
	for opType := range operations {
		types = append(types, opType)
	}
	return types
}

// Check if create and read operations have compatible schemas
func hasValidCreateReadConsistency(resource *ResourceInfo, spec OpenAPISpec) bool {
	if resource.CreateSchema == "" {
		return false
	}

	// Re-parse the spec to get raw schema data
	specData, err := json.Marshal(spec)
	if err != nil {
		return false
	}

	var rawSpec map[string]interface{}
	if err := json.Unmarshal(specData, &rawSpec); err != nil {
		return false
	}

	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	entityProps := getSchemaProperties(schemas, resource.EntityName)
	createProps := getSchemaProperties(schemas, resource.CreateSchema)

	if len(entityProps) == 0 || len(createProps) == 0 {
		return false
	}

	// Count overlapping manageable properties
	commonManageableProps := 0
	createManageableProps := 0

	for prop := range createProps {
		if !isSystemProperty(prop) {
			createManageableProps++
			if entityProps[prop] != nil {
				commonManageableProps++
			}
		}
	}

	// Need at least some manageable properties
	if createManageableProps == 0 {
		return false
	}

	// Require at least 30% overlap of create properties to exist in entity
	// This is more lenient than the 50% I had before
	overlapRatio := float64(commonManageableProps) / float64(createManageableProps)
	return overlapRatio >= 0.3
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

func generateOverlay(resources map[string]*ResourceInfo, spec OpenAPISpec, manualMappings *ManualMappings) *Overlay {
	overlay := &Overlay{
		Overlay: "1.0.0",
	}

	overlay.Info.Title = "Terraform Provider Overlay"
	overlay.Info.Version = "1.0.0"
	overlay.Info.Description = "Auto-generated overlay for Terraform resources"

	// Clean up resources by removing manually ignored operations
	cleanedResources := cleanResourcesWithManualMappings(resources, manualMappings)

	// Separate viable and non-viable resources
	viableResources := make(map[string]*ResourceInfo)
	skippedResources := make([]string, 0)

	for name, resource := range cleanedResources {
		if isTerraformViable(resource, spec) {
			viableResources[name] = resource
		} else {
			skippedResources = append(skippedResources, name)
		}
	}

	fmt.Printf("\n=== Overlay Generation Analysis ===\n")
	fmt.Printf("Total resources found: %d\n", len(cleanedResources))
	fmt.Printf("Viable for Terraform: %d\n", len(viableResources))
	fmt.Printf("Skipped (non-viable): %d\n", len(skippedResources))

	if len(skippedResources) > 0 {
		fmt.Printf("\nSkipped resources:\n")
		for _, skipped := range skippedResources {
			fmt.Printf("  - %s\n", skipped)
		}
	}

	// Filter operations with unmappable path parameters
	fmt.Printf("\n=== Operation-Level Filtering ===\n")
	filteredResources := filterOperationsWithUnmappableParameters(viableResources, spec)

	// Update description with actual count
	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(filteredResources))

	// Detect property mismatches for filtered resources only
	resourceMismatches := detectPropertyMismatches(filteredResources, spec)

	// Detect CRUD inconsistencies for filtered resources only
	resourceCRUDInconsistencies := detectCRUDInconsistencies(filteredResources, spec)

	// Track which properties already have ignore actions to avoid duplicates
	ignoreTracker := make(map[string]map[string]bool) // map[schemaName][propertyName]bool

	// Generate actions only for filtered resources
	for _, resource := range filteredResources {
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

		// Add entity operations and parameter matching
		for crudType, opInfo := range resource.Operations {
			// Double-check that this specific operation isn't in the ignore list
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) {
				fmt.Printf("    Skipping ignored operation during overlay generation: %s %s\n", opInfo.Method, opInfo.Path)
				continue
			}

			entityOp := mapCrudToEntityOperation(crudType, resource.EntityName)

			operationUpdate := map[string]interface{}{
				"x-speakeasy-entity-operation": entityOp,
			}

			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: fmt.Sprintf("$.paths[\"%s\"].%s", opInfo.Path, opInfo.Method),
				Update: operationUpdate,
			})

			// Apply parameter matching for operations that use the primary ID
			if resource.PrimaryID != "" && (crudType == "read" || crudType == "update" || crudType == "delete") {
				pathParams := extractPathParameters(opInfo.Path)
				for _, param := range pathParams {
					if param == resource.PrimaryID {
						// Check for manual parameter mapping first
						if manualMatch, hasManual := getManualParameterMatch(opInfo.Path, opInfo.Method, param, manualMappings); hasManual {
							// Only apply manual match if it's different from the parameter name
							if manualMatch != param {
								fmt.Printf("    Manual parameter mapping: %s in %s %s -> %s\n", param, opInfo.Method, opInfo.Path, manualMatch)
								overlay.Actions = append(overlay.Actions, OverlayAction{
									Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
										opInfo.Path, opInfo.Method, param),
									Update: map[string]interface{}{
										"x-speakeasy-match": manualMatch,
									},
								})
							} else {
								fmt.Printf("    Skipping manual parameter mapping: %s already matches target field %s (would create circular reference)\n", param, manualMatch)
							}
						} else {
							// Skip x-speakeasy-match when parameter name would map to itself
							// This prevents circular references like {id} -> id
							if param == "id" {
								fmt.Printf("    Skipping x-speakeasy-match: parameter %s maps to same field (avoiding circular reference)\n", param)
							} else {
								// Apply x-speakeasy-match for parameters that need mapping (e.g., change_event_id -> id)
								fmt.Printf("    Applying x-speakeasy-match to %s in %s %s -> id\n", param, opInfo.Method, opInfo.Path)
								overlay.Actions = append(overlay.Actions, OverlayAction{
									Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
										opInfo.Path, opInfo.Method, param),
									Update: map[string]interface{}{
										"x-speakeasy-match": "id",
									},
								})
							}
						}
					}
				}
			}
		}
	}

	// Process parameter matching is now handled inline above
	// No need for separate addEntityLevelParameterMatches call

	fmt.Printf("\n=== Overlay Generation Complete ===\n")
	fmt.Printf("Generated %d actions for %d viable resources\n", len(overlay.Actions), len(filteredResources))

	// Count ignore actions and match actions
	totalIgnores := 0
	totalMatches := 0
	for _, action := range overlay.Actions {
		if _, hasIgnore := action.Update["x-speakeasy-ignore"]; hasIgnore {
			totalIgnores++
		}
		if _, hasMatch := action.Update["x-speakeasy-match"]; hasMatch {
			totalMatches++
		}
	}

	if totalIgnores > 0 {
		fmt.Printf("✅ %d speakeasy ignore actions added for property issues\n", totalIgnores)
	}
	if totalMatches > 0 {
		fmt.Printf("✅ %d speakeasy match actions added for primary ID parameters\n", totalMatches)
	}

	return overlay
}

// Clean up resources by removing operations that are manually ignored
func cleanResourcesWithManualMappings(resources map[string]*ResourceInfo, manualMappings *ManualMappings) map[string]*ResourceInfo {
	cleanedResources := make(map[string]*ResourceInfo)

	fmt.Printf("\n=== Cleaning Resources with Manual Mappings ===\n")

	for name, resource := range resources {
		cleanedResource := &ResourceInfo{
			EntityName:   resource.EntityName,
			SchemaName:   resource.SchemaName,
			ResourceName: resource.ResourceName,
			Operations:   make(map[string]OperationInfo),
			CreateSchema: resource.CreateSchema,
			UpdateSchema: resource.UpdateSchema,
			PrimaryID:    resource.PrimaryID,
		}

		operationsRemoved := 0

		// Copy operations that aren't manually ignored
		for crudType, opInfo := range resource.Operations {
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) {
				fmt.Printf("  Removing manually ignored operation: %s %s (was %s for %s)\n",
					opInfo.Method, opInfo.Path, crudType, resource.EntityName)
				operationsRemoved++
			} else {
				cleanedResource.Operations[crudType] = opInfo
			}
		}

		// Only include resource if it still has operations after cleaning
		if len(cleanedResource.Operations) > 0 {
			cleanedResources[name] = cleanedResource
			if operationsRemoved > 0 {
				fmt.Printf("  Resource %s: kept %d operations, removed %d manually ignored\n",
					name, len(cleanedResource.Operations), operationsRemoved)
			}
		} else {
			fmt.Printf("  Resource %s: removed entirely (all operations were manually ignored)\n", name)
		}
	}

	fmt.Printf("Manual mapping cleanup: %d → %d resources\n", len(resources), len(cleanedResources))
	return cleanedResources
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

// Detect schema property inconsistencies (extracted from detectCRUDInconsistencies)
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
		resourceInconsistencies := detectSchemaPropertyInconsistencies(resource, schemas)

		// Check if we have fundamental validation errors that make the resource non-viable
		for _, inconsistency := range resourceInconsistencies {
			if inconsistency.PropertyName == "RESOURCE_VALIDATION" {
				fmt.Printf("⚠️  Resource %s (%s) validation failed: %s\n",
					resource.ResourceName, resource.EntityName, inconsistency.Description)
				// Mark the entire resource as having issues but don't add to inconsistencies
				// as this will be handled in the viability check
				continue
			}
		}

		// Only add property-level inconsistencies for viable resources
		var validInconsistencies []CRUDInconsistency
		for _, inconsistency := range resourceInconsistencies {
			if inconsistency.PropertyName != "RESOURCE_VALIDATION" {
				validInconsistencies = append(validInconsistencies, inconsistency)
			}
		}

		if len(validInconsistencies) > 0 {
			inconsistencies[resource.EntityName] = validInconsistencies
		}
	}

	return inconsistencies
}

// Detect schema property inconsistencies (simplified CRUD detection)
func detectSchemaPropertyInconsistencies(resource *ResourceInfo, schemas map[string]interface{}) []CRUDInconsistency {
	var inconsistencies []CRUDInconsistency

	// First, validate that we have the minimum required operations for Terraform
	_, hasCreate := resource.Operations["create"]
	_, hasRead := resource.Operations["read"]

	if !hasCreate || !hasRead {
		// Return a fundamental inconsistency - resource is not viable for Terraform
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "missing-required-operations",
			Description:       fmt.Sprintf("Resource missing required operations: Create=%v, Read=%v", hasCreate, hasRead),
			SchemasToIgnore:   []string{}, // Don't ignore anything, this makes the whole resource invalid
		}
		return []CRUDInconsistency{inconsistency}
	}

	// Validate that we have a create schema
	if resource.CreateSchema == "" {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "missing-create-schema",
			Description:       "Resource has CREATE operation but no request schema defined",
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

	// Get properties from each schema
	entityProps := getSchemaProperties(schemas, resource.EntityName)
	createProps := getSchemaProperties(schemas, resource.CreateSchema)
	updateProps := map[string]interface{}{}

	if resource.UpdateSchema != "" {
		updateProps = getSchemaProperties(schemas, resource.UpdateSchema)
	}

	// Validate that schemas exist and have properties
	if len(entityProps) == 0 {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "invalid-entity-schema",
			Description:       fmt.Sprintf("Entity schema '%s' not found or has no properties", resource.EntityName),
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

	if len(createProps) == 0 {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "invalid-create-schema",
			Description:       fmt.Sprintf("Create schema '%s' not found or has no properties", resource.CreateSchema),
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

	// Check for minimum viable overlap between create and entity schemas
	commonManageableProps := 0
	createManageableProps := 0

	for prop := range createProps {
		if !isSystemProperty(prop) {
			createManageableProps++
			if entityProps[prop] != nil {
				commonManageableProps++
			}
		}
	}

	if createManageableProps == 0 {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "no-manageable-properties",
			Description:       "Create schema has no manageable properties (all are system properties)",
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

	// Require reasonable overlap between create and entity schemas
	overlapRatio := float64(commonManageableProps) / float64(createManageableProps)
	if overlapRatio < 0.3 { // At least 30% overlap required
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "insufficient-schema-overlap",
			Description:       fmt.Sprintf("Insufficient overlap between create and entity schemas: %.1f%% (%d/%d properties)", overlapRatio*100, commonManageableProps, createManageableProps),
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

	// Now check individual property inconsistencies for viable resources
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
			inconsistencies = append(inconsistencies, inconsistency)
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

// Filter out operations with unmappable path parameters - simplified now that we handle this in viability check
func filterOperationsWithUnmappableParameters(resources map[string]*ResourceInfo, spec OpenAPISpec) map[string]*ResourceInfo {
	// The complex parameter filtering is now handled in isTerraformViable()
	// This function now just returns the resources as-is since they've already been validated
	fmt.Printf("Operation filtering handled during viability check\n")
	return resources
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
	fmt.Println("1. Load manual mappings for edge cases")
	fmt.Println("2. Identify entity schemas and match operations to entities")
	fmt.Println("3. Apply manual ignore/entity/match mappings during analysis")
	fmt.Println("4. Clean resources by removing manually ignored operations")
	fmt.Println("5. Analyze ID patterns and choose consistent primary ID per entity")
	fmt.Println("6. Filter operations with unmappable path parameters")
	fmt.Println("7. Skip annotations for non-viable resources")
	fmt.Println("8. Mark viable entity schemas with x-speakeasy-entity")
	fmt.Println("9. Tag viable operations with x-speakeasy-entity-operation")
	fmt.Println("10. Mark chosen primary ID with x-speakeasy-match")
	fmt.Println("11. Apply x-speakeasy-ignore: true to problematic properties")
}
