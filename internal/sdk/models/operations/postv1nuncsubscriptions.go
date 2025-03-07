// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1NuncSubscriptionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Subscribe to status page updates
	NuncNuncSubscription *shared.NuncNuncSubscription
}

func (o *PostV1NuncSubscriptionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1NuncSubscriptionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1NuncSubscriptionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1NuncSubscriptionsResponse) GetNuncNuncSubscription() *shared.NuncNuncSubscription {
	if o == nil {
		return nil
	}
	return o.NuncNuncSubscription
}
