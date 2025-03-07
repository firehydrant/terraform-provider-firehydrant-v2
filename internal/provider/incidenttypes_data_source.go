// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &IncidentTypesDataSource{}
var _ datasource.DataSourceWithConfigure = &IncidentTypesDataSource{}

func NewIncidentTypesDataSource() datasource.DataSource {
	return &IncidentTypesDataSource{}
}

// IncidentTypesDataSource is the data source implementation.
type IncidentTypesDataSource struct {
	client *sdk.Firehydrant
}

// IncidentTypesDataSourceModel describes the data model.
type IncidentTypesDataSourceModel struct {
	Data       []tfTypes.IncidentTypeEntity `tfsdk:"data"`
	Page       types.Int64                  `tfsdk:"page"`
	Pagination *tfTypes.PaginationEntity    `tfsdk:"pagination"`
	PerPage    types.Int64                  `tfsdk:"per_page"`
	Query      types.String                 `tfsdk:"query"`
}

// Metadata returns the data source type name.
func (r *IncidentTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_incident_types"
}

// Schema defines the schema for the data source.
func (r *IncidentTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IncidentTypes DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"template": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"custom_fields": schema.StringAttribute{
									Computed: true,
								},
								"customer_impact_summary": schema.StringAttribute{
									Computed: true,
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"impacts": schema.ListNestedAttribute{
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
												Computed: true,
											},
										},
									},
								},
								"incident_name": schema.StringAttribute{
									Computed: true,
								},
								"labels": schema.MapAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Arbitrary key:value pairs of labels for your incidents.`,
								},
								"priority": schema.StringAttribute{
									Computed: true,
								},
								"private_incident": schema.BoolAttribute{
									Computed: true,
								},
								"runbook_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"severity": schema.StringAttribute{
									Computed: true,
								},
								"summary": schema.StringAttribute{
									Computed: true,
								},
								"tag_list": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"team_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
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
												Computed: true,
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
												Computed: true,
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
												Computed: true,
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
																},
																"description": schema.StringAttribute{
																	Computed: true,
																},
																"discarded_at": schema.StringAttribute{
																	Computed: true,
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
											},
										},
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"page": schema.Int64Attribute{
				Optional: true,
			},
			"pagination": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						Computed: true,
					},
					"items": schema.Int64Attribute{
						Computed: true,
					},
					"last": schema.Int64Attribute{
						Computed: true,
					},
					"next": schema.Int64Attribute{
						Computed: true,
					},
					"page": schema.Int64Attribute{
						Computed: true,
					},
					"pages": schema.Int64Attribute{
						Computed: true,
					},
					"prev": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"per_page": schema.Int64Attribute{
				Optional: true,
			},
			"query": schema.StringAttribute{
				Optional:    true,
				Description: `A query to search incident types by their name`,
			},
		},
	}
}

func (r *IncidentTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IncidentTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IncidentTypesDataSourceModel
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

	query := new(string)
	if !data.Query.IsUnknown() && !data.Query.IsNull() {
		*query = data.Query.ValueString()
	} else {
		query = nil
	}
	page := new(int)
	if !data.Page.IsUnknown() && !data.Page.IsNull() {
		*page = int(data.Page.ValueInt64())
	} else {
		page = nil
	}
	perPage := new(int)
	if !data.PerPage.IsUnknown() && !data.PerPage.IsNull() {
		*perPage = int(data.PerPage.ValueInt64())
	} else {
		perPage = nil
	}
	request := operations.GetV1IncidentTypesRequest{
		Query:   query,
		Page:    page,
		PerPage: perPage,
	}
	res, err := r.client.IncidentTypes.List(ctx, request)
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
	if !(res.IncidentTypeEntityPaginated != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedIncidentTypeEntityPaginated(res.IncidentTypeEntityPaginated)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
