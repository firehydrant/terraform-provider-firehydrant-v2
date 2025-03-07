// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Checklists struct {
	Checks      []ChecklistCheckEntity `tfsdk:"checks"`
	CreatedAt   types.String           `tfsdk:"created_at"`
	Description types.String           `tfsdk:"description"`
	ID          types.String           `tfsdk:"id"`
	Name        types.String           `tfsdk:"name"`
	Owner       *TeamEntity            `tfsdk:"owner"`
	UpdatedAt   types.String           `tfsdk:"updated_at"`
}
