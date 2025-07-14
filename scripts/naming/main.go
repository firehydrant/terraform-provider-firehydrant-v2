package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type ConflictDetail struct {
	Schema       string
	Property     string
	ConflictType string
	Resolution   string
}

type NormalizationReport struct {
	TotalFixes   int
	SchemaFixes  int
	OverlayFixes int
}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	specPath := os.Args[1]
	overlayPath := os.Args[2]

	fmt.Printf("=== Post-Overlay Naming Normalizer ===\n")
	fmt.Printf("OpenAPI Spec: %s\n", specPath)
	fmt.Printf("Overlay File: %s\n", overlayPath)

	report := NormalizationReport{}

	// Step 1: Normalize the OpenAPI spec
	schemaRenames, err := normalizeOpenAPISpec(specPath)
	if err != nil {
		fmt.Printf("Error normalizing OpenAPI spec: %v\n", err)
		os.Exit(1)
	}
	report.SchemaFixes = len(schemaRenames)

	// Step 2: Normalize the overlay file using the schema renames
	overlayFixes, err := normalizeOverlayFile(overlayPath, schemaRenames)
	if err != nil {
		fmt.Printf("Error normalizing overlay file: %v\n", err)
		os.Exit(1)
	}
	report.OverlayFixes = overlayFixes

	report.TotalFixes = report.SchemaFixes + report.OverlayFixes

	// Print summary
	fmt.Printf("\n=== Naming Normalization Complete ===\n")
	fmt.Printf("✅ Schema renames: %d\n", report.SchemaFixes)
	fmt.Printf("✅ Overlay updates: %d\n", report.OverlayFixes)
	fmt.Printf("✅ Total fixes: %d\n", report.TotalFixes)
}

func printUsage() {
	fmt.Println("Post-Overlay Naming Normalizer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  naming <openapi-spec.json> <overlay.yaml>")
	fmt.Println()
	fmt.Println("This tool normalizes entity naming in both the OpenAPI spec and overlay file.")
	fmt.Println("It should be run AFTER overlay generation to fix entity naming issues.")
}

func normalizeOpenAPISpec(specPath string) (map[string]string, error) {
	fmt.Printf("\n=== Normalizing OpenAPI Schema Names ===\n")

	specData, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("reading spec: %w", err)
	}

	var spec map[string]interface{}
	if err := json.Unmarshal(specData, &spec); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	allEntityRenames := findAllEntityPatternsInSpec(spec)
	if len(allEntityRenames) == 0 {
		fmt.Printf("No entity references found that need normalization\n")
		return allEntityRenames, nil
	}

	fmt.Printf("Found %d entity references to normalize throughout the spec\n", len(allEntityRenames))

	if components, ok := spec["components"].(map[string]interface{}); ok {
		if schemas, ok := components["schemas"].(map[string]interface{}); ok {
			renames := identifySchemasToRename(schemas)

			renames = ensureUniqueRenamedSchemas(renames, schemas)

			for _, rename := range renames {
				if schema, exists := schemas[rename.OldName]; exists {
					schemas[rename.NewName] = schema
					delete(schemas, rename.OldName)
				}
			}
		}
	}

	updateReferences(spec, allEntityRenames)
	updateDescriptions(spec, allEntityRenames)

	normalizedData, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshaling normalized spec: %w", err)
	}

	if err := ioutil.WriteFile(specPath, normalizedData, 0644); err != nil {
		return nil, fmt.Errorf("writing normalized spec: %w", err)
	}

	fmt.Printf("Successfully normalized %d entity references\n", len(allEntityRenames))
	return allEntityRenames, nil
}

