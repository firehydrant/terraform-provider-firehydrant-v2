// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// SignalsAPIEscalationPolicy - Signals_API_EscalationPolicy model
type SignalsAPIEscalationPolicy struct {
	CreatedAt   *time.Time                                     `json:"created_at,omitempty"`
	CreatedBy   *NullableAuthor                                `json:"created_by,omitempty"`
	Default     *bool                                          `json:"default,omitempty"`
	Description *string                                        `json:"description,omitempty"`
	HandoffStep *NullableSignalsAPIEscalationPolicyHandoffStep `json:"handoff_step,omitempty"`
	ID          *string                                        `json:"id,omitempty"`
	Name        *string                                        `json:"name,omitempty"`
	// Priority-specific policies for dynamic escalation policies
	NotificationPriorityPolicies []SignalsAPINotificationPriorityPolicy `json:"notification_priority_policies,omitempty"`
	Repetitions                  *int                                   `json:"repetitions,omitempty"`
	StepStrategy                 *string                                `json:"step_strategy,omitempty"`
	Steps                        []SignalsAPIEscalationPolicyStep       `json:"steps,omitempty"`
	UpdatedAt                    *time.Time                             `json:"updated_at,omitempty"`
}

func (s SignalsAPIEscalationPolicy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SignalsAPIEscalationPolicy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SignalsAPIEscalationPolicy) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SignalsAPIEscalationPolicy) GetCreatedBy() *NullableAuthor {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *SignalsAPIEscalationPolicy) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *SignalsAPIEscalationPolicy) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SignalsAPIEscalationPolicy) GetHandoffStep() *NullableSignalsAPIEscalationPolicyHandoffStep {
	if o == nil {
		return nil
	}
	return o.HandoffStep
}

func (o *SignalsAPIEscalationPolicy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SignalsAPIEscalationPolicy) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SignalsAPIEscalationPolicy) GetNotificationPriorityPolicies() []SignalsAPINotificationPriorityPolicy {
	if o == nil {
		return nil
	}
	return o.NotificationPriorityPolicies
}

func (o *SignalsAPIEscalationPolicy) GetRepetitions() *int {
	if o == nil {
		return nil
	}
	return o.Repetitions
}

func (o *SignalsAPIEscalationPolicy) GetStepStrategy() *string {
	if o == nil {
		return nil
	}
	return o.StepStrategy
}

func (o *SignalsAPIEscalationPolicy) GetSteps() []SignalsAPIEscalationPolicyStep {
	if o == nil {
		return nil
	}
	return o.Steps
}

func (o *SignalsAPIEscalationPolicy) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
