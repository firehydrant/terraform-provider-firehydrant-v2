// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
)

// PostV1TeamsTeamIDEscalationPoliciesType - The type of target that the step will notify.
type PostV1TeamsTeamIDEscalationPoliciesType string

const (
	PostV1TeamsTeamIDEscalationPoliciesTypeOnCallSchedule PostV1TeamsTeamIDEscalationPoliciesType = "OnCallSchedule"
	PostV1TeamsTeamIDEscalationPoliciesTypeUser           PostV1TeamsTeamIDEscalationPoliciesType = "User"
	PostV1TeamsTeamIDEscalationPoliciesTypeSlackChannel   PostV1TeamsTeamIDEscalationPoliciesType = "SlackChannel"
	PostV1TeamsTeamIDEscalationPoliciesTypeEntireTeam     PostV1TeamsTeamIDEscalationPoliciesType = "EntireTeam"
	PostV1TeamsTeamIDEscalationPoliciesTypeWebhook        PostV1TeamsTeamIDEscalationPoliciesType = "Webhook"
)

func (e PostV1TeamsTeamIDEscalationPoliciesType) ToPointer() *PostV1TeamsTeamIDEscalationPoliciesType {
	return &e
}
func (e *PostV1TeamsTeamIDEscalationPoliciesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OnCallSchedule":
		fallthrough
	case "User":
		fallthrough
	case "SlackChannel":
		fallthrough
	case "EntireTeam":
		fallthrough
	case "Webhook":
		*e = PostV1TeamsTeamIDEscalationPoliciesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1TeamsTeamIDEscalationPoliciesType: %v", v)
	}
}

type Targets struct {
	// The type of target that the step will notify.
	Type PostV1TeamsTeamIDEscalationPoliciesType `json:"type"`
	// The ID of the target that the step will notify.
	ID string `json:"id"`
}

func (o *Targets) GetType() PostV1TeamsTeamIDEscalationPoliciesType {
	if o == nil {
		return PostV1TeamsTeamIDEscalationPoliciesType("")
	}
	return o.Type
}

func (o *Targets) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DistributionType - The round robin configuration for the step. One of 'unspecified', 'round_robin_by_alert', or 'round_robin_by_escalation_policy'.
type DistributionType string

const (
	DistributionTypeUnspecified                  DistributionType = "unspecified"
	DistributionTypeRoundRobinByAlert            DistributionType = "round_robin_by_alert"
	DistributionTypeRoundRobinByEscalationPolicy DistributionType = "round_robin_by_escalation_policy"
)

func (e DistributionType) ToPointer() *DistributionType {
	return &e
}
func (e *DistributionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unspecified":
		fallthrough
	case "round_robin_by_alert":
		fallthrough
	case "round_robin_by_escalation_policy":
		*e = DistributionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DistributionType: %v", v)
	}
}

type Steps struct {
	// A list of targets that the step will notify. You can specify up to 15 targets per step.
	Targets []Targets `json:"targets"`
	// An ISO8601 duration string specifying how long to wait before moving on to the next step. For the last step, this value specifies how long to wait before the escalation policy should repeat, if it repeats.
	Timeout string `json:"timeout"`
	// The round robin configuration for the step. One of 'unspecified', 'round_robin_by_alert', or 'round_robin_by_escalation_policy'.
	DistributionType *DistributionType `json:"distribution_type,omitempty"`
}

func (o *Steps) GetTargets() []Targets {
	if o == nil {
		return []Targets{}
	}
	return o.Targets
}

func (o *Steps) GetTimeout() string {
	if o == nil {
		return ""
	}
	return o.Timeout
}

func (o *Steps) GetDistributionType() *DistributionType {
	if o == nil {
		return nil
	}
	return o.DistributionType
}

// PostV1TeamsTeamIDEscalationPoliciesTargetType - The type of target to which the policy will hand off.
type PostV1TeamsTeamIDEscalationPoliciesTargetType string

const (
	PostV1TeamsTeamIDEscalationPoliciesTargetTypeEscalationPolicy PostV1TeamsTeamIDEscalationPoliciesTargetType = "EscalationPolicy"
	PostV1TeamsTeamIDEscalationPoliciesTargetTypeTeam             PostV1TeamsTeamIDEscalationPoliciesTargetType = "Team"
)

func (e PostV1TeamsTeamIDEscalationPoliciesTargetType) ToPointer() *PostV1TeamsTeamIDEscalationPoliciesTargetType {
	return &e
}
func (e *PostV1TeamsTeamIDEscalationPoliciesTargetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EscalationPolicy":
		fallthrough
	case "Team":
		*e = PostV1TeamsTeamIDEscalationPoliciesTargetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1TeamsTeamIDEscalationPoliciesTargetType: %v", v)
	}
}

// HandoffStep - A step that defines where an alert should be sent when the policy is exhausted and the alert is still unacknowledged.
type HandoffStep struct {
	// The type of target to which the policy will hand off.
	TargetType PostV1TeamsTeamIDEscalationPoliciesTargetType `json:"target_type"`
	// The ID of the target to which the policy will hand off.
	TargetID string `json:"target_id"`
}

func (o *HandoffStep) GetTargetType() PostV1TeamsTeamIDEscalationPoliciesTargetType {
	if o == nil {
		return PostV1TeamsTeamIDEscalationPoliciesTargetType("")
	}
	return o.TargetType
}

func (o *HandoffStep) GetTargetID() string {
	if o == nil {
		return ""
	}
	return o.TargetID
}

// PostV1TeamsTeamIDEscalationPolicies - Create a Signals escalation policy for a team.
type PostV1TeamsTeamIDEscalationPolicies struct {
	// The escalation policy's name.
	Name string `json:"name"`
	// A detailed description of the escalation policy.
	Description *string `json:"description,omitempty"`
	// The number of times that the escalation policy should repeat before an alert is dropped.
	Repetitions *int `default:"0" json:"repetitions"`
	// Whether this escalation policy should be the default for the team.
	Default *bool `default:"false" json:"default"`
	// A list of steps that define how an alert should escalate through the policy.
	Steps []Steps `json:"steps"`
	// A step that defines where an alert should be sent when the policy is exhausted and the alert is still unacknowledged.
	HandoffStep *HandoffStep `json:"handoff_step,omitempty"`
}

func (p PostV1TeamsTeamIDEscalationPolicies) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1TeamsTeamIDEscalationPolicies) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetRepetitions() *int {
	if o == nil {
		return nil
	}
	return o.Repetitions
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetSteps() []Steps {
	if o == nil {
		return []Steps{}
	}
	return o.Steps
}

func (o *PostV1TeamsTeamIDEscalationPolicies) GetHandoffStep() *HandoffStep {
	if o == nil {
		return nil
	}
	return o.HandoffStep
}
