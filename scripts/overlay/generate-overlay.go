package main

import (
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

		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			addIgnoreActionsForMismatches(overlay, resource, mismatches, ignoreTracker)
		}

		if inconsistencies, exists := resourceCRUDInconsistencies[resource.EntityName]; exists {
			addIgnoreActionsForInconsistencies(overlay, resource, inconsistencies, ignoreTracker)
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
							// Avoid creating circular reference by not mapping id -> id or slug -> slug
							if param != "id" && param != "slug" {
								primaryIdentifier := "id"
								if strings.Contains(resource.PrimaryID, "slug") {
									primaryIdentifier = "slug"
								}
								overlay.Actions = append(overlay.Actions, OverlayAction{
									Target: fmt.Sprintf("$.paths[\"%s\"].%s.parameters[?(@.name==\"%s\")]",
										opInfo.Path, opInfo.Method, param),
									Update: map[string]interface{}{
										"x-speakeasy-match": primaryIdentifier,
									},
								})
							}
						}
					}
				}
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
