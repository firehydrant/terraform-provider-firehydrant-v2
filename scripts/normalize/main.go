package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := inputPath
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	fmt.Printf("=== OpenAPI Schema Normalizer ===\n")
	fmt.Printf("Input: %s\n", inputPath)
	fmt.Printf("Output: %s\n\n", outputPath)

	specData, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading spec: %v\n", err)
		os.Exit(1)
	}

	var spec map[string]interface{}
	if err := json.Unmarshal(specData, &spec); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	report := normalizeSpec(spec)

	printNormalizationReport(report)

	normalizedData, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling normalized spec: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outputPath, normalizedData, 0644); err != nil {
		fmt.Printf("Error writing normalized spec: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Successfully normalized OpenAPI spec\n")
	fmt.Printf("   Total fixes applied: %d\n", report.TotalFixes)
}

func printUsage() {
	fmt.Println("OpenAPI Schema Normalizer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  openapi-normalize <input.json> [output.json]")
}

type NormalizationReport struct {
	TotalFixes      int
	MapClassFixes   int
	PropertyFixes   int
	ConflictDetails []ConflictDetail
}

type ConflictDetail struct {
	Schema       string
	Property     string
	ConflictType string
	Resolution   string
}

func printNormalizationReport(report NormalizationReport) {
	fmt.Println("\n=== Normalization Report ===")
	fmt.Printf("Total fixes applied: %d\n", report.TotalFixes)

	mapClassFixes := 0
	parameterFixes := 0
	enumFixes := 0
	otherFixes := 0

	for _, detail := range report.ConflictDetails {
		switch detail.ConflictType {
		case "map-class":
			mapClassFixes++
		case "parameter-type":
			parameterFixes++
		case "enum-normalization":
			enumFixes++
		default:
			otherFixes++
		}
	}

	fmt.Printf("Map/Class fixes: %d\n", mapClassFixes)
	fmt.Printf("Parameter type fixes: %d\n", parameterFixes)
	fmt.Printf("Enum normalization fixes: %d\n", enumFixes)
	fmt.Printf("Other fixes: %d\n", otherFixes)

	// Helpful for debugging
	// if len(report.ConflictDetails) > 0 {
	// 	fmt.Println("\nDetailed fixes:")
	// 	for _, detail := range report.ConflictDetails {
	// 		fmt.Printf("  - %s.%s [%s]: %s\n",
	// 			detail.Schema, detail.Property, detail.ConflictType, detail.Resolution)
	// 	}
	// }
}
