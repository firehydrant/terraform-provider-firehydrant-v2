resource "firehydrant_nunc_connection_entity" "my_nuncconnectionentity" {
  company_name    = "...my_company_name..."
  company_tos_url = "...my_company_tos_url..."
  company_website = "...my_company_website..."
  components_infrastructure_id = [
    "..."
  ]
  components_infrastructure_type = [
    "..."
  ]
  conditions_condition_id = [
    "..."
  ]
  conditions_nunc_condition = [
    "..."
  ]
  domain           = "...my_domain..."
  enable_histogram = true
  exposed_fields = [
    "..."
  ]
  greeting_body       = "...my_greeting_body..."
  greeting_title      = "...my_greeting_title..."
  operational_message = "...my_operational_message..."
  primary_color       = "...my_primary_color..."
  secondary_color     = "...my_secondary_color..."
  title               = "...my_title..."
  ui_version          = 2
}