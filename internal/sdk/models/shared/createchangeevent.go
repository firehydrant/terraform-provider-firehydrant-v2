// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type CreateChangeEventAttachmentsInput struct {
	Type string `json:"type"`
}

func (o *CreateChangeEventAttachmentsInput) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type AuthorsInput struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	SourceID string `json:"source_id"`
}

func (o *AuthorsInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AuthorsInput) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}

func (o *AuthorsInput) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type CreateChangeEventChangeIdentity struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (o *CreateChangeEventChangeIdentity) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *CreateChangeEventChangeIdentity) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateChangeEvent - Create a change event
type CreateChangeEvent struct {
	// JSON objects representing attachments, see attachments documentation for the schema
	AttachmentsInput []CreateChangeEventAttachmentsInput `json:"attachments,omitempty"`
	// Array of additional authors to add to the change event, the creating actor will automatically be added as an author
	AuthorsInput []AuthorsInput `json:"authors,omitempty"`
	// If provided and valid, the event will be linked to all changes that have the same identities. Identity *values* must be unique.
	ChangeIdentities []CreateChangeEventChangeIdentity `json:"change_identities,omitempty"`
	// An array of change IDs
	Changes     []string   `json:"changes,omitempty"`
	Description *string    `json:"description,omitempty"`
	EndsAt      *time.Time `json:"ends_at,omitempty"`
	// An array of environment IDs
	EnvironmentsInput []string `json:"environments,omitempty"`
	// The ID of a change event as assigned by an external provider
	ExternalID *string        `json:"external_id,omitempty"`
	Labels     map[string]any `json:"labels,omitempty"`
	// An array of service IDs
	ServicesInput []string   `json:"services,omitempty"`
	StartsAt      *time.Time `json:"starts_at,omitempty"`
	Summary       string     `json:"summary"`
}

func (c CreateChangeEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateChangeEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateChangeEvent) GetAttachmentsInput() []CreateChangeEventAttachmentsInput {
	if o == nil {
		return nil
	}
	return o.AttachmentsInput
}

func (o *CreateChangeEvent) GetAuthorsInput() []AuthorsInput {
	if o == nil {
		return nil
	}
	return o.AuthorsInput
}

func (o *CreateChangeEvent) GetChangeIdentities() []CreateChangeEventChangeIdentity {
	if o == nil {
		return nil
	}
	return o.ChangeIdentities
}

func (o *CreateChangeEvent) GetChanges() []string {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *CreateChangeEvent) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateChangeEvent) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *CreateChangeEvent) GetEnvironmentsInput() []string {
	if o == nil {
		return nil
	}
	return o.EnvironmentsInput
}

func (o *CreateChangeEvent) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *CreateChangeEvent) GetLabels() map[string]any {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *CreateChangeEvent) GetServicesInput() []string {
	if o == nil {
		return nil
	}
	return o.ServicesInput
}

func (o *CreateChangeEvent) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *CreateChangeEvent) GetSummary() string {
	if o == nil {
		return ""
	}
	return o.Summary
}
