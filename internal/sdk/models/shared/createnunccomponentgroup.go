// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateNuncComponentGroup struct {
	ComponentGroupID *string `json:"component_group_id,omitempty"`
	Name             string  `json:"name"`
	Position         *int    `json:"position,omitempty"`
}

func (o *CreateNuncComponentGroup) GetComponentGroupID() *string {
	if o == nil {
		return nil
	}
	return o.ComponentGroupID
}

func (o *CreateNuncComponentGroup) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateNuncComponentGroup) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}
