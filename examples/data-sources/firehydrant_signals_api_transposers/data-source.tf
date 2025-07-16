data "firehydrant_signals_api_transposers" "my_signals_api_transposers" {
  escalation_policy_id = "...my_escalation_policy_id..."
  on_call_schedule_id  = "...my_on_call_schedule_id..."
  team_id              = "...my_team_id..."
  user_id              = "...my_user_id..."
}