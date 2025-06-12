package main

import "fmt"

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
	mappedResources := applyManualMappings(resources, manualMappings)

	// Separate viable and non-viable resources
	viableResources := make(map[string]*ResourceInfo)
	skippedResources := make([]string, 0)

	for name, resource := range mappedResources {
		if isTerraformViable(resource, spec) {
			viableResources[name] = resource
		} else {
			skippedResources = append(skippedResources, name)
		}
	}

	fmt.Printf("\n=== Overlay Generation Analysis ===\n")
	fmt.Printf("Resources after Manual Mapping: %d\n", len(mappedResources))
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

	// Update description with actual count
	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(viableResources))

	// Detect property mismatches for filtered resources only
	resourceMismatches := detectPropertyMismatches(viableResources, spec)

	// Detect CRUD inconsistencies for filtered resources only
	resourceCRUDInconsistencies := detectCRUDInconsistencies(viableResources, spec)

	// Track which properties already have ignore actions to avoid duplicates
	ignoreTracker := make(map[string]map[string]bool) // map[schemaName][propertyName]bool

	// Generate actions only for filtered resources
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

	fmt.Printf("\n=== Overlay Generation Complete ===\n")
	fmt.Printf("Generated %d actions for %d viable resources\n", len(overlay.Actions), len(viableResources))

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
