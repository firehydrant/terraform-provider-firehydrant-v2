resource "firehydrant_signals_api_on_call_schedule" "my_signals_api_oncallschedule" {
  color       = "...my_color..."
  description = "...my_description..."
  member_ids = [
    "..."
  ]
  members_input = [
    {
      user_id = "...my_user_id..."
    }
  ]
  name = "...my_name..."
  restrictions_input = [
    {
      end_day    = "...my_end_day..."
      end_time   = "...my_end_time..."
      start_day  = "...my_start_day..."
      start_time = "...my_start_time..."
    }
  ]
  rotation_description = "...my_rotation_description..."
  rotation_name        = "...my_rotation_name..."
  slack_user_group_id  = "...my_slack_user_group_id..."
  start_time           = "...my_start_time..."
  strategy_input = {
    handoff_day    = "...my_handoff_day..."
    handoff_time   = "...my_handoff_time..."
    shift_duration = "...my_shift_duration..."
    type           = "...my_type..."
  }
  team_id   = "...my_team_id..."
  time_zone = "...my_time_zone..."
}