func normalizeOverlayFile(overlayPath string, schemaRenames map[string]string) (int, error) {
	fmt.Printf("\n=== Normalizing Overlay File ===\n")

	overlayData, err := ioutil.ReadFile(overlayPath)
	if err != nil {
		return 0, fmt.Errorf("reading overlay: %w", err)
	}

	var overlay map[string]interface{}
	if err := yaml.Unmarshal(overlayData, &overlay); err != nil {
		return 0, fmt.Errorf("parsing YAML: %w", err)
	}

	fixCount := 0

	if actions, ok := overlay["actions"].([]interface{}); ok {
		for _, actionInterface := range actions {
			if action, ok := actionInterface.(map[string]interface{}); ok {
				if target, ok := action["target"].(string); ok {
					newTarget := normalizeOverlayTarget(target, schemaRenames)
					if newTarget != target {
						action["target"] = newTarget
						fixCount++
					}
				}

				if update, ok := action["update"].(map[string]interface{}); ok {
					updateFixes := normalizeOverlayUpdate(update)
					fixCount += updateFixes
				}
			}
		}
	}

	if fixCount == 0 {
		fmt.Printf("No overlay updates needed\n")
		return 0, nil
	}

	normalizedOverlay := createOrderedOverlayYAML(overlay)

	if err := ioutil.WriteFile(overlayPath, []byte(normalizedOverlay), 0644); err != nil {
		return 0, fmt.Errorf("writing normalized overlay: %w", err)
	}

	fmt.Printf("Successfully updated %d overlay references\n", fixCount)
	return fixCount, nil
}

func createOrderedOverlayYAML(overlay map[string]interface{}) string {
	var result strings.Builder

	if overlayVersion, ok := overlay["overlay"]; ok {
		result.WriteString(fmt.Sprintf("overlay: %v\n", overlayVersion))
	}

	if info, ok := overlay["info"].(map[string]interface{}); ok {
		result.WriteString("info:\n")
		if title, ok := info["title"]; ok {
			result.WriteString(fmt.Sprintf("  title: %s\n", title))
		}
		if version, ok := info["version"]; ok {
			result.WriteString(fmt.Sprintf("  version: %s\n", version))
		}
		if description, ok := info["description"]; ok {
			result.WriteString(fmt.Sprintf("  description: %s\n", description))
		}
	}

	if actions, ok := overlay["actions"]; ok {
		actionsYAML, err := yaml.Marshal(map[string]interface{}{"actions": actions})
		if err == nil {
			result.WriteString(string(actionsYAML))
		}
	}

	return result.String()
}

func normalizeOverlayTarget(target string, schemaRenames map[string]string) string {
	schemaPattern := regexp.MustCompile(`\$\.components\.schemas\.([^.]+)(.*)`)

	return schemaPattern.ReplaceAllStringFunc(target, func(match string) string {
		matches := schemaPattern.FindStringSubmatch(match)
		if len(matches) >= 2 {
			schemaName := matches[1]
			remainder := ""
			if len(matches) > 2 {
				remainder = matches[2]
			}

			if isIdentityOnlySchema(schemaName) {
				return match
			}

			if newName, exists := schemaRenames[schemaName]; exists {
				return "$.components.schemas." + newName + remainder
			}

			normalizedName := normalizeSchemaName(schemaName)
			if normalizedName != schemaName {
				return "$.components.schemas." + normalizedName + remainder
			}
		}
		return match
	})
}

func normalizeOverlayUpdate(update map[string]interface{}) int {
	fixCount := 0

	if entity, ok := update["x-speakeasy-entity"].(string); ok {
		normalizedEntity := normalizeEntityName(entity)
		if normalizedEntity != entity {
			update["x-speakeasy-entity"] = normalizedEntity
			fixCount++
		}
	}

	if entityOp, ok := update["x-speakeasy-entity-operation"].(string); ok {
		normalizedEntityOp := normalizeEntityOperation(entityOp)
		if normalizedEntityOp != entityOp {
			update["x-speakeasy-entity-operation"] = normalizedEntityOp
			fixCount++
		}
	}

	return fixCount
}

func normalizeEntityName(name string) string {
	if isIdentityOnlySchema(name) {
		return name
	}
	return normalizeSchemaName(name)
}

func normalizeEntityOperation(entityOp string) string {
	parts := strings.Split(entityOp, "#")
	if len(parts) == 2 {
		entityName := parts[0]
		operation := parts[1]

		normalizedEntity := normalizeEntityName(entityName)
		return normalizedEntity + "#" + operation
	}
	return entityOp
}

