// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDRequest struct {
	// Slack Connection UUID
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connection_id"`
	EmojiActionID string `pathParam:"style=simple,explode=false,name=emoji_action_id"`
}

func (o *DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDRequest) GetEmojiActionID() string {
	if o == nil {
		return ""
	}
	return o.EmojiActionID
}

type DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
