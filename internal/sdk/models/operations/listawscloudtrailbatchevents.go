// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListAwsCloudtrailBatchEventsRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ListAwsCloudtrailBatchEventsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListAwsCloudtrailBatchEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List events for an AWS CloudTrail batch
	ChangeEvent *shared.ChangeEvent
}

func (o *ListAwsCloudtrailBatchEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAwsCloudtrailBatchEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAwsCloudtrailBatchEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAwsCloudtrailBatchEventsResponse) GetChangeEvent() *shared.ChangeEvent {
	if o == nil {
		return nil
	}
	return o.ChangeEvent
}
