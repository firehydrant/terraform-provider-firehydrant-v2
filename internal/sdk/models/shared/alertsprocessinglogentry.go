// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// Context - An unstructured representation of this log entry's context.
type Context struct {
}

type AlertsProcessingLogEntry struct {
	// An unstructured representation of this log entry's context.
	Context     *Context   `json:"context,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ID          *string    `json:"id,omitempty"`
	Level       *string    `json:"level,omitempty"`
	Message     *string    `json:"message,omitempty"`
	MessageType *string    `json:"message_type,omitempty"`
}

func (a AlertsProcessingLogEntry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlertsProcessingLogEntry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AlertsProcessingLogEntry) GetContext() *Context {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *AlertsProcessingLogEntry) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AlertsProcessingLogEntry) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AlertsProcessingLogEntry) GetLevel() *string {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *AlertsProcessingLogEntry) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AlertsProcessingLogEntry) GetMessageType() *string {
	if o == nil {
		return nil
	}
	return o.MessageType
}
