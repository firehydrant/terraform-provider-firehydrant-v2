data "firehydrant_ticketing_projects_project_list_items_entities" "my_ticketing_projects_projectlistitemsentities" {
  configured_projects   = true
  connection_ids        = "...my_connection_ids..."
  page                  = 0
  per_page              = 7
  providers             = "...my_providers..."
  query                 = "...my_query..."
  supports_ticket_types = "...my_supports_ticket_types..."
}