resource "firehydrant_ticketing_ticket_entity" "my_ticketing_ticketentity" {
  description = "...my_description..."
  priority_id = "...my_priority_id..."
  project_id  = "...my_project_id..."
  related_to  = "...my_related_to..."
  remote_url  = "...my_remote_url..."
  state       = "...my_state..."
  summary     = "...my_summary..."
  tag_list = [
    "..."
  ]
  type = "...my_type..."
}