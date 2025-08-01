// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetAlertRequest struct {
	AlertID string `pathParam:"style=simple,explode=false,name=alert_id"`
}

func (o *GetAlertRequest) GetAlertID() string {
	if o == nil {
		return ""
	}
	return o.AlertID
}

type GetAlertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a single alert
	Alerts *shared.Alerts
}

func (o *GetAlertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAlertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAlertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAlertResponse) GetAlerts() *shared.Alerts {
	if o == nil {
		return nil
	}
	return o.Alerts
}
