// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateIncidentAlertRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// Array of alert IDs to be assigned to the incident
	RequestBody []string `request:"mediaType=application/json"`
}

func (o *CreateIncidentAlertRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentAlertRequest) GetRequestBody() []string {
	if o == nil {
		return []string{}
	}
	return o.RequestBody
}

type CreateIncidentAlertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateIncidentAlertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIncidentAlertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIncidentAlertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
