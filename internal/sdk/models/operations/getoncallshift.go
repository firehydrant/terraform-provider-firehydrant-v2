// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type GetOnCallShiftRequest struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TeamID     string `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule_id"`
}

func (o *GetOnCallShiftRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetOnCallShiftRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetOnCallShiftRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

type GetOnCallShiftResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get a Signals on-call shift by ID
	SignalsAPIOnCallShift *shared.SignalsAPIOnCallShift
}

func (o *GetOnCallShiftResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOnCallShiftResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOnCallShiftResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOnCallShiftResponse) GetSignalsAPIOnCallShift() *shared.SignalsAPIOnCallShift {
	if o == nil {
		return nil
	}
	return o.SignalsAPIOnCallShift
}
