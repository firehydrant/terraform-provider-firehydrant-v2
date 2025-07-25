// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateScimGroupMember struct {
	// String that represents the user's UUID to assign to the team
	Value string `json:"value"`
}

func (o *CreateScimGroupMember) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateScimGroup - SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.
type CreateScimGroup struct {
	// The name of the team being created
	DisplayName string                  `json:"displayName"`
	Members     []CreateScimGroupMember `json:"members"`
}

func (o *CreateScimGroup) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *CreateScimGroup) GetMembers() []CreateScimGroupMember {
	if o == nil {
		return []CreateScimGroupMember{}
	}
	return o.Members
}
