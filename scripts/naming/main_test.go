package main

import (
	"testing"
)

func TestNormalizeOverlayTarget(t *testing.T) {
	schemaRenames := map[string]string{
		"IncidentTypeEntity":   "IncidentType",
		"NuncConnectionEntity": "NuncConnection",
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"$.components.schemas.IncidentTypeEntity.properties.name", "$.components.schemas.IncidentType.properties.name"},
		{"$.components.schemas.NuncConnectionEntity", "$.components.schemas.NuncConnection"},
		{"$.components.schemas.SomeNewEntity.properties.field", "$.components.schemas.SomeNew.properties.field"},
		{"$.components.schemas.UserIdentityEntity.properties.field", "$.components.schemas.UserIdentity.properties.field"},
		{"$.components.schemas.UserIdentity.properties.field", "$.components.schemas.UserIdentity.properties.field"},
		{"$.components.schemas.ChangeIdentityEntity.properties.field", "$.components.schemas.ChangeIdentity.properties.field"},
		{"$.paths[\"/v1/incidents\"].get", "$.paths[\"/v1/incidents\"].get"},
		{"$.components.schemas.SomeOtherSchema.properties.field", "$.components.schemas.SomeOtherSchema.properties.field"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeOverlayTarget(tt.input, schemaRenames)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestNormalizeOverlayUpdate(t *testing.T) {
	tests := []struct {
		name           string
		update         map[string]interface{}
		expectedFixes  int
		expectedEntity string
	}{
		{
			name: "x-speakeasy-entity normalization",
			update: map[string]interface{}{
				"x-speakeasy-entity": "IncidentTypeEntity",
			},
			expectedFixes:  1,
			expectedEntity: "IncidentType",
		},
		{
			name: "no entity references",
			update: map[string]interface{}{
				"additionalProperties": true,
			},
			expectedFixes:  0,
			expectedEntity: "",
		},
		{
			name: "identity with entity should normalize",
			update: map[string]interface{}{
				"x-speakeasy-entity": "ChangeIdentityEntity",
			},
			expectedFixes:  1,
			expectedEntity: "ChangeIdentity",
		},
		{
			name: "identity without entity should not change",
			update: map[string]interface{}{
				"x-speakeasy-entity": "UserIdentity",
			},
			expectedFixes:  0,
			expectedEntity: "UserIdentity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateCopy := make(map[string]interface{})
			for k, v := range tt.update {
				updateCopy[k] = v
			}

			fixes := normalizeOverlayUpdate(updateCopy)

			if fixes != tt.expectedFixes {
				t.Errorf("expected %d fixes, got %d", tt.expectedFixes, fixes)
			}

			if tt.expectedEntity != "" {
				if actual, ok := updateCopy["x-speakeasy-entity"].(string); !ok || actual != tt.expectedEntity {
					t.Errorf("expected entity %s, got %s", tt.expectedEntity, actual)
				}
			}
		})
	}
}

func TestNormalizeEntityOperation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"IncidentTypeEntity#read", "IncidentType#read"},
		{"Webhooks_Entities_WebhookEntity#list", "Webhooks#list"},
		{"NullableIncidentTypeEntity#read", "NullableIncidentType#read"},
		{"IncidentTypeEntityPaginated#list", "IncidentTypePaginated#list"},
		{"ChangeIdentityEntity#read", "ChangeIdentity#read"},
		{"UserIdentity#read", "UserIdentity#read"},
		{"IncidentType#update", "IncidentType#update"},
		{"InvalidFormat", "InvalidFormat"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeEntityOperation(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestNormalizeEntityName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic entity removal
		{"UserEntity", "User"},
		{"Ticketing_ProjectFieldMapEntity", "Ticketing_ProjectFieldMap"},
		{"RunbooksAction_entity", "RunbooksAction"},

		// Complex patterns
		{"Webhooks_Entities_WebhookEntity", "Webhooks"},
		{"ServiceEntityLite", "ServiceLite"},
		{"NullableIncidentTypeEntity", "NullableIncidentType"},
		{"IncidentTypeEntityPaginated", "IncidentTypePaginated"},

		// Identity cases
		{"ChangeIdentityEntity", "ChangeIdentity"},
		{"ChangeIdentityEntityPaginated", "ChangeIdentityPaginated"},
		{"UserIdentityEntity", "UserIdentity"},
		{"UserIdentity", "UserIdentity"},
		{"Identity", "Identity"},
		{"IdentityProvider", "IdentityProvider"},

		// No change cases
		{"IncidentType", "IncidentType"},
		{"SimpleSchema", "SimpleSchema"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeEntityName(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestRemoveEmbeddedEntityReferences(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"UserEntity", "User"},
		{"Signals_API_GroupingEntity_Strategy", "Signals_API_Grouping_Strategy"},
		{"TeamEntityLite", "TeamLite"},
		{"ServiceEntityChecklist", "ServiceChecklist"},
		{"RunbooksAction_entity", "RunbooksAction"},
		{"UserIdentityEntity", "UserIdentity"},
		{"ChangeIdentityEntity", "ChangeIdentity"},
		{"ChangeIdentityEntityPaginated", "ChangeIdentityPaginated"},
		{"ServiceIdentity", "ServiceIdentity"},
		{"UserIdentity", "UserIdentity"},
		{"Identity", "Identity"},
		{"IdentityProvider", "IdentityProvider"},
		{"SimpleSchema", "SimpleSchema"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := removeEmbeddedEntityReferences(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestIsIdentityOnlySchema(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Should be protected (identity without entity patterns)
		{"Identity", true},
		{"IdentityProvider", true},
		{"UserIdentity", true},
		{"IdentityToken", true},

		// Should NOT be protected (identity with entity patterns)
		{"ChangeIdentityEntity", false},
		{"ChangeIdentityEntityPaginated", false},
		{"UserIdentityEntity", false},
		{"IdentityEntities", false},

		// Should NOT be protected (no identity)
		{"UserEntity", false},
		{"SimpleSchema", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isIdentityOnlySchema(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
