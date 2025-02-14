// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PutV1PostMortemsQuestionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update the questions configured to be provided and filled out on future retrospective reports.
	PostMortemsQuestionTypeEntity *shared.PostMortemsQuestionTypeEntity
}

func (o *PutV1PostMortemsQuestionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutV1PostMortemsQuestionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutV1PostMortemsQuestionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutV1PostMortemsQuestionsResponse) GetPostMortemsQuestionTypeEntity() *shared.PostMortemsQuestionTypeEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsQuestionTypeEntity
}
