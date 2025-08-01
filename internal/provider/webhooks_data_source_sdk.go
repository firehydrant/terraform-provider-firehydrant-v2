// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (r *WebhooksDataSourceModel) ToOperationsGetWebhookRequest(ctx context.Context) (*operations.GetWebhookRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var webhookID string
	webhookID = r.ID.ValueString()

	out := operations.GetWebhookRequest{
		WebhookID: webhookID,
	}

	return &out, diags
}
