// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type ListAwsConnectionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// AWS account ID containing the role to be assumed
	AwsAccountID *string `queryParam:"style=form,explode=true,name=aws_account_id"`
	// ARN of the role to be assumed
	TargetArn *string `queryParam:"style=form,explode=true,name=target_arn"`
	// The external ID supplied when assuming the role
	ExternalID *string `queryParam:"style=form,explode=true,name=external_id"`
}

func (o *ListAwsConnectionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAwsConnectionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListAwsConnectionsRequest) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *ListAwsConnectionsRequest) GetTargetArn() *string {
	if o == nil {
		return nil
	}
	return o.TargetArn
}

func (o *ListAwsConnectionsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

type ListAwsConnectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Lists the available and configured AWS integration connections for the authenticated organization.
	IntegrationsAwsConnectionPaginated *shared.IntegrationsAwsConnectionPaginated
}

func (o *ListAwsConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAwsConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAwsConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAwsConnectionsResponse) GetIntegrationsAwsConnectionPaginated() *shared.IntegrationsAwsConnectionPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsConnectionPaginated
}
