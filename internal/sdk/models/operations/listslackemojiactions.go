// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListSlackEmojiActionsRequest struct {
	// Slack Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Page         *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage      *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListSlackEmojiActionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListSlackEmojiActionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListSlackEmojiActionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListSlackEmojiActionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSlackEmojiActionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSlackEmojiActionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSlackEmojiActionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
