data "firehydrant-terraform-sdk_services" "my_services" {
  available_downstream_dependencies_for_id = "...my_available_downstream_dependencies_for_id..."
  available_upstream_dependencies_for_id   = "...my_available_upstream_dependencies_for_id..."
  functionalities                          = "...my_functionalities..."
  impacted                                 = "...my_impacted..."
  include = [
    "..."
  ]
  labels           = "...my_labels..."
  lite             = true
  name             = "...my_name..."
  owner            = "...my_owner..."
  page             = 3
  per_page         = 7
  query            = "...my_query..."
  responding_teams = "...my_responding_teams..."
  tiers            = "...my_tiers..."
}