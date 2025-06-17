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

	// Clean up resources by removing manually ignored operations

	// In general, we don't need to alter the spec at this point
	//   We can simply not apply x-speakeasy extentions to these entities
	//   Without the extensions, they will not be registered in the Terraform provider
	cleanedResources := applyManualMappings(resources, manualMappings)

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
	fmt.Printf("Resources after Manual Mappings: %d\n", len(cleanedResources))
	fmt.Printf("Viable for Terraform: %d\n", len(viableResources))
	fmt.Printf("Skipped (non-viable): %d\n", len(skippedResources))

	// Helpful for debugging
	// if len(skippedResources) > 0 {
	// 	fmt.Printf("\nSkipped resources:\n")
	// 	for _, skipped := range skippedResources {
	// 		fmt.Printf("  - %s\n", skipped)
	// 	}
	// }

	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(viableResources))

	resourceMismatches := detectPropertyMismatches(viableResources, spec)

	resourceCRUDInconsistencies := detectCRUDInconsistencies(viableResources, spec)

	ignoreTracker := make(map[string]map[string]bool)

	requiredFieldsMap := make(map[string]map[string]bool)
	specData, _ := json.Marshal(spec)
	var rawSpec map[string]interface{}
	json.Unmarshal(specData, &rawSpec)
	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

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
		if resource.CreateSchema != "" && ignoreTracker[resource.CreateSchema] == nil {
			ignoreTracker[resource.CreateSchema] = make(map[string]bool)
		}
		if resource.UpdateSchema != "" && ignoreTracker[resource.UpdateSchema] == nil {
			ignoreTracker[resource.UpdateSchema] = make(map[string]bool)
		}

		requiredFields := requiredFieldsMap[resource.EntityName]

		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			for _, mismatch := range mismatches {
				if requiredFields[mismatch.PropertyName] {
					if resource.CreateSchema != "" {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.CreateSchema, mismatch.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-param-optional": true,
							},
						})
					}
				} else {
					addIgnoreActionsForMismatches(overlay, resource, mismatches, ignoreTracker)
				}
			}
		}

		if inconsistencies, exists := resourceCRUDInconsistencies[resource.EntityName]; exists {
			for _, inconsistency := range inconsistencies {
				if requiredFields[inconsistency.PropertyName] {
					if resource.CreateSchema != "" {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.CreateSchema, inconsistency.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-param-optional": true,
							},
						})
					}
				} else {
					addIgnoreActionsForInconsistencies(overlay, inconsistencies, ignoreTracker)
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

							fmt.Printf("    Mapped parameter %s -> %s for %s %s\n",
								param, targetField, opInfo.Method, opInfo.Path)
						}
					}
				}
			}
		}
	}

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
