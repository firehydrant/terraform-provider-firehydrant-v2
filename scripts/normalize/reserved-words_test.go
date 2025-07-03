package main

import (
	"testing"
)

func TestNormalizeTerraformKeywords(t *testing.T) {
	tests := []struct {
		name     string
		schemas  map[string]interface{}
		expected int // number of conflicts expected
	}{
		{
			name: "single terraform keyword",
			schemas: map[string]interface{}{
				"Config": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"count": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"value": map[string]interface{}{"type": "integer"},
							},
						},
					},
				},
			},
			expected: 1,
		},
		{
			name: "multiple terraform keywords",
			schemas: map[string]interface{}{
				"TerraformConfig": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"count": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"value": map[string]interface{}{"type": "integer"},
							},
						},
						"provider": map[string]interface{}{
							"$ref": "#/components/schemas/Provider",
						},
						"variable": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"name": map[string]interface{}{"type": "string"},
							},
						},
					},
				},
			},
			expected: 3,
		},
		{
			name: "no terraform keywords",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"name": map[string]interface{}{
							"type": "string",
						},
						"email": map[string]interface{}{
							"type": "string",
						},
					},
				},
			},
			expected: 0,
		},
		{
			name: "terraform keyword but not object type",
			schemas: map[string]interface{}{
				"Config": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"count": map[string]interface{}{
							"type": "integer",
						},
					},
				},
			},
			expected: 0, // Should not be replaced because it's not an object
		},
		{
			name: "nested terraform keywords",
			schemas: map[string]interface{}{
				"NestedConfig": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"config": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"terraform": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"version": map[string]interface{}{"type": "string"},
									},
								},
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
			// Make a deep copy to avoid modifying the original
			schemasCopy := deepCopyMap(tt.schemas)
			conflicts := normalizeTerraformKeywords(schemasCopy)

			if len(conflicts) != tt.expected {
				t.Errorf("expected %d conflicts, got %d", tt.expected, len(conflicts))
			}

			// Verify that terraform keywords were replaced with empty objects
			if tt.expected > 0 {
				if !hasEmptyObjectReplacements(schemasCopy) {
					t.Error("expected terraform keywords to be replaced with empty objects")
				}
			}
		})
	}
}

