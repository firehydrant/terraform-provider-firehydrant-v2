// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NullableSignalsAPITransposerExamplePayload struct {
}

// NullableSignalsAPITransposer - Signals_API_Transposer model
type NullableSignalsAPITransposer struct {
	CreatedBy      *NullableAuthor                             `json:"created_by,omitempty"`
	Description    *string                                     `json:"description,omitempty"`
	Editable       *bool                                       `json:"editable,omitempty"`
	ExamplePayload *NullableSignalsAPITransposerExamplePayload `json:"example_payload,omitempty"`
	Expected       *string                                     `json:"expected,omitempty"`
	Expression     *string                                     `json:"expression,omitempty"`
	IngestURL      *string                                     `json:"ingest_url,omitempty"`
	Name           *string                                     `json:"name,omitempty"`
	Slug           *string                                     `json:"slug,omitempty"`
	Tags           []string                                    `json:"tags,omitempty"`
	UpdatedBy      *NullableAuthor                             `json:"updated_by,omitempty"`
	Website        *string                                     `json:"website,omitempty"`
}

func (o *NullableSignalsAPITransposer) GetCreatedBy() *NullableAuthor {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *NullableSignalsAPITransposer) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *NullableSignalsAPITransposer) GetEditable() *bool {
	if o == nil {
		return nil
	}
	return o.Editable
}

func (o *NullableSignalsAPITransposer) GetExamplePayload() *NullableSignalsAPITransposerExamplePayload {
	if o == nil {
		return nil
	}
	return o.ExamplePayload
}

func (o *NullableSignalsAPITransposer) GetExpected() *string {
	if o == nil {
		return nil
	}
	return o.Expected
}

func (o *NullableSignalsAPITransposer) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}

func (o *NullableSignalsAPITransposer) GetIngestURL() *string {
	if o == nil {
		return nil
	}
	return o.IngestURL
}

func (o *NullableSignalsAPITransposer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NullableSignalsAPITransposer) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *NullableSignalsAPITransposer) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *NullableSignalsAPITransposer) GetUpdatedBy() *NullableAuthor {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *NullableSignalsAPITransposer) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}
