data "firehydrant_ticketing_tickets_entities" "my_ticketing_ticketsentities" {
  assigned_user      = "...my_assigned_user..."
  page               = 0
  per_page           = 10
  state              = "...my_state..."
  tag_match_strategy = "...my_tag_match_strategy..."
  tags               = "...my_tags..."
}