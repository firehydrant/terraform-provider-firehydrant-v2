package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func generateOverlay(resources map[string]*ResourceInfo, spec OpenAPISpec, manualMappings *ManualMappings) *Overlay {
	overlay := &Overlay{
		Overlay: "1.0.0",
	}

	overlay.Info.Title = "Terraform Provider Overlay"
	overlay.Info.Version = "1.0.0"
	overlay.Info.Description = "Auto-generated overlay for Terraform resources"

	viableResources := make(map[string]*ResourceInfo)
	skippedResources := make([]string, 0)

	for name, resource := range resources {
		if isTerraformViable(resource, spec, manualMappings) {
			viableResources[name] = resource
		} else {
			skippedResources = append(skippedResources, name)
		}
	}

	fmt.Printf("\n=== Overlay Generation Analysis ===\n")
	fmt.Printf("Viable for Terraform: %d\n", len(viableResources))
	fmt.Printf("Skipped (non-viable): %d\n", len(skippedResources))

	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(viableResources))

	specData, _ := json.Marshal(spec)
	var rawSpec map[string]interface{}
	json.Unmarshal(specData, &rawSpec)
	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	resourceMismatches := detectPropertyMismatches(viableResources, spec)
	resourceCRUDInconsistencies := detectCRUDInconsistencies(viableResources, spec)

	ignoreTracker := make(map[string]map[string]bool)
	readonlyTracker := make(map[string]map[string]bool)
	additionalPropsTracker := make(map[string]map[string]bool) // New tracker

	requiredFieldsMap := make(map[string]map[string]bool)
	for _, resource := range viableResources {
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

	additionalPropsMappings := getAdditionalPropertiesMappings(manualMappings)

	fmt.Printf("\n=== Applying Additional Properties Mappings ===\n")
	for schemaName, properties := range additionalPropsMappings {
		if additionalPropsTracker[schemaName] == nil {
			additionalPropsTracker[schemaName] = make(map[string]bool)
		}

		for _, propertyPath := range properties {
			// Skip if already processed
			if additionalPropsTracker[schemaName][propertyPath] {
				continue
			}

			targetPath := buildAdditionalPropertiesPath(schemaName, propertyPath)

			// Add additionalProperties: true
			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: targetPath,
				Update: map[string]interface{}{
					"additionalProperties": true,
				},
			})

			// Check if this is an entity schema (not a request schema)
			isEntitySchema := false
			for _, resource := range viableResources {
				if resource.EntityName == schemaName {
					isEntitySchema = true
					break
				}
			}

			// ALWAYS mark as readonly when applying to any entity schema
			// This is to prevent idempedance errors during generation, as the request schema frequently has a different type for the given field
			if isEntitySchema {
				// Initialize readonly tracker for this schema if needed
				if readonlyTracker[schemaName] == nil {
					readonlyTracker[schemaName] = make(map[string]bool)
				}

				// For the actual property that has additionalProperties
				// We need to mark it as readonly at the exact same path
				readonlyTarget := targetPath

				// Also track the top-level property for ignore prevention
				propParts := strings.Split(propertyPath, ".")
				topLevelProp := propParts[0]

				// Mark the actual property (at same path as additionalProperties) as readonly
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: readonlyTarget,
					Update: map[string]interface{}{
						"x-speakeasy-param-readonly": true,
					},
				})

				// Track both the full path and top-level property
				readonlyTracker[schemaName][propertyPath] = true
				readonlyTracker[schemaName][topLevelProp] = true
			}

			additionalPropsTracker[schemaName][propertyPath] = true
		}
	}

	for _, resource := range viableResources {
		entityUpdate := map[string]interface{}{
			"x-speakeasy-entity": resource.EntityName,
		}

		overlay.Actions = append(overlay.Actions, OverlayAction{
			Target: fmt.Sprintf("$.components.schemas.%s", resource.SchemaName),
			Update: entityUpdate,
		})

		if ignoreTracker[resource.EntityName] == nil {
			ignoreTracker[resource.EntityName] = make(map[string]bool)
		}
		if readonlyTracker[resource.EntityName] == nil {
			readonlyTracker[resource.EntityName] = make(map[string]bool)
		}
		if resource.CreateSchema != "" && ignoreTracker[resource.CreateSchema] == nil {
			ignoreTracker[resource.CreateSchema] = make(map[string]bool)
		}
		if resource.UpdateSchema != "" && ignoreTracker[resource.UpdateSchema] == nil {
			ignoreTracker[resource.UpdateSchema] = make(map[string]bool)
		}

		requiredFields := requiredFieldsMap[resource.EntityName]

		entitySchema, _ := schemas[resource.EntityName].(map[string]interface{})
		var createSchema map[string]interface{}
		var updateSchema map[string]interface{}

		if resource.CreateSchema != "" {
			createSchema, _ = schemas[resource.CreateSchema].(map[string]interface{})
		}
		if resource.UpdateSchema != "" {
			updateSchema, _ = schemas[resource.UpdateSchema].(map[string]interface{})
		}

		readonlyActions := detectReadonlyFields(resource.EntityName, entitySchema, createSchema, updateSchema, schemas)
		for _, action := range readonlyActions {
			// Extract property name from target for tracking
			parts := strings.Split(action.Target, ".")
			if len(parts) >= 4 && parts[len(parts)-2] == "properties" {
				propName := parts[len(parts)-1]
				if !readonlyTracker[resource.EntityName][propName] {
					overlay.Actions = append(overlay.Actions, action)
					readonlyTracker[resource.EntityName][propName] = true
				}
			}
		}

		// Handle mismatches - but skip if property was normalized
		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			for _, mismatch := range mismatches {
				// Check if this property should be skipped (was normalized)
				if shouldSkipIgnore(resource.EntityName, mismatch.PropertyName, viableResources, spec) {
					continue
				}

				if requiredFields[mismatch.PropertyName] {
					if resource.CreateSchema != "" {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s",
								resource.CreateSchema, mismatch.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-param-optional": true,
							},
						})
					}
				} else {
					addIgnoreActionsForMismatches(overlay, resource, []PropertyMismatch{mismatch}, ignoreTracker)
				}
			}
		}

		if inconsistencies, exists := resourceCRUDInconsistencies[resource.EntityName]; exists {
			for _, inconsistency := range inconsistencies {
				if shouldSkipIgnore(resource.EntityName, inconsistency.PropertyName, viableResources, spec) {
					continue
				}

				if requiredFields[inconsistency.PropertyName] {
					if resource.CreateSchema != "" {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s",
								resource.CreateSchema, inconsistency.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-param-optional": true,
							},
						})
					}
				} else {
					addIgnoreActionsForInconsistencies(overlay, []CRUDInconsistency{inconsistency}, ignoreTracker)
				}
			}
		}

		for crudType, opInfo := range resource.Operations {
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) {
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

			if resource.PrimaryID != "" && (crudType == "read" || crudType == "update" || crudType == "delete") {
				pathParams := extractPathParameters(opInfo.Path)
				for _, param := range pathParams {
					if param == resource.PrimaryID {
						if manualMatch, hasManual := getManualParameterMatch(opInfo.Path, opInfo.Method, param, manualMappings); hasManual {
							if manualMatch != param {
								overlay.Actions = append(overlay.Actions, OverlayAction{
									Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
										opInfo.Path, opInfo.Method, param),
									Update: map[string]interface{}{
										"x-speakeasy-match": manualMatch,
									},
								})
							}
						} else {
							var targetField string

							// Don't create circular mappings
							if param == "id" || param == "slug" {
								continue
							}

							// Get entity properties to check what fields exist
							entityProps := getEntityProperties(resource.EntityName, spec)
							_, hasID := entityProps["id"]
							_, hasSlug := entityProps["slug"]

							// Determine the target field based on parameter name and entity fields
							if strings.HasSuffix(param, "_slug") && hasSlug {
								targetField = "slug"
							} else if strings.Contains(param, "slug") && hasSlug && !hasID {
								targetField = "slug"
							} else if strings.HasSuffix(param, "_id") && hasID {
								targetField = "id"
							} else if hasID {
								targetField = "id"
							} else if hasSlug {
								targetField = "slug"
							} else {
								fmt.Printf("    Warning: Cannot map parameter %s - entity has neither id nor slug field\n", param)
								continue
							}

							overlay.Actions = append(overlay.Actions, OverlayAction{
								Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
									opInfo.Path, opInfo.Method, param),
								Update: map[string]interface{}{
									"x-speakeasy-match": targetField,
								},
							})
						}
					}
				}
			}
		}
	}

	// Apply manual property ignores
	manualPropertyIgnores := getManualPropertyIgnores(manualMappings)
	for schemaName, properties := range manualPropertyIgnores {
		for _, propertyName := range properties {
			// Initialize ignore tracker if needed
			if ignoreTracker[schemaName] == nil {
				ignoreTracker[schemaName] = make(map[string]bool)
			}

			// Only add if not already ignored
			if !ignoreTracker[schemaName][propertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", schemaName, propertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[schemaName][propertyName] = true
				fmt.Printf("✅ Added manual property ignore: %s.%s\n", schemaName, propertyName)
			}
		}
	}

	fmt.Printf("\n=== Overlay Generation Complete ===\n")
	fmt.Printf("Generated %d actions for %d viable resources\n", len(overlay.Actions), len(viableResources))

	// Count different types of actions
	totalIgnores := 0
	totalMatches := 0
	totalReadonly := 0
	totalAdditionalProps := 0

	for _, action := range overlay.Actions {
		if _, hasIgnore := action.Update["x-speakeasy-ignore"]; hasIgnore {
			totalIgnores++
		}
		if _, hasMatch := action.Update["x-speakeasy-match"]; hasMatch {
			totalMatches++
		}
		if _, hasReadonly := action.Update["x-speakeasy-param-readonly"]; hasReadonly {
			totalReadonly++
		}
		if val, hasAdditional := action.Update["additionalProperties"]; hasAdditional && val == true {
			totalAdditionalProps++
		}
	}

	if totalAdditionalProps > 0 {
		fmt.Printf("✅ %d additionalProperties actions added for flexible schema fields\n", totalAdditionalProps)
	}
	if totalIgnores > 0 {
		fmt.Printf("✅ %d speakeasy ignore actions added for unresolved property issues\n", totalIgnores)
	}
	if totalMatches > 0 {
		fmt.Printf("✅ %d speakeasy match actions added for primary ID parameters\n", totalMatches)
	}
	if totalReadonly > 0 {
		fmt.Printf("✅ %d speakeasy readonly actions added (includes computed fields and additional properties)\n", totalReadonly)
	}

	return overlay
}
