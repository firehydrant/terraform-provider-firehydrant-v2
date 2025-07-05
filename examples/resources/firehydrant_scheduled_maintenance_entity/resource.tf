resource "firehydrant_scheduled_maintenance_entity" "my_scheduledmaintenanceentity" {
  description = "...my_description..."
  ends_at     = "2022-04-07T08:13:54.801Z"
  impacts_input = [
    {
      condition_id = "...my_condition_id..."
      id           = "...my_id..."
      type         = "...my_type..."
    }
  ]
  labels = {
    # ...
  }
  name      = "...my_name..."
  starts_at = "2022-01-27T15:03:02.842Z"
  status_pages_input = [
    {
      connection_id    = "...my_connection_id..."
      integration_slug = "...my_integration_slug..."
    }
  ]
  summary = "...my_summary..."
}