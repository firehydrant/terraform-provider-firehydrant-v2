resource "firehydrant_runbook_entity" "my_runbookentity" {
  attachment_rule_input = {
    logic     = "...my_logic..."
    user_data = "...my_user_data..."
  }
  auto_attach_to_restricted_incidents = true
  description                         = "...my_description..."
  name                                = "...my_name..."
  owner_input = {
    id = "...my_id..."
  }
  steps_input = [
    {
      action_id = "...my_action_id..."
      name      = "...my_name..."
      rule = {
        logic     = "...my_logic..."
        user_data = "...my_user_data..."
      }
    }
  ]
  summary  = "...my_summary..."
  tutorial = false
  type     = "...my_type..."
}