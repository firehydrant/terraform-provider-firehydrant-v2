// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateStatusUpdateTemplate - Create a status update template for your organization
type CreateStatusUpdateTemplate struct {
	Body string `json:"body"`
	Name string `json:"name"`
}

func (o *CreateStatusUpdateTemplate) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}

func (o *CreateStatusUpdateTemplate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
