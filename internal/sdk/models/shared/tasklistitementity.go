// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TaskListItemEntity struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (o *TaskListItemEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *TaskListItemEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
