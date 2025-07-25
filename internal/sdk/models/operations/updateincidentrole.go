// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type UpdateIncidentRoleRequest struct {
	IncidentRoleID     string                    `pathParam:"style=simple,explode=false,name=incident_role_id"`
	UpdateIncidentRole shared.UpdateIncidentRole `request:"mediaType=application/json"`
}

func (o *UpdateIncidentRoleRequest) GetIncidentRoleID() string {
	if o == nil {
		return ""
	}
	return o.IncidentRoleID
}

func (o *UpdateIncidentRoleRequest) GetUpdateIncidentRole() shared.UpdateIncidentRole {
	if o == nil {
		return shared.UpdateIncidentRole{}
	}
	return o.UpdateIncidentRole
}

type UpdateIncidentRoleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a single incident role from its ID
	IncidentRole *shared.IncidentRole
}

func (o *UpdateIncidentRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIncidentRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIncidentRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIncidentRoleResponse) GetIncidentRole() *shared.IncidentRole {
	if o == nil {
		return nil
	}
	return o.IncidentRole
}
