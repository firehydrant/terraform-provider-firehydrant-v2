// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (r *PriorityResourceModel) ToOperationsDeletePriorityRequest(ctx context.Context) (*operations.DeletePriorityRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var prioritySlug string
	prioritySlug = r.Slug.ValueString()

	out := operations.DeletePriorityRequest{
		PrioritySlug: prioritySlug,
	}

	return &out, diags
}

func (r *PriorityResourceModel) ToOperationsGetPriorityRequest(ctx context.Context) (*operations.GetPriorityRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var prioritySlug string
	prioritySlug = r.Slug.ValueString()

	out := operations.GetPriorityRequest{
		PrioritySlug: prioritySlug,
	}

	return &out, diags
}

func (r *PriorityResourceModel) ToOperationsUpdatePriorityRequest(ctx context.Context) (*operations.UpdatePriorityRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var prioritySlug string
	prioritySlug = r.Slug.ValueString()

	updatePriority, updatePriorityDiags := r.ToSharedUpdatePriority(ctx)
	diags.Append(updatePriorityDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdatePriorityRequest{
		PrioritySlug:   prioritySlug,
		UpdatePriority: *updatePriority,
	}

	return &out, diags
}

func (r *PriorityResourceModel) ToSharedCreatePriority(ctx context.Context) (*shared.CreatePriority, diag.Diagnostics) {
	var diags diag.Diagnostics

	defaultVar := new(bool)
	if !r.Default.IsUnknown() && !r.Default.IsNull() {
		*defaultVar = r.Default.ValueBool()
	} else {
		defaultVar = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	var slug string
	slug = r.Slug.ValueString()

	out := shared.CreatePriority{
		Default:     defaultVar,
		Description: description,
		Slug:        slug,
	}

	return &out, diags
}

func (r *PriorityResourceModel) ToSharedUpdatePriority(ctx context.Context) (*shared.UpdatePriority, diag.Diagnostics) {
	var diags diag.Diagnostics

	defaultVar := new(bool)
	if !r.Default.IsUnknown() && !r.Default.IsNull() {
		*defaultVar = r.Default.ValueBool()
	} else {
		defaultVar = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	out := shared.UpdatePriority{
		Default:     defaultVar,
		Description: description,
		Slug:        slug,
	}

	return &out, diags
}
