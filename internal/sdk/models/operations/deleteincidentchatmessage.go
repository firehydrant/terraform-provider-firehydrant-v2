// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type DeleteIncidentChatMessageRequest struct {
	MessageID  string `pathParam:"style=simple,explode=false,name=message_id"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteIncidentChatMessageRequest) GetMessageID() string {
	if o == nil {
		return ""
	}
	return o.MessageID
}

func (o *DeleteIncidentChatMessageRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteIncidentChatMessageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete an existing generic chat message on an incident.
	EventGenericChatMessage *shared.EventGenericChatMessage
}

func (o *DeleteIncidentChatMessageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteIncidentChatMessageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteIncidentChatMessageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteIncidentChatMessageResponse) GetEventGenericChatMessage() *shared.EventGenericChatMessage {
	if o == nil {
		return nil
	}
	return o.EventGenericChatMessage
}
