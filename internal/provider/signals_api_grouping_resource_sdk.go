// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SignalsAPIGroupingResourceModel) RefreshFromSharedSignalsAPIGrouping(ctx context.Context, resp *shared.SignalsAPIGrouping) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Action == nil {
			r.Action = nil
		} else {
			r.Action = &tfTypes.NullableSignalsAPIGroupingAction{}
			if resp.Action.Fyi == nil {
				r.Action.Fyi = nil
			} else {
				r.Action.Fyi = &tfTypes.NullableSignalsAPIGroupingActionFyi{}
				if resp.Action.Fyi.SlackChannels != nil {
					r.Action.Fyi.SlackChannels = []tfTypes.IntegrationsSlackSlackChannel{}
					if len(r.Action.Fyi.SlackChannels) > len(resp.Action.Fyi.SlackChannels) {
						r.Action.Fyi.SlackChannels = r.Action.Fyi.SlackChannels[:len(resp.Action.Fyi.SlackChannels)]
					}
					for slackChannelsCount, slackChannelsItem := range resp.Action.Fyi.SlackChannels {
						var slackChannels tfTypes.IntegrationsSlackSlackChannel
						slackChannels.ID = types.StringPointerValue(slackChannelsItem.ID)
						slackChannels.Name = types.StringPointerValue(slackChannelsItem.Name)
						slackChannels.SlackChannelID = types.StringPointerValue(slackChannelsItem.SlackChannelID)
						if slackChannelsCount+1 > len(r.Action.Fyi.SlackChannels) {
							r.Action.Fyi.SlackChannels = append(r.Action.Fyi.SlackChannels, slackChannels)
						} else {
							r.Action.Fyi.SlackChannels[slackChannelsCount].ID = slackChannels.ID
							r.Action.Fyi.SlackChannels[slackChannelsCount].Name = slackChannels.Name
							r.Action.Fyi.SlackChannels[slackChannelsCount].SlackChannelID = slackChannels.SlackChannelID
						}
					}
				}
			}
			r.Action.Link = types.BoolPointerValue(resp.Action.Link)
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.ReferenceAlertTimePeriod = types.StringPointerValue(resp.ReferenceAlertTimePeriod)
		if resp.Strategy != nil {
			if resp.Strategy.Substring == nil {
				r.Strategy.Substring = nil
			} else {
				r.Strategy.Substring = &tfTypes.CreateSignalsAlertGroupingConfigurationSubstring{}
				r.Strategy.Substring.FieldName = types.StringPointerValue(resp.Strategy.Substring.FieldName)
				r.Strategy.Substring.Value = types.StringPointerValue(resp.Strategy.Substring.Value)
			}
		}
	}

	return diags
}

func (r *SignalsAPIGroupingResourceModel) ToOperationsDeleteSignalsAlertGroupingConfigurationRequest(ctx context.Context) (*operations.DeleteSignalsAlertGroupingConfigurationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var id string
	id = r.ID.ValueString()

	out := operations.DeleteSignalsAlertGroupingConfigurationRequest{
		ID: id,
	}

	return &out, diags
}

func (r *SignalsAPIGroupingResourceModel) ToOperationsGetSignalsAlertGroupingConfigurationRequest(ctx context.Context) (*operations.GetSignalsAlertGroupingConfigurationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var id string
	id = r.ID.ValueString()

	out := operations.GetSignalsAlertGroupingConfigurationRequest{
		ID: id,
	}

	return &out, diags
}

