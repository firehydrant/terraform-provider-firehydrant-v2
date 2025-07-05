resource "firehydrant_change_event_entity" "my_changeevententity" {
  attachments_input = [
    {
      type = "...my_type..."
    }
  ]
  authors_input = [
    {
      name      = "...my_name..."
      source    = "...my_source..."
      source_id = "...my_source_id..."
    }
  ]
  change_identities = [
    {
      type  = "...my_type..."
      value = "...my_value..."
    }
  ]
  changes = [
    "..."
  ]
  description = "...my_description..."
  ends_at     = "2022-03-21T09:48:23.407Z"
  environments_input = [
    "..."
  ]
  external_id = "...my_external_id..."
  services_input = [
    "..."
  ]
  starts_at = "2022-09-04T13:34:30.636Z"
  summary   = "...my_summary..."
}