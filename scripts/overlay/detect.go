package main

import (
	"fmt"
	"strings"

	"github.com/firehydrant/terraform-provider-firehydrant/scripts/common"
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

func detectPropertyMismatches(resources map[string]*ResourceInfo, schemas map[string]interface{}, requiredFieldsMap map[string]map[string]bool) map[string][]PropertyMismatch {
	mismatches := make(map[string][]PropertyMismatch)

	for _, resource := range resources {
		var resourceMismatches []PropertyMismatch
		requiredFields := requiredFieldsMap[resource.EntityName]

		if resource.CreateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.CreateSchema].(map[string]interface{}); exists {
					createMismatches := findPropertyMismatches(entitySchema, requestSchema, "create", requiredFields)
					resourceMismatches = append(resourceMismatches, createMismatches...)
				}
			}
		}

		if resource.UpdateSchema != "" {
			if entitySchema, exists := schemas[resource.EntityName].(map[string]interface{}); exists {
				if requestSchema, exists := schemas[resource.UpdateSchema].(map[string]interface{}); exists {
					updateMismatches := findPropertyMismatches(entitySchema, requestSchema, "update", requiredFields)
					resourceMismatches = append(resourceMismatches, updateMismatches...)
				}
			}
		}

		if len(resourceMismatches) > 0 {
			mismatches[resource.EntityName] = resourceMismatches
		}
	}

	for _, schema := range schemas {
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			delete(schemaMap, "__spec")
		}
	}

	return mismatches
}

