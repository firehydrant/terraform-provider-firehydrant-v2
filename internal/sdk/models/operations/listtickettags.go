// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListTicketTagsRequest struct {
	Prefix *string `queryParam:"style=form,explode=true,name=prefix"`
}

func (o *ListTicketTagsRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

type ListTicketTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all of the ticket tags in the organization
	TagPaginated *shared.TagPaginated
}

func (o *ListTicketTagsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTicketTagsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTicketTagsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTicketTagsResponse) GetTagPaginated() *shared.TagPaginated {
	if o == nil {
		return nil
	}
	return o.TagPaginated
}
