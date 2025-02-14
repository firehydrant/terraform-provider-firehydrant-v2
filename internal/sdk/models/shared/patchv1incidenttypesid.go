// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PatchV1IncidentTypesIDImpacts struct {
	// The id of impact
	ID string `json:"id"`
	// The id of the condition
	ConditionID string `json:"condition_id"`
}

func (o *PatchV1IncidentTypesIDImpacts) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1IncidentTypesIDImpacts) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type PatchV1IncidentTypesIDTemplate struct {
	Description           *string `json:"description,omitempty"`
	CustomerImpactSummary *string `json:"customer_impact_summary,omitempty"`
	// A labels hash of keys and values
	Labels   map[string]string `json:"labels,omitempty"`
	Severity *string           `json:"severity,omitempty"`
	Priority *string           `json:"priority,omitempty"`
	// List of tags for the incident
	TagList []string `json:"tag_list,omitempty"`
	// List of ids of Runbooks to attach to incidents created from this type
	RunbookIds      []string `json:"runbook_ids,omitempty"`
	PrivateIncident *bool    `json:"private_incident,omitempty"`
	// List of ids of teams to be assigned to incidents
	TeamIds []string `json:"team_ids,omitempty"`
	// An array of impact/condition combinations
	Impacts []PatchV1IncidentTypesIDImpacts `json:"impacts,omitempty"`
}

func (o *PatchV1IncidentTypesIDTemplate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1IncidentTypesIDTemplate) GetCustomerImpactSummary() *string {
	if o == nil {
		return nil
	}
	return o.CustomerImpactSummary
}

func (o *PatchV1IncidentTypesIDTemplate) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PatchV1IncidentTypesIDTemplate) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *PatchV1IncidentTypesIDTemplate) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *PatchV1IncidentTypesIDTemplate) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *PatchV1IncidentTypesIDTemplate) GetRunbookIds() []string {
	if o == nil {
		return nil
	}
	return o.RunbookIds
}

func (o *PatchV1IncidentTypesIDTemplate) GetPrivateIncident() *bool {
	if o == nil {
		return nil
	}
	return o.PrivateIncident
}

func (o *PatchV1IncidentTypesIDTemplate) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

func (o *PatchV1IncidentTypesIDTemplate) GetImpacts() []PatchV1IncidentTypesIDImpacts {
	if o == nil {
		return nil
	}
	return o.Impacts
}

// PatchV1IncidentTypesID - Update a single incident type from its ID
type PatchV1IncidentTypesID struct {
	Name     string                         `json:"name"`
	Template PatchV1IncidentTypesIDTemplate `json:"template"`
}

func (o *PatchV1IncidentTypesID) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PatchV1IncidentTypesID) GetTemplate() PatchV1IncidentTypesIDTemplate {
	if o == nil {
		return PatchV1IncidentTypesIDTemplate{}
	}
	return o.Template
}
