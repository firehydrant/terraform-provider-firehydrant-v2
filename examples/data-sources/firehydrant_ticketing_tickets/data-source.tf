data "firehydrant_ticketing_tickets" "my_ticketing_tickets" {
  assigned_user      = "...my_assigned_user..."
  page               = 3
  per_page           = 5
  state              = "...my_state..."
  tag_match_strategy = "...my_tag_match_strategy..."
  tags               = "...my_tags..."
}