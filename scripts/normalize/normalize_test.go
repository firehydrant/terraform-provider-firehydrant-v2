package main

import (
	"testing"
)

func TestNormalizeSpec(t *testing.T) {
	tests := []struct {
		name     string
		spec     map[string]interface{}
		expected int // expected total fixes
	}{
		{
			name: "complete spec with all normalizations",
			spec: map[string]interface{}{
				"components": map[string]interface{}{
					"schemas": map[string]interface{}{
						"User": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"status": map[string]interface{}{
									"type": "string",
									"enum": []interface{}{"active", "inactive"},
								},
								"count": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"value": map[string]interface{}{"type": "integer"},
									},
								},
							},
							"additionalProperties": true,
						},
					},
				},
				"paths": map[string]interface{}{
					"/users/{id}": map[string]interface{}{
						"get": map[string]interface{}{
							"parameters": []interface{}{
								map[string]interface{}{
									"name": "id",
									"in":   "path",
									"schema": map[string]interface{}{
										"type": "integer",
									},
								},
							},
						},
					},
				},
			},
			expected: 4, // terraform keyword + enum + additionalProperties + path parameter
		},
		{
			name: "spec with no components",
			spec: map[string]interface{}{
				"openapi": "3.0.0",
			},
			expected: 0,
		},
		{
			name: "spec with empty components",
			spec: map[string]interface{}{
				"components": map[string]interface{}{},
			},
			expected: 0,
		},
		{
			name: "spec with schemas but no paths",
			spec: map[string]interface{}{
				"components": map[string]interface{}{
					"schemas": map[string]interface{}{
						"User": map[string]interface{}{
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
			expected: 1, // only terraform keyword fix
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := normalizeSpec(tt.spec)

			if report.TotalFixes != tt.expected {
				t.Errorf("expected %d total fixes, got %d", tt.expected, report.TotalFixes)
			}

			if len(report.ConflictDetails) != tt.expected {
				t.Errorf("expected %d conflict details, got %d", tt.expected, len(report.ConflictDetails))
			}
		})
	}
}

func TestApplyGlobalNormalizations(t *testing.T) {
	tests := []struct {
		name     string
		schemas  map[string]interface{}
		expected int
	}{
		{
			name: "schema with additionalProperties",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type":                 "object",
					"additionalProperties": true,
					"properties": map[string]interface{}{
						"name": map[string]interface{}{"type": "string"},
					},
				},
			},
			expected: 1,
		},
		{
			name: "schema without additionalProperties",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"name": map[string]interface{}{"type": "string"},
					},
				},
			},
			expected: 0,
		},
		{
			name: "multiple schemas with additionalProperties",
			schemas: map[string]interface{}{
				"User": map[string]interface{}{
					"type":                 "object",
					"additionalProperties": true,
				},
				"Product": map[string]interface{}{
					"type":                 "object",
					"additionalProperties": map[string]interface{}{"type": "string"},
				},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conflicts := applyGlobalNormalizations(tt.schemas)

			if len(conflicts) != tt.expected {
				t.Errorf("expected %d conflicts, got %d", tt.expected, len(conflicts))
			}
		})
	}
}
