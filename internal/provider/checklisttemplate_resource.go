// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/operations"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/validators"
	speakeasy_objectvalidators "github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ChecklistTemplateResource{}
var _ resource.ResourceWithImportState = &ChecklistTemplateResource{}

func NewChecklistTemplateResource() resource.Resource {
	return &ChecklistTemplateResource{}
}

// ChecklistTemplateResource defines the resource implementation.
type ChecklistTemplateResource struct {
	client *sdk.FirehydrantTerraformSDK
}

// ChecklistTemplateResourceModel describes the resource data model.
type ChecklistTemplateResourceModel struct {
	Checks      []tfTypes.Checks    `tfsdk:"checks"`
	CreatedAt   types.String        `tfsdk:"created_at"`
	Description types.String        `tfsdk:"description"`
	ID          types.String        `tfsdk:"id"`
	Name        types.String        `tfsdk:"name"`
	Owner       *tfTypes.TeamEntity `tfsdk:"owner"`
	TeamID      types.String        `tfsdk:"team_id"`
	UpdatedAt   types.String        `tfsdk:"updated_at"`
}

func (r *ChecklistTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_checklist_template"
}

func (r *ChecklistTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ChecklistTemplate Resource",
		Attributes: map[string]schema.Attribute{
			"checks": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `The description of the check`,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `The name of the check. Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"status": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
				Description: `An array of checks for the checklist template`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Checklist Template UUID`,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"owner": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"created_at": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"created_by": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"email": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"source": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"functionalities": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"active_incidents": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `List of active incident guids`,
								},
								"alert_on_add": schema.BoolAttribute{
									Computed: true,
								},
								"auto_add_responding_team": schema.BoolAttribute{
									Computed: true,
								},
								"created_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"external_resources": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Computed: true,
											},
											"connection_name": schema.StringAttribute{
												Computed: true,
											},
											"connection_type": schema.StringAttribute{
												Computed: true,
											},
											"created_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"remote_id": schema.StringAttribute{
												Computed: true,
											},
											"remote_url": schema.StringAttribute{
												Computed: true,
											},
											"updated_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
										},
									},
									Description: `Information about known linkages to representations of services outside of FireHydrant.`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"labels": schema.MapAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `An object of label key and values`,
								},
								"links": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"href_url": schema.StringAttribute{
												Computed: true,
											},
											"icon_url": schema.StringAttribute{
												Computed: true,
											},
											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `List of links attached to this functionality.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"owner": schema.SingleNestedAttribute{
									Computed:    true,
									Description: `TeamEntity model`,
								},
								"slug": schema.StringAttribute{
									Computed: true,
								},
								"updated_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"updated_by": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"email": schema.StringAttribute{
											Computed: true,
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
										"source": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"memberships": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default_incident_role": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"created_at": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
										"description": schema.StringAttribute{
											Computed: true,
										},
										"discarded_at": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
										"summary": schema.StringAttribute{
											Computed: true,
										},
										"updated_at": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
									},
									Description: `IncidentRoleEntity model`,
								},
								"schedule": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"discarded": schema.BoolAttribute{
											Computed: true,
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"integration": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"user": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"created_at": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
										"email": schema.StringAttribute{
											Computed: true,
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
										"signals_enabled_notification_types": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"slack_linked": schema.BoolAttribute{
											Computed: true,
										},
										"slack_user_id": schema.StringAttribute{
											Computed: true,
										},
										"updated_at": schema.StringAttribute{
											Computed: true,
											Validators: []validator.String{
												validators.IsRFC3339(),
											},
										},
									},
								},
							},
						},
					},
					"ms_teams_channel": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"channel_id": schema.StringAttribute{
								Computed: true,
							},
							"channel_name": schema.StringAttribute{
								Computed: true,
							},
							"channel_url": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"ms_team_id": schema.StringAttribute{
								Computed: true,
							},
							"status": schema.StringAttribute{
								Computed: true,
							},
							"team_name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"owned_functionalities": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"active_incidents": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `List of active incident guids`,
								},
								"alert_on_add": schema.BoolAttribute{
									Computed: true,
								},
								"auto_add_responding_team": schema.BoolAttribute{
									Computed: true,
								},
								"created_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"external_resources": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Computed: true,
											},
											"connection_name": schema.StringAttribute{
												Computed: true,
											},
											"connection_type": schema.StringAttribute{
												Computed: true,
											},
											"created_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"remote_id": schema.StringAttribute{
												Computed: true,
											},
											"remote_url": schema.StringAttribute{
												Computed: true,
											},
											"updated_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
										},
									},
									Description: `Information about known linkages to representations of services outside of FireHydrant.`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"labels": schema.MapAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `An object of label key and values`,
								},
								"links": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"href_url": schema.StringAttribute{
												Computed: true,
											},
											"icon_url": schema.StringAttribute{
												Computed: true,
											},
											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `List of links attached to this functionality.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"owner": schema.SingleNestedAttribute{
									Computed:    true,
									Description: `TeamEntity model`,
								},
								"slug": schema.StringAttribute{
									Computed: true,
								},
								"updated_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"updated_by": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"email": schema.StringAttribute{
											Computed: true,
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
										"source": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"owned_runbooks": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attachment_rule": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"logic": schema.SingleNestedAttribute{
											Computed:    true,
											Description: `An unstructured object of key/value pairs describing the logic for applying the rule.`,
										},
										"user_data": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"label": schema.StringAttribute{
													Computed: true,
												},
												"type": schema.StringAttribute{
													Computed: true,
												},
												"value": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
								},
								"categories": schema.StringAttribute{
									Computed:    true,
									Description: `categories the runbook applies to`,
								},
								"created_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"owner": schema.SingleNestedAttribute{
									Computed:    true,
									Description: `TeamEntity model`,
								},
								"summary": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"updated_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
							},
						},
					},
					"signals_ical_url": schema.StringAttribute{
						Computed: true,
					},
					"slack_channel": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"slack_channel_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"slug": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `TeamEntity model`,
			},
			"team_id": schema.StringAttribute{
				Optional:    true,
				Description: `The ID of the Team that owns the checklist template`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *ChecklistTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.FirehydrantTerraformSDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.FirehydrantTerraformSDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ChecklistTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ChecklistTemplateResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedPostV1ChecklistTemplates()
	res, err := r.client.ChecklistTemplates.Create(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ChecklistTemplateEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedChecklistTemplateEntity(res.ChecklistTemplateEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id string
	id = data.ID.ValueString()

	request1 := operations.GetV1ChecklistTemplatesIDRequest{
		ID: id,
	}
	res1, err := r.client.ChecklistTemplates.Get(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.ChecklistTemplateEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedChecklistTemplateEntity(res1.ChecklistTemplateEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChecklistTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ChecklistTemplateResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var id string
	id = data.ID.ValueString()

	request := operations.GetV1ChecklistTemplatesIDRequest{
		ID: id,
	}
	res, err := r.client.ChecklistTemplates.Get(ctx, request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ChecklistTemplateEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedChecklistTemplateEntity(res.ChecklistTemplateEntity)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChecklistTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ChecklistTemplateResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	patchV1ChecklistTemplatesID := *data.ToSharedPatchV1ChecklistTemplatesID()
	request := operations.PatchV1ChecklistTemplatesIDRequest{
		ID:                          id,
		PatchV1ChecklistTemplatesID: patchV1ChecklistTemplatesID,
	}
	res, err := r.client.ChecklistTemplates.Update(ctx, request)
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
	if !(res.ChecklistTemplateEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedChecklistTemplateEntity(res.ChecklistTemplateEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var id1 string
	id1 = data.ID.ValueString()

	request1 := operations.GetV1ChecklistTemplatesIDRequest{
		ID: id1,
	}
	res1, err := r.client.ChecklistTemplates.Get(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.ChecklistTemplateEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedChecklistTemplateEntity(res1.ChecklistTemplateEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChecklistTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ChecklistTemplateResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var id string
	id = data.ID.ValueString()

	request := operations.DeleteV1ChecklistTemplatesIDRequest{
		ID: id,
	}
	res, err := r.client.ChecklistTemplates.Archive(ctx, request)
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

}

func (r *ChecklistTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
