// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &TicketingProjectsProjectListItemsDataSource{}
var _ datasource.DataSourceWithConfigure = &TicketingProjectsProjectListItemsDataSource{}

func NewTicketingProjectsProjectListItemsDataSource() datasource.DataSource {
	return &TicketingProjectsProjectListItemsDataSource{}
}

// TicketingProjectsProjectListItemsDataSource is the data source implementation.
type TicketingProjectsProjectListItemsDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// TicketingProjectsProjectListItemsDataSourceModel describes the data model.
type TicketingProjectsProjectListItemsDataSourceModel struct {
	Attribute            types.String                                   `tfsdk:"attribute"`
	ConfiguredProjects   types.Bool                                     `queryParam:"style=form,explode=true,name=configured_projects" tfsdk:"configured_projects"`
	ConnectionID         types.String                                   `tfsdk:"connection_id"`
	ConnectionIds        types.String                                   `queryParam:"style=form,explode=true,name=connection_ids" tfsdk:"connection_ids"`
	ConnectionSlug       types.String                                   `tfsdk:"connection_slug"`
	ConnectionType       types.String                                   `tfsdk:"connection_type"`
	Details              *tfTypes.NullableTicketingProjectConfigDetails `tfsdk:"details"`
	ExternalField        types.String                                   `tfsdk:"external_field"`
	ID                   types.String                                   `tfsdk:"id"`
	Logic                map[string]types.String                        `tfsdk:"logic"`
	Name                 types.String                                   `tfsdk:"name"`
	Page                 types.Int32                                    `queryParam:"style=form,explode=true,name=page" tfsdk:"page"`
	PerPage              types.Int32                                    `queryParam:"style=form,explode=true,name=per_page" tfsdk:"per_page"`
	Presentation         types.String                                   `tfsdk:"presentation"`
	Providers            types.String                                   `queryParam:"style=form,explode=true,name=providers" tfsdk:"providers"`
	Query                types.String                                   `queryParam:"style=form,explode=true,name=query" tfsdk:"query"`
	Strategy             types.String                                   `tfsdk:"strategy"`
	SupportsTicketTypes  types.String                                   `queryParam:"style=form,explode=true,name=supports_ticket_types" tfsdk:"supports_ticket_types"`
	TicketingProjectID   types.String                                   `tfsdk:"ticketing_project_id"`
	TicketingProjectName types.String                                   `tfsdk:"ticketing_project_name"`
	Type                 types.String                                   `tfsdk:"type"`
	UpdatedAt            types.String                                   `tfsdk:"updated_at"`
	UserData             map[string]types.String                        `tfsdk:"user_data"`
	Value                types.String                                   `tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *TicketingProjectsProjectListItemsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ticketing_projects_project_list_items"
}

// Schema defines the schema for the data source.
func (r *TicketingProjectsProjectListItemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "TicketingProjectsProjectListItems DataSource",

		Attributes: map[string]schema.Attribute{
			"attribute": schema.StringAttribute{
				Computed: true,
			},
			"configured_projects": schema.BoolAttribute{
				Optional: true,
			},
			"connection_id": schema.StringAttribute{
				Computed: true,
			},
			"connection_ids": schema.StringAttribute{
				Optional: true,
			},
			"connection_slug": schema.StringAttribute{
				Computed: true,
			},
			"connection_type": schema.StringAttribute{
				Computed: true,
			},
			"details": schema.SingleNestedAttribute{
				Computed:    true,
				Description: `A config object containing details about the project config. Can be one of: Ticketing::JiraCloud::ProjectConfig, Ticketing::JiraOnprem::ProjectConfig, or Ticketing::Shortcut::ProjectConfig`,
			},
			"external_field": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"logic": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An unstructured object of key/value pairs describing the logic for applying the rule.`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"page": schema.Int32Attribute{
				Optional: true,
			},
			"per_page": schema.Int32Attribute{
				Optional: true,
			},
			"presentation": schema.StringAttribute{
				Computed: true,
			},
			"providers": schema.StringAttribute{
				Optional: true,
			},
			"query": schema.StringAttribute{
				Optional: true,
			},
			"strategy": schema.StringAttribute{
				Computed: true,
			},
			"supports_ticket_types": schema.StringAttribute{
				Optional: true,
			},
			"ticketing_project_id": schema.StringAttribute{
				Computed: true,
			},
			"ticketing_project_name": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"user_data": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"value": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *TicketingProjectsProjectListItemsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TicketingProjectsProjectListItemsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TicketingProjectsProjectListItemsDataSourceModel
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

	request, requestDiags := data.ToOperationsListTicketingProjectsRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Ticketing.ListTicketingProjects(ctx, *request)
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
