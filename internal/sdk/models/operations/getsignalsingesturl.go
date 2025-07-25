// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetSignalsIngestURLRequest struct {
	// Team ID to send signals to directly
	TeamID *string `queryParam:"style=form,explode=true,name=team_id"`
	// Escalation policy ID to send signals to directly. `team_id` is required if this is provided.
	EscalationPolicyID *string `queryParam:"style=form,explode=true,name=escalation_policy_id"`
	// On-call schedule ID to send signals to directly. `team_id` is required if this is provided.
	OnCallScheduleID *string `queryParam:"style=form,explode=true,name=on_call_schedule_id"`
	// User ID to send signals to directly
	UserID *string `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *GetSignalsIngestURLRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetSignalsIngestURLRequest) GetEscalationPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.EscalationPolicyID
}

func (o *GetSignalsIngestURLRequest) GetOnCallScheduleID() *string {
	if o == nil {
		return nil
	}
	return o.OnCallScheduleID
}

func (o *GetSignalsIngestURLRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type GetSignalsIngestURLResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve the url for ingesting signals for your organization
	SignalsAPIIngestKey *shared.SignalsAPIIngestKey
}

func (o *GetSignalsIngestURLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSignalsIngestURLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSignalsIngestURLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSignalsIngestURLResponse) GetSignalsAPIIngestKey() *shared.SignalsAPIIngestKey {
	if o == nil {
		return nil
	}
	return o.SignalsAPIIngestKey
}
