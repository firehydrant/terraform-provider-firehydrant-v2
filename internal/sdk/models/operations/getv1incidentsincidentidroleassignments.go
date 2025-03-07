// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

// Status - Filter on status of the role assignment
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type GetV1IncidentsIncidentIDRoleAssignmentsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// Filter on status of the role assignment
	Status *Status `queryParam:"style=form,explode=true,name=status"`
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetV1IncidentsIncidentIDRoleAssignmentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a list of all of the current role assignments for the incident
	IncidentsRoleAssignmentEntityPaginated *shared.IncidentsRoleAssignmentEntityPaginated
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsResponse) GetIncidentsRoleAssignmentEntityPaginated() *shared.IncidentsRoleAssignmentEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsRoleAssignmentEntityPaginated
}
