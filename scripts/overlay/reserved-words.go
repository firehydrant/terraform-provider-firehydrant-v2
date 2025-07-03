package main

import (
	"fmt"
	"strings"
)

func applyTerraformReservedWordIgnores(overlay *Overlay, resources map[string]*ResourceInfo, datasources map[string]*ResourceInfo, schemas map[string]interface{}, ignoreTracker map[string]map[string]bool) int {
	// Comprehensive list of Terraform reserved root attribute names
	// Most recent information found at time of writing is from: https://medium.com/@d3vpasha/reserved-keywords-in-terraform-f37a4cbf3a81
	// We removed data from this list as we can successfully compile without ignoring
	terraformReservedWords := []string{
		"count", "connection", "for_each", "lifecycle", "depends_on",
		"provider", "provisioner", "locals", "resource",
		"variable", "output", "module", "terraform", "import",
		"moved", "removed", "check", "precondition", "postcondition",
	}

	reservedWordCount := 0

	// Combine all resources and data sources for processing
	allResources := make(map[string]*ResourceInfo)
	for k, v := range resources {
		allResources[k] = v
	}
	for k, v := range datasources {
		allResources[k] = v
	}

	for _, resource := range allResources {
		entitySchema, exists := schemas[resource.EntityName].(map[string]interface{})
		if !exists {
			continue
		}

		entityProps, _ := entitySchema["properties"].(map[string]interface{})
		if entityProps == nil {
			continue
		}

		if ignoreTracker[resource.EntityName] == nil {
			ignoreTracker[resource.EntityName] = make(map[string]bool)
		}

		for propName := range entityProps {
			if isReservedTerraformWord(propName, terraformReservedWords) {
				// Only add if not already ignored
				if !ignoreTracker[resource.EntityName][propName] {
					overlay.Actions = append(overlay.Actions, OverlayAction{
						Target: fmt.Sprintf("$.components.schemas.%s.properties.%s", resource.EntityName, propName),
						Update: map[string]interface{}{
							"x-speakeasy-ignore": true,
						},
					})
					ignoreTracker[resource.EntityName][propName] = true
					reservedWordCount++
					fmt.Printf("🚫 Auto-ignored Terraform reserved word: %s.%s\n", resource.EntityName, propName)
				}
			}
		}
	}

	return reservedWordCount
}

func isReservedTerraformWord(propName string, reservedWords []string) bool {
	lowerProp := strings.ToLower(propName)
	for _, reserved := range reservedWords {
		if lowerProp == strings.ToLower(reserved) {
			return true
		}
	}
	return false
}
