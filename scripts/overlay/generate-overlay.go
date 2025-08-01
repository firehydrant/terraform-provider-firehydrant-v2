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

	// Entity configuration contains information about which entities are enabled
	// or if we should process all entities by default
	entityConfig := buildEntityConfig(manualMappings)

	specData, _ := json.Marshal(spec)
	var rawSpec map[string]interface{}
	json.Unmarshal(specData, &rawSpec)
	components, _ := rawSpec["components"].(map[string]interface{})
	schemas, _ := components["schemas"].(map[string]interface{})

	viableResources := make(map[string]*ResourceInfo)
	nonViableResources := make([]string, 0)
	disabledResources := make([]string, 0)

	for name, resource := range resources {
		if !entityConfig.ShouldProcessEntity(resource.EntityName) {
			// We track disabled resources but still include them as viable so
			// they receive the appropriate speakeasy extensions incase they are related to our enabled entities via sub/parent relationships
			disabledResources = append(disabledResources, name)
		}

		if isTerraformViable(resource, manualMappings, schemas) {
			viableResources[name] = resource
		} else {
			nonViableResources = append(nonViableResources, name)
		}
	}

	fmt.Printf("\n=== Overlay Generation Analysis ===\n")
	if len(entityConfig.EnabledEntities) > 0 {
		fmt.Printf("Using Explicitly Enabled Entities\n")
		fmt.Printf("Enabled entities: %d\n", len(entityConfig.EnabledEntities))
	} else if entityConfig.HasExplicitEnabled {
		fmt.Printf("All entities enabled by default\n")
	}

	fmt.Printf("Viable for Terraform: %d\n", len(viableResources))
	fmt.Printf("Skipped (non-viable): %d\n", len(nonViableResources))

	overlay.Info.Description = fmt.Sprintf("Auto-generated overlay for %d viable Terraform resources", len(viableResources))

	requiredFieldsMap := extractRequiredFields(viableResources, schemas)

	resourceMismatches := detectPropertyMismatches(viableResources, schemas, requiredFieldsMap)
	resourceCRUDInconsistencies := detectCRUDInconsistencies(viableResources, schemas)

	ignoreTracker := make(map[string]map[string]bool)
	entityReadonlyTracker := make(map[string]map[string]bool) // Only for entity schema readonly properties
	nameOverrideTracker := make(map[string]map[string]bool)   // For request schema name overrides
	additionalPropsTracker := make(map[string]map[string]bool)

	additionalPropsMappings := getAdditionalPropertiesMappings(manualMappings)

	// For empty objects in our entities, we apply additionalProperties: true and mark as readonly
	// This allows terraform to track the response without needing to know the exact properties
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

			// We always mark as readonly when applying additionalProperties to any entity schema as terraform can't manage any property without a defined type/shape
			if strings.HasSuffix(schemaName, "Entity") {
				if entityReadonlyTracker[schemaName] == nil {
					entityReadonlyTracker[schemaName] = make(map[string]bool)
				}

				// For the actual property that has additionalProperties
				// We need to mark it as readonly at the exact same path
				readonlyTarget := targetPath

				// Also track the top-level property for ignore prevention
				propParts := strings.Split(propertyPath, ".")
				topLevelProp := propParts[0]

				// Only add readonly if not already applied (prevent duplicates)
				if !entityReadonlyTracker[schemaName][propertyPath] {
					// Mark the actual property (at same path as additionalProperties) as readonly
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: readonlyTarget,
						Update: map[string]interface{}{
							"x-speakeasy-param-readonly": true,
						},
					})

					// Track both the full path and top-level property
					entityReadonlyTracker[schemaName][propertyPath] = true
					entityReadonlyTracker[schemaName][topLevelProp] = true
				}
			}

			additionalPropsTracker[schemaName][propertyPath] = true
		}
	}

	fmt.Printf("\n=== Structural Mismatches & Type Inconsistencies ===\n")
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
		if entityReadonlyTracker[resource.EntityName] == nil {
			entityReadonlyTracker[resource.EntityName] = make(map[string]bool)
		}
		if resource.CreateSchema != "" {
			if ignoreTracker[resource.CreateSchema] == nil {
				ignoreTracker[resource.CreateSchema] = make(map[string]bool)
			}
			if nameOverrideTracker[resource.CreateSchema] == nil {
				nameOverrideTracker[resource.CreateSchema] = make(map[string]bool)
			}
		}
		if resource.UpdateSchema != "" {
			if ignoreTracker[resource.UpdateSchema] == nil {
				ignoreTracker[resource.UpdateSchema] = make(map[string]bool)
			}
			if nameOverrideTracker[resource.UpdateSchema] == nil {
				nameOverrideTracker[resource.UpdateSchema] = make(map[string]bool)
			}
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
				if !entityReadonlyTracker[resource.EntityName][propName] {
					overlay.Actions = append(overlay.Actions, action)
					entityReadonlyTracker[resource.EntityName][propName] = true
				}
			}
		}

		// Handle mismatches - these should apply name overrides to REQUEST schemas
		if mismatches, exists := resourceMismatches[resource.EntityName]; exists {
			for _, mismatch := range mismatches {

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
					// Apply name override to CREATE schema if not already applied
					if resource.CreateSchema != "" && !nameOverrideTracker[resource.CreateSchema][mismatch.PropertyName] {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s",
								resource.CreateSchema, mismatch.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-name-override": mismatch.PropertyName + "_input",
							},
						})
						nameOverrideTracker[resource.CreateSchema][mismatch.PropertyName] = true
					}

					// Apply name override to UPDATE schema if different from create and not already applied
					if resource.UpdateSchema != "" && resource.UpdateSchema != resource.CreateSchema &&
						!nameOverrideTracker[resource.UpdateSchema][mismatch.PropertyName] {
						overlay.Actions = append(overlay.Actions, OverlayAction{
							Target: fmt.Sprintf("$.components.schemas.%s.properties.%s",
								resource.UpdateSchema, mismatch.PropertyName),
							Update: map[string]interface{}{
								"x-speakeasy-name-override": mismatch.PropertyName + "_input",
							},
						})
						nameOverrideTracker[resource.UpdateSchema][mismatch.PropertyName] = true
					}
				}
			}
		}

		if inconsistencies, exists := resourceCRUDInconsistencies[resource.EntityName]; exists {
			for _, inconsistency := range inconsistencies {

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
					handleCRUDInconsistency(overlay, resource, inconsistency, entityReadonlyTracker)
				}
			}
		}

		for crudType, opInfo := range resource.Operations {
			// manual ignores are the only ignore operations that we should have in our overlay
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) || !entityConfig.ShouldProcessEntity(resource.EntityName) {
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
							entityProps := getSchemaProperties(schemas, resource.EntityName)
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

	// Process non-viable resources as potential data sources
	fmt.Printf("\n=== Data Source Generation for Non-Viable Resources ===\n")

	viableDatasources := make(map[string]*ResourceInfo)
	for name, resource := range resources {
		if !isTerraformViable(resource, manualMappings, schemas) {
			// Check if this resource can be a data source (has read operation)
			if isViableDatasource(resource) {
				viableDatasources[name] = resource
			}
		}
	}

	fmt.Printf("Potential data sources: %d\n", len(viableDatasources))

	manualPropertyIgnores := getManualPropertyIgnores(manualMappings)

	for _, resource := range viableDatasources {
		fmt.Printf("Processing data source: %s (%s)\n", resource.ResourceName, resource.EntityName)

		if ignoreTracker[resource.EntityName] == nil {
			ignoreTracker[resource.EntityName] = make(map[string]bool)
		}

		entityUpdate := map[string]interface{}{
			"x-speakeasy-entity": resource.EntityName,
		}

		overlay.Actions = append(overlay.Actions, OverlayAction{
			Target: fmt.Sprintf("$.components.schemas.%s", resource.SchemaName),
			Update: entityUpdate,
		})

		if properties, exists := manualPropertyIgnores[resource.EntityName]; exists {
			for _, propertyName := range properties {
				if !ignoreTracker[resource.EntityName][propertyName] {
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, propertyName),
						Update: map[string]interface{}{
							"x-speakeasy-ignore": true,
						},
					})
					ignoreTracker[resource.EntityName][propertyName] = true
					fmt.Printf("✅ Added manual property ignore for data source: %s.%s\n", resource.EntityName, propertyName)
				}
			}
		}

		applyAdditionalPropertiesForDataSource(overlay, resource, additionalPropsMappings, additionalPropsTracker, entityReadonlyTracker)

		applyReadonlyForDataSource(overlay, resource, schemas, entityReadonlyTracker)

		for crudType, opInfo := range resource.Operations {
			if crudType != "read" && crudType != "list" {
				continue
			}

			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) || !entityConfig.ShouldProcessEntity(resource.EntityName) {
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

			if resource.PrimaryID != "" && crudType == "read" {
				applyParameterMatchingForDataSource(overlay, resource, opInfo, schemas, manualMappings)
			}
		}
	}

	fmt.Printf("Applied data source extensions to %d non-viable resources\n", len(viableDatasources))

	fmt.Printf("\n=== Terraform Reserved Words Auto-Ignore ===\n")

	reservedWordCount := applyTerraformReservedWordIgnores(overlay, viableResources, viableDatasources, schemas, ignoreTracker)

	if reservedWordCount > 0 {
		fmt.Printf("✅ Auto-ignored %d Terraform reserved word properties\n", reservedWordCount)
	} else {
		fmt.Printf("✅ No Terraform reserved words found in entity properties\n")
	}

	manualIgnoreCount := 0
	for schemaName, properties := range manualPropertyIgnores {
		for _, propertyName := range properties {
			// Initialize ignore tracker if needed
			if ignoreTracker[schemaName] == nil {
				ignoreTracker[schemaName] = make(map[string]bool)
			}

			// Only add if not already ignored (by reserved word logid or previous ignore)
			if !ignoreTracker[schemaName][propertyName] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", schemaName, propertyName),
					Update: map[string]interface{}{
						"x-speakeasy-ignore": true,
					},
				})
				ignoreTracker[schemaName][propertyName] = true
				manualIgnoreCount++
				fmt.Printf("✅ Added manual property ignore: %s.%s\n", schemaName, propertyName)
			}
		}
	}
	fmt.Println("Total manual property ignores applied:", manualIgnoreCount)

	fmt.Printf("\n=== Overlay Generation Complete ===\n")
	if len(entityConfig.EnabledEntities) > 0 {
		fmt.Printf("Generated %d actions for %d enabled viable resources and %d enabled data sources\n",
			len(overlay.Actions), len(viableResources), len(viableDatasources))
		fmt.Printf("Disabled %d entities due to explicit enable configuration\n", len(disabledResources))
	} else {
		fmt.Printf("Generated %d actions for %d viable resources and %d data sources (all entities enabled by default)\n",
			len(overlay.Actions), len(viableResources), len(viableDatasources))
	}

	totalIgnores := 0
	totalMatches := 0
	totalReadonly := 0
	totalAdditionalProps := 0
	totalNameOverrides := 0

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
		if _, hasOverride := action.Update["x-speakeasy-name-override"]; hasOverride {
			totalNameOverrides++
		}
	}

	if totalAdditionalProps > 0 {
		fmt.Printf("✅ %d additionalProperties actions added for flexible schema fields\n", totalAdditionalProps)
	}
	if totalNameOverrides > 0 {
		fmt.Printf("✅ %d speakeasy name override actions added for type mismatches\n", totalNameOverrides)
	}
	if totalIgnores > 0 {
		fmt.Printf("✅ %d speakeasy ignore actions added\n", totalIgnores)
	}
	if totalMatches > 0 {
		fmt.Printf("✅ %d speakeasy match actions added for primary ID parameters\n", totalMatches)
	}
	if totalReadonly > 0 {
		fmt.Printf("✅ %d speakeasy readonly actions added (includes computed fields and additional properties)\n", totalReadonly)
	}

	return overlay
}

