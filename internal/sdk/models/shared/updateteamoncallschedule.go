// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateTeamOnCallScheduleMembersInput struct {
	// The ID of a user who should be added to the schedule's rotation. You can add a user to the rotation
	// multiple times to construct more complex rotations, and you can specify a `null` user ID to create
	// unassigned slots in the rotation.
	//
	UserID *string `json:"user_id,omitempty"`
}

func (o *UpdateTeamOnCallScheduleMembersInput) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type UpdateTeamOnCallScheduleRestrictionsInput struct {
	// The day of the week on which the restriction should end, as its long-form name (e.g. "monday", "tuesday", etc).
	EndDay string `json:"end_day"`
	// An ISO8601 time string specifying when the restriction should end.
	EndTime string `json:"end_time"`
	// The day of the week on which the restriction should start, as its long-form name (e.g. "monday", "tuesday", etc).
	StartDay string `json:"start_day"`
	// An ISO8601 time string specifying when the restriction should start.
	StartTime string `json:"start_time"`
}

func (o *UpdateTeamOnCallScheduleRestrictionsInput) GetEndDay() string {
	if o == nil {
		return ""
	}
	return o.EndDay
}

func (o *UpdateTeamOnCallScheduleRestrictionsInput) GetEndTime() string {
	if o == nil {
		return ""
	}
	return o.EndTime
}

func (o *UpdateTeamOnCallScheduleRestrictionsInput) GetStartDay() string {
	if o == nil {
		return ""
	}
	return o.StartDay
}

func (o *UpdateTeamOnCallScheduleRestrictionsInput) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

// UpdateTeamOnCallScheduleStrategyInput - An object that specifies how the rotation's on-call shifts should be generated.
type UpdateTeamOnCallScheduleStrategyInput struct {
	// The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".
	HandoffDay *string `json:"handoff_day,omitempty"`
	// An ISO8601 time string specifying when on-call shifts should hand off. This value is only used if the strategy type is "daily" or "weekly".
	HandoffTime *string `json:"handoff_time,omitempty"`
	// An ISO8601 duration string specifying how long each shift should last. This value is only used if the strategy type is "custom".
	ShiftDuration *string `json:"shift_duration,omitempty"`
	// The type of strategy. Must be one of "daily", "weekly", or "custom".
	Type string `json:"type"`
}

func (o *UpdateTeamOnCallScheduleStrategyInput) GetHandoffDay() *string {
	if o == nil {
		return nil
	}
	return o.HandoffDay
}

func (o *UpdateTeamOnCallScheduleStrategyInput) GetHandoffTime() *string {
	if o == nil {
		return nil
	}
	return o.HandoffTime
}

func (o *UpdateTeamOnCallScheduleStrategyInput) GetShiftDuration() *string {
	if o == nil {
		return nil
	}
	return o.ShiftDuration
}

func (o *UpdateTeamOnCallScheduleStrategyInput) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// UpdateTeamOnCallSchedule - Update a Signals on-call schedule by ID. For backwards compatibility, all parameters except for
// `name` and `description` will be ignored if the schedule has more than one rotation. If the schedule
// has only one rotation, you can continue to update that rotation using the rotation-specific parameters.
type UpdateTeamOnCallSchedule struct {
	// A hex color code that will be used to represent the schedule's rotation in FireHydrant's UI.
	Color *string `json:"color,omitempty"`
	// A new, detailed description for the on-call schedule.
	Description *string `json:"description,omitempty"`
	// An ISO8601 time string specifying when the updated schedule should take effect. This
	// value must be provided if editing an attribute that would affect how the schedule's
	// shifts are generated, such as the time zone, members, strategy, or restrictions.
	//
	EffectiveAt *string `json:"effective_at,omitempty"`
	// This parameter is deprecated; use `members` instead.
	MemberIds []string `json:"member_ids,omitempty"`
	// An ordered list of objects that specify members of the schedule's rotation.
	MembersInput []UpdateTeamOnCallScheduleMembersInput `json:"members,omitempty"`
	// A new name for the on-call schedule.
	Name *string `json:"name,omitempty"`
	// A list of objects that restrict the schedule's rotation to specific on-call periods.
	RestrictionsInput []UpdateTeamOnCallScheduleRestrictionsInput `json:"restrictions,omitempty"`
	// A new, detailed description for the schedule's rotation.
	RotationDescription *string `json:"rotation_description,omitempty"`
	// A new name for the schedule's rotation.
	RotationName *string `json:"rotation_name,omitempty"`
	// The ID of a Slack user group to sync the rotation's on-call members to.
	SlackUserGroupID *string `json:"slack_user_group_id,omitempty"`
	// An object that specifies how the rotation's on-call shifts should be generated.
	StrategyInput *UpdateTeamOnCallScheduleStrategyInput `json:"strategy,omitempty"`
	// The time zone in which the on-call schedule's rotation will operate. This value must be a valid IANA time zone name.
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o *UpdateTeamOnCallSchedule) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *UpdateTeamOnCallSchedule) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateTeamOnCallSchedule) GetEffectiveAt() *string {
	if o == nil {
		return nil
	}
	return o.EffectiveAt
}

func (o *UpdateTeamOnCallSchedule) GetMemberIds() []string {
	if o == nil {
		return nil
	}
	return o.MemberIds
}

func (o *UpdateTeamOnCallSchedule) GetMembersInput() []UpdateTeamOnCallScheduleMembersInput {
	if o == nil {
		return nil
	}
	return o.MembersInput
}

func (o *UpdateTeamOnCallSchedule) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateTeamOnCallSchedule) GetRestrictionsInput() []UpdateTeamOnCallScheduleRestrictionsInput {
	if o == nil {
		return nil
	}
	return o.RestrictionsInput
}

func (o *UpdateTeamOnCallSchedule) GetRotationDescription() *string {
	if o == nil {
		return nil
	}
	return o.RotationDescription
}

func (o *UpdateTeamOnCallSchedule) GetRotationName() *string {
	if o == nil {
		return nil
	}
	return o.RotationName
}

func (o *UpdateTeamOnCallSchedule) GetSlackUserGroupID() *string {
	if o == nil {
		return nil
	}
	return o.SlackUserGroupID
}

func (o *UpdateTeamOnCallSchedule) GetStrategyInput() *UpdateTeamOnCallScheduleStrategyInput {
	if o == nil {
		return nil
	}
	return o.StrategyInput
}

func (o *UpdateTeamOnCallSchedule) GetTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.TimeZone
}
