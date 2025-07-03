// File: scripts/normalize/enums_test.go
package main

import (
	"testing"
)

func TestNormalizeEnums(t *testing.T) {
	tests := []struct {
		name     string
		schemas  map[string]interface{}
		expected int // number of conflicts expected
	}{
		{
			name: "single enum property",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"status": map[string]interface{}{
							"type": "string",
							"enum": []interface{}{"active", "inactive", "pending"},
						},
					},
				},
			},
			expected: 1,
		},
		{
			name: "nested enum properties",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"status": map[string]interface{}{
							"type": "string",
							"enum": []interface{}{"active", "inactive"},
						},
						"preferences": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"theme": map[string]interface{}{
									"type": "string",
									"enum": []interface{}{"light", "dark"},
								},
							},
						},
					},
				},
			},
			expected: 2,
		},
		{
			name:     "no enum properties",
			schemas:  map[string]interface{}{},
			expected: 0,
		},
		{
			name: "array with enum items",
			schemas: map[string]interface{}{
				"Config": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"tags": map[string]interface{}{
							"type": "array",
							"items": map[string]interface{}{
								"type": "string",
								"enum": []interface{}{"urgent", "normal", "low"},
							},
						},
					},
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conflicts := normalizeEnums(tt.schemas)
			if len(conflicts) != tt.expected {
				t.Errorf("expected %d conflicts, got %d", tt.expected, len(conflicts))
			}

			// Verify transformation occurred
			if tt.expected > 0 {
				// Check that at least one schema was transformed
				found := false
				for _, schema := range tt.schemas {
					if hasXSpeakeasyEnums(schema) {
						found = true
						break
					}
				}
				if !found {
					t.Error("expected at least one schema to have x-speakeasy-enums")
				}
			}
		})
	}
}

func TestFindAllEnumProperties(t *testing.T) {
	schemas := map[string]interface{}{
		"User": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"status": map[string]interface{}{
					"type": "string",
					"enum": []interface{}{"active", "inactive"},
				},
				"role": map[string]interface{}{
					"type": "string",
					"enum": []interface{}{"admin", "user"},
				},
			},
		},
		"Product": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"category": map[string]interface{}{
					"type": "string",
					"enum": []interface{}{"electronics", "clothing"},
				},
			},
		},
	}

	enums := findAllEnumProperties(schemas)

	if len(enums) != 3 {
		t.Errorf("expected 3 enum properties, got %d", len(enums))
	}

	// Verify enum info structure
	for _, enumInfo := range enums {
		if enumInfo.SchemaName == "" {
			t.Error("expected non-empty schema name")
		}
		if enumInfo.PropertyName == "" {
			t.Error("expected non-empty property name")
		}
		if len(enumInfo.EnumValues) == 0 {
			t.Error("expected non-empty enum values")
		}
		if enumInfo.Target == nil {
			t.Error("expected non-nil target")
		}
	}
}

