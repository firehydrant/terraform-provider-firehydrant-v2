// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateNuncSubscription - Subscribe to status page updates
type CreateNuncSubscription struct {
	Email string `json:"email"`
}

func (o *CreateNuncSubscription) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}
