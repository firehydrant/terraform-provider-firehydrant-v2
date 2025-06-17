package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type OpenAPISpec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]PathItem    `json:"paths"`
	Components Components             `json:"components"`
}

type Components struct {
	Schemas         map[string]Schema      `json:"schemas"`
	SecuritySchemes map[string]interface{} `json:"securitySchemes,omitempty"`
}

type Schema struct {
	Type       string                 `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
	AllOf      []interface{}          `json:"allOf,omitempty"`
	Nullable   bool                   `json:"nullable,omitempty"`
	Items      interface{}            `json:"items,omitempty"`
	Raw        map[string]interface{} `json:"-"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
}

type Operation struct {
	OperationID string                 `json:"operationId,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Parameters  []Parameter            `json:"parameters,omitempty"`
	RequestBody map[string]interface{} `json:"requestBody,omitempty"`
	Responses   map[string]interface{} `json:"responses,omitempty"`
}

type Parameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required,omitempty"`
	Schema   Schema `json:"schema,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	specPath := os.Args[1]
	var mappingsPath string
	if len(os.Args) > 2 {
		mappingsPath = os.Args[2]
	} else {
		mappingsPath = "manual-mappings.yaml"
	}

	fmt.Printf("=== Terraform Overlay Generator ===\n")
	fmt.Printf("Input: %s\n", specPath)

	manualMappings := loadManualMappings(mappingsPath)

	specData, err := ioutil.ReadFile(specPath)
	if err != nil {
		fmt.Printf("Error reading spec file: %v\n", err)
		os.Exit(1)
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(specData, &spec); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d paths and %d schemas\n\n", len(spec.Paths), len(spec.Components.Schemas))

	resources := analyzeSpec(spec, manualMappings)

	overlay := generateOverlay(resources, spec, manualMappings)

	if err := writeOverlay(overlay); err != nil {
		fmt.Printf("Error writing overlay: %v\n", err)
		os.Exit(1)
	}

	printOverlaySummary(overlay)
}

func printUsage() {
	fmt.Println("OpenAPI Terraform Overlay Generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-overlay <input.json>")
}

func (s *Schema) UnmarshalJSON(data []byte) error {
	type Alias Schema
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Also unmarshal into raw map
	if err := json.Unmarshal(data, &s.Raw); err != nil {
		return err
	}

	return nil
}

func writeOverlay(overlay *Overlay) error {
	data, err := yaml.Marshal(overlay)
	if err != nil {
		return fmt.Errorf("marshaling overlay: %w", err)
	}

	overlayPath := "terraform-overlay.yaml"
	if err := ioutil.WriteFile(overlayPath, data, 0644); err != nil {
		return fmt.Errorf("writing overlay file: %w", err)
	}

	fmt.Printf("Overlay written to: %s\n", overlayPath)
	return nil
}

func printOverlaySummary(overlay *Overlay) {
	fmt.Println("\n=== Summary ===")
	fmt.Printf("✅ Successfully generated overlay with %d actions\n", len(overlay.Actions))

	fmt.Println("\nOverlay approach:")
	fmt.Println("1. Load manual mappings for edge cases")
	fmt.Println("2. Identify entity schemas and match operations to entities")
	fmt.Println("3. Apply manual ignore/entity/match mappings during analysis")
	fmt.Println("4. Clean resources by removing manually ignored operations")
	fmt.Println("5. Analyze ID patterns and choose consistent primary ID per entity")
	fmt.Println("6. Filter operations with unmappable path parameters")
	fmt.Println("7. Skip annotations for non-viable resources")
	fmt.Println("8. Mark viable entity schemas with x-speakeasy-entity")
	fmt.Println("9. Tag viable operations with x-speakeasy-entity-operation")
	fmt.Println("10. Mark chosen primary ID with x-speakeasy-match")
	fmt.Println("11. Apply x-speakeasy-ignore: true to problematic properties")
}
