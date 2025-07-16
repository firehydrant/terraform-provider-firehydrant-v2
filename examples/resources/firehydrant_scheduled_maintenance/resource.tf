resource "firehydrant_scheduled_maintenance" "my_scheduledmaintenance" {
  description = "...my_description..."
  ends_at     = "2022-06-27T06:25:09.256Z"
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
  starts_at = "2022-04-02T12:05:43.660Z"
  status_pages_input = [
    {
      connection_id    = "...my_connection_id..."
      integration_slug = "...my_integration_slug..."
    }
  ]
  summary = "...my_summary..."
}