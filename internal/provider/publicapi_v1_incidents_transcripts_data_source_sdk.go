// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (r *PublicAPIV1IncidentsTranscriptsDataSourceModel) ToOperationsListTranscriptEntriesRequest(ctx context.Context) (*operations.ListTranscriptEntriesRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	after := new(string)
	if !r.After.IsUnknown() && !r.After.IsNull() {
		*after = r.After.ValueString()
	} else {
		after = nil
	}
	before := new(string)
	if !r.Before.IsUnknown() && !r.Before.IsNull() {
		*before = r.Before.ValueString()
	} else {
		before = nil
	}
	sort := new(string)
	if !r.Sort.IsUnknown() && !r.Sort.IsNull() {
		*sort = r.Sort.ValueString()
	} else {
		sort = nil
	}
	var incidentID string
	incidentID = r.IncidentID.ValueString()

	out := operations.ListTranscriptEntriesRequest{
		After:      after,
		Before:     before,
		Sort:       sort,
		IncidentID: incidentID,
	}

	return &out, diags
}
