// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// AlertsSirenAlertEntityLabels - Arbitrary key:value pairs of labels.
type AlertsSirenAlertEntityLabels struct {
}

type AlertsSirenAlertEntity struct {
	ID          *string    `json:"id,omitempty"`
	Summary     *string    `json:"summary,omitempty"`
	Description *string    `json:"description,omitempty"`
	StartsAt    *time.Time `json:"starts_at,omitempty"`
	EndsAt      *time.Time `json:"ends_at,omitempty"`
	Status      *string    `json:"status,omitempty"`
	RemoteID    *string    `json:"remote_id,omitempty"`
	RemoteURL   *string    `json:"remote_url,omitempty"`
	// Arbitrary key:value pairs of labels.
	Labels     *AlertsSirenAlertEntityLabels `json:"labels,omitempty"`
	Tags       []string                      `json:"tags,omitempty"`
	SignalID   *string                       `json:"signal_id,omitempty"`
	SignalRule *SignalsAPIRuleEntity         `json:"signal_rule,omitempty"`
}

func (a AlertsSirenAlertEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlertsSirenAlertEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AlertsSirenAlertEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AlertsSirenAlertEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *AlertsSirenAlertEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AlertsSirenAlertEntity) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *AlertsSirenAlertEntity) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *AlertsSirenAlertEntity) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AlertsSirenAlertEntity) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *AlertsSirenAlertEntity) GetRemoteURL() *string {
	if o == nil {
		return nil
	}
	return o.RemoteURL
}

func (o *AlertsSirenAlertEntity) GetLabels() *AlertsSirenAlertEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *AlertsSirenAlertEntity) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AlertsSirenAlertEntity) GetSignalID() *string {
	if o == nil {
		return nil
	}
	return o.SignalID
}

func (o *AlertsSirenAlertEntity) GetSignalRule() *SignalsAPIRuleEntity {
	if o == nil {
		return nil
	}
	return o.SignalRule
}
