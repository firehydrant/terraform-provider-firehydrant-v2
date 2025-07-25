// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SignalsAPINotificationPolicyItemDataSource{}
var _ datasource.DataSourceWithConfigure = &SignalsAPINotificationPolicyItemDataSource{}

func NewSignalsAPINotificationPolicyItemDataSource() datasource.DataSource {
	return &SignalsAPINotificationPolicyItemDataSource{}
}

// SignalsAPINotificationPolicyItemDataSource is the data source implementation.
type SignalsAPINotificationPolicyItemDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// SignalsAPINotificationPolicyItemDataSourceModel describes the data model.
type SignalsAPINotificationPolicyItemDataSourceModel struct {
	CreatedAt               types.String `tfsdk:"created_at"`
	ID                      types.String `tfsdk:"id"`
	MaxDelay                types.String `tfsdk:"max_delay"`
	NotificationGroupMethod types.String `tfsdk:"notification_group_method"`
	Priority                types.String `tfsdk:"priority"`
	UpdatedAt               types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *SignalsAPINotificationPolicyItemDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_signals_api_notification_policy_item"
}

// Schema defines the schema for the data source.
func (r *SignalsAPINotificationPolicyItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SignalsAPINotificationPolicyItem DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"max_delay": schema.StringAttribute{
				Computed:    true,
				Description: `The maximum delay for notifications`,
			},
			"notification_group_method": schema.StringAttribute{
				Computed: true,
			},
			"priority": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SignalsAPINotificationPolicyItemDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Firehydrant)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Firehydrant, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SignalsAPINotificationPolicyItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SignalsAPINotificationPolicyItemDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsGetNotificationPolicyRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Signals.GetNotificationPolicy(ctx, *request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SignalsAPINotificationPolicyItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedSignalsAPINotificationPolicyItem(ctx, res.SignalsAPINotificationPolicyItem)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
