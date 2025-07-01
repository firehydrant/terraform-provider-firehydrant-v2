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
		return report
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		fmt.Println("Warning: No schemas found in components")
		return report
	}

	paths, pathsOk := spec["paths"].(map[string]interface{})
	if !pathsOk {
		fmt.Println("Warning: No paths found in spec")
	}

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
		if detail.ConflictType == "map-class" {
			report.MapClassFixes++
		} else {
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
