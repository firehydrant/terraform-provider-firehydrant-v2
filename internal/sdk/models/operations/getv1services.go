// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1ServicesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A comma separated list of label key / values in the format of 'key=value,key2=value2'. To filter change events that have a key (with no specific value), omit the value
	Labels *string `queryParam:"style=form,explode=true,name=labels"`
	// A query to search services by their name or description
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search services by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A query to search services by their tier
	Tiers *string `queryParam:"style=form,explode=true,name=tiers"`
	// A query to search services by if they are impacted with active incidents
	Impacted *string `queryParam:"style=form,explode=true,name=impacted"`
	// A query to search services by their owner
	Owner *string `queryParam:"style=form,explode=true,name=owner"`
	// A comma separated list of team ids
	RespondingTeams *string `queryParam:"style=form,explode=true,name=responding_teams"`
	// A comma separated list of functionality ids
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A query to find services that are available to be downstream dependencies for the passed service ID
	AvailableDownstreamDependenciesForID *string `queryParam:"style=form,explode=true,name=available_downstream_dependencies_for_id"`
	// A query to find services that are available to be upstream dependencies for the passed service ID
	AvailableUpstreamDependenciesForID *string `queryParam:"style=form,explode=true,name=available_upstream_dependencies_for_id"`
	// Boolean to determine whether to return a slimified version of the services object
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
	// Use in conjunction with lite param to specify additional attributes to include
	Include []string `queryParam:"style=form,explode=false,name=include"`
}

func (o *GetV1ServicesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1ServicesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1ServicesRequest) GetLabels() *string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *GetV1ServicesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1ServicesRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetV1ServicesRequest) GetTiers() *string {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *GetV1ServicesRequest) GetImpacted() *string {
	if o == nil {
		return nil
	}
	return o.Impacted
}

func (o *GetV1ServicesRequest) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *GetV1ServicesRequest) GetRespondingTeams() *string {
	if o == nil {
		return nil
	}
	return o.RespondingTeams
}

func (o *GetV1ServicesRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *GetV1ServicesRequest) GetAvailableDownstreamDependenciesForID() *string {
	if o == nil {
		return nil
	}
	return o.AvailableDownstreamDependenciesForID
}

func (o *GetV1ServicesRequest) GetAvailableUpstreamDependenciesForID() *string {
	if o == nil {
		return nil
	}
	return o.AvailableUpstreamDependenciesForID
}

func (o *GetV1ServicesRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

func (o *GetV1ServicesRequest) GetInclude() []string {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetV1ServicesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all of the services that have been added to the organization.
	ServiceEntityPaginated *shared.ServiceEntityPaginated
}

func (o *GetV1ServicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1ServicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1ServicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1ServicesResponse) GetServiceEntityPaginated() *shared.ServiceEntityPaginated {
	if o == nil {
		return nil
	}
	return o.ServiceEntityPaginated
}
