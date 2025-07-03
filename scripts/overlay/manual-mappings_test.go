// File: scripts/overlay/manual_mappings_test.go
package main

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLoadManualMappings(t *testing.T) {
	// Suppress output during tests
	restore := suppressOutput()
	defer restore()

	t.Run("non-existent file", func(t *testing.T) {
		mappings := loadManualMappings("non-existent-file.yaml")
		if mappings == nil {
			t.Error("expected non-nil mappings for non-existent file")
		}
		if mappings == nil || len(mappings.Operations) != 0 {
			t.Error("expected empty operations for non-existent file")
		}
	})

	t.Run("valid yaml file", func(t *testing.T) {
		// Create temporary test file
		testMappings := ManualMappings{
			Operations: []ManualMapping{
				{
					Path:   "/users/{id}",
					Method: "get",
					Action: "match",
					Value:  "id:user_id",
				},
				{
					Path:   "/admin/debug",
					Method: "get",
					Action: "ignore",
				},
				{
					Path:   "/special/endpoint",
					Method: "post",
					Action: "entity",
					Value:  "CustomEntity",
				},
				{
					Action:   "ignore_property",
					Schema:   "UserEntity",
					Property: "internal_field",
				},
				{
					Action:   "additional_properties",
					Schema:   "ConfigEntity",
					Property: "metadata",
				},
			},
		}

		yamlData, err := yaml.Marshal(testMappings)
		if err != nil {
			t.Fatalf("failed to marshal test data: %v", err)
		}

		tmpFile, err := ioutil.TempFile("", "test-mappings-*.yaml")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.Write(yamlData); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}
		tmpFile.Close()

		// Test loading
		mappings := loadManualMappings(tmpFile.Name())
		if mappings == nil {
			t.Error("expected non-nil mappings")
		}
		if mappings == nil || len(mappings.Operations) != 5 {
			t.Errorf("expected 5 operations, got %d", len(mappings.Operations))
		}

		// Verify first operation
		if mappings.Operations[0].Path != "/users/{id}" {
			t.Errorf("expected path '/users/{id}', got '%s'", mappings.Operations[0].Path)
		}
		if mappings.Operations[0].Action != "match" {
			t.Errorf("expected action 'match', got '%s'", mappings.Operations[0].Action)
		}
		if mappings.Operations[0].Value != "id:user_id" {
			t.Errorf("expected value 'id:user_id', got '%s'", mappings.Operations[0].Value)
		}

		// Verify ignore operation
		if mappings.Operations[1].Action != "ignore" {
			t.Errorf("expected action 'ignore', got '%s'", mappings.Operations[1].Action)
		}

		// Verify entity operation
		if mappings.Operations[2].Action != "entity" {
			t.Errorf("expected action 'entity', got '%s'", mappings.Operations[2].Action)
		}
		if mappings.Operations[2].Value != "CustomEntity" {
			t.Errorf("expected value 'CustomEntity', got '%s'", mappings.Operations[2].Value)
		}

		// Verify property ignore operation
		if mappings.Operations[3].Action != "ignore_property" {
			t.Errorf("expected action 'ignore_property', got '%s'", mappings.Operations[3].Action)
		}
		if mappings.Operations[3].Schema != "UserEntity" {
			t.Errorf("expected schema 'UserEntity', got '%s'", mappings.Operations[3].Schema)
		}
		if mappings.Operations[3].Property != "internal_field" {
			t.Errorf("expected property 'internal_field', got '%s'", mappings.Operations[3].Property)
		}

		// Verify additional properties operation
		if mappings.Operations[4].Action != "additional_properties" {
			t.Errorf("expected action 'additional_properties', got '%s'", mappings.Operations[4].Action)
		}
		if mappings.Operations[4].Schema != "ConfigEntity" {
			t.Errorf("expected schema 'ConfigEntity', got '%s'", mappings.Operations[4].Schema)
		}
		if mappings.Operations[4].Property != "metadata" {
			t.Errorf("expected property 'metadata', got '%s'", mappings.Operations[4].Property)
		}
	})

	t.Run("invalid yaml file", func(t *testing.T) {
		// Create temporary file with invalid YAML
		tmpFile, err := ioutil.TempFile("", "test-invalid-*.yaml")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		invalidYaml := "invalid: yaml: content: [unclosed bracket"
		if _, err := tmpFile.WriteString(invalidYaml); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}
		tmpFile.Close()

		// Should return empty mappings on parse error
		mappings := loadManualMappings(tmpFile.Name())
		if mappings == nil {
			t.Error("expected non-nil mappings even for invalid YAML")
		}
		if mappings == nil || len(mappings.Operations) != 0 {
			t.Error("expected empty operations for invalid YAML")
		}
	})

	t.Run("empty yaml file", func(t *testing.T) {
		tmpFile, err := ioutil.TempFile("", "test-empty-*.yaml")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		mappings := loadManualMappings(tmpFile.Name())
		if mappings == nil {
			t.Error("expected non-nil mappings for empty file")
		}
		if mappings == nil || len(mappings.Operations) != 0 {
			t.Error("expected empty operations for empty file")
		}
	})
}

