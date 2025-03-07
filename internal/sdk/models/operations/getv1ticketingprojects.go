// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1TicketingProjectsRequest struct {
	SupportsTicketTypes *string `queryParam:"style=form,explode=true,name=supports_ticket_types"`
	Providers           *string `queryParam:"style=form,explode=true,name=providers"`
	ConnectionIds       *string `queryParam:"style=form,explode=true,name=connection_ids"`
	ConfiguredProjects  *bool   `queryParam:"style=form,explode=true,name=configured_projects"`
	Query               *string `queryParam:"style=form,explode=true,name=query"`
	Page                *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage             *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1TicketingProjectsRequest) GetSupportsTicketTypes() *string {
	if o == nil {
		return nil
	}
	return o.SupportsTicketTypes
}

func (o *GetV1TicketingProjectsRequest) GetProviders() *string {
	if o == nil {
		return nil
	}
	return o.Providers
}

func (o *GetV1TicketingProjectsRequest) GetConnectionIds() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionIds
}

func (o *GetV1TicketingProjectsRequest) GetConfiguredProjects() *bool {
	if o == nil {
		return nil
	}
	return o.ConfiguredProjects
}

func (o *GetV1TicketingProjectsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1TicketingProjectsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1TicketingProjectsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1TicketingProjectsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all ticketing projects available to the organization
	TicketingProjectsProjectListItemEntity *shared.TicketingProjectsProjectListItemEntity
}

func (o *GetV1TicketingProjectsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TicketingProjectsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TicketingProjectsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1TicketingProjectsResponse) GetTicketingProjectsProjectListItemEntity() *shared.TicketingProjectsProjectListItemEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectsProjectListItemEntity
}
