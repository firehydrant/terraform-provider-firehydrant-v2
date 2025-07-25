// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateIncidentRoleAssignment - Assign a role to a user for this incident
type CreateIncidentRoleAssignment struct {
	// The Incident Role ID to assign the role to
	IncidentRoleID string `json:"incident_role_id"`
	// The user ID to assign the role to
	UserID string `json:"user_id"`
}

func (o *CreateIncidentRoleAssignment) GetIncidentRoleID() string {
	if o == nil {
		return ""
	}
	return o.IncidentRoleID
}

func (o *CreateIncidentRoleAssignment) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
