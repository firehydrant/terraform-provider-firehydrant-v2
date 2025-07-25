// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// PublicAPIV1IncidentsTranscript - PublicAPI_V1_Incidents_Transcript model
type PublicAPIV1IncidentsTranscript struct {
	Author *NullableAuthor `json:"author,omitempty"`
	// The time the transcript entry was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The unique identifier for the transcript entry
	ID *string `json:"id,omitempty"`
	// The speaker for the transcript entry
	Speaker *string `json:"speaker,omitempty"`
	// The start time for the transcript entry
	Start *int `json:"start,omitempty"`
	// The end time for the transcript entry
	Until *int `json:"until,omitempty"`
	// The time the transcript entry was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The words spoken for the transcript entry
	Words *string `json:"words,omitempty"`
}

func (p PublicAPIV1IncidentsTranscript) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PublicAPIV1IncidentsTranscript) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PublicAPIV1IncidentsTranscript) GetAuthor() *NullableAuthor {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *PublicAPIV1IncidentsTranscript) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PublicAPIV1IncidentsTranscript) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PublicAPIV1IncidentsTranscript) GetSpeaker() *string {
	if o == nil {
		return nil
	}
	return o.Speaker
}

func (o *PublicAPIV1IncidentsTranscript) GetStart() *int {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *PublicAPIV1IncidentsTranscript) GetUntil() *int {
	if o == nil {
		return nil
	}
	return o.Until
}

func (o *PublicAPIV1IncidentsTranscript) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PublicAPIV1IncidentsTranscript) GetWords() *string {
	if o == nil {
		return nil
	}
	return o.Words
}
