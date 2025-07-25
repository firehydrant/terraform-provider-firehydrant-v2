// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// Audiences model
type Audiences struct {
	// When the audience was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Whether this is the organization's default audience
	Default *bool `json:"default,omitempty"`
	// Description of the audience and its purpose (maximum 4000 characters)
	Description *string `json:"description,omitempty"`
	// List of incident details for this audience
	Details []AudiencesDetail `json:"details,omitempty"`
	// When the audience was discarded (soft deleted)
	DiscardedAt *time.Time `json:"discarded_at,omitempty"`
	// Unique identifier for the audience
	ID *string `json:"id,omitempty"`
	// Name of the audience (maximum 255 characters)
	Name *string `json:"name,omitempty"`
	// Slug of the audience, unique and autogenerated
	Slug *string `json:"slug,omitempty"`
	// When the audience was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (a Audiences) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Audiences) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Audiences) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Audiences) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *Audiences) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Audiences) GetDetails() []AudiencesDetail {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Audiences) GetDiscardedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DiscardedAt
}

func (o *Audiences) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Audiences) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Audiences) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Audiences) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
