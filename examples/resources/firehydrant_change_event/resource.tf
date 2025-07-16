resource "firehydrant_change_event" "my_changeevent" {
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
  ends_at     = "2022-06-18T18:47:10.816Z"
  environments_input = [
    "..."
  ]
  external_id = "...my_external_id..."
  services_input = [
    "..."
  ]
  starts_at = "2020-01-28T14:07:12.119Z"
  summary   = "...my_summary..."
}