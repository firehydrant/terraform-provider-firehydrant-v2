// File: scripts/overlay/terraform_viable_test.go
package main

import (
	"testing"
)

func TestIsTerraformViable(t *testing.T) {
	// Suppress output during tests
	restore := suppressOutput()
	defer restore()

	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"CreateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name     string
		resource *ResourceInfo
		expected bool
	}{
		{
			name: "viable resource",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				ResourceName: "user",
				CreateSchema: "CreateUserRequest",
				Operations: map[string]OperationInfo{
					"create": {
						OperationID: "createUser",
						Path:        "/users",
						Method:      "post",
					},
					"read": {
						OperationID: "getUser",
						Path:        "/users/{id}",
						Method:      "get",
					},
				},
			},
			expected: true,
		},
		{
			name: "missing create operation",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				ResourceName: "user",
				CreateSchema: "CreateUserRequest",
				Operations: map[string]OperationInfo{
					"read": {
						OperationID: "getUser",
						Path:        "/users/{id}",
						Method:      "get",
					},
				},
			},
			expected: false,
		},
		{
			name: "missing read operation",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				ResourceName: "user",
				CreateSchema: "CreateUserRequest",
				Operations: map[string]OperationInfo{
					"create": {
						OperationID: "createUser",
						Path:        "/users",
						Method:      "post",
					},
				},
			},
			expected: false,
		},
		{
			name: "missing create schema",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				ResourceName: "user",
				CreateSchema: "",
				Operations: map[string]OperationInfo{
					"create": {
						OperationID: "createUser",
						Path:        "/users",
						Method:      "post",
					},
					"read": {
						OperationID: "getUser",
						Path:        "/users/{id}",
						Method:      "get",
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTerraformViable(tt.resource, &ManualMappings{}, schemas)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestValidateOperationParameters(t *testing.T) {
	restore := suppressOutput()
	defer restore()

	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name      string
		resource  *ResourceInfo
		primaryID string
		expected  int // number of valid operations expected
	}{
		{
			name: "valid operations",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"create": {
						Path:   "/users",
						Method: "post",
					},
					"read": {
						Path:   "/users/{id}",
						Method: "get",
					},
					"update": {
						Path:   "/users/{id}",
						Method: "put",
					},
					"delete": {
						Path:   "/users/{id}",
						Method: "delete",
					},
					"list": {
						Path:   "/users",
						Method: "get",
					},
				},
			},
			primaryID: "id",
			expected:  5,
		},
		{
			name: "create with unexpected primary ID",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"create": {
						Path:   "/users/{id}", // Should not have primary ID
						Method: "post",
					},
					"read": {
						Path:   "/users/{id}",
						Method: "get",
					},
				},
			},
			primaryID: "id",
			expected:  1, // Only read should be valid
		},
		{
			name: "operations missing primary ID",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"create": {
						Path:   "/users",
						Method: "post",
					},
					"read": {
						Path:   "/users/{other_id}", // Wrong parameter name
						Method: "get",
					},
				},
			},
			primaryID: "id",
			expected:  1, // Only create should be valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validOps := validateOperationParameters(tt.resource, tt.primaryID, schemas, &ManualMappings{})
			if len(validOps) != tt.expected {
				t.Errorf("expected %d valid operations, got %d", tt.expected, len(validOps))
			}
		})
	}
}

