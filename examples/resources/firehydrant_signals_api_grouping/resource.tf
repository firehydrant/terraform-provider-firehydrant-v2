resource "firehydrant_signals_api_grouping" "my_signals_api_grouping" {
  action_input = {
    fyi = {
      slack_channel_ids = [
        "..."
      ]
    }
    link = false
  }
  reference_alert_time_period = "...my_reference_alert_time_period..."
  strategy = {
    substring = {
      field_name = "...my_field_name..."
      value      = "...my_value..."
    }
  }
}