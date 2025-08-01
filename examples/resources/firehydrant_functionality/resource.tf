resource "firehydrant_functionality" "my_functionality" {
  alert_on_add             = true
  auto_add_responding_team = true
  description              = "...my_description..."
  external_resources_input = [
    {
      connection_type = "...my_connection_type..."
      remote_id       = "...my_remote_id..."
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
  services_input = [
    {
      id = "...my_id..."
    }
  ]
  teams_input = [
    {
      id = "...my_id..."
    }
  ]
}