package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// Manual mapping configuration
type ManualMapping struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Action string `yaml:"action"` // "ignore", "entity", "match"
	Value  string `yaml:"value,omitempty"`
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

	fmt.Printf("Loaded %d manual mappings from %s\n", len(mappings.Operations), mappingsPath)
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

		// Copy operations that aren't manually ignored
		for crudType, opInfo := range resource.Operations {
			if shouldIgnoreOperation(opInfo.Path, opInfo.Method, manualMappings) {
				fmt.Printf("  Removing manually ignored operation: %s %s (was %s for %s)\n",
					opInfo.Method, opInfo.Path, crudType, resource.EntityName)
				operationsRemoved++
			} else {
				cleanedResource.Operations[crudType] = opInfo
			}
		}

		// Only include resource if it still has operations after cleaning
		if len(cleanedResource.Operations) > 0 {
			cleanedResources[name] = cleanedResource
			if operationsRemoved > 0 {
				fmt.Printf("  Resource %s: kept %d operations, removed %d manually ignored\n",
					name, len(cleanedResource.Operations), operationsRemoved)
			}
		} else {
			fmt.Printf("  Resource %s: removed entirely (all operations were manually ignored)\n", name)
		}
	}

	fmt.Printf("Manual mapping cleanup: %d → %d resources\n", len(resources), len(cleanedResources))
	return cleanedResources
}

func getManualParameterMatch(path, method, paramName string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.ToLower(mapping.Method) == strings.ToLower(method) && mapping.Action == "match" {
			// For match mappings, we expect the value to be in format "param_name:field_name"
			parts := strings.SplitN(mapping.Value, ":", 2)
			if len(parts) == 2 && parts[0] == paramName {
				fmt.Printf("    Manual mapping: Parameter %s in %s %s -> %s\n", paramName, method, path, parts[1])
				return parts[1], true
			}
		}
	}
	return "", false
}

func shouldIgnoreOperation(path, method string, manualMappings *ManualMappings) bool {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == "ignore" {
			fmt.Printf("    Manual mapping: Ignoring operation %s %s\n", method, path)
			return true
		}
	}
	return false
}

func getManualEntityMapping(path, method string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == "entity" {
			fmt.Printf("    Manual mapping: Operation %s %s -> Entity %s\n", method, path, mapping.Value)
			return mapping.Value, true
		}
	}
	return "", false
}
