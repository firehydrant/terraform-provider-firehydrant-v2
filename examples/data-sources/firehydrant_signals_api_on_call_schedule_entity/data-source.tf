data "firehydrant_signals_api_on_call_schedule_entity" "my_signals_api_oncallscheduleentity" {
  schedule_id             = "...my_schedule_id..."
  shift_time_window_end   = "...my_shift_time_window_end..."
  shift_time_window_start = "...my_shift_time_window_start..."
  team_id                 = "...my_team_id..."
}