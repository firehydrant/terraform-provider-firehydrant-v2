// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListAvailableInboundFieldMapsRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *ListAvailableInboundFieldMapsRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type ListAvailableInboundFieldMapsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns metadata for the fields that are available for inbound field mapping.
	TicketingProjectInboundMappableField *shared.TicketingProjectInboundMappableField
}

func (o *ListAvailableInboundFieldMapsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAvailableInboundFieldMapsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAvailableInboundFieldMapsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAvailableInboundFieldMapsResponse) GetTicketingProjectInboundMappableField() *shared.TicketingProjectInboundMappableField {
	if o == nil {
		return nil
	}
	return o.TicketingProjectInboundMappableField
}
