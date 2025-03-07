// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1IntegrationsStatuspageConnectionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1IntegrationsStatuspageConnectionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IntegrationsStatuspageConnectionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1IntegrationsStatuspageConnectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists the available and configured Statuspage integrations connections for the authenticated organization.
	IntegrationsStatuspageConnectionEntityPaginated *shared.IntegrationsStatuspageConnectionEntityPaginated
}

func (o *GetV1IntegrationsStatuspageConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IntegrationsStatuspageConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IntegrationsStatuspageConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1IntegrationsStatuspageConnectionsResponse) GetIntegrationsStatuspageConnectionEntityPaginated() *shared.IntegrationsStatuspageConnectionEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionEntityPaginated
}
