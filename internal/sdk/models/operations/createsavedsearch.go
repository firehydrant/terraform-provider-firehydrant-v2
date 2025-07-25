// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type CreateSavedSearchRequest struct {
	ResourceType      string                   `pathParam:"style=simple,explode=false,name=resource_type"`
	CreateSavedSearch shared.CreateSavedSearch `request:"mediaType=application/json"`
}

func (o *CreateSavedSearchRequest) GetResourceType() string {
	if o == nil {
		return ""
	}
	return o.ResourceType
}

func (o *CreateSavedSearchRequest) GetCreateSavedSearch() shared.CreateSavedSearch {
	if o == nil {
		return shared.CreateSavedSearch{}
	}
	return o.CreateSavedSearch
}

type CreateSavedSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new saved search for a particular resource type
	SavedSearch *shared.SavedSearch
}

func (o *CreateSavedSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSavedSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSavedSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSavedSearchResponse) GetSavedSearch() *shared.SavedSearch {
	if o == nil {
		return nil
	}
	return o.SavedSearch
}
