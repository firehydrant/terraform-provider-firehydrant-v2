// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// TicketingProjectsProjectListItem - Ticketing_Projects_ProjectListItem model
type TicketingProjectsProjectListItem struct {
	Config         *NullableTicketingProjectConfig   `json:"config,omitempty"`
	ConnectionSlug *string                           `json:"connection_slug,omitempty"`
	FieldMap       *NullableTicketingProjectFieldMap `json:"field_map,omitempty"`
	ID             *string                           `json:"id,omitempty"`
	Name           *string                           `json:"name,omitempty"`
	UpdatedAt      *time.Time                        `json:"updated_at,omitempty"`
}

func (t TicketingProjectsProjectListItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TicketingProjectsProjectListItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TicketingProjectsProjectListItem) GetConfig() *NullableTicketingProjectConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *TicketingProjectsProjectListItem) GetConnectionSlug() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionSlug
}

func (o *TicketingProjectsProjectListItem) GetFieldMap() *NullableTicketingProjectFieldMap {
	if o == nil {
		return nil
	}
	return o.FieldMap
}

func (o *TicketingProjectsProjectListItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TicketingProjectsProjectListItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TicketingProjectsProjectListItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
