// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateEnvironment - Creates an environment for the organization
type CreateEnvironment struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

func (o *CreateEnvironment) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateEnvironment) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
