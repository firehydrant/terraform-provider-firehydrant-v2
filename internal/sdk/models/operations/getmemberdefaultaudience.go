// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetMemberDefaultAudienceRequest struct {
	MemberID string `pathParam:"style=simple,explode=false,name=member_id"`
}

func (o *GetMemberDefaultAudienceRequest) GetMemberID() string {
	if o == nil {
		return ""
	}
	return o.MemberID
}

type GetMemberDefaultAudienceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get member's default audience
	Audiences *shared.Audiences
}

func (o *GetMemberDefaultAudienceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMemberDefaultAudienceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMemberDefaultAudienceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMemberDefaultAudienceResponse) GetAudiences() *shared.Audiences {
	if o == nil {
		return nil
	}
	return o.Audiences
}
