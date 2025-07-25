// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomFieldsFieldValue struct {
	Description types.String   `tfsdk:"description"`
	DisplayName types.String   `tfsdk:"display_name"`
	FieldID     types.String   `tfsdk:"field_id"`
	Name        types.String   `tfsdk:"name"`
	Slug        types.String   `tfsdk:"slug"`
	Value       types.String   `tfsdk:"value"`
	ValueArray  []types.String `tfsdk:"value_array"`
	ValueString types.String   `tfsdk:"value_string"`
	ValueType   types.String   `tfsdk:"value_type"`
}
