// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type FilterParams struct {
}

type MetricsMttxGroup struct {
	Count            *int          `json:"count,omitempty"`
	CountDiff        *int          `json:"count_diff,omitempty"`
	CountPercentDiff *float32      `json:"count_percent_diff,omitempty"`
	FilterParams     *FilterParams `json:"filter_params,omitempty"`
	GroupAttributes  *string       `json:"group_attributes,omitempty"`
	Healthiness      *float32      `json:"healthiness,omitempty"`
	Mtta             *float32      `json:"mtta,omitempty"`
	MttaDiff         *float32      `json:"mtta_diff,omitempty"`
	MttaPercentDiff  *float32      `json:"mtta_percent_diff,omitempty"`
	Mttd             *float32      `json:"mttd,omitempty"`
	MttdDiff         *float32      `json:"mttd_diff,omitempty"`
	MttdPercentDiff  *float32      `json:"mttd_percent_diff,omitempty"`
	Mttm             *float32      `json:"mttm,omitempty"`
	MttmDiff         *float32      `json:"mttm_diff,omitempty"`
	MttmPercentDiff  *float32      `json:"mttm_percent_diff,omitempty"`
	Mttr             *float32      `json:"mttr,omitempty"`
	MttrDiff         *float32      `json:"mttr_diff,omitempty"`
	MttrPercentDiff  *float32      `json:"mttr_percent_diff,omitempty"`
}

func (o *MetricsMttxGroup) GetCount() *int {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MetricsMttxGroup) GetCountDiff() *int {
	if o == nil {
		return nil
	}
	return o.CountDiff
}

func (o *MetricsMttxGroup) GetCountPercentDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.CountPercentDiff
}

func (o *MetricsMttxGroup) GetFilterParams() *FilterParams {
	if o == nil {
		return nil
	}
	return o.FilterParams
}

func (o *MetricsMttxGroup) GetGroupAttributes() *string {
	if o == nil {
		return nil
	}
	return o.GroupAttributes
}

func (o *MetricsMttxGroup) GetHealthiness() *float32 {
	if o == nil {
		return nil
	}
	return o.Healthiness
}

func (o *MetricsMttxGroup) GetMtta() *float32 {
	if o == nil {
		return nil
	}
	return o.Mtta
}

func (o *MetricsMttxGroup) GetMttaDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttaDiff
}

func (o *MetricsMttxGroup) GetMttaPercentDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttaPercentDiff
}

func (o *MetricsMttxGroup) GetMttd() *float32 {
	if o == nil {
		return nil
	}
	return o.Mttd
}

func (o *MetricsMttxGroup) GetMttdDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttdDiff
}

func (o *MetricsMttxGroup) GetMttdPercentDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttdPercentDiff
}

func (o *MetricsMttxGroup) GetMttm() *float32 {
	if o == nil {
		return nil
	}
	return o.Mttm
}

func (o *MetricsMttxGroup) GetMttmDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttmDiff
}

func (o *MetricsMttxGroup) GetMttmPercentDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttmPercentDiff
}

func (o *MetricsMttxGroup) GetMttr() *float32 {
	if o == nil {
		return nil
	}
	return o.Mttr
}

func (o *MetricsMttxGroup) GetMttrDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttrDiff
}

func (o *MetricsMttxGroup) GetMttrPercentDiff() *float32 {
	if o == nil {
		return nil
	}
	return o.MttrPercentDiff
}
