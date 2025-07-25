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
var _ datasource.DataSource = &IncidentsChannelDataSource{}
var _ datasource.DataSourceWithConfigure = &IncidentsChannelDataSource{}

func NewIncidentsChannelDataSource() datasource.DataSource {
	return &IncidentsChannelDataSource{}
}

// IncidentsChannelDataSource is the data source implementation.
type IncidentsChannelDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// IncidentsChannelDataSourceModel describes the data model.
type IncidentsChannelDataSourceModel struct {
	IconURL    types.String `tfsdk:"icon_url"`
	ID         types.String `tfsdk:"id"`
	IncidentID types.String `tfsdk:"incident_id"`
	Name       types.String `tfsdk:"name"`
	Source     types.String `tfsdk:"source"`
	SourceID   types.String `tfsdk:"source_id"`
	SourceName types.String `tfsdk:"source_name"`
	Status     types.String `tfsdk:"status"`
	URL        types.String `tfsdk:"url"`
}

// Metadata returns the data source type name.
func (r *IncidentsChannelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_incidents_channel"
}

// Schema defines the schema for the data source.
func (r *IncidentsChannelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IncidentsChannel DataSource",

		Attributes: map[string]schema.Attribute{
			"icon_url": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"incident_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"source_name": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *IncidentsChannelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IncidentsChannelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IncidentsChannelDataSourceModel
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

	request, requestDiags := data.ToOperationsGetIncidentChannelRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Incidents.GetIncidentChannel(ctx, *request)
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
	if !(res.IncidentsChannel != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedIncidentsChannel(ctx, res.IncidentsChannel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