func (r *SignalsAPIGroupingResourceModel) ToOperationsUpdateSignalsAlertGroupingConfigurationRequest(ctx context.Context) (*operations.UpdateSignalsAlertGroupingConfigurationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var id string
	id = r.ID.ValueString()

	updateSignalsAlertGroupingConfiguration, updateSignalsAlertGroupingConfigurationDiags := r.ToSharedUpdateSignalsAlertGroupingConfiguration(ctx)
	diags.Append(updateSignalsAlertGroupingConfigurationDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateSignalsAlertGroupingConfigurationRequest{
		ID:                                      id,
		UpdateSignalsAlertGroupingConfiguration: *updateSignalsAlertGroupingConfiguration,
	}

	return &out, diags
}

func (r *SignalsAPIGroupingResourceModel) ToSharedCreateSignalsAlertGroupingConfiguration(ctx context.Context) (*shared.CreateSignalsAlertGroupingConfiguration, diag.Diagnostics) {
	var diags diag.Diagnostics

	var actionInput *shared.CreateSignalsAlertGroupingConfigurationActionInput
	if r.ActionInput != nil {
		var fyi *shared.CreateSignalsAlertGroupingConfigurationFyi
		if r.ActionInput.Fyi != nil {
			slackChannelIds := make([]string, 0, len(r.ActionInput.Fyi.SlackChannelIds))
			for _, slackChannelIdsItem := range r.ActionInput.Fyi.SlackChannelIds {
				slackChannelIds = append(slackChannelIds, slackChannelIdsItem.ValueString())
			}
			fyi = &shared.CreateSignalsAlertGroupingConfigurationFyi{
				SlackChannelIds: slackChannelIds,
			}
		}
		link := new(bool)
		if !r.ActionInput.Link.IsUnknown() && !r.ActionInput.Link.IsNull() {
			*link = r.ActionInput.Link.ValueBool()
		} else {
			link = nil
		}
		actionInput = &shared.CreateSignalsAlertGroupingConfigurationActionInput{
			Fyi:  fyi,
			Link: link,
		}
	}
	var referenceAlertTimePeriod string
	referenceAlertTimePeriod = r.ReferenceAlertTimePeriod.ValueString()

	var substring *shared.CreateSignalsAlertGroupingConfigurationSubstring
	if r.Strategy.Substring != nil {
		var fieldName string
		fieldName = r.Strategy.Substring.FieldName.ValueString()

		var value string
		value = r.Strategy.Substring.Value.ValueString()

		substring = &shared.CreateSignalsAlertGroupingConfigurationSubstring{
			FieldName: fieldName,
			Value:     value,
		}
	}
	strategy := shared.CreateSignalsAlertGroupingConfigurationStrategy{
		Substring: substring,
	}
	out := shared.CreateSignalsAlertGroupingConfiguration{
		ActionInput:              actionInput,
		ReferenceAlertTimePeriod: referenceAlertTimePeriod,
		Strategy:                 strategy,
	}

	return &out, diags
}

func (r *SignalsAPIGroupingResourceModel) ToSharedUpdateSignalsAlertGroupingConfiguration(ctx context.Context) (*shared.UpdateSignalsAlertGroupingConfiguration, diag.Diagnostics) {
	var diags diag.Diagnostics

	var actionInput *shared.UpdateSignalsAlertGroupingConfigurationActionInput
	if r.ActionInput != nil {
		var fyi *shared.UpdateSignalsAlertGroupingConfigurationFyi
		if r.ActionInput.Fyi != nil {
			slackChannelIds := make([]string, 0, len(r.ActionInput.Fyi.SlackChannelIds))
			for _, slackChannelIdsItem := range r.ActionInput.Fyi.SlackChannelIds {
				slackChannelIds = append(slackChannelIds, slackChannelIdsItem.ValueString())
			}
			fyi = &shared.UpdateSignalsAlertGroupingConfigurationFyi{
				SlackChannelIds: slackChannelIds,
			}
		}
		link := new(bool)
		if !r.ActionInput.Link.IsUnknown() && !r.ActionInput.Link.IsNull() {
			*link = r.ActionInput.Link.ValueBool()
		} else {
			link = nil
		}
		actionInput = &shared.UpdateSignalsAlertGroupingConfigurationActionInput{
			Fyi:  fyi,
			Link: link,
		}
	}
	referenceAlertTimePeriod := new(string)
	if !r.ReferenceAlertTimePeriod.IsUnknown() && !r.ReferenceAlertTimePeriod.IsNull() {
		*referenceAlertTimePeriod = r.ReferenceAlertTimePeriod.ValueString()
	} else {
		referenceAlertTimePeriod = nil
	}
	var strategy *shared.UpdateSignalsAlertGroupingConfigurationStrategy
	var substring *shared.UpdateSignalsAlertGroupingConfigurationSubstring
	if r.Strategy.Substring != nil {
		var fieldName string
		fieldName = r.Strategy.Substring.FieldName.ValueString()

		var value string
		value = r.Strategy.Substring.Value.ValueString()

		substring = &shared.UpdateSignalsAlertGroupingConfigurationSubstring{
			FieldName: fieldName,
			Value:     value,
		}
	}
	strategy = &shared.UpdateSignalsAlertGroupingConfigurationStrategy{
		Substring: substring,
	}
	out := shared.UpdateSignalsAlertGroupingConfiguration{
		ActionInput:              actionInput,
		ReferenceAlertTimePeriod: referenceAlertTimePeriod,
		Strategy:                 strategy,
	}

	return &out, diags
}
