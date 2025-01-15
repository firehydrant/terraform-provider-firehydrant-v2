// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// WebhooksEntitiesWebhookEntity - Webhooks_Entities_WebhookEntity model
type WebhooksEntitiesWebhookEntity struct {
	ID            *string       `json:"id,omitempty"`
	URL           *string       `json:"url,omitempty"`
	State         *string       `json:"state,omitempty"`
	CreatedBy     *AuthorEntity `json:"created_by,omitempty"`
	CreatedAt     *time.Time    `json:"created_at,omitempty"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
	Subscriptions *string       `json:"subscriptions,omitempty"`
}

func (w WebhooksEntitiesWebhookEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhooksEntitiesWebhookEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebhooksEntitiesWebhookEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebhooksEntitiesWebhookEntity) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *WebhooksEntitiesWebhookEntity) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WebhooksEntitiesWebhookEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *WebhooksEntitiesWebhookEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WebhooksEntitiesWebhookEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *WebhooksEntitiesWebhookEntity) GetSubscriptions() *string {
	if o == nil {
		return nil
	}
	return o.Subscriptions
}
