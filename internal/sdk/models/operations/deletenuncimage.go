// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type DeleteNuncImageRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	Type             string `pathParam:"style=simple,explode=false,name=type"`
}

func (o *DeleteNuncImageRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *DeleteNuncImageRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type DeleteNuncImageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete an image attached to a FireHydrant status page
	NuncConnection *shared.NuncConnection
}

func (o *DeleteNuncImageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNuncImageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNuncImageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteNuncImageResponse) GetNuncConnection() *shared.NuncConnection {
	if o == nil {
		return nil
	}
	return o.NuncConnection
}
