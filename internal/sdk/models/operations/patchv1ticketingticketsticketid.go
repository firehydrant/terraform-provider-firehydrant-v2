// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1TicketingTicketsTicketIDRequest struct {
	TicketID                        string                                 `pathParam:"style=simple,explode=false,name=ticket_id"`
	PatchV1TicketingTicketsTicketID shared.PatchV1TicketingTicketsTicketID `request:"mediaType=application/json"`
}

func (o *PatchV1TicketingTicketsTicketIDRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *PatchV1TicketingTicketsTicketIDRequest) GetPatchV1TicketingTicketsTicketID() shared.PatchV1TicketingTicketsTicketID {
	if o == nil {
		return shared.PatchV1TicketingTicketsTicketID{}
	}
	return o.PatchV1TicketingTicketsTicketID
}

type PatchV1TicketingTicketsTicketIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a ticket's attributes
	TicketingTicketEntity *shared.TicketingTicketEntity
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetTicketingTicketEntity() *shared.TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.TicketingTicketEntity
}
