// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type UpdateScimGroupRequest struct {
	ID              string                 `pathParam:"style=simple,explode=false,name=id"`
	UpdateScimGroup shared.UpdateScimGroup `request:"mediaType=application/scim+json"`
}

func (o *UpdateScimGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateScimGroupRequest) GetUpdateScimGroup() shared.UpdateScimGroup {
	if o == nil {
		return shared.UpdateScimGroup{}
	}
	return o.UpdateScimGroup
}

type UpdateScimGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateScimGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateScimGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateScimGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
