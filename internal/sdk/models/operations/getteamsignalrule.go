// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetTeamSignalRuleRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	ID     string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeamSignalRuleRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamSignalRuleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTeamSignalRuleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get a Signals rule by ID.
	SignalsAPIRule *shared.SignalsAPIRule
}

func (o *GetTeamSignalRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamSignalRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamSignalRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamSignalRuleResponse) GetSignalsAPIRule() *shared.SignalsAPIRule {
	if o == nil {
		return nil
	}
	return o.SignalsAPIRule
}
