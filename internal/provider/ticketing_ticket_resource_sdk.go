// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/operations"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (r *TicketingTicketResourceModel) ToOperationsGetTicketRequest(ctx context.Context) (*operations.GetTicketRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var ticketID string
	ticketID = r.ID.ValueString()

	out := operations.GetTicketRequest{
		TicketID: ticketID,
	}

	return &out, diags
}

func (r *TicketingTicketResourceModel) ToOperationsUpdateTicketRequest(ctx context.Context) (*operations.UpdateTicketRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var ticketID string
	ticketID = r.ID.ValueString()

	updateTicket, updateTicketDiags := r.ToSharedUpdateTicket(ctx)
	diags.Append(updateTicketDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateTicketRequest{
		TicketID:     ticketID,
		UpdateTicket: *updateTicket,
	}

	return &out, diags
}

func (r *TicketingTicketResourceModel) ToSharedCreateTicket(ctx context.Context) (*shared.CreateTicket, diag.Diagnostics) {
	var diags diag.Diagnostics

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	priorityID := new(string)
	if !r.PriorityID.IsUnknown() && !r.PriorityID.IsNull() {
		*priorityID = r.PriorityID.ValueString()
	} else {
		priorityID = nil
	}
	projectID := new(string)
	if !r.ProjectID.IsUnknown() && !r.ProjectID.IsNull() {
		*projectID = r.ProjectID.ValueString()
	} else {
		projectID = nil
	}
	relatedTo := new(string)
	if !r.RelatedTo.IsUnknown() && !r.RelatedTo.IsNull() {
		*relatedTo = r.RelatedTo.ValueString()
	} else {
		relatedTo = nil
	}
	remoteURL := new(string)
	if !r.RemoteURL.IsUnknown() && !r.RemoteURL.IsNull() {
		*remoteURL = r.RemoteURL.ValueString()
	} else {
		remoteURL = nil
	}
	state := new(string)
	if !r.State.IsUnknown() && !r.State.IsNull() {
		*state = r.State.ValueString()
	} else {
		state = nil
	}
	var summary string
	summary = r.Summary.ValueString()

	var tagList []string
	if r.TagList != nil {
		tagList = make([]string, 0, len(r.TagList))
		for _, tagListItem := range r.TagList {
			tagList = append(tagList, tagListItem.ValueString())
		}
	}
	typeVar := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = r.Type.ValueString()
	} else {
		typeVar = nil
	}
	out := shared.CreateTicket{
		Description: description,
		PriorityID:  priorityID,
		ProjectID:   projectID,
		RelatedTo:   relatedTo,
		RemoteURL:   remoteURL,
		State:       state,
		Summary:     summary,
		TagList:     tagList,
		Type:        typeVar,
	}

	return &out, diags
}

func (r *TicketingTicketResourceModel) ToSharedUpdateTicket(ctx context.Context) (*shared.UpdateTicket, diag.Diagnostics) {
	var diags diag.Diagnostics

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	priorityID := new(string)
	if !r.PriorityID.IsUnknown() && !r.PriorityID.IsNull() {
		*priorityID = r.PriorityID.ValueString()
	} else {
		priorityID = nil
	}
	state := new(string)
	if !r.State.IsUnknown() && !r.State.IsNull() {
		*state = r.State.ValueString()
	} else {
		state = nil
	}
	summary := new(string)
	if !r.Summary.IsUnknown() && !r.Summary.IsNull() {
		*summary = r.Summary.ValueString()
	} else {
		summary = nil
	}
	var tagList []string
	if r.TagList != nil {
		tagList = make([]string, 0, len(r.TagList))
		for _, tagListItem := range r.TagList {
			tagList = append(tagList, tagListItem.ValueString())
		}
	}
	typeVar := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = r.Type.ValueString()
	} else {
		typeVar = nil
	}
	out := shared.UpdateTicket{
		Description: description,
		PriorityID:  priorityID,
		State:       state,
		Summary:     summary,
		TagList:     tagList,
		Type:        typeVar,
	}

	return &out, diags
}
