// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDRequest struct {
	RetrospectiveID                                         string                                                         `pathParam:"style=simple,explode=false,name=retrospective_id"`
	IncidentID                                              string                                                         `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID shared.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDRequest) GetRetrospectiveID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveID
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDRequest) GetPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID() shared.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID {
	if o == nil {
		return shared.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID{}
	}
	return o.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveID
}

type PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a retrospective attached to an incident
	IncidentsRetrospectiveEntity *shared.IncidentsRetrospectiveEntity
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDResponse) GetIncidentsRetrospectiveEntity() *shared.IncidentsRetrospectiveEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRetrospectiveEntity
}