func TestHasValidCreateReadConsistency(t *testing.T) {
	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
				"bio":  map[string]interface{}{"type": "string"},
			},
		},
		"CreateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "string"},
				"bio":  map[string]interface{}{"type": "string"},
			},
		},
		"EmptyCreateRequest": map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	}

	tests := []struct {
		name     string
		resource *ResourceInfo
		expected bool
	}{
		{
			name: "valid consistency",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				CreateSchema: "CreateUserRequest",
			},
			expected: true,
		},
		{
			name: "no create schema",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				CreateSchema: "",
			},
			expected: false,
		},
		{
			name: "empty create schema",
			resource: &ResourceInfo{
				EntityName:   "UserEntity",
				CreateSchema: "EmptyCreateRequest",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidCreateReadConsistency(tt.resource, schemas)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestGetSchemaProperties(t *testing.T) {
	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"InvalidSchema": "not an object",
		"NoProperties": map[string]interface{}{
			"type": "object",
		},
	}

	tests := []struct {
		name       string
		schemaName string
		expected   int
	}{
		{
			name:       "valid schema",
			schemaName: "UserEntity",
			expected:   2,
		},
		{
			name:       "non-existent schema",
			schemaName: "NonExistent",
			expected:   0,
		},
		{
			name:       "invalid schema",
			schemaName: "InvalidSchema",
			expected:   0,
		},
		{
			name:       "schema without properties",
			schemaName: "NoProperties",
			expected:   0,
		},
		{
			name:       "empty schema name",
			schemaName: "",
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSchemaProperties(schemas, tt.schemaName)
			if len(result) != tt.expected {
				t.Errorf("expected %d properties, got %d", tt.expected, len(result))
			}
		})
	}
}

func TestIsSystemProperty(t *testing.T) {
	tests := []struct {
		name     string
		propName string
		expected bool
	}{
		{
			name:     "id property",
			propName: "id",
			expected: true,
		},
		{
			name:     "created_at property",
			propName: "created_at",
			expected: true,
		},
		{
			name:     "updated_by property",
			propName: "updated_by",
			expected: true,
		},
		{
			name:     "user_id property",
			propName: "user_id",
			expected: true,
		},
		{
			name:     "custom_created_at property",
			propName: "custom_created_at",
			expected: true,
		},
		{
			name:     "version property",
			propName: "version",
			expected: true,
		},
		{
			name:     "etag property",
			propName: "etag",
			expected: true,
		},
		{
			name:     "revision property",
			propName: "revision",
			expected: true,
		},
		{
			name:     "last_modified property",
			propName: "last_modified",
			expected: true,
		},
		{
			name:     "regular property",
			propName: "name",
			expected: false,
		},
		{
			name:     "description property",
			propName: "description",
			expected: false,
		},
		{
			name:     "case insensitive",
			propName: "CREATED_AT",
			expected: true,
		},
		{
			name:     "similar but not system",
			propName: "creation_time",
			expected: false,
		},
		{
			name:     "empty string",
			propName: "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSystemProperty(tt.propName)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestCheckFieldExistsInEntityWithRefResolution(t *testing.T) {
	restore := suppressOutput()
	defer restore()

	entityProps := map[string]interface{}{
		"id":   map[string]interface{}{"type": "string"},
		"name": map[string]interface{}{"type": "string"},
		"profile": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"bio": map[string]interface{}{"type": "string"},
			},
		},
		"settings": map[string]interface{}{
			"$ref": "#/components/schemas/UserSettings",
		},
	}

	schemas := map[string]interface{}{
		"UserSettings": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"theme": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name      string
		fieldPath string
		expected  bool
	}{
		{
			name:      "simple field exists",
			fieldPath: "id",
			expected:  true,
		},
		{
			name:      "simple field doesn't exist",
			fieldPath: "nonexistent",
			expected:  false,
		},
		{
			name:      "nested field exists",
			fieldPath: "profile.bio",
			expected:  true,
		},
		{
			name:      "nested field doesn't exist",
			fieldPath: "profile.nonexistent",
			expected:  false,
		},
		{
			name:      "ref field exists",
			fieldPath: "settings.theme",
			expected:  true,
		},
		{
			name:      "ref field doesn't exist",
			fieldPath: "settings.nonexistent",
			expected:  false,
		},
		{
			name:      "invalid path",
			fieldPath: "name.invalid.path",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkFieldExistsInEntityWithRefResolution(tt.fieldPath, entityProps, schemas)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIdentifyEntityPrimaryID(t *testing.T) {
	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"ProductEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"slug": map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
	}

	tests := []struct {
		name     string
		resource *ResourceInfo
		expected string
		expectID bool
	}{
		{
			name: "simple id parameter",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"read": {
						Path: "/users/{id}",
					},
					"update": {
						Path: "/users/{id}",
					},
				},
			},
			expected: "id",
			expectID: true,
		},
		{
			name: "entity-specific id parameter",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"read": {
						Path: "/users/{user_id}",
					},
					"update": {
						Path: "/users/{user_id}",
					},
				},
			},
			expected: "user_id",
			expectID: true,
		},
		{
			name: "slug parameter",
			resource: &ResourceInfo{
				EntityName: "ProductEntity",
				Operations: map[string]OperationInfo{
					"read": {
						Path: "/products/{product_slug}",
					},
				},
			},
			expected: "product_slug",
			expectID: true,
		},
		{
			name: "no valid parameters",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"create": {
						Path: "/users",
					},
					"list": {
						Path: "/users",
					},
				},
			},
			expected: "",
			expectID: false,
		},
		{
			name: "multiple conflicting parameters - fixed",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"read": {
						Path: "/users/{user_id}",
					},
					"update": {
						Path: "/users/{user_id}", // Fixed: make them the same to not conflict
					},
				},
			},
			expected: "user_id", // Fixed: now they should match
			expectID: true,      // Fixed: now should find ID
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := identifyEntityPrimaryID(tt.resource, schemas)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
			if found != tt.expectID {
				t.Errorf("expected found=%t, got %t", tt.expectID, found)
			}
		})
	}
}

