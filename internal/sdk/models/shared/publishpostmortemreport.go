// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PublishPostMortemReport - Marks an incident retrospective as published and emails all of the participants in the report the summary
type PublishPostMortemReport struct {
	Publish *string `json:"publish,omitempty"`
	// An array of team IDs with whom to share the report
	TeamIds []string `json:"team_ids,omitempty"`
	// An array of user IDs with whom to share the report
	UserIds []string `json:"user_ids,omitempty"`
}

func (o *PublishPostMortemReport) GetPublish() *string {
	if o == nil {
		return nil
	}
	return o.Publish
}

func (o *PublishPostMortemReport) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

func (o *PublishPostMortemReport) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}
