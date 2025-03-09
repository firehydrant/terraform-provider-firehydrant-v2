// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// TicketingFieldMapsMappableFieldEntity - Ticketing_FieldMaps_MappableFieldEntity model
type TicketingFieldMapsMappableFieldEntity struct {
	// The ID of the field
	Value *string `json:"value,omitempty"`
	// The human-readable name of the field
	Label *string `json:"label,omitempty"`
	// The allowed type of the field
	Type *string `json:"type,omitempty"`
	// The allowed values of the field
	AllowedValues []string `json:"allowed_values,omitempty"`
	// If the field is required to be mapped
	Required *string `json:"required,omitempty"`
}

func (o *TicketingFieldMapsMappableFieldEntity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *TicketingFieldMapsMappableFieldEntity) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *TicketingFieldMapsMappableFieldEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TicketingFieldMapsMappableFieldEntity) GetAllowedValues() []string {
	if o == nil {
		return nil
	}
	return o.AllowedValues
}

func (o *TicketingFieldMapsMappableFieldEntity) GetRequired() *string {
	if o == nil {
		return nil
	}
	return o.Required
}
