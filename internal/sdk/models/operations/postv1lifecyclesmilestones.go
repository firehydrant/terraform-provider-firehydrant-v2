// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1LifecyclesMilestonesRequestBody struct {
	// The name of the milestone
	Name string `form:"name=name"`
	// A long-form description of the milestone's purpose
	Description string `form:"name=description"`
	// A unique identifier for the milestone. If not provided, one will be generated from the name.
	Slug *string `form:"name=slug"`
	// The ID of the phase to which the milestone should belong
	PhaseID string `form:"name=phase_id"`
	// The position of the milestone within the phase. If not provided, the milestone will be added as the last milestone in the phase.
	Position *int `form:"name=position"`
	// The ID of a later milestone that cannot be started until this milestone has a timestamp populated
	RequiredAtMilestoneID *string `form:"name=required_at_milestone_id"`
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetPhaseID() string {
	if o == nil {
		return ""
	}
	return o.PhaseID
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PostV1LifecyclesMilestonesRequestBody) GetRequiredAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.RequiredAtMilestoneID
}

type PostV1LifecyclesMilestonesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new milestone
	LifecyclesPhaseEntityList *shared.LifecyclesPhaseEntityList
}

func (o *PostV1LifecyclesMilestonesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1LifecyclesMilestonesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1LifecyclesMilestonesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1LifecyclesMilestonesResponse) GetLifecyclesPhaseEntityList() *shared.LifecyclesPhaseEntityList {
	if o == nil {
		return nil
	}
	return o.LifecyclesPhaseEntityList
}
