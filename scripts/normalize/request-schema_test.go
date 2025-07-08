package main

import (
	"testing"
)

func TestNormalizeRequestSchemasWithPaths(t *testing.T) {
	t.Run("extracts inline request schema", func(t *testing.T) {
		paths := map[string]interface{}{
			"/v1/test": map[string]interface{}{
				"post": map[string]interface{}{
					"operationId": "create_test",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"name": map[string]interface{}{
											"type": "string",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		schemas := map[string]interface{}{}

		conflicts := normalizeRequestSchemasWithPaths(paths, schemas)

		// Should have one conflict
		if len(conflicts) != 1 {
			t.Errorf("expected 1 conflict, got %d", len(conflicts))
		}

		// Should create the schema
		if _, exists := schemas["create_test"]; !exists {
			t.Error("expected create_test schema to be created")
		}

		// Should replace with $ref
		postOp := paths["/v1/test"].(map[string]interface{})["post"].(map[string]interface{})
		requestBody := postOp["requestBody"].(map[string]interface{})
		content := requestBody["content"].(map[string]interface{})
		jsonContent := content["application/json"].(map[string]interface{})
		schema := jsonContent["schema"].(map[string]interface{})

		if ref, hasRef := schema["$ref"]; !hasRef || ref != "#/components/schemas/create_test" {
			t.Errorf("expected $ref to be #/components/schemas/create_test, got %v", ref)
		}
	})

	t.Run("ignores schemas with existing $ref", func(t *testing.T) {
		paths := map[string]interface{}{
			"/v1/test": map[string]interface{}{
				"post": map[string]interface{}{
					"operationId": "create_test",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/ExistingSchema",
								},
							},
						},
					},
				},
			},
		}
		schemas := map[string]interface{}{}

		conflicts := normalizeRequestSchemasWithPaths(paths, schemas)

		if len(conflicts) != 0 {
			t.Errorf("expected 0 conflicts, got %d", len(conflicts))
		}
	})

	t.Run("handles multiple inline schemas", func(t *testing.T) {
		paths := map[string]interface{}{
			"/v1/users": map[string]interface{}{
				"post": map[string]interface{}{
					"operationId": "create_user",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
								},
							},
						},
					},
				},
				"put": map[string]interface{}{
					"operationId": "update_user",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
								},
							},
						},
					},
				},
			},
		}
		schemas := map[string]interface{}{}

		conflicts := normalizeRequestSchemasWithPaths(paths, schemas)

		if len(conflicts) != 2 {
			t.Errorf("expected 2 conflicts, got %d", len(conflicts))
		}

		if len(schemas) != 2 {
			t.Errorf("expected 2 schemas, got %d", len(schemas))
		}
	})

	t.Run("handles no inline schemas", func(t *testing.T) {
		paths := map[string]interface{}{
			"/v1/users": map[string]interface{}{
				"get": map[string]interface{}{
					"operationId": "list_users",
				},
			},
		}
		schemas := map[string]interface{}{}

		conflicts := normalizeRequestSchemasWithPaths(paths, schemas)

		if len(conflicts) != 0 {
			t.Errorf("expected 0 conflicts, got %d", len(conflicts))
		}
	})
}

func TestGenerateRequestSchemaName(t *testing.T) {
	tests := []struct {
		name        string
		operationId string
		contentType string
		expected    string
	}{
		{
			name:        "uses operation id with json content",
			operationId: "create_user",
			contentType: "application/json",
			expected:    "create_user",
		},
		{
			name:        "adds suffix for form data",
			operationId: "upload_file",
			contentType: "multipart/form-data",
			expected:    "upload_file_form",
		},
		{
			name:        "adds suffix for url encoded",
			operationId: "submit_form",
			contentType: "application/x-www-form-urlencoded",
			expected:    "submit_form_form_encoded",
		},
		{
			name:        "adds suffix for xml",
			operationId: "update_data",
			contentType: "application/xml",
			expected:    "update_data_xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRequestSchemaName(tt.operationId, tt.contentType)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestEnsureUniqueSchemaName(t *testing.T) {
	t.Run("returns base name when unique", func(t *testing.T) {
		schemas := map[string]interface{}{}
		result := ensureUniqueSchemaName("create_user", schemas)
		if result != "create_user" {
			t.Errorf("expected create_user, got %s", result)
		}
	})

	t.Run("adds suffix when name exists", func(t *testing.T) {
		schemas := map[string]interface{}{
			"create_user": map[string]interface{}{},
		}
		result := ensureUniqueSchemaName("create_user", schemas)
		if result != "create_user_1" {
			t.Errorf("expected create_user_1, got %s", result)
		}
	})

	t.Run("finds next available number", func(t *testing.T) {
		schemas := map[string]interface{}{
			"create_user":   map[string]interface{}{},
			"create_user_1": map[string]interface{}{},
			"create_user_2": map[string]interface{}{},
		}
		result := ensureUniqueSchemaName("create_user", schemas)
		if result != "create_user_3" {
			t.Errorf("expected create_user_3, got %s", result)
		}
	})
}
