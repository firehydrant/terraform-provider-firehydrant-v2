// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServiceResource{}
var _ resource.ResourceWithImportState = &ServiceResource{}

func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

// ServiceResource defines the resource implementation.
type ServiceResource struct {
	// Provider configured SDK client.
	client *sdk.Firehydrant
}

// ServiceResourceModel describes the resource data model.
type ServiceResourceModel struct {
	ActiveIncidents           []types.String                                `tfsdk:"active_incidents"`
	AlertOnAdd                types.Bool                                    `tfsdk:"alert_on_add"`
	AllowedParams             []types.String                                `tfsdk:"allowed_params"`
	AutoAddRespondingTeam     types.Bool                                    `tfsdk:"auto_add_responding_team"`
	Checklists                []tfTypes.ChecklistTemplate                   `tfsdk:"checklists"`
	CompletedChecks           types.Int32                                   `tfsdk:"completed_checks"`
	CreatedAt                 types.String                                  `tfsdk:"created_at"`
	Description               types.String                                  `tfsdk:"description"`
	ExternalResources         []tfTypes.ExternalResource                    `tfsdk:"external_resources"`
	ExternalResourcesInput    []tfTypes.CreateServiceExternalResourcesInput `tfsdk:"external_resources_input"`
	Functionalities           []tfTypes.Functionality                       `tfsdk:"functionalities"`
	FunctionalitiesInput      []tfTypes.CreateServiceFunctionalitiesInput   `tfsdk:"functionalities_input"`
	ID                        types.String                                  `tfsdk:"id"`
	Labels                    map[string]types.String                       `tfsdk:"labels"`
	LastImport                *tfTypes.NullableImportsImportableResource    `tfsdk:"last_import"`
	Links                     []tfTypes.Links                               `tfsdk:"links"`
	LinksInput                []tfTypes.CreateServiceLinksInput             `tfsdk:"links_input"`
	ManagedBy                 types.String                                  `tfsdk:"managed_by"`
	ManagedBySettings         *tfTypes.ServiceManagedBySettings             `tfsdk:"managed_by_settings"`
	Name                      types.String                                  `tfsdk:"name"`
	Owner                     *tfTypes.NullableTeamLite                     `tfsdk:"owner"`
	OwnerInput                *tfTypes.CreateServiceOwnerInput              `tfsdk:"owner_input"`
	ServiceChecklistUpdatedAt types.String                                  `tfsdk:"service_checklist_updated_at"`
	ServiceTier               types.Int32                                   `tfsdk:"service_tier"`
	Slug                      types.String                                  `tfsdk:"slug"`
	Teams                     []tfTypes.TeamLite                            `tfsdk:"teams"`
	TeamsInput                []tfTypes.CreateServiceTeamsInput             `tfsdk:"teams_input"`
	UpdatedAt                 types.String                                  `tfsdk:"updated_at"`
	UpdatedBy                 *tfTypes.NullableAuthor                       `tfsdk:"updated_by"`
}

