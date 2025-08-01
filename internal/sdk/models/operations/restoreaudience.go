// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type RestoreAudienceRequest struct {
	// Unique identifier of the audience
	AudienceID string `pathParam:"style=simple,explode=false,name=audience_id"`
}

func (o *RestoreAudienceRequest) GetAudienceID() string {
	if o == nil {
		return ""
	}
	return o.AudienceID
}

type RestoreAudienceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Restore a previously archived audience
	Audiences *shared.Audiences
}

func (o *RestoreAudienceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RestoreAudienceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RestoreAudienceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RestoreAudienceResponse) GetAudiences() *shared.Audiences {
	if o == nil {
		return nil
	}
	return o.Audiences
}
