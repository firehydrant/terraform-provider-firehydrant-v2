// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NullableSignalsAPIEscalationPolicyHandoffStep struct {
	ID     types.String              `tfsdk:"id"`
	Target *NullableSignalsAPITarget `tfsdk:"target"`
}
