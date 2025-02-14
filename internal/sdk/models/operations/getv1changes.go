// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1ChangesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Filter changes by summary
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *GetV1ChangesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1ChangesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1ChangesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetV1ChangesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1ChangesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1ChangesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1ChangesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
