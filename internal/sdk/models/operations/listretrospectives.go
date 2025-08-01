// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
	"time"
)

type ListRetrospectivesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Filter the reports by an incident ID
	IncidentID *string `queryParam:"style=form,explode=true,name=incident_id"`
	// Filter for reports updated after the given ISO8601 timestamp
	UpdatedSince *time.Time `queryParam:"style=form,explode=true,name=updated_since"`
}

func (l ListRetrospectivesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRetrospectivesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRetrospectivesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRetrospectivesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListRetrospectivesRequest) GetIncidentID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentID
}

func (o *ListRetrospectivesRequest) GetUpdatedSince() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedSince
}

type ListRetrospectivesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all retrospective reports
	IncidentsRetrospectivePaginated *shared.IncidentsRetrospectivePaginated
}

func (o *ListRetrospectivesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRetrospectivesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRetrospectivesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListRetrospectivesResponse) GetIncidentsRetrospectivePaginated() *shared.IncidentsRetrospectivePaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsRetrospectivePaginated
}
