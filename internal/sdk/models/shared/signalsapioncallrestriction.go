// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SignalsAPIOnCallRestriction struct {
	EndDay    *string `json:"end_day,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	StartDay  *string `json:"start_day,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

func (o *SignalsAPIOnCallRestriction) GetEndDay() *string {
	if o == nil {
		return nil
	}
	return o.EndDay
}

func (o *SignalsAPIOnCallRestriction) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *SignalsAPIOnCallRestriction) GetStartDay() *string {
	if o == nil {
		return nil
	}
	return o.StartDay
}

func (o *SignalsAPIOnCallRestriction) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}
