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
var _ datasource.DataSource = &IntegrationsAwsConnectionDataSource{}
var _ datasource.DataSourceWithConfigure = &IntegrationsAwsConnectionDataSource{}

func NewIntegrationsAwsConnectionDataSource() datasource.DataSource {
	return &IntegrationsAwsConnectionDataSource{}
}

// IntegrationsAwsConnectionDataSource is the data source implementation.
type IntegrationsAwsConnectionDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// IntegrationsAwsConnectionDataSourceModel describes the data model.
type IntegrationsAwsConnectionDataSourceModel struct {
	AwsAccountID      types.String   `tfsdk:"aws_account_id"`
	ConnectionStatus  types.String   `tfsdk:"connection_status"`
	EnvironmentID     types.String   `tfsdk:"environment_id"`
	EnvironmentName   types.String   `tfsdk:"environment_name"`
	ExternalID        types.String   `tfsdk:"external_id"`
	ID                types.String   `tfsdk:"id"`
	Regions           []types.String `tfsdk:"regions"`
	StatusDescription types.String   `tfsdk:"status_description"`
	StatusText        types.String   `tfsdk:"status_text"`
	TargetArn         types.String   `tfsdk:"target_arn"`
}

// Metadata returns the data source type name.
func (r *IntegrationsAwsConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integrations_aws_connection"
}

// Schema defines the schema for the data source.
func (r *IntegrationsAwsConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IntegrationsAwsConnection DataSource",

		Attributes: map[string]schema.Attribute{
			"aws_account_id": schema.StringAttribute{
				Computed: true,
			},
			"connection_status": schema.StringAttribute{
				Computed: true,
			},
			"environment_id": schema.StringAttribute{
				Computed: true,
			},
			"environment_name": schema.StringAttribute{
				Computed: true,
			},
			"external_id": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"regions": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"status_description": schema.StringAttribute{
				Computed: true,
			},
			"status_text": schema.StringAttribute{
				Computed: true,
			},
			"target_arn": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *IntegrationsAwsConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IntegrationsAwsConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IntegrationsAwsConnectionDataSourceModel
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

	request, requestDiags := data.ToOperationsGetAwsConnectionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Integrations.GetAwsConnection(ctx, *request)
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
	if !(res.IntegrationsAwsConnection != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedIntegrationsAwsConnection(ctx, res.IntegrationsAwsConnection)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
