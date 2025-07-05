resource "firehydrant_signals_api_email_target_entity" "my_signals_api_emailtargetentity" {
  allowed_senders = [
    "..."
  ]
  description            = "...my_description..."
  level_cel              = "...my_level_cel..."
  name                   = "...my_name..."
  rule_matching_strategy = "...my_rule_matching_strategy..."
  rules = [
    "..."
  ]
  slug       = "...my_slug..."
  status_cel = "...my_status_cel..."
  target_input = {
    id   = "...my_id..."
    type = "...my_type..."
  }
}