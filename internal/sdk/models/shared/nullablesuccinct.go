// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NullableSuccinct struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *NullableSuccinct) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableSuccinct) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
