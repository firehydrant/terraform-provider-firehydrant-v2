// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetInboundFieldMapRequest struct {
	MapID              string `pathParam:"style=simple,explode=false,name=map_id"`
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *GetInboundFieldMapRequest) GetMapID() string {
	if o == nil {
		return ""
	}
	return o.MapID
}

func (o *GetInboundFieldMapRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type GetInboundFieldMapResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve inbound field map for a ticketing project
	TicketingProjectInboundFieldMap *shared.TicketingProjectInboundFieldMap
}

func (o *GetInboundFieldMapResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInboundFieldMapResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInboundFieldMapResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetInboundFieldMapResponse) GetTicketingProjectInboundFieldMap() *shared.TicketingProjectInboundFieldMap {
	if o == nil {
		return nil
	}
	return o.TicketingProjectInboundFieldMap
}
