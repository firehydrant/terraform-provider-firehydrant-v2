// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1TaskListsTaskListIDRequest struct {
	TaskListID                 string                            `pathParam:"style=simple,explode=false,name=task_list_id"`
	PatchV1TaskListsTaskListID shared.PatchV1TaskListsTaskListID `request:"mediaType=application/json"`
}

func (o *PatchV1TaskListsTaskListIDRequest) GetTaskListID() string {
	if o == nil {
		return ""
	}
	return o.TaskListID
}

func (o *PatchV1TaskListsTaskListIDRequest) GetPatchV1TaskListsTaskListID() shared.PatchV1TaskListsTaskListID {
	if o == nil {
		return shared.PatchV1TaskListsTaskListID{}
	}
	return o.PatchV1TaskListsTaskListID
}

type PatchV1TaskListsTaskListIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updates a task list's attributes and task list items
	TaskListEntity *shared.TaskListEntity
}

func (o *PatchV1TaskListsTaskListIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1TaskListsTaskListIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1TaskListsTaskListIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1TaskListsTaskListIDResponse) GetTaskListEntity() *shared.TaskListEntity {
	if o == nil {
		return nil
	}
	return o.TaskListEntity
}
