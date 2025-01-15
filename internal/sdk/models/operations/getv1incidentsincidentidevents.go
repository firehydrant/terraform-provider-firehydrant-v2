// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1IncidentsIncidentIDEventsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// A comma separated list of types of events to filter by
	Types   *string `queryParam:"style=form,explode=true,name=types"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetTypes() *string {
	if o == nil {
		return nil
	}
	return o.Types
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1IncidentsIncidentIDEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all events for an incident. An event is a timeline entry. This can be filtered with params to retrieve events of a certain type.
	IncidentEventEntityPaginated *shared.IncidentEventEntityPaginated
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetIncidentEventEntityPaginated() *shared.IncidentEventEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentEventEntityPaginated
}