func (r *ServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (r *ServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Service Resource",
		Attributes: map[string]schema.Attribute{
			"active_incidents": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `List of active incident guids`,
			},
			"alert_on_add": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"allowed_params": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"auto_add_responding_team": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"checklists": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"checks": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Computed: true,
									},
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"status": schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
						"connected_services": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"alert_on_add": schema.BoolAttribute{
										Computed: true,
									},
									"allowed_params": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"auto_add_responding_team": schema.BoolAttribute{
										Computed: true,
									},
									"completed_checks": schema.Int32Attribute{
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
									"id": schema.StringAttribute{
										Computed: true,
									},
									"labels": schema.SingleNestedAttribute{
										Computed:    true,
										Description: `An object of label key and values`,
									},
									"name": schema.StringAttribute{
										Computed: true,
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
											"id": schema.StringAttribute{
												Computed: true,
											},
											"in_support_hours": schema.BoolAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"signals_ical_url": schema.StringAttribute{
												Computed: true,
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
									"service_checklist_updated_at": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
									},
									"service_tier": schema.Int32Attribute{
										Computed: true,
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
							Description: `List of services that use this checklist`,
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
								"id": schema.StringAttribute{
									Computed: true,
								},
								"in_support_hours": schema.BoolAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"signals_ical_url": schema.StringAttribute{
									Computed: true,
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
						"updated_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
					},
				},
				Description: `List of checklists associated with a service`,
			},
			"completed_checks": schema.Int32Attribute{
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
				Optional: true,
			},
			"external_resources": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"connection_full_favicon_url": schema.StringAttribute{
							Computed: true,
						},
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
			"external_resources_input": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"connection_type": schema.StringAttribute{
							Optional:    true,
							Description: `The integration slug for the external resource. Can be one of: github, opsgenie, pager_duty, victorops. Not required if the resource has already been imported.`,
						},
						"remote_id": schema.StringAttribute{
							Required: true,
						},
					},
				},
				Description: `An array of external resources to attach to this service.`,
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
									"connection_full_favicon_url": schema.StringAttribute{
										Computed: true,
									},
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
							Validators: []validator.Map{
								mapvalidator.ValueStringsAre(validators.IsValidJSON()),
							},
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
								"id": schema.StringAttribute{
									Computed: true,
								},
								"in_support_hours": schema.BoolAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"signals_ical_url": schema.StringAttribute{
									Computed: true,
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
						"services": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"alert_on_add": schema.BoolAttribute{
										Computed: true,
									},
									"allowed_params": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
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
									"id": schema.StringAttribute{
										Computed: true,
									},
									"labels": schema.SingleNestedAttribute{
										Computed:    true,
										Description: `An object of label key and values`,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"service_tier": schema.Int32Attribute{
										Computed: true,
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
							Description: `Services this functionality provides`,
						},
						"slug": schema.StringAttribute{
							Computed: true,
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
									"id": schema.StringAttribute{
										Computed: true,
									},
									"in_support_hours": schema.BoolAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"signals_ical_url": schema.StringAttribute{
										Computed: true,
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
							Description: `List of teams attached to the functionality`,
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
				Description: `List of functionalities attached to the service`,
			},
			"functionalities_input": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional:    true,
							Description: `If you are trying to reuse a functionality, you may set the ID to attach it to the service`,
						},
						"summary": schema.StringAttribute{
							Optional:    true,
							Description: `If you are trying to create a new functionality and attach it to this service, set the summary key`,
						},
					},
				},
				Description: `An array of functionalities`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A hash of label keys and values`,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"last_import": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"import_errors": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_at": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
								},
								"data": schema.SingleNestedAttribute{
									Computed:    true,
									Description: `Additional error data`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"message": schema.StringAttribute{
									Computed: true,
								},
								"resource": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Computed: true,
										},
										"resource_id": schema.StringAttribute{
											Computed: true,
										},
										"resource_type": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"imported_at": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"remote_id": schema.StringAttribute{
						Computed: true,
					},
					"state": schema.StringAttribute{
						Computed: true,
					},
				},
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
				Description: `List of links attached to this service.`,
			},
			"links_input": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"href_url": schema.StringAttribute{
							Required:    true,
							Description: `URL`,
						},
						"icon_url": schema.StringAttribute{
							Optional:    true,
							Description: `An optional URL to an icon representing this link`,
						},
						"name": schema.StringAttribute{
							Required:    true,
							Description: `Short name used to display and identify this link`,
						},
					},
				},
				Description: `An array of links to associate with this service`,
			},
			"managed_by": schema.StringAttribute{
				Computed:    true,
				Description: `If set, this field indicates that the service is managed by an integration and thus cannot be set manually`,
			},
			"managed_by_settings": schema.SingleNestedAttribute{
				Computed:    true,
				Description: `Indicates the settings of the catalog that manages this service`,
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
					"id": schema.StringAttribute{
						Computed: true,
					},
					"in_support_hours": schema.BoolAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"signals_ical_url": schema.StringAttribute{
						Computed: true,
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
			"owner_input": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
				},
				Description: `An object representing a Team that owns the service`,
			},
			"service_checklist_updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"service_tier": schema.Int32Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Integer representing service tier. Lower values represent higher criticality. If not specified the default value will be 5. must be one of ["0", "1", "2", "3", "4", "5"]`,
				Validators: []validator.Int32{
					int32validator.OneOf(
						0,
						1,
						2,
						3,
						4,
						5,
					),
				},
			},
			"slug": schema.StringAttribute{
				Computed: true,
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
						"id": schema.StringAttribute{
							Computed: true,
						},
						"in_support_hours": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"signals_ical_url": schema.StringAttribute{
							Computed: true,
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
				Description: `List of teams attached to the service`,
			},
			"teams_input": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
				Description: `An array of teams to attach to this service.`,
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
	}
}

func (r *ServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Firehydrant)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Firehydrant, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ServiceResourceModel
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

	request, requestDiags := data.ToSharedCreateService(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.CatalogEntries.CreateService(ctx, *request)
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
	if !(res.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedService(ctx, res.Service)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetServiceRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.CatalogEntries.GetService(ctx, *request1)
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
	if !(res1.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedService(ctx, res1.Service)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ServiceResourceModel
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

	request, requestDiags := data.ToOperationsGetServiceRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.CatalogEntries.GetService(ctx, *request)
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
	if !(res.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedService(ctx, res.Service)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ServiceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsUpdateServiceRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.CatalogEntries.UpdateService(ctx, *request)
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
	if !(res.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedService(ctx, res.Service)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetServiceRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.CatalogEntries.GetService(ctx, *request1)
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
	if !(res1.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedService(ctx, res1.Service)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ServiceResourceModel
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

	request, requestDiags := data.ToOperationsDeleteServiceRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.CatalogEntries.DeleteService(ctx, *request)
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

func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
