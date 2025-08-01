// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostMortemsQuestionTypePaginated - PostMortems_QuestionTypePaginated model
type PostMortemsQuestionTypePaginated struct {
	Data       []PostMortemsQuestionType `json:"data,omitempty"`
	Pagination *NullablePagination       `json:"pagination,omitempty"`
}

func (o *PostMortemsQuestionTypePaginated) GetData() []PostMortemsQuestionType {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PostMortemsQuestionTypePaginated) GetPagination() *NullablePagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
