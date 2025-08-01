// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateChangeIdentity - Update an identity for the change entry
type UpdateChangeIdentity struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (o *UpdateChangeIdentity) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateChangeIdentity) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}
