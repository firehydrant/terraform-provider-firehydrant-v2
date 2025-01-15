// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/types"
	"net/http"
)

type BucketSize string

const (
	BucketSizeWeek    BucketSize = "week"
	BucketSizeDay     BucketSize = "day"
	BucketSizeMonth   BucketSize = "month"
	BucketSizeAllTime BucketSize = "all_time"
)

func (e BucketSize) ToPointer() *BucketSize {
	return &e
}
func (e *BucketSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "week":
		fallthrough
	case "day":
		fallthrough
	case "month":
		fallthrough
	case "all_time":
		*e = BucketSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BucketSize: %v", v)
	}
}

type By string

const (
	ByTotal           By = "total"
	BySeverity        By = "severity"
	ByPriority        By = "priority"
	ByFunctionality   By = "functionality"
	ByService         By = "service"
	ByEnvironment     By = "environment"
	ByUser            By = "user"
	ByUserInvolvement By = "user_involvement"
)

func (e By) ToPointer() *By {
	return &e
}
func (e *By) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total":
		fallthrough
	case "severity":
		fallthrough
	case "priority":
		fallthrough
	case "functionality":
		fallthrough
	case "service":
		fallthrough
	case "environment":
		fallthrough
	case "user":
		fallthrough
	case "user_involvement":
		*e = By(v)
		return nil
	default:
		return fmt.Errorf("invalid value for By: %v", v)
	}
}

type QueryParamSortField string

const (
	QueryParamSortFieldMttd      QueryParamSortField = "mttd"
	QueryParamSortFieldMtta      QueryParamSortField = "mtta"
	QueryParamSortFieldMttm      QueryParamSortField = "mttm"
	QueryParamSortFieldMttr      QueryParamSortField = "mttr"
	QueryParamSortFieldCount     QueryParamSortField = "count"
	QueryParamSortFieldTotalTime QueryParamSortField = "total_time"
)

func (e QueryParamSortField) ToPointer() *QueryParamSortField {
	return &e
}
func (e *QueryParamSortField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mttd":
		fallthrough
	case "mtta":
		fallthrough
	case "mttm":
		fallthrough
	case "mttr":
		fallthrough
	case "count":
		fallthrough
	case "total_time":
		*e = QueryParamSortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSortField: %v", v)
	}
}

type QueryParamSortDirection string

const (
	QueryParamSortDirectionAsc  QueryParamSortDirection = "asc"
	QueryParamSortDirectionDesc QueryParamSortDirection = "desc"
)

func (e QueryParamSortDirection) ToPointer() *QueryParamSortDirection {
	return &e
}
func (e *QueryParamSortDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = QueryParamSortDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSortDirection: %v", v)
	}
}

type GetV1MetricsIncidentsRequest struct {
	// The start date to return metrics from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate       *types.Date              `queryParam:"style=form,explode=true,name=end_date"`
	BucketSize    *BucketSize              `queryParam:"style=form,explode=true,name=bucket_size"`
	By            *By                      `queryParam:"style=form,explode=true,name=by"`
	SortField     *QueryParamSortField     `queryParam:"style=form,explode=true,name=sort_field"`
	SortDirection *QueryParamSortDirection `queryParam:"style=form,explode=true,name=sort_direction"`
	SortLimit     *int                     `queryParam:"style=form,explode=true,name=sort_limit"`
	Conditions    *string                  `queryParam:"style=form,explode=true,name=conditions"`
}

func (g GetV1MetricsIncidentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1MetricsIncidentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1MetricsIncidentsRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1MetricsIncidentsRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *GetV1MetricsIncidentsRequest) GetBucketSize() *BucketSize {
	if o == nil {
		return nil
	}
	return o.BucketSize
}

func (o *GetV1MetricsIncidentsRequest) GetBy() *By {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *GetV1MetricsIncidentsRequest) GetSortField() *QueryParamSortField {
	if o == nil {
		return nil
	}
	return o.SortField
}

func (o *GetV1MetricsIncidentsRequest) GetSortDirection() *QueryParamSortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *GetV1MetricsIncidentsRequest) GetSortLimit() *int {
	if o == nil {
		return nil
	}
	return o.SortLimit
}

func (o *GetV1MetricsIncidentsRequest) GetConditions() *string {
	if o == nil {
		return nil
	}
	return o.Conditions
}

type GetV1MetricsIncidentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with time bucketed analytics data
	MetricsMetricsEntity *shared.MetricsMetricsEntity
}

func (o *GetV1MetricsIncidentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1MetricsIncidentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1MetricsIncidentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1MetricsIncidentsResponse) GetMetricsMetricsEntity() *shared.MetricsMetricsEntity {
	if o == nil {
		return nil
	}
	return o.MetricsMetricsEntity
}
