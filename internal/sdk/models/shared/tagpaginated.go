// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// TagPaginated model
type TagPaginated struct {
	Data       []Tag               `json:"data,omitempty"`
	Pagination *NullablePagination `json:"pagination,omitempty"`
}

func (o *TagPaginated) GetData() []Tag {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TagPaginated) GetPagination() *NullablePagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
