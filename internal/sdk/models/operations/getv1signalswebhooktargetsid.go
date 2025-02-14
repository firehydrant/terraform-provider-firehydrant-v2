// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1SignalsWebhookTargetsIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetV1SignalsWebhookTargetsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetV1SignalsWebhookTargetsIDCreatedBy struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
	Email  *string `json:"email,omitempty"`
}

func (o *GetV1SignalsWebhookTargetsIDCreatedBy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetV1SignalsWebhookTargetsIDCreatedBy) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetV1SignalsWebhookTargetsIDCreatedBy) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *GetV1SignalsWebhookTargetsIDCreatedBy) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

// GetV1SignalsWebhookTargetsIDResponseBody - Get a Signals webhook target by ID
type GetV1SignalsWebhookTargetsIDResponseBody struct {
	ID            *string                                `json:"id,omitempty"`
	URL           *string                                `json:"url,omitempty"`
	State         *string                                `json:"state,omitempty"`
	CreatedBy     *GetV1SignalsWebhookTargetsIDCreatedBy `json:"created_by,omitempty"`
	CreatedAt     *string                                `json:"created_at,omitempty"`
	UpdatedAt     *string                                `json:"updated_at,omitempty"`
	Subscriptions *string                                `json:"subscriptions,omitempty"`
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetCreatedBy() *GetV1SignalsWebhookTargetsIDCreatedBy {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetV1SignalsWebhookTargetsIDResponseBody) GetSubscriptions() *string {
	if o == nil {
		return nil
	}
	return o.Subscriptions
}

type GetV1SignalsWebhookTargetsIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get a Signals webhook target by ID
	Object *GetV1SignalsWebhookTargetsIDResponseBody
}

func (o *GetV1SignalsWebhookTargetsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1SignalsWebhookTargetsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1SignalsWebhookTargetsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1SignalsWebhookTargetsIDResponse) GetObject() *GetV1SignalsWebhookTargetsIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
