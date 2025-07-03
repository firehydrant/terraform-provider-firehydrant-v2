// File: scripts/overlay/reserved_words_test.go
package main

import (
	"testing"
)

func TestApplyTerraformReservedWordIgnores(t *testing.T) {
	// Suppress output during tests
	restore := suppressOutput()
	defer restore()

	resources := map[string]*ResourceInfo{
		"user": {
			EntityName: "UserEntity",
		},
		"config": {
			EntityName: "ConfigEntity",
		},
	}

	datasources := map[string]*ResourceInfo{
		"settings": {
			EntityName: "SettingsEntity",
		},
	}

	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":       map[string]interface{}{"type": "string"},
				"name":     map[string]interface{}{"type": "string"},
				"count":    map[string]interface{}{"type": "integer"}, // Reserved word
				"provider": map[string]interface{}{"type": "string"},  // Reserved word
				"normal":   map[string]interface{}{"type": "string"},  // Not reserved
			},
		},
		"ConfigEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"terraform": map[string]interface{}{"type": "object"}, // Reserved word
				"variable":  map[string]interface{}{"type": "string"}, // Reserved word
				"module":    map[string]interface{}{"type": "object"}, // Reserved word
				"value":     map[string]interface{}{"type": "string"}, // Not reserved
			},
		},
		"SettingsEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"locals":     map[string]interface{}{"type": "object"}, // Reserved word
				"depends_on": map[string]interface{}{"type": "array"},  // Reserved word
				"config":     map[string]interface{}{"type": "object"}, // Not reserved
			},
		},
	}

	overlay := &Overlay{}
	ignoreTracker := make(map[string]map[string]bool)

	count := applyTerraformReservedWordIgnores(overlay, resources, datasources, schemas, ignoreTracker)

	expectedCount := 7 // count, provider, terraform, variable, module, locals, depends_on
	if count != expectedCount {
		t.Errorf("expected %d reserved words ignored, got %d", expectedCount, count)
	}

	if len(overlay.Actions) != expectedCount {
		t.Errorf("expected %d overlay actions, got %d", expectedCount, len(overlay.Actions))
	}

	// Check that reserved words were marked for ignore
	expectedIgnores := map[string][]string{
		"UserEntity":     {"count", "provider"},
		"ConfigEntity":   {"terraform", "variable", "module"},
		"SettingsEntity": {"locals", "depends_on"},
	}

	for schema, expectedProps := range expectedIgnores {
		if ignoreTracker[schema] == nil {
			t.Errorf("expected ignore tracker for %s", schema)
			continue
		}
		for _, prop := range expectedProps {
			if !ignoreTracker[schema][prop] {
				t.Errorf("expected %s.%s to be marked for ignore", schema, prop)
			}
		}
	}

	// Verify overlay actions were created correctly
	foundActions := make(map[string]map[string]bool)
	for _, action := range overlay.Actions {
		// Extract schema and property from target
		// Expected format: $.components.schemas.SchemaName.properties.propertyName
		if action.Target != "" {
			parts := splitTarget(action.Target)
			if len(parts) >= 2 {
				schema := parts[0]
				property := parts[1]

				if foundActions[schema] == nil {
					foundActions[schema] = make(map[string]bool)
				}
				foundActions[schema][property] = true
			}
		}

		// Verify the action has the correct update
		if ignore, ok := action.Update["x-speakeasy-ignore"]; !ok || ignore != true {
			t.Errorf("expected action to have x-speakeasy-ignore=true, got %v", action.Update)
		}
	}

	// Verify all expected actions were created
	for schema, expectedProps := range expectedIgnores {
		for _, prop := range expectedProps {
			if !foundActions[schema][prop] {
				t.Errorf("expected to find overlay action for %s.%s", schema, prop)
			}
		}
	}
}

