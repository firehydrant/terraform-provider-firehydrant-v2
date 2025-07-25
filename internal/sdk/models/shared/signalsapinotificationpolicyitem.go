// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// SignalsAPINotificationPolicyItem - Signals_API_NotificationPolicyItem model
type SignalsAPINotificationPolicyItem struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ID        *string    `json:"id,omitempty"`
	// The maximum delay for notifications
	MaxDelay                *string    `json:"max_delay,omitempty"`
	NotificationGroupMethod *string    `json:"notification_group_method,omitempty"`
	Priority                *string    `json:"priority,omitempty"`
	UpdatedAt               *time.Time `json:"updated_at,omitempty"`
}

func (s SignalsAPINotificationPolicyItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SignalsAPINotificationPolicyItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SignalsAPINotificationPolicyItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SignalsAPINotificationPolicyItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SignalsAPINotificationPolicyItem) GetMaxDelay() *string {
	if o == nil {
		return nil
	}
	return o.MaxDelay
}

func (o *SignalsAPINotificationPolicyItem) GetNotificationGroupMethod() *string {
	if o == nil {
		return nil
	}
	return o.NotificationGroupMethod
}

func (o *SignalsAPINotificationPolicyItem) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *SignalsAPINotificationPolicyItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
