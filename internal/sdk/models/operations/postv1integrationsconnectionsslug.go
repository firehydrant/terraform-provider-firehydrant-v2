// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostV1IntegrationsConnectionsSlugRequest struct {
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (o *PostV1IntegrationsConnectionsSlugRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type PostV1IntegrationsConnectionsSlugResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostV1IntegrationsConnectionsSlugResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1IntegrationsConnectionsSlugResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1IntegrationsConnectionsSlugResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
