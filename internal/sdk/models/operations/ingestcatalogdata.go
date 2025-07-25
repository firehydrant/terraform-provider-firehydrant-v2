// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type IngestCatalogDataRequest struct {
	CatalogID         string                   `pathParam:"style=simple,explode=false,name=catalog_id"`
	IngestCatalogData shared.IngestCatalogData `request:"mediaType=application/json"`
}

func (o *IngestCatalogDataRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *IngestCatalogDataRequest) GetIngestCatalogData() shared.IngestCatalogData {
	if o == nil {
		return shared.IngestCatalogData{}
	}
	return o.IngestCatalogData
}

type IngestCatalogDataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
	Imports *shared.Imports
}

func (o *IngestCatalogDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *IngestCatalogDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *IngestCatalogDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *IngestCatalogDataResponse) GetImports() *shared.Imports {
	if o == nil {
		return nil
	}
	return o.Imports
}