func TestIsReservedKeyword(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		expected     bool
	}{
		{
			name:         "terraform keyword",
			propertyName: "terraform",
			expected:     true,
		},
		{
			name:         "count keyword",
			propertyName: "count",
			expected:     true,
		},
		{
			name:         "provider keyword",
			propertyName: "provider",
			expected:     true,
		},
		{
			name:         "case insensitive - uppercase",
			propertyName: "TERRAFORM",
			expected:     true,
		},
		{
			name:         "case insensitive - mixed case",
			propertyName: "Provider",
			expected:     true,
		},
		{
			name:         "non-reserved keyword",
			propertyName: "user",
			expected:     false,
		},
		{
			name:         "similar but not reserved",
			propertyName: "counter",
			expected:     false,
		},
		{
			name:         "empty string",
			propertyName: "",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isReservedKeyword(tt.propertyName)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestShouldReplaceProperty(t *testing.T) {
	tests := []struct {
		name      string
		propValue interface{}
		expected  bool
	}{
		{
			name: "object type property",
			propValue: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"value": map[string]interface{}{"type": "string"},
				},
			},
			expected: true,
		},
		{
			name: "reference property",
			propValue: map[string]interface{}{
				"$ref": "#/components/schemas/SomeSchema",
			},
			expected: true,
		},
		{
			name: "implicit object with properties",
			propValue: map[string]interface{}{
				"properties": map[string]interface{}{
					"name": map[string]interface{}{"type": "string"},
				},
			},
			expected: true,
		},
		{
			name: "string property",
			propValue: map[string]interface{}{
				"type": "string",
			},
			expected: false,
		},
		{
			name: "integer property",
			propValue: map[string]interface{}{
				"type": "integer",
			},
			expected: false,
		},
		{
			name: "array property",
			propValue: map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "string",
				},
			},
			expected: false,
		},
		{
			name:      "non-map value",
			propValue: "string",
			expected:  false,
		},
		{
			name:      "nil value",
			propValue: nil,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldReplaceProperty(tt.propValue)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestReplaceTerraformKeywordsInSchema(t *testing.T) {
	t.Run("replace simple terraform keyword", func(t *testing.T) {
		schema := map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"count": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"value": map[string]interface{}{"type": "integer"},
					},
				},
				"name": map[string]interface{}{
					"type": "string",
				},
			},
		}

		conflicts := replaceTerraformKeywordsInSchema("TestSchema", schema, "")

		if len(conflicts) != 1 {
			t.Errorf("expected 1 conflict, got %d", len(conflicts))
		}

		// Verify the count property was replaced
		properties := schema["properties"].(map[string]interface{})
		countProp := properties["count"].(map[string]interface{})

		if countProp["type"] != "object" {
			t.Error("expected count property to remain object type")
		}

		if props, ok := countProp["properties"].(map[string]interface{}); !ok || len(props) != 0 {
			t.Error("expected count property to have empty properties")
		}

		// Verify name property was not affected
		nameProp := properties["name"].(map[string]interface{})
		if nameProp["type"] != "string" {
			t.Error("expected name property to remain unchanged")
		}
	})

	t.Run("nested terraform keywords", func(t *testing.T) {
		schema := map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"config": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"terraform": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"version": map[string]interface{}{"type": "string"},
							},
						},
					},
				},
			},
		}

		conflicts := replaceTerraformKeywordsInSchema("TestSchema", schema, "")

		if len(conflicts) != 1 {
			t.Errorf("expected 1 conflict, got %d", len(conflicts))
		}

		// Verify nested terraform keyword was replaced
		configProp := schema["properties"].(map[string]interface{})["config"].(map[string]interface{})
		terraformProp := configProp["properties"].(map[string]interface{})["terraform"].(map[string]interface{})

		if props, ok := terraformProp["properties"].(map[string]interface{}); !ok || len(props) != 0 {
			t.Error("expected terraform property to have empty properties")
		}
	})
}

// Helper function to create a deep copy of a map for testing
func deepCopyMap(original map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})
	for key, value := range original {
		switch v := value.(type) {
		case map[string]interface{}:
			copy[key] = deepCopyMap(v)
		case []interface{}:
			copySlice := make([]interface{}, len(v))
			for i, item := range v {
				if itemMap, ok := item.(map[string]interface{}); ok {
					copySlice[i] = deepCopyMap(itemMap)
				} else {
					copySlice[i] = item
				}
			}
			copy[key] = copySlice
		default:
			copy[key] = value
		}
	}
	return copy
}

// Helper function to check if schemas have empty object replacements
func hasEmptyObjectReplacements(schemas map[string]interface{}) bool {
	for _, schema := range schemas {
		if hasEmptyObjectInSchema(schema) {
			return true
		}
	}
	return false
}

func hasEmptyObjectInSchema(obj interface{}) bool {
	switch v := obj.(type) {
	case map[string]interface{}:
		if properties, hasProps := v["properties"].(map[string]interface{}); hasProps {
			for _, prop := range properties {
				if propMap, ok := prop.(map[string]interface{}); ok {
					if propMap["type"] == "object" {
						if props, hasNestedProps := propMap["properties"].(map[string]interface{}); hasNestedProps && len(props) == 0 {
							return true
						}
					}
				}
			}
			// Check nested properties
			for _, prop := range properties {
				if hasEmptyObjectInSchema(prop) {
					return true
				}
			}
		}
		// Check other nested objects
		for key, value := range v {
			if key != "properties" && hasEmptyObjectInSchema(value) {
				return true
			}
		}
	case []interface{}:
		for _, item := range v {
			if hasEmptyObjectInSchema(item) {
				return true
			}
		}
	}
	return false
}
