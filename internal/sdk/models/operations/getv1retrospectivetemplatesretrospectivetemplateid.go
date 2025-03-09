// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetV1RetrospectiveTemplatesRetrospectiveTemplateIDRequest struct {
	RetrospectiveTemplateID string `pathParam:"style=simple,explode=false,name=retrospective_template_id"`
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDRequest) GetRetrospectiveTemplateID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveTemplateID
}

type GetV1RetrospectiveTemplatesRetrospectiveTemplateIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a single retrospective template by ID
	RetrospectivesTemplateEntity *shared.RetrospectivesTemplateEntity
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDResponse) GetRetrospectivesTemplateEntity() *shared.RetrospectivesTemplateEntity {
	if o == nil {
		return nil
	}
	return o.RetrospectivesTemplateEntity
}
