// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
var _ resource.Resource = &IncidentTypeResource{}
var _ resource.ResourceWithImportState = &IncidentTypeResource{}

func NewIncidentTypeResource() resource.Resource {
	return &IncidentTypeResource{}
}

// IncidentTypeResource defines the resource implementation.
type IncidentTypeResource struct {
	client *sdk.FirehydrantTerraformSDK
}

// IncidentTypeResourceModel describes the resource data model.
type IncidentTypeResourceModel struct {
	CreatedAt      types.String                                    `tfsdk:"created_at"`
	ID             types.String                                    `tfsdk:"id"`
	Name           types.String                                    `tfsdk:"name"`
	Template       tfTypes.Template                                `tfsdk:"template"`
	TemplateValues *tfTypes.IncidentTypeEntityTemplateValuesEntity `tfsdk:"template_values"`
	UpdatedAt      types.String                                    `tfsdk:"updated_at"`
}

func (r *IncidentTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_incident_type"
}

func (r *IncidentTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IncidentType Resource",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"template": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"custom_fields": schema.StringAttribute{
						Computed: true,
					},
					"customer_impact_summary": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"impacts": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"condition_id": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The id of the condition. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"condition_name": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The id of impact. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["environment", "functionality", "service"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"environment",
											"functionality",
											"service",
										),
									},
								},
							},
						},
						Description: `An array of impact/condition combinations`,
					},
					"incident_name": schema.StringAttribute{
						Computed: true,
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `A labels hash of keys and values`,
					},
					"priority": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"private_incident": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"runbook_ids": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of ids of Runbooks to attach to incidents created from this type`,
					},
					"severity": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"summary": schema.StringAttribute{
						Computed: true,
					},
					"tag_list": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of tags for the incident`,
					},
					"team_ids": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of ids of teams to be assigned to incidents`,
					},
				},
			},
			"template_values": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"environments": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_id": schema.StringAttribute{
									Computed: true,
								},
								"condition_name": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["environment", "functionality", "service"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"environment",
											"functionality",
											"service",
										),
									},
								},
							},
						},
					},
					"functionalities": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_id": schema.StringAttribute{
									Computed: true,
								},
								"condition_name": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["environment", "functionality", "service"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"environment",
											"functionality",
											"service",
										),
									},
								},
							},
						},
					},
					"runbooks": schema.SingleNestedAttribute{
						Computed:    true,
						Description: `A hash mapping runbook IDs to runbook names.`,
					},
					"services": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_id": schema.StringAttribute{
									Computed: true,
								},
								"condition_name": schema.StringAttribute{
									Computed: true,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["environment", "functionality", "service"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"environment",
											"functionality",
											"service",
										),
									},
								},
							},
						},
					},
					"teams": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
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
						},
					},
				},
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

func (r *IncidentTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IncidentTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *IncidentTypeResourceModel
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

	request := *data.ToSharedPostV1IncidentTypes()
	res, err := r.client.IncidentTypes.Create(ctx, request)
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
	if !(res.IncidentTypeEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedIncidentTypeEntity(res.IncidentTypeEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IncidentTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *IncidentTypeResourceModel
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

	request := operations.GetV1IncidentTypesIDRequest{
		ID: id,
	}
	res, err := r.client.IncidentTypes.Get(ctx, request)
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
	if !(res.IncidentTypeEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedIncidentTypeEntity(res.IncidentTypeEntity)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IncidentTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *IncidentTypeResourceModel
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

	patchV1IncidentTypesID := *data.ToSharedPatchV1IncidentTypesID()
	request := operations.PatchV1IncidentTypesIDRequest{
		ID:                     id,
		PatchV1IncidentTypesID: patchV1IncidentTypesID,
	}
	res, err := r.client.IncidentTypes.Update(ctx, request)
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
	if !(res.IncidentTypeEntity != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedIncidentTypeEntity(res.IncidentTypeEntity)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IncidentTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *IncidentTypeResourceModel
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

	request := operations.DeleteV1IncidentTypesIDRequest{
		ID: id,
	}
	res, err := r.client.IncidentTypes.Delete(ctx, request)
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

func (r *IncidentTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
