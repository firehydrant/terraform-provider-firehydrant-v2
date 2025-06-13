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
		if mapping.Path == path && strings.ToLower(mapping.Method) == strings.ToLower(method) && mapping.Action == "match" {
			// For match mappings, we expect the value to be in format "param_name:field_name"
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
