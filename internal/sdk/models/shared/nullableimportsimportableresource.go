// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type NullableImportsImportableResource struct {
	ImportErrors []ImportsImportError `json:"import_errors,omitempty"`
	ImportedAt   *time.Time           `json:"imported_at,omitempty"`
	RemoteID     *string              `json:"remote_id,omitempty"`
	State        *string              `json:"state,omitempty"`
}

func (n NullableImportsImportableResource) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NullableImportsImportableResource) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NullableImportsImportableResource) GetImportErrors() []ImportsImportError {
	if o == nil {
		return nil
	}
	return o.ImportErrors
}

func (o *NullableImportsImportableResource) GetImportedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ImportedAt
}

func (o *NullableImportsImportableResource) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *NullableImportsImportableResource) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
