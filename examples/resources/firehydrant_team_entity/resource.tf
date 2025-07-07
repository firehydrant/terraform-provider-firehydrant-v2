resource "firehydrant_team_entity" "my_teamentity" {
  description = "...my_description..."
  invite_emails = [
    "..."
  ]
  memberships_input = [
    {
      incident_role_id = "...my_incident_role_id..."
      schedule_id      = "...my_schedule_id..."
      user_id          = "...my_user_id..."
    }
  ]
  ms_teams_channel_input = {
    channel_id = "...my_channel_id..."
    ms_team_id = "...my_ms_team_id..."
  }
  name             = "...my_name..."
  slack_channel_id = "...my_slack_channel_id..."
  slug             = "...my_slug..."
}