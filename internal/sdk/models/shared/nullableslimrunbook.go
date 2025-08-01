// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type NullableSlimRunbook struct {
	AttachmentRule *NullableRules `json:"attachment_rule,omitempty"`
	// categories the runbook applies to
	Categories  []string          `json:"categories,omitempty"`
	CreatedAt   *time.Time        `json:"created_at,omitempty"`
	Description *string           `json:"description,omitempty"`
	ID          *string           `json:"id,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Owner       *NullableTeamLite `json:"owner,omitempty"`
	Summary     *string           `json:"summary,omitempty"`
	Type        *string           `json:"type,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
}

func (n NullableSlimRunbook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NullableSlimRunbook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NullableSlimRunbook) GetAttachmentRule() *NullableRules {
	if o == nil {
		return nil
	}
	return o.AttachmentRule
}

func (o *NullableSlimRunbook) GetCategories() []string {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *NullableSlimRunbook) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *NullableSlimRunbook) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *NullableSlimRunbook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableSlimRunbook) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NullableSlimRunbook) GetOwner() *NullableTeamLite {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *NullableSlimRunbook) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *NullableSlimRunbook) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NullableSlimRunbook) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
