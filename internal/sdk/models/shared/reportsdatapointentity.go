// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ReportsDataPointEntity struct {
	Key   *string `json:"key,omitempty"`
	Value *int    `json:"value,omitempty"`
}

func (o *ReportsDataPointEntity) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *ReportsDataPointEntity) GetValue() *int {
	if o == nil {
		return nil
	}
	return o.Value
}
