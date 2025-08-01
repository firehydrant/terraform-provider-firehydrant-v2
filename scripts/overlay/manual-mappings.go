package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

type ManualMappingAction string

const (
	Ignore               ManualMappingAction = "ignore"
	Entity               ManualMappingAction = "entity"
	Match                ManualMappingAction = "match"
	Enable               ManualMappingAction = "enable"
	IgnoreProperty       ManualMappingAction = "ignore_property"
	AdditionalProperties ManualMappingAction = "additional_properties"
)

// Manual mapping configuration, corresponds with the manual-mappings.yaml file
type ManualMapping struct {
	Path   string              `yaml:"path"`
	Method string              `yaml:"method"`
	Action ManualMappingAction `yaml:"action"`
	Value  string              `yaml:"value,omitempty"`

	// For entity mappings
	Schema   string `yaml:"schema,omitempty"`
	Property string `yaml:"property,omitempty"`

	// For entity enabling
	Entity string `yaml:"entity,omitempty"`
}

type EntityConfig struct {
	EnabledEntities    map[string]bool
	HasExplicitEnabled bool
}

type ManualMappings struct {
	Operations []ManualMapping `yaml:"operations"`
}

func buildEntityConfig(manualMappings *ManualMappings) *EntityConfig {
	enabled := make(map[string]bool)
	hasExplicit := false

	for _, mapping := range manualMappings.Operations {
		if mapping.Action == Enable && mapping.Entity != "" {
			enabled[mapping.Entity] = true
			hasExplicit = true
		}
	}

	HasExplicitEnabled := (len(enabled) == 0 && !hasExplicit)

	return &EntityConfig{
		EnabledEntities:    enabled,
		HasExplicitEnabled: HasExplicitEnabled,
	}
}

func (ec *EntityConfig) ShouldProcessEntity(entityName string) bool {
	if ec.HasExplicitEnabled {
		return true
	}
	return ec.EnabledEntities[entityName]
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

func getManualParameterMatch(path, method, paramName string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == Match {
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
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == Ignore {
			fmt.Println("  Ignoring operation:", method, path)
			return true
		}
	}
	return false
}

func getManualEntityMapping(path, method string, manualMappings *ManualMappings) (string, bool) {
	for _, mapping := range manualMappings.Operations {
		if mapping.Path == path && strings.EqualFold(mapping.Method, method) && mapping.Action == Entity {
			return mapping.Value, true
		}
	}
	return "", false
}

func getManualPropertyIgnores(manualMappings *ManualMappings) map[string][]string {
	ignores := make(map[string][]string)

	for _, mapping := range manualMappings.Operations {
		if mapping.Action == IgnoreProperty && mapping.Schema != "" && mapping.Property != "" {
			ignores[mapping.Schema] = append(ignores[mapping.Schema], mapping.Property)
		}
	}

	return ignores
}

// Add this function in the same file as your other manual mapping functions
func getAdditionalPropertiesMappings(manualMappings *ManualMappings) map[string][]string {
	additionalProps := make(map[string][]string)

	for _, mapping := range manualMappings.Operations {
		if mapping.Action == AdditionalProperties && mapping.Schema != "" && mapping.Property != "" {
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