func findPropertyMismatches(entitySchema, requestSchema map[string]interface{}, operation string, requiredFields map[string]bool) []PropertyMismatch {
	var mismatches []PropertyMismatch

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	requestProps, _ := requestSchema["properties"].(map[string]interface{})

	if entityProps == nil || requestProps == nil {
		return mismatches
	}

	spec, _ := entitySchema["__spec"].(map[string]interface{})
	components, _ := spec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	for propName, entityProp := range entityProps {
		if requestProp, exists := requestProps[propName]; exists {
			// Skip required fields, they should already be handled in normalization
			// Ensures we do not accidentally ignore any required fields, prevents generating unusable resources
			if requiredFields[propName] {
				continue
			}

			// Use the SAME structural comparison logic that works in detectReadonlyFields
			if common.HasTopLevelStructuralMismatch(entityProp, requestProp, schemas) {
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

// Describe the structural difference for reporting/debugging
func describeStructuralDifference(entityProp, requestProp interface{}) string {
	entityStructure := common.GetPropertyStructure(entityProp)
	requestStructure := common.GetPropertyStructure(requestProp)
	return fmt.Sprintf("request structure '%s' != response structure '%s'", requestStructure, entityStructure)
}

func detectCRUDInconsistencies(resources map[string]*ResourceInfo, schemas map[string]interface{}) map[string][]CRUDInconsistency {
	inconsistencies := make(map[string][]CRUDInconsistency)

	requiredFieldsMap := make(map[string]map[string]bool)
	for _, resource := range resources {
		requiredFields := make(map[string]bool)

		if resource.CreateSchema != "" {
			if createSchema, ok := schemas[resource.CreateSchema].(map[string]interface{}); ok {
				if required, ok := createSchema["required"].([]interface{}); ok {
					for _, field := range required {
						if fieldName, ok := field.(string); ok {
							requiredFields[fieldName] = true
						}
					}
				}
			}
		}
		requiredFieldsMap[resource.EntityName] = requiredFields
	}

	for _, resource := range resources {
		requiredFields := requiredFieldsMap[resource.EntityName]
		resourceInconsistencies := detectSchemaPropertyInconsistencies(resource, schemas, requiredFields)

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
func detectSchemaPropertyInconsistencies(resource *ResourceInfo, schemas map[string]interface{}, requiredFields map[string]bool) []CRUDInconsistency {
	var inconsistencies []CRUDInconsistency

	// First, validate that we have the minimum required operations for Terraform
	_, hasCreate := resource.Operations["create"] // in absence of a POST operation, PUT operations are registered as create
	_, hasRead := resource.Operations["read"]

	if !hasCreate || !hasRead {
		// Return a fundamental inconsistency - resource is not viable for Terraform
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "missing-required-operations",
			Description:       fmt.Sprintf("Resource missing required operations: Create=%v, Read=%v", hasCreate, hasRead),
			SchemasToIgnore:   []string{}, // Don't ignore anything, this makes the whole resource invalid since Terraform needs a create and a read, at minimum
		}
		return []CRUDInconsistency{inconsistency}
	}

	if resource.CreateSchema == "" {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "missing-create-schema",
			Description:       "Resource has CREATE operation but no request schema defined",
			SchemasToIgnore:   []string{},
		}
		return []CRUDInconsistency{inconsistency}
	}

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

	// Need at least some managable properties
	// if there is aleast one common manageable property, we won't rule it out at this stage
	if createManageableProps == 0 {
		inconsistency := CRUDInconsistency{
			PropertyName:      "RESOURCE_VALIDATION",
			InconsistencyType: "no-manageable-properties",
			Description:       "Create schema has no manageable properties (all are system properties)",
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

		if requiredFields[propName] {
			// Skip required fields - they should already be handled in normalization
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

// isComputedField checks if a field name appears to be computed/readonly
func isComputedField(fieldName string) bool {
	// Most of these are specific to runbook steps
	computedPatterns := []string{
		"created_at", "updated_at", "created_by", "updated_by",
		"is_editable", "votes", "categories", "runbook_template_id",
		"action_elements", "step_elements", "automatic", "repeats",
		"repeats_duration", "delay_duration", "reruns",
	}

	lowerField := strings.ToLower(fieldName)
	for _, pattern := range computedPatterns {
		if lowerField == pattern || strings.HasSuffix(lowerField, "_"+pattern) {
			return true
		}
	}

	return false
}

// detectReadonlyFields finds type differences between entity and request schemas
// and marks properties as readonly if they don't exist in request schemas or have structural mismatches
// this includes nested properties and will mark them as readonly if they are not present in the request schemas
func detectReadonlyFields(entityName string, entitySchema, createSchema, updateSchema map[string]interface{},
	schemas map[string]interface{}) []OverlayAction {

	var actions []OverlayAction

	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	createProps := make(map[string]interface{})
	updateProps := make(map[string]interface{})

	if createSchema != nil {
		createProps, _ = createSchema["properties"].(map[string]interface{})
	}
	if updateSchema != nil {
		updateProps, _ = updateSchema["properties"].(map[string]interface{})
	}

	for propName, entityProp := range entityProps {
		_, inCreate := createProps[propName]
		_, inUpdate := updateProps[propName]

		// Case 1: Property doesn't exist in any request schema - mark as readonly
		if !inCreate && !inUpdate {
			// Skip if it's a system field that we expect to be readonly
			if propName == "id" || propName == "slug" {
				continue
			}

			actions = append(actions, OverlayAction{
				Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", entityName, propName),
				Update: map[string]interface{}{
					"x-speakeasy-param-readonly": true,
				},
			})
		} else {
			// Case 2: Property exists in request schema(s) - check for structural mismatches
			var createProp, updateProp interface{}
			hasStructuralMismatch := false

			if inCreate {
				createProp = createProps[propName]
				if common.HasTopLevelStructuralMismatch(entityProp, createProp, schemas) {
					hasStructuralMismatch = true
				}
			}

			if inUpdate {
				updateProp = updateProps[propName]
				if common.HasTopLevelStructuralMismatch(entityProp, updateProp, schemas) {
					hasStructuralMismatch = true
				}
			}

			// If there's a structural mismatch at the top level, mark the entity property as readonly
			if hasStructuralMismatch {
				actions = append(actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", entityName, propName),
					Update: map[string]interface{}{
						"x-speakeasy-param-readonly": true,
					},
				})
			} else {
				// No top-level mismatch - check nested properties for partial readonly
				actions = append(actions, detectNestedReadonlyFields(entityName, propName,
					entityProp, createProp, updateProp, schemas)...)
			}
		}
	}

	return actions
}

// detectNestedReadonlyFields checks for readonly nested properties within manageable fields
func detectNestedReadonlyFields(entityName, propName string, entityProp, createProp, updateProp interface{},
	schemas map[string]interface{}) []OverlayAction {

	var actions []OverlayAction

	// Resolve any refs
	entityResolved, _ := common.GetResolvedPropertyType(entityProp, schemas)
	createResolved, _ := common.GetResolvedPropertyType(createProp, schemas)
	updateResolved, _ := common.GetResolvedPropertyType(updateProp, schemas)

	// For objects, check nested properties
	if entityResolved != nil && entityResolved["type"] == "object" {
		entityNestedProps, _ := entityResolved["properties"].(map[string]interface{})
		createNestedProps := make(map[string]interface{})
		updateNestedProps := make(map[string]interface{})

		if createResolved != nil {
			createNestedProps, _ = createResolved["properties"].(map[string]interface{})
		}
		if updateResolved != nil {
			updateNestedProps, _ = updateResolved["properties"].(map[string]interface{})
		}

		// Check each nested property
		for nestedPropName := range entityNestedProps {
			_, inCreate := createNestedProps[nestedPropName]
			_, inUpdate := updateNestedProps[nestedPropName]

			if !inCreate && !inUpdate {
				actions = append(actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s.properties.%s",
						entityName, propName, nestedPropName),
					Update: map[string]interface{}{
						"x-speakeasy-param-readonly": true,
					},
				})
			}
		}
	}

	// For arrays of objects, check item properties
	if entityResolved != nil && entityResolved["type"] == "array" {
		if items, ok := entityResolved["items"].(map[string]interface{}); ok {
			if itemProps, ok := items["properties"].(map[string]interface{}); ok {
				// Get request item properties
				createItemProps := make(map[string]interface{})
				updateItemProps := make(map[string]interface{})

				if createResolved != nil {
					if createItems, ok := createResolved["items"].(map[string]interface{}); ok {
						createItemProps, _ = createItems["properties"].(map[string]interface{})
					}
				}
				if updateResolved != nil {
					if updateItems, ok := updateResolved["items"].(map[string]interface{}); ok {
						updateItemProps, _ = updateItems["properties"].(map[string]interface{})
					}
				}

				// Check each item property
				for itemPropName := range itemProps {
					_, inCreate := createItemProps[itemPropName]
					_, inUpdate := updateItemProps[itemPropName]

					if !inCreate && !inUpdate || isComputedField(itemPropName) {
						actions = append(actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s.items.properties.%s",
								entityName, propName, itemPropName),
							Update: map[string]interface{}{
								"x-speakeasy-param-readonly": true,
							},
						})
					}
				}
			}
		}
	}

	return actions
}
