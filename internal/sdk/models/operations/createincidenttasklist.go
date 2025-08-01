// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type CreateIncidentTaskListRequest struct {
	IncidentID             string                        `pathParam:"style=simple,explode=false,name=incident_id"`
	CreateIncidentTaskList shared.CreateIncidentTaskList `request:"mediaType=application/json"`
}

func (o *CreateIncidentTaskListRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentTaskListRequest) GetCreateIncidentTaskList() shared.CreateIncidentTaskList {
	if o == nil {
		return shared.CreateIncidentTaskList{}
	}
	return o.CreateIncidentTaskList
}

type CreateIncidentTaskListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Add all tasks from list to incident
	Task *shared.Task
}

func (o *CreateIncidentTaskListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIncidentTaskListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIncidentTaskListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIncidentTaskListResponse) GetTask() *shared.Task {
	if o == nil {
		return nil
	}
	return o.Task
}
