// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type MetricsMilestonesFunnelDataBucket struct {
	FilterParams    *NullableMetricsMilestonesFunnelDataBucketFilterParams `json:"filter_params,omitempty"`
	MilestoneCounts []MetricsMilestonesFunnelDataBucketMilestoneCount      `json:"milestone_counts,omitempty"`
	// The start datetime for the period
	TimeBucket *time.Time `json:"time_bucket,omitempty"`
}

func (m MetricsMilestonesFunnelDataBucket) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsMilestonesFunnelDataBucket) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetricsMilestonesFunnelDataBucket) GetFilterParams() *NullableMetricsMilestonesFunnelDataBucketFilterParams {
	if o == nil {
		return nil
	}
	return o.FilterParams
}

func (o *MetricsMilestonesFunnelDataBucket) GetMilestoneCounts() []MetricsMilestonesFunnelDataBucketMilestoneCount {
	if o == nil {
		return nil
	}
	return o.MilestoneCounts
}

func (o *MetricsMilestonesFunnelDataBucket) GetTimeBucket() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimeBucket
}
