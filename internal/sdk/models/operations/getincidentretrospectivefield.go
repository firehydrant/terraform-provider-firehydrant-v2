// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetIncidentRetrospectiveFieldRequest struct {
	RetrospectiveID string `pathParam:"style=simple,explode=false,name=retrospective_id"`
	FieldID         string `pathParam:"style=simple,explode=false,name=field_id"`
	IncidentID      string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetIncidentRetrospectiveFieldRequest) GetRetrospectiveID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveID
}

func (o *GetIncidentRetrospectiveFieldRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *GetIncidentRetrospectiveFieldRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetIncidentRetrospectiveFieldResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a field on an incident retrospective
	IncidentsRetrospectiveField *shared.IncidentsRetrospectiveField
}

func (o *GetIncidentRetrospectiveFieldResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIncidentRetrospectiveFieldResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIncidentRetrospectiveFieldResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetIncidentRetrospectiveFieldResponse) GetIncidentsRetrospectiveField() *shared.IncidentsRetrospectiveField {
	if o == nil {
		return nil
	}
	return o.IncidentsRetrospectiveField
}
