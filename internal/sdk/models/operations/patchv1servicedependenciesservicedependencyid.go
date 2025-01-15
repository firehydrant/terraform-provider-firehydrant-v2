// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1ServiceDependenciesServiceDependencyIDRequest struct {
	ServiceDependencyID                           string                                               `pathParam:"style=simple,explode=false,name=service_dependency_id"`
	PatchV1ServiceDependenciesServiceDependencyID shared.PatchV1ServiceDependenciesServiceDependencyID `request:"mediaType=application/json"`
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDRequest) GetServiceDependencyID() string {
	if o == nil {
		return ""
	}
	return o.ServiceDependencyID
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDRequest) GetPatchV1ServiceDependenciesServiceDependencyID() shared.PatchV1ServiceDependenciesServiceDependencyID {
	if o == nil {
		return shared.PatchV1ServiceDependenciesServiceDependencyID{}
	}
	return o.PatchV1ServiceDependenciesServiceDependencyID
}

type PatchV1ServiceDependenciesServiceDependencyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update the notes of the service dependency
	ServiceDependencyEntity *shared.ServiceDependencyEntity
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDResponse) GetServiceDependencyEntity() *shared.ServiceDependencyEntity {
	if o == nil {
		return nil
	}
	return o.ServiceDependencyEntity
}
