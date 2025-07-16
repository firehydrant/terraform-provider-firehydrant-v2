data "firehydrant_ticketing_projects_project_list_items" "my_ticketing_projects_projectlistitems" {
  configured_projects   = false
  connection_ids        = "...my_connection_ids..."
  page                  = 5
  per_page              = 7
  providers             = "...my_providers..."
  query                 = "...my_query..."
  supports_ticket_types = "...my_supports_ticket_types..."
}