func findAllEntityPatternsInSpec(spec map[string]interface{}) map[string]string {
	allRenames := make(map[string]string)

	if components, ok := spec["components"].(map[string]interface{}); ok {
		if schemas, ok := components["schemas"].(map[string]interface{}); ok {
			for schemaName := range schemas {
				if shouldNormalizeSchemaName(schemaName) {
					normalizedName := normalizeSchemaName(schemaName)
					if normalizedName != schemaName && normalizedName != "" {
						allRenames[schemaName] = normalizedName
					}
				}
			}
		}
	}

	findEntityReferencesRecursive(spec, allRenames)

	return allRenames
}

func findEntityReferencesRecursive(obj interface{}, allRenames map[string]string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		for _, value := range v {
			switch val := value.(type) {
			case string:
				if shouldNormalizeString(val) {
					normalizedValue := normalizeSchemaName(val)
					if normalizedValue != val && normalizedValue != "" {
						allRenames[val] = normalizedValue
					}
				}
			}
			findEntityReferencesRecursive(value, allRenames)
		}
	case []interface{}:
		for _, item := range v {
			findEntityReferencesRecursive(item, allRenames)
		}
	}
}

func shouldNormalizeSchemaName(name string) bool {
	if isIdentityOnlySchema(name) {
		return false
	}

	return containsEntityPattern(name)
}

func containsEntityPattern(s string) bool {
	return strings.Contains(s, "Entity") ||
		strings.Contains(s, "Entities") ||
		strings.Contains(s, "entity") ||
		strings.Contains(s, "entities")
}

func shouldNormalizeString(s string) bool {
	if isIdentityOnlySchema(s) {
		return false
	}

	if strings.Contains(s, " ") ||
		strings.Contains(s, "http") ||
		strings.Contains(s, "://") ||
		strings.Contains(s, "/") ||
		len(s) < 3 {
		return false
	}

	return containsEntityPattern(s) && (isCapitalized(s) || strings.Contains(s, "_"))
}

func isCapitalized(s string) bool {
	return len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z'
}

func isIdentityOnlySchema(name string) bool {
	lowerName := strings.ToLower(name)

	if !strings.Contains(lowerName, "identity") {
		return false
	}

	hasEntityPattern := false

	if strings.Contains(lowerName, "entity") {
		withoutIdentity := strings.ReplaceAll(lowerName, "identity", "")
		if strings.Contains(withoutIdentity, "entity") || strings.Contains(withoutIdentity, "entities") {
			hasEntityPattern = true
		}
	}

	return !hasEntityPattern
}

type SchemaRenameInfo struct {
	OldName string
	NewName string
}

func identifySchemasToRename(schemas map[string]interface{}) []SchemaRenameInfo {
	var renames []SchemaRenameInfo

	for schemaName := range schemas {
		if shouldNormalizeSchemaName(schemaName) {
			newName := normalizeSchemaName(schemaName)
			if newName != schemaName && newName != "" {
				renames = append(renames, SchemaRenameInfo{
					OldName: schemaName,
					NewName: newName,
				})
			}
		}
	}

	return renames
}

func normalizeSchemaName(name string) string {
	if isIdentityOnlySchema(name) {
		return name
	}

	original := name

	if strings.HasPrefix(name, "Nullable") {
		inner := strings.TrimPrefix(name, "Nullable")
		normalizedInner := normalizeEntityPart(inner)
		if normalizedInner != "" {
			return "Nullable" + normalizedInner
		}
		return name
	}

	if strings.HasSuffix(name, "Paginated") {
		inner := strings.TrimSuffix(name, "Paginated")
		normalizedInner := normalizeEntityPart(inner)
		if normalizedInner != "" {
			return normalizedInner + "Paginated"
		}
		return name
	}

	normalized := normalizeEntityPart(name)
	if normalized == "" {
		return original
	}

	return normalized
}