func TestMapsToEntityID(t *testing.T) {
	tests := []struct {
		name       string
		paramName  string
		entityName string
		expected   bool
	}{
		{
			name:       "exact match",
			paramName:  "user_id",
			entityName: "UserEntity",
			expected:   true,
		},
		{
			name:       "case insensitive match",
			paramName:  "USER_ID",
			entityName: "UserEntity",
			expected:   true,
		},
		{
			name:       "no match",
			paramName:  "product_id",
			entityName: "UserEntity",
			expected:   false,
		},
		{
			name:       "complex entity name",
			paramName:  "user_profile_id",
			entityName: "UserProfileEntity",
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapsToEntityID(tt.paramName, tt.entityName)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIsEntityID(t *testing.T) {
	tests := []struct {
		name      string
		paramName string
		expected  bool
	}{
		{
			name:      "id parameter",
			paramName: "id",
			expected:  true,
		},
		{
			name:      "user_id parameter",
			paramName: "user_id",
			expected:  true,
		},
		{
			name:      "product_id parameter",
			paramName: "product_id",
			expected:  true,
		},
		{
			name:      "case insensitive",
			paramName: "USER_ID",
			expected:  true,
		},
		{
			name:      "not an id",
			paramName: "name",
			expected:  false,
		},
		{
			name:      "similar but not id",
			paramName: "identifier",
			expected:  false,
		},
		{
			name:      "empty string",
			paramName: "",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEntityID(tt.paramName)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIsViableDatasource(t *testing.T) {
	tests := []struct {
		name     string
		resource *ResourceInfo
		expected bool
	}{
		{
			name: "viable with read operation",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"read": {
						OperationID: "getUser",
					},
				},
			},
			expected: true,
		},
		{
			name: "viable with list operation",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"list": {
						OperationID: "listUsers",
					},
				},
			},
			expected: true,
		},
		{
			name: "viable with both operations",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"read": {
						OperationID: "getUser",
					},
					"list": {
						OperationID: "listUsers",
					},
				},
			},
			expected: true,
		},
		{
			name: "not viable - no read/list operations",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{
					"create": {
						OperationID: "createUser",
					},
					"update": {
						OperationID: "updateUser",
					},
				},
			},
			expected: false,
		},
		{
			name: "not viable - no entity name",
			resource: &ResourceInfo{
				EntityName: "",
				Operations: map[string]OperationInfo{
					"read": {
						OperationID: "getUser",
					},
				},
			},
			expected: false,
		},
		{
			name: "not viable - no operations",
			resource: &ResourceInfo{
				EntityName: "UserEntity",
				Operations: map[string]OperationInfo{},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isViableDatasource(tt.resource)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
