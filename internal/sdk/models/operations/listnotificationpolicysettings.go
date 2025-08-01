// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListNotificationPolicySettingsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListNotificationPolicySettingsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListNotificationPolicySettingsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListNotificationPolicySettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all Signals notification policies.
	SignalsAPINotificationPolicyItemPaginated *shared.SignalsAPINotificationPolicyItemPaginated
}

func (o *ListNotificationPolicySettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListNotificationPolicySettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListNotificationPolicySettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListNotificationPolicySettingsResponse) GetSignalsAPINotificationPolicyItemPaginated() *shared.SignalsAPINotificationPolicyItemPaginated {
	if o == nil {
		return nil
	}
	return o.SignalsAPINotificationPolicyItemPaginated
}
