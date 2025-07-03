// File: scripts/overlay/detect_test.go
package main

import (
	"testing"
)

func TestDetectPropertyMismatches(t *testing.T) {
	resources := map[string]*ResourceInfo{
		"user": {
			EntityName:   "UserEntity",
			CreateSchema: "CreateUserRequest",
			UpdateSchema: "UpdateUserRequest",
		},
	}

	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type": "string",
				},
				"name": map[string]interface{}{
					"type": "string",
				},
				"profile": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"bio": map[string]interface{}{"type": "string"},
					},
				},
			},
			"__spec": map[string]interface{}{
				"components": map[string]interface{}{
					"schemas": map[string]interface{}{},
				},
			},
		},
		"CreateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"profile": map[string]interface{}{
					"type": "array", // Structural mismatch: object vs array
					"items": map[string]interface{}{
						"type": "string",
					},
				},
			},
			"required": []interface{}{"name"},
		},
		"UpdateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

	requiredFields := map[string]map[string]bool{
		"UserEntity": {
			"name": true,
		},
	}

	mismatches := detectPropertyMismatches(resources, schemas, requiredFields)

	if len(mismatches) != 1 {
		t.Errorf("expected 1 resource with mismatches, got %d", len(mismatches))
	}

	userMismatches := mismatches["UserEntity"]
	if len(userMismatches) != 1 {
		t.Errorf("expected 1 mismatch for UserEntity, got %d", len(userMismatches))
	}

	if userMismatches[0].PropertyName != "profile" {
		t.Errorf("expected mismatch on 'profile', got '%s'", userMismatches[0].PropertyName)
	}

	if userMismatches[0].MismatchType != "structural-mismatch" {
		t.Errorf("expected 'structural-mismatch', got '%s'", userMismatches[0].MismatchType)
	}
}

func TestFindPropertyMismatches(t *testing.T) {
	entitySchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"age": map[string]interface{}{
				"type": "integer",
			},
			"profile": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"bio": map[string]interface{}{"type": "string"},
				},
			},
		},
		"__spec": map[string]interface{}{
			"components": map[string]interface{}{
				"schemas": map[string]interface{}{},
			},
		},
	}

	requestSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string", // Same type - no mismatch
			},
			"age": map[string]interface{}{
				"type": "string", // Different type - mismatch
			},
			"profile": map[string]interface{}{
				"type": "array", // Different structure - mismatch
				"items": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

	requiredFields := map[string]bool{
		"name": true, // This should be skipped
	}

	mismatches := findPropertyMismatches(entitySchema, requestSchema, "create", requiredFields)

	if len(mismatches) != 2 {
		t.Errorf("expected 2 mismatches, got %d", len(mismatches))
	}

	// Check that required field "name" was skipped
	for _, mismatch := range mismatches {
		if mismatch.PropertyName == "name" {
			t.Error("expected required field 'name' to be skipped")
		}
	}

	// Check that we found the expected mismatches
	foundAge := false
	foundProfile := false
	for _, mismatch := range mismatches {
		if mismatch.PropertyName == "age" {
			foundAge = true
		}
		if mismatch.PropertyName == "profile" {
			foundProfile = true
		}
	}

	if !foundAge {
		t.Error("expected to find mismatch on 'age' property")
	}
	if !foundProfile {
		t.Error("expected to find mismatch on 'profile' property")
	}
}

func TestDetectCRUDInconsistencies(t *testing.T) {
	resources := map[string]*ResourceInfo{
		"user": {
			EntityName:   "UserEntity",
			CreateSchema: "CreateUserRequest",
			UpdateSchema: "UpdateUserRequest",
			Operations: map[string]OperationInfo{
				"create": {OperationID: "createUser"},
				"read":   {OperationID: "getUser"},
			},
		},
		"product": {
			EntityName:   "ProductEntity",
			CreateSchema: "CreateProductRequest",
			Operations: map[string]OperationInfo{
				"create": {OperationID: "createProduct"},
				"read":   {OperationID: "getProduct"},
			},
		},
	}

	schemas := map[string]interface{}{
		"UserEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
				"bio":  map[string]interface{}{"type": "string"}, // Only in entity
			},
		},
		"CreateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name":  map[string]interface{}{"type": "string"},
				"email": map[string]interface{}{"type": "string"}, // Only in create
			},
			"required": []interface{}{"name"},
		},
		"UpdateUserRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "string"},
				"age":  map[string]interface{}{"type": "integer"}, // Only in update
			},
		},
		"ProductEntity": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "string"},
				"name": map[string]interface{}{"type": "string"},
			},
		},
		"CreateProductRequest": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "string"},
			},
		},
	}

	inconsistencies := detectCRUDInconsistencies(resources, schemas)

	if len(inconsistencies) != 1 {
		t.Errorf("expected 1 resource with inconsistencies, got %d", len(inconsistencies))
	}

	userInconsistencies := inconsistencies["UserEntity"]
	if len(userInconsistencies) == 0 {
		t.Error("expected inconsistencies for UserEntity")
	}

	// Should find inconsistencies for properties that aren't in all CRUD operations
	expectedInconsistencies := []string{"bio", "email", "age"}
	foundInconsistencies := make(map[string]bool)

	for _, inconsistency := range userInconsistencies {
		foundInconsistencies[inconsistency.PropertyName] = true
	}

	for _, expected := range expectedInconsistencies {
		if !foundInconsistencies[expected] {
			t.Errorf("expected to find inconsistency for property '%s'", expected)
		}
	}
}

