resource "firehydrant_incident" "my_incident" {
  alert_ids = [
    "..."
  ]
  custom_fields_input = [
    {
      field_id = "...my_field_id..."
      value_array = [
        "..."
      ]
      value_string = "...my_value_string..."
    }
  ]
  customer_impact_summary = "...my_customer_impact_summary..."
  description             = "...my_description..."
  external_links          = "...my_external_links..."
  impacts_input = [
    {
      condition_id = "...my_condition_id..."
      id           = "...my_id..."
      type         = "...my_type..."
    }
  ]
  incident_type_id = "...my_incident_type_id..."
  labels = {
    # ...
  }
  milestones_input = [
    {
      occurred_at = "2022-06-05T16:55:15.357Z"
      type        = "...my_type..."
    }
  ]
  name       = "...my_name..."
  priority   = "...my_priority..."
  restricted = false
  runbook_ids = [
    "..."
  ]
  severity                  = "...my_severity..."
  severity_condition_id     = "...my_severity_condition_id..."
  severity_impact_id        = "...my_severity_impact_id..."
  skip_incident_type_values = false
  summary                   = "...my_summary..."
  tag_list = [
    "..."
  ]
  team_ids = [
    "..."
  ]
}