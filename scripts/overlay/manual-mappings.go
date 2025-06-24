package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// Manual mapping configuration, corresponds with the manual-mappings.yaml file
type ManualMapping struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Action string `yaml:"action"` // "ignore", "entity", "match"
	Value  string `yaml:"value,omitempty"`

	// For entity mappings
	Schema   string `yaml:"schema,omitempty"`
	Property string `yaml:"property,omitempty"`
}

type ManualMappings struct {
	Operations []ManualMapping `yaml:"operations"`
}

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

	return &mappings
}

func applyManualMappings(resources map[string]*ResourceInfo, manualMappings *ManualMappings) map[string]*ResourceInfo {
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

		for crudType, opInfo := range resource.Operations {
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) {
				operationsRemoved++
			} else {
				cleanedResource.Operations[crudType] = opInfo
			}
		}

		if len(cleanedResource.Operations) > 0 {
			cleanedResources[name] = cleanedResource
			if operationsRemoved > 0 {
			}
		}
	}

	fmt.Printf("Manual mapping cleanup: %d → %d resources\n", len(resources), len(cleanedResources))
	return cleanedResources
}

func getManualParameterMatch(path, method, paramName string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == "match" {
			parts := strings.SplitN(mapping.Value, ":", 2)
			if len(parts) == 2 && parts[0] == paramName {
				return parts[1], true
			}
		}
	}
	return "", false
}

func shouldIgnoreOperation(path, method string, manualMappings *ManualMappings) bool {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == "ignore" {
			return true
		}
	}
	return false
}

func getManualEntityMapping(path, method string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == "entity" {
			return mapping.Value, true
		}
	}
	return "", false
}

func getManualPropertyIgnores(manualMappings *ManualMappings) map[string][]string {
	ignores := make(map[string][]string)

	for _, mapping := range manualMappings.Operations {
		if mapping.Action == "ignore_property" && mapping.Schema != "" && mapping.Property != "" {
			fmt.Printf("    Manual property ignore: %s.%s\n", mapping.Schema, mapping.Property)
			ignores[mapping.Schema] = append(ignores[mapping.Schema], mapping.Property)
		}
	}

	return ignores
}

// Add this function in the same file as your other manual mapping functions
func getAdditionalPropertiesMappings(manualMappings *ManualMappings) map[string][]string {
	additionalProps := make(map[string][]string)

	for _, mapping := range manualMappings.Operations {
		if mapping.Action == "additional_properties" && mapping.Schema != "" && mapping.Property != "" {
			additionalProps[mapping.Schema] = append(additionalProps[mapping.Schema], mapping.Property)
		}
	}

	return additionalProps
}

// Add this helper function for building paths
func buildAdditionalPropertiesPath(schemaName string, propertyPath string) string {
	parts := strings.Split(propertyPath, ".")
	path := fmt.Sprintf("$.components.schemas.%s", schemaName)

	for _, part := range parts {
		if part == "items" {
			// For array items, we need to set additionalProperties on the items object
			path += ".items"
		} else {
			path += fmt.Sprintf(".properties.%s", part)
		}
	}

	return path
}
