// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1IntegrationsStatuspageConnectionsConnectionIDRequest struct {
	// Connection UUID
	ConnectionID                                         string                                                      `pathParam:"style=simple,explode=false,name=connection_id"`
	PatchV1IntegrationsStatuspageConnectionsConnectionID shared.PatchV1IntegrationsStatuspageConnectionsConnectionID `request:"mediaType=application/json"`
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDRequest) GetPatchV1IntegrationsStatuspageConnectionsConnectionID() shared.PatchV1IntegrationsStatuspageConnectionsConnectionID {
	if o == nil {
		return shared.PatchV1IntegrationsStatuspageConnectionsConnectionID{}
	}
	return o.PatchV1IntegrationsStatuspageConnectionsConnectionID
}

type PatchV1IntegrationsStatuspageConnectionsConnectionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update the given Statuspage integration connection.
	IntegrationsStatuspageConnectionEntity *shared.IntegrationsStatuspageConnectionEntity
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetIntegrationsStatuspageConnectionEntity() *shared.IntegrationsStatuspageConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionEntity
}
