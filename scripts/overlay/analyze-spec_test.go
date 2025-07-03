// File: scripts/overlay/analyze_test.go
package main

import (
	"testing"
)

func TestAnalyzeSpec(t *testing.T) {
	spec := OpenAPISpec{
		Components: Components{
			Schemas: map[string]Schema{
				"UserEntity": {
					Type: "object",
					Properties: map[string]interface{}{
						"id":   map[string]interface{}{"type": "string"},
						"name": map[string]interface{}{"type": "string"},
					},
				},
				"CreateUserRequest": {
					Type: "object",
					Properties: map[string]interface{}{
						"name": map[string]interface{}{"type": "string"},
					},
				},
			},
		},
		Paths: map[string]PathItem{
			"/users": {
				Post: &Operation{
					OperationID: "createUser",
					RequestBody: map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CreateUserRequest",
								},
							},
						},
					},
					Responses: map[string]interface{}{
						"201": map[string]interface{}{
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/UserEntity",
									},
								},
							},
						},
					},
				},
			},
			"/users/{id}": {
				Get: &Operation{
					OperationID: "getUser",
					Responses: map[string]interface{}{
						"200": map[string]interface{}{
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/UserEntity",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	resources := analyzeSpec(spec, &ManualMappings{})

	if len(resources) != 1 {
		t.Errorf("expected 1 resource, got %d", len(resources))
	}

	if resource, exists := resources["user"]; exists {
		if resource.EntityName != "UserEntity" {
			t.Errorf("expected entity name 'UserEntity', got '%s'", resource.EntityName)
		}
		if resource.ResourceName != "user" {
			t.Errorf("expected resource name 'user', got '%s'", resource.ResourceName)
		}
		if len(resource.Operations) != 2 {
			t.Errorf("expected 2 operations, got %d", len(resource.Operations))
		}
	} else {
		t.Error("expected 'user' resource not found")
	}
}

func TestIdentifyEntitySchemas(t *testing.T) {
	schemas := map[string]Schema{
		"UserEntity": {
			Type: "object",
			Properties: map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"CreateUserRequest": {
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"ProductEntity": {
			Type: "object",
			Properties: map[string]interface{}{
				"slug": map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"NullableUser": {
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string"},
			},
		},
	}

	entities := identifyEntitySchemas(schemas)

	if len(entities) != 2 {
		t.Errorf("expected 2 entities, got %d", len(entities))
	}

	if !entities["UserEntity"] {
		t.Error("expected UserEntity to be identified as entity")
	}
	if !entities["ProductEntity"] {
		t.Error("expected ProductEntity to be identified as entity")
	}
	if entities["CreateUserRequest"] {
		t.Error("expected CreateUserRequest to NOT be identified as entity")
	}
	if entities["NullableUser"] {
		t.Error("expected NullableUser to NOT be identified as entity")
	}
}

func TestIsEntitySchema(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		expected bool
	}{
		{
			name: "valid entity with id",
			schema: Schema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":   map[string]interface{}{"type": "string"},
					"name": map[string]interface{}{"type": "string"},
				},
			},
			expected: true,
		},
		{
			name: "valid entity with slug",
			schema: Schema{
				Type: "object",
				Properties: map[string]interface{}{
					"slug": map[string]interface{}{"type": "string"},
					"name": map[string]interface{}{"type": "string"},
				},
			},
			expected: true,
		},
		{
			name: "invalid - no identifier",
			schema: Schema{
				Type: "object",
				Properties: map[string]interface{}{
					"name": map[string]interface{}{"type": "string"},
				},
			},
			expected: false,
		},
		{
			name: "invalid - not object type",
			schema: Schema{
				Type: "string",
			},
			expected: false,
		},
		{
			name: "invalid - no properties",
			schema: Schema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEntitySchema("TestEntity", tt.schema)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}

	// Test name-based exclusions
	excludedNames := []string{
		"CreateUserRequest",
		"UpdateUserRequest",
		"DeleteUserRequest",
		"UserResponse",
		"UserPaginated",
		"NullableUser",
	}

	validSchema := Schema{
		Type: "object",
		Properties: map[string]interface{}{
			"id": map[string]interface{}{"type": "string"},
		},
	}

	for _, name := range excludedNames {
		if isEntitySchema(name, validSchema) {
			t.Errorf("expected %s to be excluded as entity", name)
		}
	}
}

func TestDetermineCrudType(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		method      string
		operationID string
		expected    string
	}{
		{
			name:        "create by operation id",
			path:        "/users",
			method:      "post",
			operationID: "createUser",
			expected:    "create",
		},
		{
			name:        "update by operation id",
			path:        "/users/{id}",
			method:      "patch",
			operationID: "updateUser",
			expected:    "update",
		},
		{
			name:        "delete by operation id",
			path:        "/users/{id}",
			method:      "delete",
			operationID: "deleteUser",
			expected:    "delete",
		},
		{
			name:        "read by operation id",
			path:        "/users/{id}",
			method:      "get",
			operationID: "getUser",
			expected:    "read",
		},
		{
			name:        "list by operation id",
			path:        "/users",
			method:      "get",
			operationID: "listUsers",
			expected:    "list",
		},
		{
			name:        "create by method - post without path param",
			path:        "/users",
			method:      "post",
			operationID: "someOperation",
			expected:    "create",
		},
		{
			name:        "read by method - get with path param",
			path:        "/users/{id}",
			method:      "get",
			operationID: "someOperation",
			expected:    "read",
		},
		{
			name:        "list by method - get without path param",
			path:        "/users",
			method:      "get",
			operationID: "someOperation",
			expected:    "list",
		},
		{
			name:        "update by method - put with path param",
			path:        "/users/{id}",
			method:      "put",
			operationID: "someOperation",
			expected:    "update",
		},
		{
			name:        "create by method - put without path param",
			path:        "/users",
			method:      "put",
			operationID: "someOperation",
			expected:    "create",
		},
		{
			name:        "update by method - patch",
			path:        "/users/{id}",
			method:      "patch",
			operationID: "someOperation",
			expected:    "update",
		},
		{
			name:        "delete by method",
			path:        "/users/{id}",
			method:      "delete",
			operationID: "someOperation",
			expected:    "delete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := determineCrudType(tt.path, tt.method, tt.operationID)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestFindEntityFromOperation(t *testing.T) {
	entitySchemas := map[string]bool{
		"UserEntity":    true,
		"ProductEntity": true,
	}

	tests := []struct {
		name      string
		operation *Operation
		expected  string
	}{
		{
			name: "find entity from response ref",
			operation: &Operation{
				Responses: map[string]interface{}{
					"200": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/UserEntity",
								},
							},
						},
					},
				},
			},
			expected: "UserEntity",
		},
		{
			name: "find entity from response array",
			operation: &Operation{
				Responses: map[string]interface{}{
					"200": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"properties": map[string]interface{}{
										"data": map[string]interface{}{
											"type": "array",
											"items": map[string]interface{}{
												"$ref": "#/components/schemas/ProductEntity",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: "ProductEntity",
		},
		{
			name: "find entity from tags",
			operation: &Operation{
				Tags: []string{"User"},
			},
			expected: "UserEntity",
		},
		{
			name: "no entity found",
			operation: &Operation{
				Responses: map[string]interface{}{
					"200": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/UnknownEntity",
								},
							},
						},
					},
				},
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findEntityFromOperation(tt.operation, entitySchemas)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestExtractSchemaName(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		expected string
	}{
		{
			name:     "standard ref",
			ref:      "#/components/schemas/UserEntity",
			expected: "UserEntity",
		},
		{
			name:     "nested path",
			ref:      "#/definitions/nested/UserEntity",
			expected: "UserEntity",
		},
		{
			name:     "single segment",
			ref:      "UserEntity",
			expected: "UserEntity",
		},
		{
			name:     "empty ref",
			ref:      "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractSchemaName(tt.ref)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestDeriveResourceName(t *testing.T) {
	tests := []struct {
		name       string
		entityName string
		expected   string
	}{
		{
			name:       "simple entity",
			entityName: "UserEntity",
			expected:   "user",
		},
		{
			name:       "multi-word entity",
			entityName: "UserProfileEntity",
			expected:   "user_profile",
		},
		{
			name:       "with duplicate prefix",
			entityName: "IncidentsIncidentEntity",
			expected:   "incidents_incident", // Fixed expectation
		},
		{
			name:       "already snake case",
			entityName: "user_settingsEntity",
			expected:   "user_settings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deriveResourceName(tt.entityName)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "CamelCase",
			input:    "UserProfile",
			expected: "user_profile",
		},
		{
			name:     "PascalCase",
			input:    "XMLHttpRequest",
			expected: "xml_http_request", // Fixed expectation
		},
		{
			name:     "single word",
			input:    "User",
			expected: "user",
		},
		{
			name:     "already snake_case",
			input:    "user_profile",
			expected: "user_profile",
		},
		{
			name:     "with numbers",
			input:    "User123Profile",
			expected: "user123_profile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestExtractPathParameters(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "single parameter",
			path:     "/users/{id}",
			expected: []string{"id"},
		},
		{
			name:     "multiple parameters",
			path:     "/users/{user_id}/posts/{post_id}",
			expected: []string{"user_id", "post_id"},
		},
		{
			name:     "no parameters",
			path:     "/users",
			expected: []string{},
		},
		{
			name:     "complex path",
			path:     "/orgs/{org_id}/users/{user_id}/settings/{setting_name}",
			expected: []string{"org_id", "user_id", "setting_name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPathParameters(tt.path)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d parameters, got %d", len(tt.expected), len(result))
				return
			}
			for i, param := range result {
				if param != tt.expected[i] {
					t.Errorf("expected parameter %d to be %s, got %s", i, tt.expected[i], param)
				}
			}
		})
	}
}

func TestExtractRequiredFields(t *testing.T) {
	resources := map[string]*ResourceInfo{
		"user": {
			EntityName:   "UserEntity",
			CreateSchema: "CreateUserRequest",
		},
	}

	schemas := map[string]interface{}{
		"CreateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name":     map[string]interface{}{"type": "string"},
				"email":    map[string]interface{}{"type": "string"},
				"optional": map[string]interface{}{"type": "string"},
			},
			"required": []interface{}{"name", "email"},
		},
	}

	result := extractRequiredFields(resources, schemas)

	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}

	userFields := result["UserEntity"]
	if len(userFields) != 2 {
		t.Errorf("expected 2 required fields, got %d", len(userFields))
	}

	if !userFields["name"] {
		t.Error("expected 'name' to be required")
	}
	if !userFields["email"] {
		t.Error("expected 'email' to be required")
	}
	if userFields["optional"] {
		t.Error("expected 'optional' to NOT be required")
	}
}
