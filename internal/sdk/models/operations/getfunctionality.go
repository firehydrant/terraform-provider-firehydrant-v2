// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetFunctionalityRequest struct {
	FunctionalityID string `pathParam:"style=simple,explode=false,name=functionality_id"`
}

func (o *GetFunctionalityRequest) GetFunctionalityID() string {
	if o == nil {
		return ""
	}
	return o.FunctionalityID
}

type GetFunctionalityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves a single functionality by ID
	Functionality *shared.Functionality
}

func (o *GetFunctionalityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFunctionalityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFunctionalityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFunctionalityResponse) GetFunctionality() *shared.Functionality {
	if o == nil {
		return nil
	}
	return o.Functionality
}
