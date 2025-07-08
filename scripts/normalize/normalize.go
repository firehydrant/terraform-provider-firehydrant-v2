package main

import (
	"fmt"
)

func normalizeSpec(spec map[string]interface{}) NormalizationReport {
	report := NormalizationReport{
		ConflictDetails: make([]ConflictDetail, 0),
	}

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No components found in spec")
		// Create components if it doesn't exist
		components = map[string]interface{}{}
		spec["components"] = components
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No schemas found in components")
		// Create schemas if it doesn't exist
		schemas = map[string]interface{}{}
		components["schemas"] = schemas
	}

	paths, pathsOk := spec["paths"].(map[string]interface{})
	if !pathsOk {
		fmt.Println("Warning: No paths found in spec")
	}

	// We need to normalize request schemas first to ensure that extracted schemas are available for other normalization steps
	if pathsOk {
		requestSchemaFixes := normalizeRequestSchemasWithPaths(paths, schemas)
		report.ConflictDetails = append(report.ConflictDetails, requestSchemaFixes...)
	}

	terraformFixes := normalizeTerraformKeywords(schemas)
	report.ConflictDetails = append(report.ConflictDetails, terraformFixes...)

	// Only apply non-structural normalizations
	globalFixes := applyGlobalNormalizations(schemas)
	report.ConflictDetails = append(report.ConflictDetails, globalFixes...)

	enumFixes := normalizeEnums(schemas)
	report.ConflictDetails = append(report.ConflictDetails, enumFixes...)

	if pathsOk {
		parameterFixes := normalizePathParameters(paths)
		report.ConflictDetails = append(report.ConflictDetails, parameterFixes...)

		pathEnumFixes := normalizePathEnums(paths)
		report.ConflictDetails = append(report.ConflictDetails, pathEnumFixes...)
	}

	report.TotalFixes = len(report.ConflictDetails)
	for _, detail := range report.ConflictDetails {
		switch detail.ConflictType {
		case "map-class":
			report.MapClassFixes++
		case "terraform-keyword":
			report.TerraformKeywordFixes++
		case "request-schema-extraction":
			report.PropertyFixes++
		default:
			report.PropertyFixes++
		}
	}

	return report
}

func applyGlobalNormalizations(schemas map[string]interface{}) []ConflictDetail {
	conflicts := make([]ConflictDetail, 0)

	fmt.Printf("Applying global normalizations to %d schemas\n", len(schemas))

	for schemaName, schema := range schemas {
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			continue
		}

		// Only normalize additionalProperties (map vs class issue)
		schemaConflicts := normalizeAdditionalProperties(schemaName, schemaMap, "")
		conflicts = append(conflicts, schemaConflicts...)
	}

	return conflicts
}
