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
var _ datasource.DataSource = &IncidentsConferenceBridgesDataSource{}
var _ datasource.DataSourceWithConfigure = &IncidentsConferenceBridgesDataSource{}

func NewIncidentsConferenceBridgesDataSource() datasource.DataSource {
	return &IncidentsConferenceBridgesDataSource{}
}

// IncidentsConferenceBridgesDataSource is the data source implementation.
type IncidentsConferenceBridgesDataSource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// IncidentsConferenceBridgesDataSourceModel describes the data model.
type IncidentsConferenceBridgesDataSourceModel struct {
	Attachments              []tfTypes.IncidentsConferenceBridgeAttachment `tfsdk:"attachments"`
	HasTranslatedTranscripts types.Bool                                    `tfsdk:"has_translated_transcripts"`
	ID                       types.String                                  `tfsdk:"id"`
	IncidentID               types.String                                  `tfsdk:"incident_id"`
	LanguageCodes            []types.String                                `tfsdk:"language_codes"`
	PreviousHostAssignment   types.String                                  `tfsdk:"previous_host_assignment"`
	TranscriptionStatus      types.String                                  `tfsdk:"transcription_status"`
	TranscriptionSubCode     types.String                                  `tfsdk:"transcription_sub_code"`
}

// Metadata returns the data source type name.
func (r *IncidentsConferenceBridgesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_incidents_conference_bridges"
}

// Schema defines the schema for the data source.
func (r *IncidentsConferenceBridgesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IncidentsConferenceBridges DataSource",

		Attributes: map[string]schema.Attribute{
			"attachments": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{},
				},
				Description: `A list of objects attached to this item. Can be one of: Link, CustomerSupportIssue, or GenericAttachment`,
			},
			"has_translated_transcripts": schema.BoolAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"incident_id": schema.StringAttribute{
				Required: true,
			},
			"language_codes": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of language codes that have translated transcripts for this conference bridge`,
			},
			"previous_host_assignment": schema.StringAttribute{
				Computed: true,
			},
			"transcription_status": schema.StringAttribute{
				Computed: true,
			},
			"transcription_sub_code": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *IncidentsConferenceBridgesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IncidentsConferenceBridgesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IncidentsConferenceBridgesDataSourceModel
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

	request, requestDiags := data.ToOperationsListIncidentConferenceBridgesRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Incidents.ListIncidentConferenceBridges(ctx, *request)
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
