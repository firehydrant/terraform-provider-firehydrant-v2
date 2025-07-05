resource "firehydrant_signals_api_rule_entity" "my_signals_api_ruleentity" {
  create_incident_condition_when = "...my_create_incident_condition_when..."
  deduplication_expiry           = "PT1H"
  expression                     = "...my_expression..."
  incident_type_id               = "...my_incident_type_id..."
  name                           = "...my_name..."
  notification_priority_override = "...my_notification_priority_override..."
  target_id                      = "...my_target_id..."
  target_type                    = "...my_target_type..."
}