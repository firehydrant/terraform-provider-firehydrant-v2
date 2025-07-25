// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListTeamEscalationPoliciesRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	// A query string for searching through the list of escalation policies.
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListTeamEscalationPoliciesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *ListTeamEscalationPoliciesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListTeamEscalationPoliciesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListTeamEscalationPoliciesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListTeamEscalationPoliciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all Signals escalation policies for a team.
	SignalsAPIEscalationPolicyPaginated *shared.SignalsAPIEscalationPolicyPaginated
}

func (o *ListTeamEscalationPoliciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTeamEscalationPoliciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTeamEscalationPoliciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTeamEscalationPoliciesResponse) GetSignalsAPIEscalationPolicyPaginated() *shared.SignalsAPIEscalationPolicyPaginated {
	if o == nil {
		return nil
	}
	return o.SignalsAPIEscalationPolicyPaginated
}
