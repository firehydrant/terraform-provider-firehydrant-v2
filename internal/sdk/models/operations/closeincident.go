// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type CloseIncidentRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *CloseIncidentRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type CloseIncidentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Closes an incident and optionally close all children
	Incident *shared.Incident
}

func (o *CloseIncidentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CloseIncidentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CloseIncidentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CloseIncidentResponse) GetIncident() *shared.Incident {
	if o == nil {
		return nil
	}
	return o.Incident
}
