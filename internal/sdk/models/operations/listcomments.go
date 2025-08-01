// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"net/http"
	"time"
)

type ListCommentsRequest struct {
	// An ISO8601 timestamp that allows filtering for comments posted before the provided time.
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// An ISO8601 timestamp that allows filtering for comments posted after the provided time.
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// Allows sorting comments by the time they were posted, ascending or descending.
	Sort           *string `default:"asc" queryParam:"style=form,explode=true,name=sort"`
	ConversationID string  `pathParam:"style=simple,explode=false,name=conversation_id"`
}

func (l ListCommentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCommentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCommentsRequest) GetBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListCommentsRequest) GetAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListCommentsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListCommentsRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

type ListCommentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCommentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCommentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCommentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
