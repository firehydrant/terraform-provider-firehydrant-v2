// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UserPaginated model
type UserPaginated struct {
	Data       []User              `json:"data,omitempty"`
	Pagination *NullablePagination `json:"pagination,omitempty"`
}

func (o *UserPaginated) GetData() []User {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UserPaginated) GetPagination() *NullablePagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
