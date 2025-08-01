// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
	"time"
)

type GetSignalsGroupedMetricsRequest struct {
	// A comma separated list of signal rule IDs
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A comma separated list of user IDs
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// String that determines how records are grouped
	GroupBy *string `queryParam:"style=form,explode=true,name=group_by"`
	// String that determines how records are sorted
	SortBy *string `queryParam:"style=form,explode=true,name=sort_by"`
	// String that determines how records are sorted
	SortDirection *string `queryParam:"style=form,explode=true,name=sort_direction"`
	// The start date to return metrics from
	StartDate *time.Time `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *time.Time `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetSignalsGroupedMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSignalsGroupedMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetSignalsGroupedMetricsRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *GetSignalsGroupedMetricsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetSignalsGroupedMetricsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetSignalsGroupedMetricsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetSignalsGroupedMetricsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetSignalsGroupedMetricsRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *GetSignalsGroupedMetricsRequest) GetGroupBy() *string {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *GetSignalsGroupedMetricsRequest) GetSortBy() *string {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *GetSignalsGroupedMetricsRequest) GetSortDirection() *string {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *GetSignalsGroupedMetricsRequest) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetSignalsGroupedMetricsRequest) GetEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetSignalsGroupedMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Generate a report of grouped metrics for Signals alerts
	SignalsAPIAnalyticsGroupedMetrics *shared.SignalsAPIAnalyticsGroupedMetrics
}

func (o *GetSignalsGroupedMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSignalsGroupedMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSignalsGroupedMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSignalsGroupedMetricsResponse) GetSignalsAPIAnalyticsGroupedMetrics() *shared.SignalsAPIAnalyticsGroupedMetrics {
	if o == nil {
		return nil
	}
	return o.SignalsAPIAnalyticsGroupedMetrics
}
