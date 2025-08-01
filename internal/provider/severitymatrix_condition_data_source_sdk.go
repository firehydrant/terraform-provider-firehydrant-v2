// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (r *SeverityMatrixConditionDataSourceModel) ToOperationsGetSeverityMatrixConditionRequest(ctx context.Context) (*operations.GetSeverityMatrixConditionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var conditionID string
	conditionID = r.ID.ValueString()

	out := operations.GetSeverityMatrixConditionRequest{
		ConditionID: conditionID,
	}

	return &out, diags
}
