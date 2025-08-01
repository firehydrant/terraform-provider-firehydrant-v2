// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateSeverity - Update a specific severity
type UpdateSeverity struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Position    *int    `json:"position,omitempty"`
	Slug        *string `json:"slug,omitempty"`
}

func (o *UpdateSeverity) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *UpdateSeverity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateSeverity) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *UpdateSeverity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}
