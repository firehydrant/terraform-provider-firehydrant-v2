// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1TeamsTeamIDEscalationPoliciesRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	// A query string for searching through the list of escalation policies.
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1TeamsTeamIDEscalationPoliciesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetV1TeamsTeamIDEscalationPoliciesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1TeamsTeamIDEscalationPoliciesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1TeamsTeamIDEscalationPoliciesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1TeamsTeamIDEscalationPoliciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1TeamsTeamIDEscalationPoliciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TeamsTeamIDEscalationPoliciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TeamsTeamIDEscalationPoliciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
