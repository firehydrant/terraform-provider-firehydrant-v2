// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1NuncConnectionsNuncConnectionIDSubscribersRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

type GetV1NuncConnectionsNuncConnectionIDSubscribersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves the list of subscribers for a status page.
	NuncEmailSubscribersEntity *shared.NuncEmailSubscribersEntity
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetNuncEmailSubscribersEntity() *shared.NuncEmailSubscribersEntity {
	if o == nil {
		return nil
	}
	return o.NuncEmailSubscribersEntity
}