func applyAdditionalPropertiesForDataSource(overlay *Overlay, resource *ResourceInfo,
	additionalPropsMappings map[string][]string, additionalPropsTracker map[string]map[string]bool,
	entityReadonlyTracker map[string]map[string]bool) {

	if properties, exists := additionalPropsMappings[resource.EntityName]; exists {
		if additionalPropsTracker[resource.EntityName] == nil {
			additionalPropsTracker[resource.EntityName] = make(map[string]bool)
		}

		for _, propertyPath := range properties {
			// Skip if already processed
			if additionalPropsTracker[resource.EntityName][propertyPath] {
				continue
			}

			targetPath := buildAdditionalPropertiesPath(resource.EntityName, propertyPath)

			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: targetPath,
				Update: map[string]interface{}{
					"additionalProperties": true,
				},
			})

			// Mark as readonly (data sources are always readonly)
			if entityReadonlyTracker[resource.EntityName] == nil {
				entityReadonlyTracker[resource.EntityName] = make(map[string]bool)
			}

			if !entityReadonlyTracker[resource.EntityName][propertyPath] {
				overlay.Actions = append(overlay.Actions, OverlayAction{
					Target: targetPath,
					Update: map[string]interface{}{
						"x-speakeasy-param-readonly": true,
					},
				})

				propParts := strings.Split(propertyPath, ".")
				topLevelProp := propParts[0]
				entityReadonlyTracker[resource.EntityName][propertyPath] = true
				entityReadonlyTracker[resource.EntityName][topLevelProp] = true
			}

			additionalPropsTracker[resource.EntityName][propertyPath] = true
		}
	}
}

