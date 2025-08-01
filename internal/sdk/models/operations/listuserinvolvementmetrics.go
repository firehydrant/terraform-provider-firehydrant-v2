// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/types"
	"net/http"
)

type ListUserInvolvementMetricsRequest struct {
	// The start date to return metrics from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate       *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	BucketSize    *string     `queryParam:"style=form,explode=true,name=bucket_size"`
	By            *string     `queryParam:"style=form,explode=true,name=by"`
	SortField     *string     `queryParam:"style=form,explode=true,name=sort_field"`
	SortDirection *string     `queryParam:"style=form,explode=true,name=sort_direction"`
	SortLimit     *int        `queryParam:"style=form,explode=true,name=sort_limit"`
}

func (l ListUserInvolvementMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUserInvolvementMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUserInvolvementMetricsRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListUserInvolvementMetricsRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *ListUserInvolvementMetricsRequest) GetBucketSize() *string {
	if o == nil {
		return nil
	}
	return o.BucketSize
}

func (o *ListUserInvolvementMetricsRequest) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *ListUserInvolvementMetricsRequest) GetSortField() *string {
	if o == nil {
		return nil
	}
	return o.SortField
}

func (o *ListUserInvolvementMetricsRequest) GetSortDirection() *string {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *ListUserInvolvementMetricsRequest) GetSortLimit() *int {
	if o == nil {
		return nil
	}
	return o.SortLimit
}

type ListUserInvolvementMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with time bucketed analytics data
	Metrics *shared.Metrics
}

func (o *ListUserInvolvementMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserInvolvementMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserInvolvementMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserInvolvementMetricsResponse) GetMetrics() *shared.Metrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}
