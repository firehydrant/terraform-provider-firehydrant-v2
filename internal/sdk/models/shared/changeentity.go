// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// ChangeEntityLabels - Arbitrary key/value pairs of labels.
type ChangeEntityLabels struct {
}

// ChangeEntity model
type ChangeEntity struct {
	// UUID of the Change
	ID *string `json:"id,omitempty"`
	// Description of the Change
	Summary *string `json:"summary,omitempty"`
	// The time the change entry was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Arbitrary key/value pairs of labels.
	Labels *ChangeEntityLabels `json:"labels,omitempty"`
	// Description of the Change
	Description *string `json:"description,omitempty"`
}

func (c ChangeEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChangeEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChangeEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *ChangeEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ChangeEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ChangeEntity) GetLabels() *ChangeEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ChangeEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
