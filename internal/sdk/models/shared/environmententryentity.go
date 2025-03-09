// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// EnvironmentEntryEntity model
type EnvironmentEntryEntity struct {
	// UUID of the Environment
	ID *string `json:"id,omitempty"`
	// Name of the Environment
	Name *string `json:"name,omitempty"`
	// Slug of the Environment
	Slug *string `json:"slug,omitempty"`
	// Description of the Environment
	Description *string `json:"description,omitempty"`
	// The time the environment was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The time the environment was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// List of active incident guids
	ActiveIncidents []string `json:"active_incidents,omitempty"`
	// Information about known linkages to representations of services outside of FireHydrant.
	ExternalResources []ExternalResourceEntity `json:"external_resources,omitempty"`
}

func (e EnvironmentEntryEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EnvironmentEntryEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EnvironmentEntryEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnvironmentEntryEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EnvironmentEntryEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *EnvironmentEntryEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EnvironmentEntryEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *EnvironmentEntryEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *EnvironmentEntryEntity) GetActiveIncidents() []string {
	if o == nil {
		return nil
	}
	return o.ActiveIncidents
}

func (o *EnvironmentEntryEntity) GetExternalResources() []ExternalResourceEntity {
	if o == nil {
		return nil
	}
	return o.ExternalResources
}
