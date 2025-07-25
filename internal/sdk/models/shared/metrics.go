// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Bucket struct {
}

type DisplayInformation struct {
}

// Metrics model
type Metrics struct {
	// The size of returned buckets. Can be one of: day, week, month, or all_time.
	BucketSize *int     `json:"bucket_size,omitempty"`
	Buckets    []Bucket `json:"buckets,omitempty"`
	// The field by which the metrics are grouped. Can be one of: total, severity, priority, functionality, service, environment, or user.
	By                 *string                     `json:"by,omitempty"`
	DisplayInformation *DisplayInformation         `json:"display_information,omitempty"`
	Keys               []string                    `json:"keys,omitempty"`
	Sort               *NullableMetricsMetricsSort `json:"sort,omitempty"`
	Type               *string                     `json:"type,omitempty"`
}

func (o *Metrics) GetBucketSize() *int {
	if o == nil {
		return nil
	}
	return o.BucketSize
}

func (o *Metrics) GetBuckets() []Bucket {
	if o == nil {
		return nil
	}
	return o.Buckets
}

func (o *Metrics) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *Metrics) GetDisplayInformation() *DisplayInformation {
	if o == nil {
		return nil
	}
	return o.DisplayInformation
}

func (o *Metrics) GetKeys() []string {
	if o == nil {
		return nil
	}
	return o.Keys
}

func (o *Metrics) GetSort() *NullableMetricsMetricsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *Metrics) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
