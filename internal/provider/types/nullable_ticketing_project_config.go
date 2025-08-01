// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NullableTicketingProjectConfig struct {
	ConnectionID         types.String                           `tfsdk:"connection_id"`
	ConnectionType       types.String                           `tfsdk:"connection_type"`
	Details              *NullableTicketingProjectConfigDetails `tfsdk:"details"`
	ID                   types.String                           `tfsdk:"id"`
	TicketingProjectID   types.String                           `tfsdk:"ticketing_project_id"`
	TicketingProjectName types.String                           `tfsdk:"ticketing_project_name"`
}
