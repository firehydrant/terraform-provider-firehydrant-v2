// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ImportsImportErrorEntityResourceEntity struct {
	ResourceID   *string `json:"resource_id,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Name         *string `json:"name,omitempty"`
}

func (o *ImportsImportErrorEntityResourceEntity) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *ImportsImportErrorEntityResourceEntity) GetResourceType() *string {
	if o == nil {
		return nil
	}
	return o.ResourceType
}

func (o *ImportsImportErrorEntityResourceEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
