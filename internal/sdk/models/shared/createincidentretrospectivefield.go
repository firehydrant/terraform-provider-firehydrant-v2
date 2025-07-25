// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateIncidentRetrospectiveField struct {
	HelpText          *string  `json:"help_text,omitempty"`
	IsRequired        *bool    `json:"is_required,omitempty"`
	Label             string   `json:"label"`
	PermissibleValues []string `json:"permissible_values,omitempty"`
	Schema            []string `json:"schema,omitempty"`
	Type              string   `json:"type"`
}

func (o *CreateIncidentRetrospectiveField) GetHelpText() *string {
	if o == nil {
		return nil
	}
	return o.HelpText
}

func (o *CreateIncidentRetrospectiveField) GetIsRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsRequired
}

func (o *CreateIncidentRetrospectiveField) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *CreateIncidentRetrospectiveField) GetPermissibleValues() []string {
	if o == nil {
		return nil
	}
	return o.PermissibleValues
}

func (o *CreateIncidentRetrospectiveField) GetSchema() []string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *CreateIncidentRetrospectiveField) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
