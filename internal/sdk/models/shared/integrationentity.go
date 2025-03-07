// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type IntegrationEntity struct {
	ID              *string    `json:"id,omitempty"`
	IntegrationName *string    `json:"integration_name,omitempty"`
	IntegrationSlug *string    `json:"integration_slug,omitempty"`
	DisplayName     *string    `json:"display_name,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
}

func (i IntegrationEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IntegrationEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IntegrationEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationEntity) GetIntegrationName() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationName
}

func (o *IntegrationEntity) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *IntegrationEntity) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *IntegrationEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}
