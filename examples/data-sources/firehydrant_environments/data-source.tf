data "firehydrant_environments" "my_environments" {
  name     = "...my_name..."
  page     = 5
  per_page = 1
  query    = "...my_query..."
}