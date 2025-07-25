// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IncidentsRelationships - Incidents_Relationships model
type IncidentsRelationships struct {
	// The root incident's child incidents.
	Children []PublicAPIV1IncidentsSuccinct        `json:"children,omitempty"`
	Parent   *NullablePublicAPIV1IncidentsSuccinct `json:"parent,omitempty"`
	// A list of incidents that share the root incident's parent.
	Siblings []PublicAPIV1IncidentsSuccinct `json:"siblings,omitempty"`
}

func (o *IncidentsRelationships) GetChildren() []PublicAPIV1IncidentsSuccinct {
	if o == nil {
		return nil
	}
	return o.Children
}

func (o *IncidentsRelationships) GetParent() *NullablePublicAPIV1IncidentsSuccinct {
	if o == nil {
		return nil
	}
	return o.Parent
}

func (o *IncidentsRelationships) GetSiblings() []PublicAPIV1IncidentsSuccinct {
	if o == nil {
		return nil
	}
	return o.Siblings
}
