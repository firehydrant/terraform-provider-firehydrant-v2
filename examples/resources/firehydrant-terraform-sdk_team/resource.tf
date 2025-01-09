resource "firehydrant-terraform-sdk_team" "my_team" {
  description = "...my_description..."
  memberships = [
    {
      incident_role_id = "...my_incident_role_id..."
      schedule_id      = "...my_schedule_id..."
      user_id          = "...my_user_id..."
    }
  ]
  ms_teams_channel = {
    channel_id = "...my_channel_id..."
    ms_team_id = "...my_ms_team_id..."
  }
  name             = "...my_name..."
  slack_channel_id = "...my_slack_channel_id..."
  slug             = "...my_slug..."
}