// Apply readonly properties for data source (all properties should be readonly)
func applyReadonlyForDataSource(overlay *Overlay, resource *ResourceInfo, schemas map[string]interface{},
	entityReadonlyTracker map[string]map[string]bool) {

	entitySchema, _ := schemas[resource.EntityName].(map[string]interface{})
	if entitySchema == nil {
		return
	}

	// For data sources, we can be more aggressive about marking properties as readonly
	// since they're never used for input
	entityProps, _ := entitySchema["properties"].(map[string]interface{})
	if entityProps == nil {
		return
	}

	if entityReadonlyTracker[resource.EntityName] == nil {
		entityReadonlyTracker[resource.EntityName] = make(map[string]bool)
	}

	for propName := range entityProps {
		// Skip if already marked readonly
		if entityReadonlyTracker[resource.EntityName][propName] {
			continue
		}

		// Skip basic identifier fields
		if propName == "id" || propName == "slug" {
			continue
		}

		// Mark as readonly
		overlay.Actions = append(overlay.Actions, OverlayAction{
			Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, propName),
			Update: map[string]interface{}{
				"x-speakeasy-param-readonly": true,
			},
		})

		entityReadonlyTracker[resource.EntityName][propName] = true
	}
}

// Apply parameter matching for data source read operations
func applyParameterMatchingForDataSource(overlay *Overlay, resource *ResourceInfo, opInfo OperationInfo,
	schemas map[string]interface{}, manualMappings *ManualMappings) {

	pathParams := extractPathParameters(opInfo.Path)
	for _, param := range pathParams {
		if param == resource.PrimaryID {
			// Check for manual mappings first
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
				// Apply automatic matching logic
				var targetField string

				// Don't create circular mappings
				if param == "id" || param == "slug" {
					continue
				}

				// Get entity properties to check what fields exist
				entityProps := getSchemaProperties(schemas, resource.EntityName)
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
					fmt.Printf("    Warning: Cannot map parameter %s for data source - entity has neither id nor slug field\n", param)
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

func handleCRUDInconsistency(overlay *Overlay, resource *ResourceInfo, inconsistency CRUDInconsistency, entityReadonlyTracker map[string]map[string]bool) {
	entityHas := strings.Contains(inconsistency.Description, "Entity:true")
	createHas := strings.Contains(inconsistency.Description, "Create:true")
	updateHas := strings.Contains(inconsistency.Description, "Update:true")

	// If a property exists only in Entity (not in any request), we mark it as read-only
	if entityHas && !createHas && !updateHas {
		// This is a read-only

		if entityReadonlyTracker[resource.EntityName] == nil {
			entityReadonlyTracker[resource.EntityName] = make(map[string]bool)
		}

		if !entityReadonlyTracker[resource.EntityName][inconsistency.PropertyName] {
			overlay.Actions = append(overlay.Actions, OverlayAction{
				Target: fmt.Sprintf("$.components.schemas.%s.properties.%s",
					resource.EntityName, inconsistency.PropertyName),
				Update: map[string]interface{}{
					"x-speakeasy-param-readonly": true,
				},
			})
			entityReadonlyTracker[resource.EntityName][inconsistency.PropertyName] = true
			fmt.Printf("    Marked Entity field as readonly: %s.%s (only in response)\n",
				resource.EntityName, inconsistency.PropertyName)
		}
		return
	}
}
