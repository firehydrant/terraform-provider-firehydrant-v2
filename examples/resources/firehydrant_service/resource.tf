resource "firehydrant_service" "my_service" {
  alert_on_add             = true
  auto_add_responding_team = false
  description              = "...my_description..."
  external_resources_input = [
    {
      connection_type = "...my_connection_type..."
      remote_id       = "...my_remote_id..."
    }
  ]
  functionalities_input = [
    {
      id      = "...my_id..."
      summary = "...my_summary..."
    }
  ]
  links_input = [
    {
      href_url = "...my_href_url..."
      icon_url = "...my_icon_url..."
      name     = "...my_name..."
    }
  ]
  name = "...my_name..."
  owner_input = {
    id = "...my_id..."
  }
  service_tier = 4
  teams_input = [
    {
      id = "...my_id..."
    }
  ]
}