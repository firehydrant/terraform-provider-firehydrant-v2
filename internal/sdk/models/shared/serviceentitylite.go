// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// ServiceEntityLiteLabels - An object of label key and values
type ServiceEntityLiteLabels struct {
}

// ServiceEntityLite model
type ServiceEntityLite struct {
	ID            *string    `json:"id,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Slug          *string    `json:"slug,omitempty"`
	ServiceTier   *int       `json:"service_tier,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	AllowedParams []string   `json:"allowed_params,omitempty"`
	// An object of label key and values
	Labels                *ServiceEntityLiteLabels `json:"labels,omitempty"`
	AlertOnAdd            *bool                    `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam *bool                    `json:"auto_add_responding_team,omitempty"`
}

func (s ServiceEntityLite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceEntityLite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceEntityLite) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceEntityLite) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ServiceEntityLite) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ServiceEntityLite) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *ServiceEntityLite) GetServiceTier() *int {
	if o == nil {
		return nil
	}
	return o.ServiceTier
}

func (o *ServiceEntityLite) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ServiceEntityLite) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ServiceEntityLite) GetAllowedParams() []string {
	if o == nil {
		return nil
	}
	return o.AllowedParams
}

func (o *ServiceEntityLite) GetLabels() *ServiceEntityLiteLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ServiceEntityLite) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *ServiceEntityLite) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}
