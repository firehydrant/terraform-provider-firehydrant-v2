package main

import (
	"testing"
)

func TestNormalizeAdditionalProperties(t *testing.T) {
	tests := []struct {
		name       string
		schemaName string
		obj        interface{}
		path       string
		expected   int
	}{
		{
			name:       "object with additionalProperties true",
			schemaName: "User",
			obj: map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
				"properties": map[string]interface{}{
					"name": map[string]interface{}{"type": "string"},
				},
			},
			path:     "",
			expected: 1,
		},
		{
			name:       "object with additionalProperties schema",
			schemaName: "User",
			obj: map[string]interface{}{
				"type":                 "object",
				"additionalProperties": map[string]interface{}{"type": "string"},
			},
			path:     "",
			expected: 1,
		},
		{
			name:       "object without additionalProperties",
			schemaName: "User",
			obj: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{"type": "string"},
				},
			},
			path:     "",
			expected: 0,
		},
		{
			name:       "nested object with additionalProperties",
			schemaName: "User",
			obj: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"metadata": map[string]interface{}{
						"type":                 "object",
						"additionalProperties": true,
					},
				},
			},
			path:     "",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conflicts := normalizeAdditionalProperties(tt.schemaName, tt.obj, tt.path)

			if len(conflicts) != tt.expected {
				t.Errorf("expected %d conflicts, got %d", tt.expected, len(conflicts))
			}

			// Verify additionalProperties was removed
			if tt.expected > 0 {
				if hasAdditionalProperties(tt.obj) {
					t.Error("expected additionalProperties to be removed")
				}
			}
		})
	}
}

// Helper function to check if an object still has additionalProperties
func hasAdditionalProperties(obj interface{}) bool {
	switch v := obj.(type) {
	case map[string]interface{}:
		if _, hasAdditional := v["additionalProperties"]; hasAdditional {
			return true
		}
		for _, value := range v {
			if hasAdditionalProperties(value) {
				return true
			}
		}
	case []interface{}:
		for _, item := range v {
			if hasAdditionalProperties(item) {
				return true
			}
		}
	}
	return false
}
