// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type UpdateSignalsEmailTargetRequest struct {
	ID                       string                          `pathParam:"style=simple,explode=false,name=id"`
	UpdateSignalsEmailTarget shared.UpdateSignalsEmailTarget `request:"mediaType=application/json"`
}

func (o *UpdateSignalsEmailTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSignalsEmailTargetRequest) GetUpdateSignalsEmailTarget() shared.UpdateSignalsEmailTarget {
	if o == nil {
		return shared.UpdateSignalsEmailTarget{}
	}
	return o.UpdateSignalsEmailTarget
}

type UpdateSignalsEmailTargetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a Signals email target by ID
	SignalsAPIEmailTarget *shared.SignalsAPIEmailTarget
}

func (o *UpdateSignalsEmailTargetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSignalsEmailTargetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSignalsEmailTargetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSignalsEmailTargetResponse) GetSignalsAPIEmailTarget() *shared.SignalsAPIEmailTarget {
	if o == nil {
		return nil
	}
	return o.SignalsAPIEmailTarget
}
