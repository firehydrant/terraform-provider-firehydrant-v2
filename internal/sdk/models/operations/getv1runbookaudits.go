// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"net/http"
)

// AuditableType - A query to filter audits by type
type AuditableType string

const (
	AuditableTypeRunbooksStep    AuditableType = "Runbooks::Step"
	AuditableTypeRunbooksRunbook AuditableType = "Runbooks::Runbook"
)

func (e AuditableType) ToPointer() *AuditableType {
	return &e
}
func (e *AuditableType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Runbooks::Step":
		fallthrough
	case "Runbooks::Runbook":
		*e = AuditableType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuditableType: %v", v)
	}
}

type GetV1RunbookAuditsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A query to filter audits by type
	AuditableType *AuditableType `default:"Runbooks::Step" queryParam:"style=form,explode=true,name=auditable_type"`
	// A query to sort audits by their created_at timestamp. Options are 'asc' or 'desc'
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
}

func (g GetV1RunbookAuditsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1RunbookAuditsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1RunbookAuditsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1RunbookAuditsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1RunbookAuditsRequest) GetAuditableType() *AuditableType {
	if o == nil {
		return nil
	}
	return o.AuditableType
}

func (o *GetV1RunbookAuditsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

type GetV1RunbookAuditsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1RunbookAuditsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1RunbookAuditsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1RunbookAuditsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
