// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IncidentsMilestone struct {
	CreatedAt  types.String `tfsdk:"created_at"`
	Duration   types.String `tfsdk:"duration"`
	ID         types.String `tfsdk:"id"`
	OccurredAt types.String `tfsdk:"occurred_at"`
	Type       types.String `tfsdk:"type"`
	UpdatedAt  types.String `tfsdk:"updated_at"`
}