func TestGetManualParameterMatch(t *testing.T) {
	mappings := &ManualMappings{
		Operations: []ManualMapping{
			{
				Path:   "/users/{user_id}",
				Method: "get",
				Action: "match",
				Value:  "user_id:id",
			},
			{
				Path:   "/posts/{post_id}",
				Method: "get",
				Action: "match",
				Value:  "post_id:slug",
			},
			{
				Path:   "/complex/{param}",
				Method: "post",
				Action: "match",
				Value:  "param:nested.field.id",
			},
			{
				Path:   "/wrong-action/{id}",
				Method: "get",
				Action: "ignore", // Not a match action
				Value:  "id:something",
			},
			{
				Path:   "/malformed/{id}",
				Method: "get",
				Action: "match",
				Value:  "malformed_value", // No colon separator
			},
		},
	}

	tests := []struct {
		name      string
		path      string
		method    string
		paramName string
		expected  string
		found     bool
	}{
		{
			name:      "found match",
			path:      "/users/{user_id}",
			method:    "get",
			paramName: "user_id",
			expected:  "id",
			found:     true,
		},
		{
			name:      "found match case insensitive method",
			path:      "/users/{user_id}",
			method:    "GET",
			paramName: "user_id",
			expected:  "id",
			found:     true,
		},
		{
			name:      "found different mapping",
			path:      "/posts/{post_id}",
			method:    "get",
			paramName: "post_id",
			expected:  "slug",
			found:     true,
		},
		{
			name:      "found nested field mapping",
			path:      "/complex/{param}",
			method:    "post",
			paramName: "param",
			expected:  "nested.field.id",
			found:     true,
		},
		{
			name:      "no match - wrong param",
			path:      "/users/{user_id}",
			method:    "get",
			paramName: "wrong_param",
			expected:  "",
			found:     false,
		},
		{
			name:      "no match - wrong path",
			path:      "/wrong/path",
			method:    "get",
			paramName: "user_id",
			expected:  "",
			found:     false,
		},
		{
			name:      "no match - wrong method",
			path:      "/users/{user_id}",
			method:    "post",
			paramName: "user_id",
			expected:  "",
			found:     false,
		},
		{
			name:      "no match - wrong action type",
			path:      "/wrong-action/{id}",
			method:    "get",
			paramName: "id",
			expected:  "",
			found:     false,
		},
		{
			name:      "no match - malformed value",
			path:      "/malformed/{id}",
			method:    "get",
			paramName: "id",
			expected:  "",
			found:     false,
		},
		{
			name:      "no match - nonexistent param",
			path:      "/users/{user_id}",
			method:    "get",
			paramName: "nonexistent",
			expected:  "",
			found:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := getManualParameterMatch(tt.path, tt.method, tt.paramName, mappings)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
			if found != tt.found {
				t.Errorf("expected found=%t, got %t", tt.found, found)
			}
		})
	}
}

