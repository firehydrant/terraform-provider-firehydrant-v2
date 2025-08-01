// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListPostMortemReasonsRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage  *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListPostMortemReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *ListPostMortemReasonsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListPostMortemReasonsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListPostMortemReasonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all contributing factors to an incident
	PostMortemsReasonPaginated *shared.PostMortemsReasonPaginated
}

func (o *ListPostMortemReasonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPostMortemReasonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPostMortemReasonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPostMortemReasonsResponse) GetPostMortemsReasonPaginated() *shared.PostMortemsReasonPaginated {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonPaginated
}
