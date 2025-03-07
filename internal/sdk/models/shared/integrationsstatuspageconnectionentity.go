// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IntegrationsStatuspageConnectionEntity - Integrations_Statuspage_ConnectionEntity model
type IntegrationsStatuspageConnectionEntity struct {
	ID                *string                                        `json:"id,omitempty"`
	PageName          *string                                        `json:"page_name,omitempty"`
	PageID            *string                                        `json:"page_id,omitempty"`
	Conditions        []IntegrationsStatuspageConditionEntity        `json:"conditions,omitempty"`
	Severities        []IntegrationsStatuspageSeverityEntity         `json:"severities,omitempty"`
	MilestoneMappings []IntegrationsStatuspageMilestoneMappingEntity `json:"milestone_mappings,omitempty"`
}

func (o *IntegrationsStatuspageConnectionEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationsStatuspageConnectionEntity) GetPageName() *string {
	if o == nil {
		return nil
	}
	return o.PageName
}

func (o *IntegrationsStatuspageConnectionEntity) GetPageID() *string {
	if o == nil {
		return nil
	}
	return o.PageID
}

func (o *IntegrationsStatuspageConnectionEntity) GetConditions() []IntegrationsStatuspageConditionEntity {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *IntegrationsStatuspageConnectionEntity) GetSeverities() []IntegrationsStatuspageSeverityEntity {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *IntegrationsStatuspageConnectionEntity) GetMilestoneMappings() []IntegrationsStatuspageMilestoneMappingEntity {
	if o == nil {
		return nil
	}
	return o.MilestoneMappings
}
