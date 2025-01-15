// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PutV1RunbooksRunbookIDRequest struct {
	RunbookID              string                        `pathParam:"style=simple,explode=false,name=runbook_id"`
	PutV1RunbooksRunbookID shared.PutV1RunbooksRunbookID `request:"mediaType=application/json"`
}

func (o *PutV1RunbooksRunbookIDRequest) GetRunbookID() string {
	if o == nil {
		return ""
	}
	return o.RunbookID
}

func (o *PutV1RunbooksRunbookIDRequest) GetPutV1RunbooksRunbookID() shared.PutV1RunbooksRunbookID {
	if o == nil {
		return shared.PutV1RunbooksRunbookID{}
	}
	return o.PutV1RunbooksRunbookID
}

type PutV1RunbooksRunbookIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything
	// about a runbook, including but not limited to the steps, environments, attachment rules, and severities.
	//
	RunbookEntity *shared.RunbookEntity
}

func (o *PutV1RunbooksRunbookIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutV1RunbooksRunbookIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutV1RunbooksRunbookIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutV1RunbooksRunbookIDResponse) GetRunbookEntity() *shared.RunbookEntity {
	if o == nil {
		return nil
	}
	return o.RunbookEntity
}
