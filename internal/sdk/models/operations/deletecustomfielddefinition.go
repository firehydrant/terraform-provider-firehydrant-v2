// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type DeleteCustomFieldDefinitionRequest struct {
	FieldID string `pathParam:"style=simple,explode=false,name=field_id"`
}

func (o *DeleteCustomFieldDefinitionRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

type DeleteCustomFieldDefinitionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete a custom field definition
	OrganizationsCustomFieldDefinition *shared.OrganizationsCustomFieldDefinition
}

func (o *DeleteCustomFieldDefinitionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCustomFieldDefinitionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCustomFieldDefinitionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteCustomFieldDefinitionResponse) GetOrganizationsCustomFieldDefinition() *shared.OrganizationsCustomFieldDefinition {
	if o == nil {
		return nil
	}
	return o.OrganizationsCustomFieldDefinition
}
