data "firehydrant_scheduled_maintenances" "my_scheduledmaintenances" {
  page     = 0
  per_page = 1
  query    = "...my_query..."
}