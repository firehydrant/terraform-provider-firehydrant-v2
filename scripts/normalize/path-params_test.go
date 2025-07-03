package main

import (
	"testing"
)

func TestNormalizePathParameters(t *testing.T) {
	tests := []struct {
		name     string
		paths    map[string]interface{}
		expected int
	}{
		{
			name: "path with integer ID parameter",
			paths: map[string]interface{}{
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
			expected: 1,
		},
		{
			name: "path with string ID parameter",
			paths: map[string]interface{}{
				"/users/{id}": map[string]interface{}{
					"get": map[string]interface{}{
						"parameters": []interface{}{
							map[string]interface{}{
								"name": "id",
								"in":   "path",
								"schema": map[string]interface{}{
									"type": "string",
								},
							},
						},
					},
				},
			},
			expected: 0, // no change needed
		},
		{
			name: "multiple ID parameters",
			paths: map[string]interface{}{
				"/users/{user_id}/posts/{post_id}": map[string]interface{}{
					"get": map[string]interface{}{
						"parameters": []interface{}{
							map[string]interface{}{
								"name": "user_id",
								"in":   "path",
								"schema": map[string]interface{}{
									"type": "integer",
								},
							},
							map[string]interface{}{
								"name": "post_id",
								"in":   "path",
								"schema": map[string]interface{}{
									"type": "integer",
								},
							},
						},
					},
				},
			},
			expected: 2,
		},
		{
			name: "non-ID path parameter",
			paths: map[string]interface{}{
				"/users/{name}": map[string]interface{}{
					"get": map[string]interface{}{
						"parameters": []interface{}{
							map[string]interface{}{
								"name": "name",
								"in":   "path",
								"schema": map[string]interface{}{
									"type": "integer",
								},
							},
						},
					},
				},
			},
			expected: 0, // not an ID parameter
		},
		{
			name: "query parameter (not path)",
			paths: map[string]interface{}{
				"/users": map[string]interface{}{
					"get": map[string]interface{}{
						"parameters": []interface{}{
							map[string]interface{}{
								"name": "id",
								"in":   "query",
								"schema": map[string]interface{}{
									"type": "integer",
								},
							},
						},
					},
				},
			},
			expected: 0, // not a path parameter
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conflicts := normalizePathParameters(tt.paths)

			if len(conflicts) != tt.expected {
				t.Errorf("expected %d conflicts, got %d", tt.expected, len(conflicts))
			}

			// If we expected changes, verify they were applied
			if tt.expected > 0 {
				pathsVerified := 0
				for pathName, pathItem := range tt.paths {
					if pathMap, ok := pathItem.(map[string]interface{}); ok {
						for method, operation := range pathMap {
							if opMap, ok := operation.(map[string]interface{}); ok {
								if params, hasParams := opMap["parameters"]; hasParams {
									if paramsList, ok := params.([]interface{}); ok {
										for _, param := range paramsList {
											if paramMap, ok := param.(map[string]interface{}); ok {
												paramIn := paramMap["in"].(string)
												paramName := paramMap["name"].(string)
												if paramIn == "path" && (paramName == "id" || paramName == "user_id" || paramName == "post_id") {
													if schema, hasSchema := paramMap["schema"]; hasSchema {
														if schemaMap, ok := schema.(map[string]interface{}); ok {
															if schemaMap["type"] != "string" {
																t.Errorf("expected parameter %s to be converted to string in %s %s", paramName, method, pathName)
															} else {
																pathsVerified++
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
				if pathsVerified != tt.expected {
					t.Errorf("expected to verify %d parameter conversions, but verified %d", tt.expected, pathsVerified)
				}
			}
		})
	}
}
