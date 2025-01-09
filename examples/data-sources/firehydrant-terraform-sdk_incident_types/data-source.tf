data "firehydrant-terraform-sdk_incident_types" "my_incidenttypes" {
  page     = 10
  per_page = 5
  query    = "...my_query..."
}