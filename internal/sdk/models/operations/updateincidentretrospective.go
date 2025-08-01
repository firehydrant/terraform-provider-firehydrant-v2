// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type UpdateIncidentRetrospectiveRequest struct {
	RetrospectiveID             string                             `pathParam:"style=simple,explode=false,name=retrospective_id"`
	IncidentID                  string                             `pathParam:"style=simple,explode=false,name=incident_id"`
	UpdateIncidentRetrospective shared.UpdateIncidentRetrospective `request:"mediaType=application/json"`
}

func (o *UpdateIncidentRetrospectiveRequest) GetRetrospectiveID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveID
}

func (o *UpdateIncidentRetrospectiveRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *UpdateIncidentRetrospectiveRequest) GetUpdateIncidentRetrospective() shared.UpdateIncidentRetrospective {
	if o == nil {
		return shared.UpdateIncidentRetrospective{}
	}
	return o.UpdateIncidentRetrospective
}

type UpdateIncidentRetrospectiveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a retrospective attached to an incident
	IncidentsRetrospective *shared.IncidentsRetrospective
}

func (o *UpdateIncidentRetrospectiveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIncidentRetrospectiveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIncidentRetrospectiveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIncidentRetrospectiveResponse) GetIncidentsRetrospective() *shared.IncidentsRetrospective {
	if o == nil {
		return nil
	}
	return o.IncidentsRetrospective
}