func TestIsReservedTerraformWord(t *testing.T) {
	reservedWords := []string{"count", "provider", "terraform", "variable", "module", "locals", "depends_on", "connection", "for_each", "lifecycle", "output", "resource", "data", "provisioner", "import", "moved", "removed", "check", "precondition", "postcondition"}

	tests := []struct {
		name     string
		propName string
		expected bool
	}{
		// Basic reserved words
		{
			name:     "reserved word - count",
			propName: "count",
			expected: true,
		},
		{
			name:     "reserved word - provider",
			propName: "provider",
			expected: true,
		},
		{
			name:     "reserved word - terraform",
			propName: "terraform",
			expected: true,
		},
		{
			name:     "reserved word - variable",
			propName: "variable",
			expected: true,
		},
		{
			name:     "reserved word - module",
			propName: "module",
			expected: true,
		},
		{
			name:     "reserved word - locals",
			propName: "locals",
			expected: true,
		},
		{
			name:     "reserved word - depends_on",
			propName: "depends_on",
			expected: true,
		},
		{
			name:     "reserved word - connection",
			propName: "connection",
			expected: true,
		},
		{
			name:     "reserved word - for_each",
			propName: "for_each",
			expected: true,
		},
		{
			name:     "reserved word - lifecycle",
			propName: "lifecycle",
			expected: true,
		},
		{
			name:     "reserved word - output",
			propName: "output",
			expected: true,
		},
		{
			name:     "reserved word - provisioner",
			propName: "provisioner",
			expected: true,
		},
		{
			name:     "reserved word - import",
			propName: "import",
			expected: true,
		},
		{
			name:     "reserved word - moved",
			propName: "moved",
			expected: true,
		},
		{
			name:     "reserved word - removed",
			propName: "removed",
			expected: true,
		},
		{
			name:     "reserved word - check",
			propName: "check",
			expected: true,
		},
		{
			name:     "reserved word - precondition",
			propName: "precondition",
			expected: true,
		},
		{
			name:     "reserved word - postcondition",
			propName: "postcondition",
			expected: true,
		},

		// Case insensitive tests
		{
			name:     "reserved word - uppercase",
			propName: "TERRAFORM",
			expected: true,
		},
		{
			name:     "reserved word - mixed case",
			propName: "Provider",
			expected: true,
		},
		{
			name:     "reserved word - mixed case depends_on",
			propName: "DEPENDS_ON",
			expected: true,
		},

		// Non-reserved words
		{
			name:     "not reserved - name",
			propName: "name",
			expected: false,
		},
		{
			name:     "not reserved - value",
			propName: "value",
			expected: false,
		},
		{
			name:     "not reserved - config",
			propName: "config",
			expected: false,
		},
		{
			name:     "not reserved - user_count",
			propName: "user_count",
			expected: false,
		},
		{
			name:     "not reserved - similar but different",
			propName: "counter",
			expected: false,
		},
		{
			name:     "not reserved - providers (plural)",
			propName: "providers",
			expected: false,
		},
		{
			name:     "not reserved - terraform_version",
			propName: "terraform_version",
			expected: false,
		},
		{
			name:     "not reserved - empty string",
			propName: "",
			expected: false,
		},
		{
			name:     "not reserved - numeric",
			propName: "123",
			expected: false,
		},
		{
			name:     "not reserved - special chars",
			propName: "test-property",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isReservedTerraformWord(tt.propName, reservedWords)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestApplyTerraformReservedWordIgnoresEdgeCases(t *testing.T) {
	restore := suppressOutput()
	defer restore()

	t.Run("empty resources and datasources", func(t *testing.T) {
		resources := map[string]*ResourceInfo{}
		datasources := map[string]*ResourceInfo{}
		schemas := map[string]interface{}{}

		overlay := &Overlay{}
		ignoreTracker := make(map[string]map[string]bool)

		count := applyTerraformReservedWordIgnores(overlay, resources, datasources, schemas, ignoreTracker)

		if count != 0 {
			t.Errorf("expected 0 reserved words ignored for empty input, got %d", count)
		}

		if len(overlay.Actions) != 0 {
			t.Errorf("expected 0 overlay actions for empty input, got %d", len(overlay.Actions))
		}
	})

	t.Run("schema without properties", func(t *testing.T) {
		resources := map[string]*ResourceInfo{
			"test": {
				EntityName: "TestEntity",
			},
		}

		schemas := map[string]interface{}{
			"TestEntity": map[string]interface{}{
				"type": "object",
				// No properties
			},
		}

		overlay := &Overlay{}
		ignoreTracker := make(map[string]map[string]bool)

		count := applyTerraformReservedWordIgnores(overlay, resources, map[string]*ResourceInfo{}, schemas, ignoreTracker)

		if count != 0 {
			t.Errorf("expected 0 reserved words ignored for schema without properties, got %d", count)
		}
	})

	t.Run("invalid schema type", func(t *testing.T) {
		resources := map[string]*ResourceInfo{
			"test": {
				EntityName: "TestEntity",
			},
		}

		schemas := map[string]interface{}{
			"TestEntity": "not a map", // Invalid schema type
		}

		overlay := &Overlay{}
		ignoreTracker := make(map[string]map[string]bool)

		count := applyTerraformReservedWordIgnores(overlay, resources, map[string]*ResourceInfo{}, schemas, ignoreTracker)

		if count != 0 {
			t.Errorf("expected 0 reserved words ignored for invalid schema, got %d", count)
		}
	})

	t.Run("properties is not a map", func(t *testing.T) {
		resources := map[string]*ResourceInfo{
			"test": {
				EntityName: "TestEntity",
			},
		}

		schemas := map[string]interface{}{
			"TestEntity": map[string]interface{}{
				"type":       "object",
				"properties": "not a map", // Invalid properties type
			},
		}

		overlay := &Overlay{}
		ignoreTracker := make(map[string]map[string]bool)

		count := applyTerraformReservedWordIgnores(overlay, resources, map[string]*ResourceInfo{}, schemas, ignoreTracker)

		if count != 0 {
			t.Errorf("expected 0 reserved words ignored for invalid properties, got %d", count)
		}
	})

	t.Run("already ignored properties", func(t *testing.T) {
		resources := map[string]*ResourceInfo{
			"test": {
				EntityName: "TestEntity",
			},
		}

		schemas := map[string]interface{}{
			"TestEntity": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"count": map[string]interface{}{"type": "integer"},
				},
			},
		}

		overlay := &Overlay{}
		ignoreTracker := map[string]map[string]bool{
			"TestEntity": {
				"count": true, // Already ignored
			},
		}

		count := applyTerraformReservedWordIgnores(overlay, resources, map[string]*ResourceInfo{}, schemas, ignoreTracker)

		if count != 0 {
			t.Errorf("expected 0 reserved words ignored for already ignored property, got %d", count)
		}

		if len(overlay.Actions) != 0 {
			t.Errorf("expected 0 overlay actions for already ignored property, got %d", len(overlay.Actions))
		}
	})
}

func TestReservedWordsCompleteness(t *testing.T) {
	// Test that we have all the expected Terraform reserved words
	// This ensures we don't miss any important reserved words

	// These are the core Terraform reserved words that should be included
	coreReservedWords := []string{
		"count", "connection", "for_each", "lifecycle", "depends_on",
		"provider", "provisioner", "locals", "module", "terraform",
		"variable", "output", "import", "moved", "removed", "check",
		"precondition", "postcondition",
	}

	for _, word := range coreReservedWords {
		if !isReservedTerraformWord(word, coreReservedWords) {
			t.Errorf("core reserved word '%s' should be detected as reserved", word)
		}
	}
}

// Helper function to split target path and extract schema and property
func splitTarget(target string) []string {
	// Expected format: $.components.schemas.SchemaName.properties.PropertyName
	// Simple extraction for testing
	if target == "" {
		return []string{}
	}

	// Look for pattern like "schemas.EntityName.properties.propertyName"
	parts := []string{}

	// Find schema name
	schemaStart := "schemas."
	propStart := ".properties."

	schemaIdx := indexOf(target, schemaStart)
	if schemaIdx == -1 {
		return parts
	}

	propIdx := indexOf(target, propStart)
	if propIdx == -1 {
		return parts
	}

	// Extract schema name
	schemaName := target[schemaIdx+len(schemaStart) : propIdx]

	// Extract property name (everything after .properties.)
	propertyName := target[propIdx+len(propStart):]

	return []string{schemaName, propertyName}
}

// Helper function to find index of substring
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