func TestDetectSchemaPropertyInconsistencies(t *testing.T) {
	// Test valid resource
	t.Run("valid resource", func(t *testing.T) {
		resource := &ResourceInfo{
			EntityName:   "UserEntity",
			CreateSchema: "CreateUserRequest",
			Operations: map[string]OperationInfo{
				"create": {OperationID: "createUser"},
				"read":   {OperationID: "getUser"},
			},
		}

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

		inconsistencies := detectSchemaPropertyInconsistencies(resource, schemas, map[string]bool{})

		// Should not have validation errors
		for _, inconsistency := range inconsistencies {
			if inconsistency.PropertyName == "RESOURCE_VALIDATION" {
				t.Errorf("unexpected validation error: %s", inconsistency.Description)
			}
		}
	})

	// Test missing operations
	t.Run("missing required operations", func(t *testing.T) {
		resource := &ResourceInfo{
			EntityName:   "UserEntity",
			CreateSchema: "CreateUserRequest",
			Operations: map[string]OperationInfo{
				"create": {OperationID: "createUser"},
				// Missing read operation
			},
		}

		schemas := map[string]interface{}{
			"UserEntity": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "string"},
				},
			},
			"CreateUserRequest": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{"type": "string"},
				},
			},
		}

		inconsistencies := detectSchemaPropertyInconsistencies(resource, schemas, map[string]bool{})

		if len(inconsistencies) != 1 {
			t.Errorf("expected 1 validation error, got %d", len(inconsistencies))
		}

		if inconsistencies[0].PropertyName != "RESOURCE_VALIDATION" {
			t.Error("expected RESOURCE_VALIDATION error")
		}

		if inconsistencies[0].InconsistencyType != "missing-required-operations" {
			t.Errorf("expected 'missing-required-operations', got '%s'", inconsistencies[0].InconsistencyType)
		}
	})

	// Test missing create schema
	t.Run("missing create schema", func(t *testing.T) {
		resource := &ResourceInfo{
			EntityName:   "UserEntity",
			CreateSchema: "", // Missing create schema
			Operations: map[string]OperationInfo{
				"create": {OperationID: "createUser"},
				"read":   {OperationID: "getUser"},
			},
		}

		schemas := map[string]interface{}{
			"UserEntity": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "string"},
				},
			},
		}

		inconsistencies := detectSchemaPropertyInconsistencies(resource, schemas, map[string]bool{})

		if len(inconsistencies) != 1 {
			t.Errorf("expected 1 validation error, got %d", len(inconsistencies))
		}

		if inconsistencies[0].InconsistencyType != "missing-create-schema" {
			t.Errorf("expected 'missing-create-schema', got '%s'", inconsistencies[0].InconsistencyType)
		}
	})
}

func TestDetectReadonlyFields(t *testing.T) {
	entitySchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type": "string",
			},
			"name": map[string]interface{}{
				"type": "string",
			},
			"readonly_field": map[string]interface{}{
				"type": "string",
			},
			"mismatch_field": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"value": map[string]interface{}{"type": "string"},
				},
			},
		},
	}

	createSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"mismatch_field": map[string]interface{}{
				"type": "array", // Structural mismatch with entity
				"items": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

	schemas := map[string]interface{}{}

	actions := detectReadonlyFields("UserEntity", entitySchema, createSchema, nil, schemas)

	if len(actions) != 2 {
		t.Errorf("expected 2 readonly actions, got %d", len(actions))
	}

	// Check that readonly_field was marked as readonly (not in create schema)
	foundReadonlyField := false
	foundMismatchField := false

	for _, action := range actions {
		if action.Target == "$.components.schemas.UserEntity.properties.readonly_field" {
			foundReadonlyField = true
			if readonly, ok := action.Update["x-speakeasy-param-readonly"]; !ok || readonly != true {
				t.Error("expected readonly_field to be marked as readonly")
			}
		}
		if action.Target == "$.components.schemas.UserEntity.properties.mismatch_field" {
			foundMismatchField = true
			if readonly, ok := action.Update["x-speakeasy-param-readonly"]; !ok || readonly != true {
				t.Error("expected mismatch_field to be marked as readonly")
			}
		}
	}

	if !foundReadonlyField {
		t.Error("expected to find readonly action for readonly_field")
	}
	if !foundMismatchField {
		t.Error("expected to find readonly action for mismatch_field")
	}
}

func TestDetectNestedReadonlyFields(t *testing.T) {
	entityProp := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{"type": "string"},
			"field2": map[string]interface{}{"type": "string"},
		},
	}

	createProp := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{"type": "string"},
			// field2 missing in create - should be marked readonly
		},
	}

	schemas := map[string]interface{}{}

	actions := detectNestedReadonlyFields("UserEntity", "metadata", entityProp, createProp, nil, schemas)

	if len(actions) != 1 {
		t.Errorf("expected 1 nested readonly action, got %d", len(actions))
	}

	if actions[0].Target != "$.components.schemas.UserEntity.properties.metadata.properties.field2" {
		t.Errorf("expected nested readonly target, got %s", actions[0].Target)
	}

	if readonly, ok := actions[0].Update["x-speakeasy-param-readonly"]; !ok || readonly != true {
		t.Error("expected nested field to be marked as readonly")
	}
}

func TestDescribeStructuralDifference(t *testing.T) {
	entityProp := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"value": map[string]interface{}{"type": "string"},
		},
	}

	requestProp := map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"type": "string",
		},
	}

	description := describeStructuralDifference(entityProp, requestProp)

	expected := "request structure 'array[string]' != response structure 'object{defined}'"
	if description != expected {
		t.Errorf("expected '%s', got '%s'", expected, description)
	}
}
