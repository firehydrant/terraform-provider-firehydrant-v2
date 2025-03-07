// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type PatchV1ScheduledMaintenancesScheduledMaintenanceIDStatusPages struct {
	// The slug identifying the type of status page
	IntegrationSlug *string `json:"integration_slug,omitempty"`
	// The UUID of the status page to display this maintenance on
	ConnectionID string `json:"connection_id"`
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDStatusPages) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDStatusPages) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts struct {
	// The type of impact
	Type string `json:"type"`
	// The id of impact
	ID string `json:"id"`
	// The id of the condition
	ConditionID string `json:"condition_id"`
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

// PatchV1ScheduledMaintenancesScheduledMaintenanceID - Change the conditions of a scheduled maintenance event, including updating any status page announcements of changes.
type PatchV1ScheduledMaintenancesScheduledMaintenanceID struct {
	Name    *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	// ISO8601 timestamp for the start time of the scheduled maintenance
	StartsAt *time.Time `json:"starts_at,omitempty"`
	// ISO8601 timestamp for the end time of the scheduled maintenance
	EndsAt      *time.Time `json:"ends_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	// A json object of label keys and values
	Labels map[string]string `json:"labels,omitempty"`
	// An array of status pages to display this maintenance on
	StatusPages []PatchV1ScheduledMaintenancesScheduledMaintenanceIDStatusPages `json:"status_pages,omitempty"`
	// An array of impact/condition combinations
	Impacts []PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts `json:"impacts,omitempty"`
}

func (p PatchV1ScheduledMaintenancesScheduledMaintenanceID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PatchV1ScheduledMaintenancesScheduledMaintenanceID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetStatusPages() []PatchV1ScheduledMaintenancesScheduledMaintenanceIDStatusPages {
	if o == nil {
		return nil
	}
	return o.StatusPages
}

func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceID) GetImpacts() []PatchV1ScheduledMaintenancesScheduledMaintenanceIDImpacts {
	if o == nil {
		return nil
	}
	return o.Impacts
}
