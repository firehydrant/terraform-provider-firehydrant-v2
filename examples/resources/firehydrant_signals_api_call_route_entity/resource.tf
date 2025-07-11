resource "firehydrant_signals_api_call_route_entity" "my_signals_api_callrouteentity" {
  connect_mode     = "...my_connect_mode..."
  description      = "...my_description..."
  greeting_message = "...my_greeting_message..."
  name             = "...my_name..."
  phone_number     = "...my_phone_number..."
  routing_mode     = "...my_routing_mode..."
  steps_input = [
    {
      on_call_rotation_id = "...my_on_call_rotation_id..."
      target_id           = "...my_target_id..."
      target_type         = "...my_target_type..."
      timeout             = "...my_timeout..."
    }
  ]
  target_input = {
    id   = "...my_id..."
    type = "...my_type..."
  }
  team_id = "...my_team_id..."
}