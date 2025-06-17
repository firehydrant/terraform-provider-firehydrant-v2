package main

import "fmt"

func addIgnoreActionsForMismatches(overlay *Overlay, resource *ResourceInfo, mismatches []PropertyMismatch, ignoreTracker map[string]map[string]bool) {
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

func addIgnoreActionsForInconsistencies(overlay *Overlay, inconsistencies []CRUDInconsistency, ignoreTracker map[string]map[string]bool) {
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
