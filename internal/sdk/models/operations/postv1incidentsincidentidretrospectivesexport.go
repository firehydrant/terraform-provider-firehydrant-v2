// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

// IntegrationSlug - The name of the integration to export the retrospective to.
type IntegrationSlug string

const (
	IntegrationSlugConfluenceCloud IntegrationSlug = "confluence_cloud"
	IntegrationSlugGoogleDocs      IntegrationSlug = "google_docs"
)

func (e IntegrationSlug) ToPointer() *IntegrationSlug {
	return &e
}
func (e *IntegrationSlug) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confluence_cloud":
		fallthrough
	case "google_docs":
		*e = IntegrationSlug(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationSlug: %v", v)
	}
}

type PostV1IncidentsIncidentIDRetrospectivesExportRequestBody struct {
	// The name of the integration to export the retrospective to.
	IntegrationSlug IntegrationSlug `json:"integration_slug"`
	// The ID of the parent page to export the retrospective to.
	ParentPageID *string `json:"parent_page_id,omitempty"`
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportRequestBody) GetIntegrationSlug() IntegrationSlug {
	if o == nil {
		return IntegrationSlug("")
	}
	return o.IntegrationSlug
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportRequestBody) GetParentPageID() *string {
	if o == nil {
		return nil
	}
	return o.ParentPageID
}

type PostV1IncidentsIncidentIDRetrospectivesExportRequest struct {
	IncidentID  string                                                   `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody PostV1IncidentsIncidentIDRetrospectivesExportRequestBody `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportRequest) GetRequestBody() PostV1IncidentsIncidentIDRetrospectivesExportRequestBody {
	if o == nil {
		return PostV1IncidentsIncidentIDRetrospectivesExportRequestBody{}
	}
	return o.RequestBody
}

type PostV1IncidentsIncidentIDRetrospectivesExportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Export incident's retrospective(s) using their templates
	IncidentsExportRetrospectivesResultEntity *shared.IncidentsExportRetrospectivesResultEntity
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1IncidentsIncidentIDRetrospectivesExportResponse) GetIncidentsExportRetrospectivesResultEntity() *shared.IncidentsExportRetrospectivesResultEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsExportRetrospectivesResultEntity
}
