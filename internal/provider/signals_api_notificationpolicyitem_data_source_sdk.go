// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/provider/typeconvert"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SignalsAPINotificationPolicyItemDataSourceModel) RefreshFromSharedSignalsAPINotificationPolicyItem(ctx context.Context, resp *shared.SignalsAPINotificationPolicyItem) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		r.ID = types.StringPointerValue(resp.ID)
		r.MaxDelay = types.StringPointerValue(resp.MaxDelay)
		r.NotificationGroupMethod = types.StringPointerValue(resp.NotificationGroupMethod)
		r.Priority = types.StringPointerValue(resp.Priority)
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
	}

	return diags
}

func (r *SignalsAPINotificationPolicyItemDataSourceModel) ToOperationsGetNotificationPolicyRequest(ctx context.Context) (*operations.GetNotificationPolicyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var id string
	id = r.ID.ValueString()

	out := operations.GetNotificationPolicyRequest{
		ID: id,
	}

	return &out, diags
}
