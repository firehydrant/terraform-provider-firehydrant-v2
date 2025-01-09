resource "firehydrant-terraform-sdk_incident_type" "my_incidenttype" {
  name = "...my_name..."
  template = {
    customer_impact_summary = "...my_customer_impact_summary..."
    description             = "...my_description..."
    impacts = [
      {
        condition_id = "...my_condition_id..."
        id           = "...my_id..."
      }
    ]
    labels = {
      "see" : "documentation",
    }
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