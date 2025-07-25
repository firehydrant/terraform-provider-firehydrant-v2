// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// IncidentType model
type IncidentType struct {
	CreatedAt      *time.Time                          `json:"created_at,omitempty"`
	Description    *string                             `json:"description,omitempty"`
	ID             *string                             `json:"id,omitempty"`
	Name           *string                             `json:"name,omitempty"`
	Template       *NullableIncidentTypeTemplate       `json:"template,omitempty"`
	TemplateValues *NullableIncidentTypeTemplateValues `json:"template_values,omitempty"`
	UpdatedAt      *time.Time                          `json:"updated_at,omitempty"`
}

func (i IncidentType) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncidentType) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncidentType) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IncidentType) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentType) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentType) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IncidentType) GetTemplate() *NullableIncidentTypeTemplate {
	if o == nil {
		return nil
	}
	return o.Template
}

func (o *IncidentType) GetTemplateValues() *NullableIncidentTypeTemplateValues {
	if o == nil {
		return nil
	}
	return o.TemplateValues
}

func (o *IncidentType) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
