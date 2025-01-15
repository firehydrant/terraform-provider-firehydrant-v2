// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type LifecyclesMilestoneEntity struct {
	ID          *string       `json:"id,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Slug        *string       `json:"slug,omitempty"`
	Position    *int          `json:"position,omitempty"`
	CreatedBy   *AuthorEntity `json:"created_by,omitempty"`
	UpdatedBy   *AuthorEntity `json:"updated_by,omitempty"`
	CreatedAt   *time.Time    `json:"created_at,omitempty"`
	UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
}

func (l LifecyclesMilestoneEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LifecyclesMilestoneEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LifecyclesMilestoneEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LifecyclesMilestoneEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LifecyclesMilestoneEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LifecyclesMilestoneEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *LifecyclesMilestoneEntity) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *LifecyclesMilestoneEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *LifecyclesMilestoneEntity) GetUpdatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *LifecyclesMilestoneEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LifecyclesMilestoneEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
