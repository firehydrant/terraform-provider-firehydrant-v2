// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Checks struct {
	// The name of the check
	Name string `json:"name"`
	// The description of the check
	Description *string `json:"description,omitempty"`
}

func (o *Checks) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Checks) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type ConnectedServices struct {
	ID string `json:"id"`
}

func (o *ConnectedServices) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// PostV1ChecklistTemplates - Creates a checklist template for the organization
type PostV1ChecklistTemplates struct {
	Name string `json:"name"`
	// An array of checks for the checklist template
	Checks      []Checks `json:"checks"`
	Description *string  `json:"description,omitempty"`
	// The ID of the Team that owns the checklist template
	TeamID *string `json:"team_id,omitempty"`
	// Array of service IDs to attach checklist template to
	ConnectedServices []ConnectedServices `json:"connected_services,omitempty"`
}

func (o *PostV1ChecklistTemplates) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1ChecklistTemplates) GetChecks() []Checks {
	if o == nil {
		return []Checks{}
	}
	return o.Checks
}

func (o *PostV1ChecklistTemplates) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1ChecklistTemplates) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *PostV1ChecklistTemplates) GetConnectedServices() []ConnectedServices {
	if o == nil {
		return nil
	}
	return o.ConnectedServices
}
