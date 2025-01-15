data "firehydrant_teams" "my_teams" {
  default_incident_role = "...my_default_incident_role..."
  lite                  = true
  name                  = "...my_name..."
  page                  = 6
  per_page              = 10
  query                 = "...my_query..."
  services              = "...my_services..."
}