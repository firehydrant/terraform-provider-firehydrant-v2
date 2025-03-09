// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody struct {
	Name             *string `json:"name,omitempty"`
	ComponentGroupID *string `json:"component_group_id,omitempty"`
	Position         *int    `json:"position,omitempty"`
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody) GetComponentGroupID() *string {
	if o == nil {
		return nil
	}
	return o.ComponentGroupID
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

type PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequest struct {
	NuncConnectionID string                                                                   `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	GroupID          string                                                                   `pathParam:"style=simple,explode=false,name=group_id"`
	RequestBody      *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody `request:"mediaType=application/json"`
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequest) GetRequestBody() *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
