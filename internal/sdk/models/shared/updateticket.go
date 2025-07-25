// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateTicket - Update a ticket's attributes
type UpdateTicket struct {
	Description *string `json:"description,omitempty"`
	PriorityID  *string `json:"priority_id,omitempty"`
	State       *string `json:"state,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	// List of tags for the ticket
	TagList []string `json:"tag_list,omitempty"`
	Type    *string  `json:"type,omitempty"`
}

func (o *UpdateTicket) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateTicket) GetPriorityID() *string {
	if o == nil {
		return nil
	}
	return o.PriorityID
}

func (o *UpdateTicket) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *UpdateTicket) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *UpdateTicket) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *UpdateTicket) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