func normalizeEntityPart(part string) string {
	part = removeDuplicatedNames(part)

	part = removeEmbeddedEntityReferences(part)

	part = strings.Trim(part, "_")

	return part
}

// This is where most of the work of actually removing "Entity" and "Entities" happens
func removeEmbeddedEntityReferences(name string) string {
	if isIdentityOnlySchema(name) {
		return name
	}

	result := name

	if !strings.Contains(strings.ToLower(result), "identities") {
		result = regexp.MustCompile(`Entities([A-Z])`).ReplaceAllString(result, "$1")
		result = strings.ReplaceAll(result, "Entities_", "_")
		result = strings.ReplaceAll(result, "_Entities_", "_")
		result = strings.ReplaceAll(result, "_Entities", "")
		result = strings.TrimSuffix(result, "Entities")
	}

	if strings.Contains(strings.ToLower(result), "identity") {
		// Special handling for identity cases - only remove Entity when it's clearly separate
		// Handle cases like "ChangeIdentityEntity" and "ChangeIdentityEntityPaginated"
		if strings.Contains(result, "IdentityEntity") {
			result = strings.ReplaceAll(result, "IdentityEntity", "Identity")
		}
		// Handle underscored cases
		result = strings.ReplaceAll(result, "Identity_Entity", "Identity")
		result = strings.ReplaceAll(result, "IdentityEntity_", "Identity_")
	} else {
		// Normal entity removal when no identity is present
		result = regexp.MustCompile(`Entity([A-Z])`).ReplaceAllString(result, "$1")
		result = strings.ReplaceAll(result, "Entity_", "_")
		result = strings.ReplaceAll(result, "_Entity_", "_")
		result = strings.ReplaceAll(result, "_Entity", "")
		result = strings.TrimSuffix(result, "Entity")
	}

	// Handle lowercase variants - avoid matching within "identity"
	if !strings.Contains(strings.ToLower(result), "identity") {
		result = regexp.MustCompile(`entities([A-Z])`).ReplaceAllString(result, "$1")
		result = strings.ReplaceAll(result, "entities_", "_")
		result = strings.ReplaceAll(result, "_entities_", "_")
		result = strings.ReplaceAll(result, "_entities", "")
		result = strings.TrimSuffix(result, "entities")

		result = regexp.MustCompile(`entity([A-Z])`).ReplaceAllString(result, "$1")
		result = strings.ReplaceAll(result, "entity_", "_")
		result = strings.ReplaceAll(result, "_entity_", "_")
		result = strings.ReplaceAll(result, "_entity", "")
		result = strings.TrimSuffix(result, "entity")
	}

	// Clean up multiple underscores
	for strings.Contains(result, "__") {
		result = strings.ReplaceAll(result, "__", "_")
	}

	return strings.Trim(result, "_")
}

func removeDuplicatedNames(name string) string {
	parts := strings.Split(name, "_")
	if len(parts) < 2 {
		return name
	}

	if len(parts) >= 3 {
		for i := 0; i < len(parts)-2; i++ {
			if parts[i+1] == "Entities" && strings.HasSuffix(parts[len(parts)-1], "Entity") {
				baseLastPart := strings.TrimSuffix(parts[len(parts)-1], "Entity")
				if isPlural(parts[i], baseLastPart) || strings.EqualFold(parts[i], baseLastPart) {
					return parts[i]
				}
			}
		}
	}

	lastPart := parts[len(parts)-1]
	if strings.HasSuffix(lastPart, "Entity") {
		baseLastPart := strings.TrimSuffix(lastPart, "Entity")
		for i := 0; i < len(parts)-1; i++ {
			if isPlural(parts[i], baseLastPart) || strings.EqualFold(parts[i], baseLastPart) {
				return strings.Join(parts[:i+1], "_")
			}
		}
	}

	for i := 0; i < len(parts)-1; i++ {
		if strings.EqualFold(parts[i], lastPart) {
			return strings.Join(parts[:len(parts)-1], "_")
		}
	}

	return name
}

