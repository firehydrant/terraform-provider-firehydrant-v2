// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/provider/typeconvert"
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SignalsAPIEscalationPolicyDataSourceModel) RefreshFromSharedSignalsAPIEscalationPolicy(ctx context.Context, resp *shared.SignalsAPIEscalationPolicy) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		if resp.CreatedBy == nil {
			r.CreatedBy = nil
		} else {
			r.CreatedBy = &tfTypes.NullableAuthor{}
			r.CreatedBy.Email = types.StringPointerValue(resp.CreatedBy.Email)
			r.CreatedBy.ID = types.StringPointerValue(resp.CreatedBy.ID)
			r.CreatedBy.Name = types.StringPointerValue(resp.CreatedBy.Name)
			r.CreatedBy.Source = types.StringPointerValue(resp.CreatedBy.Source)
		}
		r.Default = types.BoolPointerValue(resp.Default)
		r.Description = types.StringPointerValue(resp.Description)
		if resp.HandoffStep == nil {
			r.HandoffStep = nil
		} else {
			r.HandoffStep = &tfTypes.NullableSignalsAPIEscalationPolicyHandoffStep{}
			r.HandoffStep.ID = types.StringPointerValue(resp.HandoffStep.ID)
			if resp.HandoffStep.Target == nil {
				r.HandoffStep.Target = nil
			} else {
				r.HandoffStep.Target = &tfTypes.NullableSignalsAPITarget{}
				r.HandoffStep.Target.ID = types.StringPointerValue(resp.HandoffStep.Target.ID)
				r.HandoffStep.Target.IsPageable = types.BoolPointerValue(resp.HandoffStep.Target.IsPageable)
				r.HandoffStep.Target.Name = types.StringPointerValue(resp.HandoffStep.Target.Name)
				r.HandoffStep.Target.TeamID = types.StringPointerValue(resp.HandoffStep.Target.TeamID)
				r.HandoffStep.Target.Type = types.StringPointerValue(resp.HandoffStep.Target.Type)
			}
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.NotificationPriorityPolicies != nil {
			r.NotificationPriorityPolicies = []tfTypes.SignalsAPINotificationPriorityPolicy{}
			if len(r.NotificationPriorityPolicies) > len(resp.NotificationPriorityPolicies) {
				r.NotificationPriorityPolicies = r.NotificationPriorityPolicies[:len(resp.NotificationPriorityPolicies)]
			}
			for notificationPriorityPoliciesCount, notificationPriorityPoliciesItem := range resp.NotificationPriorityPolicies {
				var notificationPriorityPolicies tfTypes.SignalsAPINotificationPriorityPolicy
				if notificationPriorityPoliciesItem.HandoffStep == nil {
					notificationPriorityPolicies.HandoffStep = nil
				} else {
					notificationPriorityPolicies.HandoffStep = &tfTypes.NullableSignalsAPIEscalationPolicyHandoffStep{}
					notificationPriorityPolicies.HandoffStep.ID = types.StringPointerValue(notificationPriorityPoliciesItem.HandoffStep.ID)
					if notificationPriorityPoliciesItem.HandoffStep.Target == nil {
						notificationPriorityPolicies.HandoffStep.Target = nil
					} else {
						notificationPriorityPolicies.HandoffStep.Target = &tfTypes.NullableSignalsAPITarget{}
						notificationPriorityPolicies.HandoffStep.Target.ID = types.StringPointerValue(notificationPriorityPoliciesItem.HandoffStep.Target.ID)
						notificationPriorityPolicies.HandoffStep.Target.IsPageable = types.BoolPointerValue(notificationPriorityPoliciesItem.HandoffStep.Target.IsPageable)
						notificationPriorityPolicies.HandoffStep.Target.Name = types.StringPointerValue(notificationPriorityPoliciesItem.HandoffStep.Target.Name)
						notificationPriorityPolicies.HandoffStep.Target.TeamID = types.StringPointerValue(notificationPriorityPoliciesItem.HandoffStep.Target.TeamID)
						notificationPriorityPolicies.HandoffStep.Target.Type = types.StringPointerValue(notificationPriorityPoliciesItem.HandoffStep.Target.Type)
					}
				}
				notificationPriorityPolicies.NotificationPriority = types.StringPointerValue(notificationPriorityPoliciesItem.NotificationPriority)
				notificationPriorityPolicies.Repetitions = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(notificationPriorityPoliciesItem.Repetitions))
				if notificationPriorityPoliciesItem.Steps != nil {
					notificationPriorityPolicies.Steps = []tfTypes.SignalsAPIEscalationPolicyStep{}
					for stepsCount, stepsItem := range notificationPriorityPoliciesItem.Steps {
						var steps tfTypes.SignalsAPIEscalationPolicyStep
						steps.DistributionType = types.StringPointerValue(stepsItem.DistributionType)
						steps.ID = types.StringPointerValue(stepsItem.ID)
						if stepsItem.NextTargetForRoundRobin == nil {
							steps.NextTargetForRoundRobin = nil
						} else {
							steps.NextTargetForRoundRobin = &tfTypes.NullableSignalsAPITarget{}
							steps.NextTargetForRoundRobin.ID = types.StringPointerValue(stepsItem.NextTargetForRoundRobin.ID)
							steps.NextTargetForRoundRobin.IsPageable = types.BoolPointerValue(stepsItem.NextTargetForRoundRobin.IsPageable)
							steps.NextTargetForRoundRobin.Name = types.StringPointerValue(stepsItem.NextTargetForRoundRobin.Name)
							steps.NextTargetForRoundRobin.TeamID = types.StringPointerValue(stepsItem.NextTargetForRoundRobin.TeamID)
							steps.NextTargetForRoundRobin.Type = types.StringPointerValue(stepsItem.NextTargetForRoundRobin.Type)
						}
						steps.ParentPosition = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(stepsItem.ParentPosition))
						steps.Position = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(stepsItem.Position))
						if stepsItem.Priorities != nil {
							steps.Priorities = make([]types.String, 0, len(stepsItem.Priorities))
							for _, v := range stepsItem.Priorities {
								steps.Priorities = append(steps.Priorities, types.StringValue(v))
							}
						}
						if stepsItem.Targets != nil {
							steps.Targets = []tfTypes.SignalsAPITarget{}
							for targetsCount, targetsItem := range stepsItem.Targets {
								var targets tfTypes.SignalsAPITarget
								targets.ID = types.StringPointerValue(targetsItem.ID)
								targets.IsPageable = types.BoolPointerValue(targetsItem.IsPageable)
								targets.Name = types.StringPointerValue(targetsItem.Name)
								targets.TeamID = types.StringPointerValue(targetsItem.TeamID)
								targets.Type = types.StringPointerValue(targetsItem.Type)
								if targetsCount+1 > len(steps.Targets) {
									steps.Targets = append(steps.Targets, targets)
								} else {
									steps.Targets[targetsCount].ID = targets.ID
									steps.Targets[targetsCount].IsPageable = targets.IsPageable
									steps.Targets[targetsCount].Name = targets.Name
									steps.Targets[targetsCount].TeamID = targets.TeamID
									steps.Targets[targetsCount].Type = targets.Type
								}
							}
						}
						steps.Timeout = types.StringPointerValue(stepsItem.Timeout)
						if stepsCount+1 > len(notificationPriorityPolicies.Steps) {
							notificationPriorityPolicies.Steps = append(notificationPriorityPolicies.Steps, steps)
						} else {
							notificationPriorityPolicies.Steps[stepsCount].DistributionType = steps.DistributionType
							notificationPriorityPolicies.Steps[stepsCount].ID = steps.ID
							notificationPriorityPolicies.Steps[stepsCount].NextTargetForRoundRobin = steps.NextTargetForRoundRobin
							notificationPriorityPolicies.Steps[stepsCount].ParentPosition = steps.ParentPosition
							notificationPriorityPolicies.Steps[stepsCount].Position = steps.Position
							notificationPriorityPolicies.Steps[stepsCount].Priorities = steps.Priorities
							notificationPriorityPolicies.Steps[stepsCount].Targets = steps.Targets
							notificationPriorityPolicies.Steps[stepsCount].Timeout = steps.Timeout
						}
					}
				}
				if notificationPriorityPoliciesCount+1 > len(r.NotificationPriorityPolicies) {
					r.NotificationPriorityPolicies = append(r.NotificationPriorityPolicies, notificationPriorityPolicies)
				} else {
					r.NotificationPriorityPolicies[notificationPriorityPoliciesCount].HandoffStep = notificationPriorityPolicies.HandoffStep
					r.NotificationPriorityPolicies[notificationPriorityPoliciesCount].NotificationPriority = notificationPriorityPolicies.NotificationPriority
					r.NotificationPriorityPolicies[notificationPriorityPoliciesCount].Repetitions = notificationPriorityPolicies.Repetitions
					r.NotificationPriorityPolicies[notificationPriorityPoliciesCount].Steps = notificationPriorityPolicies.Steps
				}
			}
		}
		r.Repetitions = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(resp.Repetitions))
		r.StepStrategy = types.StringPointerValue(resp.StepStrategy)
		if resp.Steps != nil {
			r.Steps = []tfTypes.SignalsAPIEscalationPolicyStep{}
			if len(r.Steps) > len(resp.Steps) {
				r.Steps = r.Steps[:len(resp.Steps)]
			}
			for stepsCount1, stepsItem1 := range resp.Steps {
				var steps1 tfTypes.SignalsAPIEscalationPolicyStep
				steps1.DistributionType = types.StringPointerValue(stepsItem1.DistributionType)
				steps1.ID = types.StringPointerValue(stepsItem1.ID)
				if stepsItem1.NextTargetForRoundRobin == nil {
					steps1.NextTargetForRoundRobin = nil
				} else {
					steps1.NextTargetForRoundRobin = &tfTypes.NullableSignalsAPITarget{}
					steps1.NextTargetForRoundRobin.ID = types.StringPointerValue(stepsItem1.NextTargetForRoundRobin.ID)
					steps1.NextTargetForRoundRobin.IsPageable = types.BoolPointerValue(stepsItem1.NextTargetForRoundRobin.IsPageable)
					steps1.NextTargetForRoundRobin.Name = types.StringPointerValue(stepsItem1.NextTargetForRoundRobin.Name)
					steps1.NextTargetForRoundRobin.TeamID = types.StringPointerValue(stepsItem1.NextTargetForRoundRobin.TeamID)
					steps1.NextTargetForRoundRobin.Type = types.StringPointerValue(stepsItem1.NextTargetForRoundRobin.Type)
				}
				steps1.ParentPosition = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(stepsItem1.ParentPosition))
				steps1.Position = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(stepsItem1.Position))
				if stepsItem1.Priorities != nil {
					steps1.Priorities = make([]types.String, 0, len(stepsItem1.Priorities))
					for _, v := range stepsItem1.Priorities {
						steps1.Priorities = append(steps1.Priorities, types.StringValue(v))
					}
				}
				if stepsItem1.Targets != nil {
					steps1.Targets = []tfTypes.SignalsAPITarget{}
					for targetsCount1, targetsItem1 := range stepsItem1.Targets {
						var targets1 tfTypes.SignalsAPITarget
						targets1.ID = types.StringPointerValue(targetsItem1.ID)
						targets1.IsPageable = types.BoolPointerValue(targetsItem1.IsPageable)
						targets1.Name = types.StringPointerValue(targetsItem1.Name)
						targets1.TeamID = types.StringPointerValue(targetsItem1.TeamID)
						targets1.Type = types.StringPointerValue(targetsItem1.Type)
						if targetsCount1+1 > len(steps1.Targets) {
							steps1.Targets = append(steps1.Targets, targets1)
						} else {
							steps1.Targets[targetsCount1].ID = targets1.ID
							steps1.Targets[targetsCount1].IsPageable = targets1.IsPageable
							steps1.Targets[targetsCount1].Name = targets1.Name
							steps1.Targets[targetsCount1].TeamID = targets1.TeamID
							steps1.Targets[targetsCount1].Type = targets1.Type
						}
					}
				}
				steps1.Timeout = types.StringPointerValue(stepsItem1.Timeout)
				if stepsCount1+1 > len(r.Steps) {
					r.Steps = append(r.Steps, steps1)
				} else {
					r.Steps[stepsCount1].DistributionType = steps1.DistributionType
					r.Steps[stepsCount1].ID = steps1.ID
					r.Steps[stepsCount1].NextTargetForRoundRobin = steps1.NextTargetForRoundRobin
					r.Steps[stepsCount1].ParentPosition = steps1.ParentPosition
					r.Steps[stepsCount1].Position = steps1.Position
					r.Steps[stepsCount1].Priorities = steps1.Priorities
					r.Steps[stepsCount1].Targets = steps1.Targets
					r.Steps[stepsCount1].Timeout = steps1.Timeout
				}
			}
		}
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
	}

	return diags
}

func (r *SignalsAPIEscalationPolicyDataSourceModel) ToOperationsGetTeamEscalationPolicyRequest(ctx context.Context) (*operations.GetTeamEscalationPolicyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var teamID string
	teamID = r.TeamID.ValueString()

	var id string
	id = r.ID.ValueString()

	out := operations.GetTeamEscalationPolicyRequest{
		TeamID: teamID,
		ID:     id,
	}

	return &out, diags
}
