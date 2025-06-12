package main

import (
	"encoding/json"
	"fmt"
)

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

func detectPropertyMismatches(resources map[string]*ResourceInfo, spec OpenAPISpec) map[string][]PropertyMismatch {
	mismatches := make(map[string][]PropertyMismatch)

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
			if hasStructuralMismatch(entityProp, requestProp) {
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
func hasStructuralMismatch(entityProp, requestProp interface{}) bool {
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
	_, hasPut := resource.Operations["put"]
	_, hasRead := resource.Operations["read"]

	createOrPut := hasCreate || hasPut
	if !createOrPut || !hasRead {
		// Return a fundamental inconsistency - resource is not viable for Terraform
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "missing-required-operations",
			Description:       fmt.Sprintf("Resource missing required operations: Create=%v, Read=%v", hasCreate, hasRead),
			SchemasToIgnore:   []string{}, // Don't ignore anything, this makes the whole resource invalid since Terraform needs a create and a read, at minimum
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
