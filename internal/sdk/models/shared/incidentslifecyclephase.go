// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type IncidentsLifecyclePhase struct {
	Description *string                       `json:"description,omitempty"`
	ID          *string                       `json:"id,omitempty"`
	Milestones  []IncidentsLifecycleMilestone `json:"milestones,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	Position    *int                          `json:"position,omitempty"`
	Type        *string                       `json:"type,omitempty"`
}

func (o *IncidentsLifecyclePhase) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentsLifecyclePhase) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentsLifecyclePhase) GetMilestones() []IncidentsLifecycleMilestone {
	if o == nil {
		return nil
	}
	return o.Milestones
}

func (o *IncidentsLifecyclePhase) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IncidentsLifecyclePhase) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *IncidentsLifecyclePhase) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
