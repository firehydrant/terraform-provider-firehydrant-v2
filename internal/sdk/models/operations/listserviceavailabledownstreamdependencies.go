// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListServiceAvailableDownstreamDependenciesRequest struct {
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *ListServiceAvailableDownstreamDependenciesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type ListServiceAvailableDownstreamDependenciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves all services that are available to be downstream dependencies
	ServiceLite *shared.ServiceLite
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetServiceLite() *shared.ServiceLite {
	if o == nil {
		return nil
	}
	return o.ServiceLite
}
