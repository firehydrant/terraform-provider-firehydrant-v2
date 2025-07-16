resource "firehydrant_incident_type" "my_incidenttype" {
  description = "...my_description..."
  name        = "...my_name..."
  template = {
    customer_impact_summary = "...my_customer_impact_summary..."
    description             = "...my_description..."
    impacts = [
      {
        condition_id = "...my_condition_id..."
        id           = "...my_id..."
      }
    ]
    priority         = "...my_priority..."
    private_incident = false
    runbook_ids = [
      "..."
    ]
    severity = "...my_severity..."
    tag_list = [
      "..."
    ]
    team_ids = [
      "..."
    ]
  }
}