func TestGenerateEnumMemberName(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		fieldName  string
		entityName string
		expected   string
	}{
		{
			name:       "basic enum member",
			value:      "active",
			fieldName:  "status",
			entityName: "User",
			expected:   "UserStatusActive",
		},
		{
			name:       "with special characters",
			value:      "in-progress",
			fieldName:  "state",
			entityName: "Task",
			expected:   "TaskStateInProgress",
		},
		{
			name:       "empty value",
			value:      "",
			fieldName:  "status",
			entityName: "User",
			expected:   "UserStatusEmpty",
		},
		{
			name:       "numeric value",
			value:      "123test",
			fieldName:  "code",
			entityName: "Product",
			expected:   "ProductCodeValue123test",
		},
		{
			name:       "request param entity",
			value:      "ascending",
			fieldName:  "sort",
			entityName: "RequestParam",
			expected:   "RequestSortAscending",
		},
		{
			name:       "request body entity",
			value:      "json",
			fieldName:  "format",
			entityName: "RequestBody",
			expected:   "RequestFormatJson",
		},
		{
			name:       "response entity",
			value:      "success",
			fieldName:  "status",
			entityName: "Response",
			expected:   "ResponseStatusSuccess",
		},
		{
			name:       "entity suffix",
			value:      "draft",
			fieldName:  "state",
			entityName: "DocumentEntity",
			expected:   "DocumentStateDraft",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateEnumMemberName(tt.value, tt.fieldName, tt.entityName)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestConvertToEnumMemberName(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "simple string",
			value:    "active",
			expected: "Active",
		},
		{
			name:     "with underscores",
			value:    "user_admin",
			expected: "UserAdmin",
		},
		{
			name:     "with hyphens",
			value:    "in-progress",
			expected: "InProgress",
		},
		{
			name:     "with spaces",
			value:    "waiting for approval",
			expected: "WaitingForApproval",
		},
		{
			name:     "mixed special chars",
			value:    "test_value-123 final",
			expected: "TestValue123Final",
		},
		{
			name:     "empty string",
			value:    "",
			expected: "Empty",
		},
		{
			name:     "only special chars",
			value:    "___---",
			expected: "Empty",
		},
		{
			name:     "starts with number",
			value:    "123abc",
			expected: "Value123abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToEnumMemberName(tt.value)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestExtractPropertyNameFromNormalizationPath(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		schemaName string
		expected   string
	}{
		{
			name:       "empty path",
			path:       "",
			schemaName: "UserEntity",
			expected:   "User",
		},
		{
			name:       "simple property path",
			path:       "properties.status",
			schemaName: "User",
			expected:   "status",
		},
		{
			name:       "nested property path",
			path:       "properties.preferences.properties.theme",
			schemaName: "User",
			expected:   "theme",
		},
		{
			name:       "array items path",
			path:       "properties.tags.items",
			schemaName: "Document",
			expected:   "tags",
		},
		{
			name:       "complex nested path",
			path:       "properties.metadata.properties.author.properties.role",
			schemaName: "Document",
			expected:   "role",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPropertyNameFromNormalizationPath(tt.path, tt.schemaName)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestTransformEnumProperty(t *testing.T) {
	enumInfo := EnumNormalizationInfo{
		SchemaName:   "User",
		PropertyPath: "properties.status",
		PropertyName: "status",
		EnumValues:   []string{"active", "inactive", "pending"},
		Target: map[string]interface{}{
			"type": "string",
			"enum": []interface{}{"active", "inactive", "pending"},
		},
	}

	conflict := transformEnumProperty(enumInfo)

	if conflict == nil {
		t.Fatal("expected conflict detail, got nil")
	}

	if conflict.Schema != "User" {
		t.Errorf("expected schema 'User', got %s", conflict.Schema)
	}

	if conflict.ConflictType != "enum-normalization" {
		t.Errorf("expected conflict type 'enum-normalization', got %s", conflict.ConflictType)
	}

	// Verify the enum was removed and x-speakeasy-enums was added
	if _, hasEnum := enumInfo.Target["enum"]; hasEnum {
		t.Error("expected enum to be removed from target")
	}

	if xEnums, hasXEnums := enumInfo.Target["x-speakeasy-enums"]; hasXEnums {
		if members, ok := xEnums.([]map[string]interface{}); ok {
			if len(members) != 3 {
				t.Errorf("expected 3 enum members, got %d", len(members))
			}
		} else {
			t.Error("expected x-speakeasy-enums to be array of maps")
		}
	} else {
		t.Error("expected x-speakeasy-enums to be added to target")
	}
}

func TestNormalizePathEnums(t *testing.T) {
	paths := map[string]interface{}{
		"/users/{id}": map[string]interface{}{
			"get": map[string]interface{}{
				"parameters": []interface{}{
					map[string]interface{}{
						"name": "sort",
						"in":   "query",
						"schema": map[string]interface{}{
							"type": "string",
							"enum": []interface{}{"asc", "desc"},
						},
					},
				},
			},
		},
	}

	conflicts := normalizePathEnums(paths)

	if len(conflicts) != 1 {
		t.Errorf("expected 1 conflict, got %d", len(conflicts))
	}

	if len(conflicts) > 0 {
		if conflicts[0].ConflictType != "enum-normalization" {
			t.Errorf("expected conflict type 'enum-normalization', got %s", conflicts[0].ConflictType)
		}
	}
}

// Helper function to check if a schema has x-speakeasy-enums
func hasXSpeakeasyEnums(obj interface{}) bool {
	switch v := obj.(type) {
	case map[string]interface{}:
		if _, hasXEnums := v["x-speakeasy-enums"]; hasXEnums {
			return true
		}
		for _, value := range v {
			if hasXSpeakeasyEnums(value) {
				return true
			}
		}
	case []interface{}:
		for _, item := range v {
			if hasXSpeakeasyEnums(item) {
				return true
			}
		}
	}
	return false
}
