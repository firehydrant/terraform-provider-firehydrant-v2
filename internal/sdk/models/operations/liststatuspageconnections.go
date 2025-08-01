// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListStatuspageConnectionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListStatuspageConnectionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListStatuspageConnectionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListStatuspageConnectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists the available and configured Statuspage integrations connections for the authenticated organization.
	IntegrationsStatuspageConnectionPaginated *shared.IntegrationsStatuspageConnectionPaginated
}

func (o *ListStatuspageConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListStatuspageConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListStatuspageConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListStatuspageConnectionsResponse) GetIntegrationsStatuspageConnectionPaginated() *shared.IntegrationsStatuspageConnectionPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionPaginated
}
