// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// Team model
type Team struct {
	CreatedAt                      *time.Time                                   `json:"created_at,omitempty"`
	CreatedBy                      *NullableAuthor                              `json:"created_by,omitempty"`
	DefaultSignalsEscalationPolicy *NullableSuccinct                            `json:"default_signals_escalation_policy,omitempty"`
	Description                    *string                                      `json:"description,omitempty"`
	Functionalities                []FunctionalityLite                          `json:"functionalities,omitempty"`
	ID                             *string                                      `json:"id,omitempty"`
	InSupportHours                 *bool                                        `json:"in_support_hours,omitempty"`
	Memberships                    []Membership                                 `json:"memberships,omitempty"`
	MsTeamsChannel                 *NullableIntegrationsMicrosoftTeamsV2Channel `json:"ms_teams_channel,omitempty"`
	Name                           *string                                      `json:"name,omitempty"`
	OwnedChecklistTemplates        []ChecklistTemplate                          `json:"owned_checklist_templates,omitempty"`
	OwnedFunctionalities           []FunctionalityLite                          `json:"owned_functionalities,omitempty"`
	OwnedRunbooks                  []SlimRunbook                                `json:"owned_runbooks,omitempty"`
	OwnedServices                  []ServiceLite                                `json:"owned_services,omitempty"`
	RespondingServices             []ServiceLite                                `json:"responding_services,omitempty"`
	Services                       []ServiceLite                                `json:"services,omitempty"`
	SignalsIcalURL                 *string                                      `json:"signals_ical_url,omitempty"`
	SlackChannel                   *NullableIntegrationsSlackSlackChannel       `json:"slack_channel,omitempty"`
	Slug                           *string                                      `json:"slug,omitempty"`
	UpdatedAt                      *time.Time                                   `json:"updated_at,omitempty"`
}

func (t Team) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Team) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Team) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Team) GetCreatedBy() *NullableAuthor {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Team) GetDefaultSignalsEscalationPolicy() *NullableSuccinct {
	if o == nil {
		return nil
	}
	return o.DefaultSignalsEscalationPolicy
}

func (o *Team) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Team) GetFunctionalities() []FunctionalityLite {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *Team) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Team) GetInSupportHours() *bool {
	if o == nil {
		return nil
	}
	return o.InSupportHours
}

func (o *Team) GetMemberships() []Membership {
	if o == nil {
		return nil
	}
	return o.Memberships
}

func (o *Team) GetMsTeamsChannel() *NullableIntegrationsMicrosoftTeamsV2Channel {
	if o == nil {
		return nil
	}
	return o.MsTeamsChannel
}

func (o *Team) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Team) GetOwnedChecklistTemplates() []ChecklistTemplate {
	if o == nil {
		return nil
	}
	return o.OwnedChecklistTemplates
}

func (o *Team) GetOwnedFunctionalities() []FunctionalityLite {
	if o == nil {
		return nil
	}
	return o.OwnedFunctionalities
}

func (o *Team) GetOwnedRunbooks() []SlimRunbook {
	if o == nil {
		return nil
	}
	return o.OwnedRunbooks
}

func (o *Team) GetOwnedServices() []ServiceLite {
	if o == nil {
		return nil
	}
	return o.OwnedServices
}

func (o *Team) GetRespondingServices() []ServiceLite {
	if o == nil {
		return nil
	}
	return o.RespondingServices
}

func (o *Team) GetServices() []ServiceLite {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *Team) GetSignalsIcalURL() *string {
	if o == nil {
		return nil
	}
	return o.SignalsIcalURL
}

func (o *Team) GetSlackChannel() *NullableIntegrationsSlackSlackChannel {
	if o == nil {
		return nil
	}
	return o.SlackChannel
}

func (o *Team) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Team) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