func TestShouldIgnoreOperation(t *testing.T) {
	// Suppress output during tests
	restore := suppressOutput()
	defer restore()

	mappings := &ManualMappings{
		Operations: []ManualMapping{
			{
				Path:   "/internal/debug",
				Method: "get",
				Action: "ignore",
			},
			{
				Path:   "/admin/reset",
				Method: "post",
				Action: "ignore",
			},
			{
				Path:   "/not-ignored",
				Method: "get",
				Action: "match", // Different action, should not ignore
				Value:  "param:field",
			},
		},
	}

	tests := []struct {
		name     string
		path     string
		method   string
		expected bool
	}{
		{
			name:     "should ignore",
			path:     "/internal/debug",
			method:   "get",
			expected: true,
		},
		{
			name:     "should ignore case insensitive",
			path:     "/admin/reset",
			method:   "POST",
			expected: true,
		},
		{
			name:     "should not ignore - no mapping",
			path:     "/users",
			method:   "get",
			expected: false,
		},
		{
			name:     "should not ignore - wrong method",
			path:     "/internal/debug",
			method:   "post",
			expected: false,
		},
		{
			name:     "should not ignore - wrong action",
			path:     "/not-ignored",
			method:   "get",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldIgnoreOperation(tt.path, tt.method, mappings)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestGetManualEntityMapping(t *testing.T) {
	mappings := &ManualMappings{
		Operations: []ManualMapping{
			{
				Path:   "/special/endpoint",
				Method: "get",
				Action: "entity",
				Value:  "SpecialEntity",
			},
			{
				Path:   "/custom/resource",
				Method: "post",
				Action: "entity",
				Value:  "CustomResourceEntity",
			},
			{
				Path:   "/not-entity",
				Method: "get",
				Action: "ignore", // Different action
			},
		},
	}

	tests := []struct {
		name     string
		path     string
		method   string
		expected string
		found    bool
	}{
		{
			name:     "found entity mapping",
			path:     "/special/endpoint",
			method:   "get",
			expected: "SpecialEntity",
			found:    true,
		},
		{
			name:     "found entity mapping case insensitive",
			path:     "/custom/resource",
			method:   "POST",
			expected: "CustomResourceEntity",
			found:    true,
		},
		{
			name:     "no entity mapping - no such path",
			path:     "/normal/endpoint",
			method:   "get",
			expected: "",
			found:    false,
		},
		{
			name:     "no entity mapping - wrong method",
			path:     "/special/endpoint",
			method:   "post",
			expected: "",
			found:    false,
		},
		{
			name:     "no entity mapping - wrong action",
			path:     "/not-entity",
			method:   "get",
			expected: "",
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := getManualEntityMapping(tt.path, tt.method, mappings)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
			if found != tt.found {
				t.Errorf("expected found=%t, got %t", tt.found, found)
			}
		})
	}
}

func TestGetManualPropertyIgnores(t *testing.T) {
	mappings := &ManualMappings{
		Operations: []ManualMapping{
			{
				Action:   "ignore_property",
				Schema:   "UserEntity",
				Property: "internal_field",
			},
			{
				Action:   "ignore_property",
				Schema:   "UserEntity",
				Property: "debug_info",
			},
			{
				Action:   "ignore_property",
				Schema:   "ProductEntity",
				Property: "admin_notes",
			},
			{
				Action:   "ignore_property",
				Schema:   "UserEntity",
				Property: "temp_data",
			},
			{
				// Missing schema - should be ignored
				Action:   "ignore_property",
				Property: "orphaned_property",
			},
			{
				// Missing property - should be ignored
				Action: "ignore_property",
				Schema: "EmptyEntity",
			},
			{
				// Different action - should be ignored
				Action:   "match",
				Schema:   "UserEntity",
				Property: "not_ignored",
			},
		},
	}

	result := getManualPropertyIgnores(mappings)

	// Check number of schemas
	if len(result) != 2 {
		t.Errorf("expected 2 schemas with ignores, got %d", len(result))
	}

	// Check UserEntity ignores
	userIgnores := result["UserEntity"]
	if len(userIgnores) != 3 {
		t.Errorf("expected 3 ignored properties for UserEntity, got %d", len(userIgnores))
	}

	expectedUserIgnores := []string{"internal_field", "debug_info", "temp_data"}
	for _, expected := range expectedUserIgnores {
		found := false
		for _, ignore := range userIgnores {
			if ignore == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected to find ignored property %s for UserEntity", expected)
		}
	}

	// Check ProductEntity ignores
	productIgnores := result["ProductEntity"]
	if len(productIgnores) != 1 {
		t.Errorf("expected 1 ignored property for ProductEntity, got %d", len(productIgnores))
	}
	if productIgnores[0] != "admin_notes" {
		t.Errorf("expected 'admin_notes' for ProductEntity, got %s", productIgnores[0])
	}

	// Check that invalid entries were filtered out
	if _, exists := result["EmptyEntity"]; exists {
		t.Error("expected EmptyEntity to be filtered out (missing property)")
	}
}

func TestGetAdditionalPropertiesMappings(t *testing.T) {
	mappings := &ManualMappings{
		Operations: []ManualMapping{
			{
				Action:   "additional_properties",
				Schema:   "UserEntity",
				Property: "metadata",
			},
			{
				Action:   "additional_properties",
				Schema:   "UserEntity",
				Property: "settings.preferences",
			},
			{
				Action:   "additional_properties",
				Schema:   "ProductEntity",
				Property: "custom_fields",
			},
			{
				Action:   "additional_properties",
				Schema:   "UserEntity",
				Property: "dynamic.config.values",
			},
			{
				// Missing schema - should be ignored
				Action:   "additional_properties",
				Property: "orphaned",
			},
			{
				// Missing property - should be ignored
				Action: "additional_properties",
				Schema: "EmptyEntity",
			},
			{
				// Different action - should be ignored
				Action:   "ignore",
				Schema:   "UserEntity",
				Property: "not_additional",
			},
		},
	}

	result := getAdditionalPropertiesMappings(mappings)

	// Check number of schemas
	if len(result) != 2 {
		t.Errorf("expected 2 schemas with additional properties, got %d", len(result))
	}

	// Check UserEntity properties
	userProps := result["UserEntity"]
	if len(userProps) != 3 {
		t.Errorf("expected 3 additional properties for UserEntity, got %d", len(userProps))
	}

	expectedUserProps := []string{"metadata", "settings.preferences", "dynamic.config.values"}
	for _, expected := range expectedUserProps {
		found := false
		for _, prop := range userProps {
			if prop == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected to find additional property %s for UserEntity", expected)
		}
	}

	// Check ProductEntity properties
	productProps := result["ProductEntity"]
	if len(productProps) != 1 {
		t.Errorf("expected 1 additional property for ProductEntity, got %d", len(productProps))
	}
	if productProps[0] != "custom_fields" {
		t.Errorf("expected 'custom_fields' for ProductEntity, got %s", productProps[0])
	}
}

func TestBuildAdditionalPropertiesPath(t *testing.T) {
	tests := []struct {
		name         string
		schemaName   string
		propertyPath string
		expected     string
	}{
		{
			name:         "simple property",
			schemaName:   "UserEntity",
			propertyPath: "metadata",
			expected:     "$.components.schemas.UserEntity.properties.metadata",
		},
		{
			name:         "nested property",
			schemaName:   "UserEntity",
			propertyPath: "settings.preferences",
			expected:     "$.components.schemas.UserEntity.properties.settings.properties.preferences",
		},
		{
			name:         "deeply nested property",
			schemaName:   "ConfigEntity",
			propertyPath: "app.database.connection.pool",
			expected:     "$.components.schemas.ConfigEntity.properties.app.properties.database.properties.connection.properties.pool",
		},
		{
			name:         "array items",
			schemaName:   "UserEntity",
			propertyPath: "tags.items",
			expected:     "$.components.schemas.UserEntity.properties.tags.items",
		},
		{
			name:         "complex nested with items",
			schemaName:   "UserEntity",
			propertyPath: "collections.items.metadata",
			expected:     "$.components.schemas.UserEntity.properties.collections.items.properties.metadata",
		},
		{
			name:         "multiple items levels",
			schemaName:   "DataEntity",
			propertyPath: "matrix.items.rows.items",
			expected:     "$.components.schemas.DataEntity.properties.matrix.items.properties.rows.items",
		},
		{
			name:         "items at the end",
			schemaName:   "ListEntity",
			propertyPath: "values.items",
			expected:     "$.components.schemas.ListEntity.properties.values.items",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildAdditionalPropertiesPath(tt.schemaName, tt.propertyPath)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// Test the ManualMapping struct fields
func TestManualMappingStruct(t *testing.T) {
	mapping := ManualMapping{
		Path:     "/test/path",
		Method:   "GET",
		Action:   "match",
		Value:    "param:field",
		Schema:   "TestEntity",
		Property: "test_property",
	}

	if mapping.Path != "/test/path" {
		t.Errorf("expected path '/test/path', got '%s'", mapping.Path)
	}
	if mapping.Method != "GET" {
		t.Errorf("expected method 'GET', got '%s'", mapping.Method)
	}
	if mapping.Action != "match" {
		t.Errorf("expected action 'match', got '%s'", mapping.Action)
	}
	if mapping.Value != "param:field" {
		t.Errorf("expected value 'param:field', got '%s'", mapping.Value)
	}
	if mapping.Schema != "TestEntity" {
		t.Errorf("expected schema 'TestEntity', got '%s'", mapping.Schema)
	}
	if mapping.Property != "test_property" {
		t.Errorf("expected property 'test_property', got '%s'", mapping.Property)
	}
}

// Test ManualMappings struct
func TestManualMappingsStruct(t *testing.T) {
	mappings := ManualMappings{
		Operations: []ManualMapping{
			{
				Path:   "/test1",
				Method: "GET",
				Action: "ignore",
			},
			{
				Path:   "/test2",
				Method: "POST",
				Action: "match",
				Value:  "param:field",
			},
		},
	}

	if len(mappings.Operations) != 2 {
		t.Errorf("expected 2 operations, got %d", len(mappings.Operations))
	}

	if mappings.Operations[0].Path != "/test1" {
		t.Errorf("expected first operation path '/test1', got '%s'", mappings.Operations[0].Path)
	}

	if mappings.Operations[1].Action != "match" {
		t.Errorf("expected second operation action 'match', got '%s'", mappings.Operations[1].Action)
	}
}
