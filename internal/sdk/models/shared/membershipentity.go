// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MembershipEntity struct {
	User     *UserEntity     `json:"user,omitempty"`
	Schedule *ScheduleEntity `json:"schedule,omitempty"`
	// IncidentRoleEntity model
	DefaultIncidentRole *IncidentRoleEntity `json:"default_incident_role,omitempty"`
}

func (o *MembershipEntity) GetUser() *UserEntity {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *MembershipEntity) GetSchedule() *ScheduleEntity {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *MembershipEntity) GetDefaultIncidentRole() *IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.DefaultIncidentRole
}
