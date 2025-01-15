// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PostV1SeverityMatrixImpactsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new impact
	SeverityMatrixImpactEntity *shared.SeverityMatrixImpactEntity
}

func (o *PostV1SeverityMatrixImpactsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1SeverityMatrixImpactsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1SeverityMatrixImpactsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1SeverityMatrixImpactsResponse) GetSeverityMatrixImpactEntity() *shared.SeverityMatrixImpactEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixImpactEntity
}
