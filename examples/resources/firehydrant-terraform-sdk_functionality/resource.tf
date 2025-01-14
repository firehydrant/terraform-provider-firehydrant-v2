resource "firehydrant-terraform-sdk_functionality" "my_functionality" {
  alert_on_add             = true
  auto_add_responding_team = true
  description              = "...my_description..."
  external_resources = [
    {
      connection_type = "...my_connection_type..."
      remote_id       = "...my_remote_id..."
    }
  ]
  labels = {
    "see" : "documentation",
  }
  links = [
    {
      href_url = "...my_href_url..."
      icon_url = "...my_icon_url..."
      name     = "...my_name..."
    }
  ]
  name = "...my_name..."
  owner = {
    id = "...my_id..."
  }
  services = [
    {
      id = "...my_id..."
    }
  ]
}