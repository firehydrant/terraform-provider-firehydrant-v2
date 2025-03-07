// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchV1IntegrationsConnectionsSlugConnectionIDRequest struct {
	Slug         string `pathParam:"style=simple,explode=false,name=slug"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PatchV1IntegrationsConnectionsSlugConnectionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
