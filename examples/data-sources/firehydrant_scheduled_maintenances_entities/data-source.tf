data "firehydrant_scheduled_maintenances_entities" "my_scheduledmaintenancesentities" {
  page     = 0
  per_page = 6
  query    = "...my_query..."
}