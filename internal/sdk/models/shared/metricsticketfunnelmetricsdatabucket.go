// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type MetricsTicketFunnelMetricsDataBucket struct {
	FilterParams *NullableMetricsTicketFunnelMetricsDataBucketFilterParams `json:"filter_params,omitempty"`
	// The number of follow ups created
	FollowUpsCreated *int `json:"follow_ups_created,omitempty"`
	// The number of follow ups completed
	FollowUpsDone *int `json:"follow_ups_done,omitempty"`
	// The number of tasks created
	TasksCreated *int `json:"tasks_created,omitempty"`
	// The number of tasks completed
	TasksDone *int `json:"tasks_done,omitempty"`
	// The start datetime for the period
	TimeBucket *time.Time `json:"time_bucket,omitempty"`
}

func (m MetricsTicketFunnelMetricsDataBucket) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsTicketFunnelMetricsDataBucket) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetFilterParams() *NullableMetricsTicketFunnelMetricsDataBucketFilterParams {
	if o == nil {
		return nil
	}
	return o.FilterParams
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetFollowUpsCreated() *int {
	if o == nil {
		return nil
	}
	return o.FollowUpsCreated
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetFollowUpsDone() *int {
	if o == nil {
		return nil
	}
	return o.FollowUpsDone
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetTasksCreated() *int {
	if o == nil {
		return nil
	}
	return o.TasksCreated
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetTasksDone() *int {
	if o == nil {
		return nil
	}
	return o.TasksDone
}

func (o *MetricsTicketFunnelMetricsDataBucket) GetTimeBucket() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimeBucket
}
