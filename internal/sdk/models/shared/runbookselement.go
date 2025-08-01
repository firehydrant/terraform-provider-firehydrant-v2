// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type RunbooksElement struct {
	DynamicSelect *NullableRunbooksElementDynamicSelect `json:"dynamic_select,omitempty"`
	ID            *string                               `json:"id,omitempty"`
	Input         *NullableRunbooksElementInput         `json:"input,omitempty"`
	Markdown      *NullableRunbooksElementMarkdown      `json:"markdown,omitempty"`
	PlainText     *NullableRunbooksElementMarkdown      `json:"plain_text,omitempty"`
	Textarea      *NullableRunbooksElementTextarea      `json:"textarea,omitempty"`
	Type          *string                               `json:"type,omitempty"`
}

func (o *RunbooksElement) GetDynamicSelect() *NullableRunbooksElementDynamicSelect {
	if o == nil {
		return nil
	}
	return o.DynamicSelect
}

func (o *RunbooksElement) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RunbooksElement) GetInput() *NullableRunbooksElementInput {
	if o == nil {
		return nil
	}
	return o.Input
}

func (o *RunbooksElement) GetMarkdown() *NullableRunbooksElementMarkdown {
	if o == nil {
		return nil
	}
	return o.Markdown
}

func (o *RunbooksElement) GetPlainText() *NullableRunbooksElementMarkdown {
	if o == nil {
		return nil
	}
	return o.PlainText
}

func (o *RunbooksElement) GetTextarea() *NullableRunbooksElementTextarea {
	if o == nil {
		return nil
	}
	return o.Textarea
}

func (o *RunbooksElement) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