func isPlural(pluralForm, singular string) bool {
	plural := strings.ToLower(pluralForm)
	sing := strings.ToLower(singular)

	if plural == sing {
		return true
	}

	if plural == sing+"s" || plural == sing+"es" {
		return true
	}

	if strings.HasSuffix(sing, "y") && plural == strings.TrimSuffix(sing, "y")+"ies" {
		return true
	}

	// Common irregular plurals
	irregularPlurals := map[string]string{
		"person": "people",
		"user":   "users",
		"book":   "books",
	}

	if expected, exists := irregularPlurals[sing]; exists {
		return plural == expected
	}

	return false
}

func ensureUniqueRenamedSchemas(renames []SchemaRenameInfo, schemas map[string]interface{}) []SchemaRenameInfo {
	uniqueRenames := make([]SchemaRenameInfo, 0)

	for _, rename := range renames {
		finalName := rename.NewName
		counter := 1

		for {
			if _, exists := schemas[finalName]; !exists {
				nameInUse := false
				for _, r := range renames {
					if r.OldName == finalName {
						nameInUse = true
						break
					}
				}
				if !nameInUse {
					break
				}
			}

			finalName = fmt.Sprintf("%s_%d", rename.NewName, counter)
			counter++
		}

		uniqueRenames = append(uniqueRenames, SchemaRenameInfo{
			OldName: rename.OldName,
			NewName: finalName,
		})
	}

	return uniqueRenames
}

func updateReferences(spec map[string]interface{}, renameMap map[string]string) {
	updateReferencesRecursive(spec, renameMap)
}

func updateReferencesRecursive(obj interface{}, renameMap map[string]string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		if ref, hasRef := v["$ref"].(string); hasRef {
			if strings.HasPrefix(ref, "#/components/schemas/") {
				schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
				if newName, shouldRename := renameMap[schemaName]; shouldRename {
					v["$ref"] = "#/components/schemas/" + newName
				}
			}
		}

		for _, value := range v {
			updateReferencesRecursive(value, renameMap)
		}

	case []interface{}:
		for _, item := range v {
			updateReferencesRecursive(item, renameMap)
		}
	}
}

func updateDescriptions(spec map[string]interface{}, renameMap map[string]string) {
	updateDescriptionsRecursive(spec, renameMap)
}

func updateDescriptionsRecursive(obj interface{}, renameMap map[string]string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		if desc, hasDesc := v["description"].(string); hasDesc {
			updatedDesc := updateDescriptionText(desc, renameMap)
			if updatedDesc != desc {
				v["description"] = updatedDesc
			}
		}

		for _, value := range v {
			updateDescriptionsRecursive(value, renameMap)
		}

	case []interface{}:
		for _, item := range v {
			updateDescriptionsRecursive(item, renameMap)
		}
	}
}

func updateDescriptionText(description string, renameMap map[string]string) string {
	updatedDesc := description

	for oldName, newName := range renameMap {
		updatedDesc = strings.ReplaceAll(updatedDesc, oldName, newName)
	}

	updatedDesc = normalizeEntityReferencesInText(updatedDesc)

	return updatedDesc
}

func normalizeEntityReferencesInText(text string) string {
	// Handle "SomethingEntity model" -> "Something model"
	entityModelPattern := regexp.MustCompile(`(\w+)Entity\s+model`)
	text = entityModelPattern.ReplaceAllString(text, "$1 model")

	// Handle "SomethingEntity entity" -> "Something entity"
	entityEntityPattern := regexp.MustCompile(`(\w+)Entity\s+entity`)
	text = entityEntityPattern.ReplaceAllString(text, "$1 entity")

	// Handle standalone entity references
	standaloneEntityPattern := regexp.MustCompile(`\b(\w+)Entity\b`)
	text = standaloneEntityPattern.ReplaceAllStringFunc(text, func(match string) string {
		// Skip if this is just the word "entity"
		if strings.ToLower(match) == "entity" {
			return match
		}
		return strings.TrimSuffix(match, "Entity")
	})

	return text
}
