package common

import (
	"testing"
)

func TestResolveRef(t *testing.T) {
	schemas := map[string]interface{}{
		"User": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"NullableUser": map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{
					"$ref": "#/components/schemas/User",
				},
			},
		},
	}

	tests := []struct {
		name           string
		ref            string
		expectedSchema map[string]interface{}
		expectedName   string
	}{
		{
			name:           "resolve direct reference",
			ref:            "#/components/schemas/User",
			expectedSchema: schemas["User"].(map[string]interface{}),
			expectedName:   "User",
		},
		{
			name:           "resolve nullable wrapper",
			ref:            "#/components/schemas/NullableUser",
			expectedSchema: schemas["User"].(map[string]interface{}),
			expectedName:   "User",
		},
		{
			name:           "empty reference",
			ref:            "",
			expectedSchema: nil,
			expectedName:   "",
		},
		{
			name:           "non-existent reference",
			ref:            "#/components/schemas/NonExistent",
			expectedSchema: nil,
			expectedName:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema, name := ResolveRef(tt.ref, schemas)

			if tt.expectedSchema == nil && schema != nil {
				t.Errorf("expected nil schema, got %v", schema)
			}
			if tt.expectedSchema != nil && schema == nil {
				t.Errorf("expected schema %v, got nil", tt.expectedSchema)
			}
			if name != tt.expectedName {
				t.Errorf("expected name %s, got %s", tt.expectedName, name)
			}
		})
	}
}

func TestGetResolvedPropertyType(t *testing.T) {
	schemas := map[string]interface{}{
		"User": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name         string
		prop         interface{}
		expectedType string
	}{
		{
			name: "inline property",
			prop: map[string]interface{}{
				"type": "string",
			},
			expectedType: "inline",
		},
		{
			name: "reference property",
			prop: map[string]interface{}{
				"$ref": "#/components/schemas/User",
			},
			expectedType: "ref",
		},
		{
			name:         "invalid property",
			prop:         "invalid",
			expectedType: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, propType := GetResolvedPropertyType(tt.prop, schemas)
			if propType != tt.expectedType {
				t.Errorf("expected type %s, got %s", tt.expectedType, propType)
			}
		})
	}
}

func TestGetPropertyStructure(t *testing.T) {
	tests := []struct {
		name     string
		prop     interface{}
		expected string
	}{
		{
			name: "string property",
			prop: map[string]interface{}{
				"type": "string",
			},
			expected: "string",
		},
		{
			name: "array with string items",
			prop: map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "string",
				},
			},
			expected: "array[string]",
		},
		{
			name: "object with properties",
			prop: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "string"},
				},
			},
			expected: "object{defined}",
		},
		{
			name: "object with additional properties",
			prop: map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
			},
			expected: "object{additional}",
		},
		{
			name: "reference",
			prop: map[string]interface{}{
				"$ref": "#/components/schemas/User",
			},
			expected: "$ref:#/components/schemas/User",
		},
		{
			name: "implicit object",
			prop: map[string]interface{}{
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "string"},
				},
			},
			expected: "implicit-object",
		},
		{
			name:     "invalid property",
			prop:     "invalid",
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPropertyStructure(tt.prop)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestGetArrayItemStructure(t *testing.T) {
	tests := []struct {
		name      string
		arrayProp map[string]interface{}
		expected  string
	}{
		{
			name: "array with string items",
			arrayProp: map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "string",
				},
			},
			expected: "string",
		},
		{
			name: "array without items",
			arrayProp: map[string]interface{}{
				"type": "array",
			},
			expected: "unknown",
		},
		{
			name:      "nil array",
			arrayProp: nil,
			expected:  "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetArrayItemStructure(tt.arrayProp)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHasTopLevelStructuralMismatch(t *testing.T) {
	schemas := map[string]interface{}{
		"User": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name        string
		entityProp  interface{}
		requestProp interface{}
		expected    bool
	}{
		{
			name: "matching string properties",
			entityProp: map[string]interface{}{
				"type": "string",
			},
			requestProp: map[string]interface{}{
				"type": "string",
			},
			expected: false,
		},
		{
			name: "mismatched types",
			entityProp: map[string]interface{}{
				"type": "string",
			},
			requestProp: map[string]interface{}{
				"type": "integer",
			},
			expected: true,
		},
		{
			name: "array item mismatch",
			entityProp: map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "string",
				},
			},
			requestProp: map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "integer",
				},
			},
			expected: true,
		},
		{
			name:        "nil properties",
			entityProp:  nil,
			requestProp: nil,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasTopLevelStructuralMismatch(tt.entityProp, tt.requestProp, schemas)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIsComputedField(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		expected  bool
	}{
		{
			name:      "created_at field",
			fieldName: "created_at",
			expected:  true,
		},
		{
			name:      "updated_by field",
			fieldName: "updated_by",
			expected:  true,
		},
		{
			name:      "is_editable field",
			fieldName: "is_editable",
			expected:  true,
		},
		{
			name:      "user_created_at field",
			fieldName: "user_created_at",
			expected:  true,
		},
		{
			name:      "regular field",
			fieldName: "name",
			expected:  false,
		},
		{
			name:      "case insensitive",
			fieldName: "CREATED_AT",
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsComputedField(tt.fieldName)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

// Benchmark tests for performance critical functions
func BenchmarkResolveRef(b *testing.B) {
	schemas := map[string]interface{}{
		"User": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{"type": "string"},
			},
		},
	}
	ref := "#/components/schemas/User"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ResolveRef(ref, schemas)
	}
}

func BenchmarkGetPropertyStructure(b *testing.B) {
	prop := map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{"type": "string"},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetPropertyStructure(prop)
	}
}
