// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NullableRunbooksActionConfig struct {
	DocumentationURL types.String      `tfsdk:"documentation_url"`
	Elements         []RunbooksElement `tfsdk:"elements"`
}
