// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetTeamEscalationPolicyRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	ID     string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeamEscalationPolicyRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamEscalationPolicyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTeamEscalationPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get a Signals escalation policy by ID
	SignalsAPIEscalationPolicy *shared.SignalsAPIEscalationPolicy
}

func (o *GetTeamEscalationPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamEscalationPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamEscalationPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamEscalationPolicyResponse) GetSignalsAPIEscalationPolicy() *shared.SignalsAPIEscalationPolicy {
	if o == nil {
		return nil
	}
	return o.SignalsAPIEscalationPolicy
}
