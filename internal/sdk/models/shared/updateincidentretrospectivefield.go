// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateIncidentRetrospectiveField - Update retrospective field value
type UpdateIncidentRetrospectiveField struct {
	// The ID of the dynamic input field to update.
	DynamicInputFieldID *string `json:"dynamic_input_field_id,omitempty"`
	// The value to set for the field.
	Value int64 `json:"value"`
}

func (o *UpdateIncidentRetrospectiveField) GetDynamicInputFieldID() *string {
	if o == nil {
		return nil
	}
	return o.DynamicInputFieldID
}

func (o *UpdateIncidentRetrospectiveField) GetValue() int64 {
	if o == nil {
		return 0
	}
	return o.Value
}
