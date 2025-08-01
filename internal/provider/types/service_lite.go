// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ServiceLite struct {
	AlertOnAdd            types.Bool         `tfsdk:"alert_on_add"`
	AllowedParams         []types.String     `tfsdk:"allowed_params"`
	AutoAddRespondingTeam types.Bool         `tfsdk:"auto_add_responding_team"`
	CreatedAt             types.String       `tfsdk:"created_at"`
	Description           types.String       `tfsdk:"description"`
	ID                    types.String       `tfsdk:"id"`
	Labels                *ServiceLiteLabels `tfsdk:"labels"`
	Name                  types.String       `tfsdk:"name"`
	ServiceTier           types.Int32        `tfsdk:"service_tier"`
	Slug                  types.String       `tfsdk:"slug"`
	UpdatedAt             types.String       `tfsdk:"updated_at"`
}
