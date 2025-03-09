// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IncidentTypeEntityTemplateEntityLabels - Arbitrary key:value pairs of labels for your incidents.
type IncidentTypeEntityTemplateEntityLabels struct {
}

type IncidentTypeEntityTemplateEntity struct {
	IncidentName          *string `json:"incident_name,omitempty"`
	Summary               *string `json:"summary,omitempty"`
	Description           *string `json:"description,omitempty"`
	CustomerImpactSummary *string `json:"customer_impact_summary,omitempty"`
	// Arbitrary key:value pairs of labels for your incidents.
	Labels          *IncidentTypeEntityTemplateEntityLabels  `json:"labels,omitempty"`
	Severity        *string                                  `json:"severity,omitempty"`
	Priority        *string                                  `json:"priority,omitempty"`
	TagList         []string                                 `json:"tag_list,omitempty"`
	RunbookIds      []string                                 `json:"runbook_ids,omitempty"`
	TeamIds         []string                                 `json:"team_ids,omitempty"`
	PrivateIncident *bool                                    `json:"private_incident,omitempty"`
	CustomFields    *string                                  `json:"custom_fields,omitempty"`
	Impacts         []IncidentTypeEntityTemplateImpactEntity `json:"impacts,omitempty"`
}

func (o *IncidentTypeEntityTemplateEntity) GetIncidentName() *string {
	if o == nil {
		return nil
	}
	return o.IncidentName
}

func (o *IncidentTypeEntityTemplateEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *IncidentTypeEntityTemplateEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentTypeEntityTemplateEntity) GetCustomerImpactSummary() *string {
	if o == nil {
		return nil
	}
	return o.CustomerImpactSummary
}

func (o *IncidentTypeEntityTemplateEntity) GetLabels() *IncidentTypeEntityTemplateEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *IncidentTypeEntityTemplateEntity) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *IncidentTypeEntityTemplateEntity) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *IncidentTypeEntityTemplateEntity) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *IncidentTypeEntityTemplateEntity) GetRunbookIds() []string {
	if o == nil {
		return nil
	}
	return o.RunbookIds
}

func (o *IncidentTypeEntityTemplateEntity) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

func (o *IncidentTypeEntityTemplateEntity) GetPrivateIncident() *bool {
	if o == nil {
		return nil
	}
	return o.PrivateIncident
}

func (o *IncidentTypeEntityTemplateEntity) GetCustomFields() *string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *IncidentTypeEntityTemplateEntity) GetImpacts() []IncidentTypeEntityTemplateImpactEntity {
	if o == nil {
		return nil
	}
	return o.Impacts
}
