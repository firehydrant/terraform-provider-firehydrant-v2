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
var _ datasource.DataSource = &IntegrationsStatuspagePagesDataSource{}
var _ datasource.DataSourceWithConfigure = &IntegrationsStatuspagePagesDataSource{}

func NewIntegrationsStatuspagePagesDataSource() datasource.DataSource {
	return &IntegrationsStatuspagePagesDataSource{}
}

// IntegrationsStatuspagePagesDataSource is the data source implementation.
type IntegrationsStatuspagePagesDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// IntegrationsStatuspagePagesDataSourceModel describes the data model.
type IntegrationsStatuspagePagesDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id"`
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (r *IntegrationsStatuspagePagesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integrations_statuspage_pages"
}

// Schema defines the schema for the data source.
func (r *IntegrationsStatuspagePagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IntegrationsStatuspagePages DataSource",

		Attributes: map[string]schema.Attribute{
			"connection_id": schema.StringAttribute{
				Required:    true,
				Description: `Connection UUID`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *IntegrationsStatuspagePagesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IntegrationsStatuspagePagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IntegrationsStatuspagePagesDataSourceModel
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

	request, requestDiags := data.ToOperationsListStatuspageConnectionPagesRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Integrations.ListStatuspageConnectionPages(ctx, *request)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
