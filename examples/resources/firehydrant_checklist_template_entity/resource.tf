resource "firehydrant_checklist_template_entity" "my_checklisttemplateentity" {
  connected_services_input = [
    {
      id = "...my_id..."
    }
  ]
  description = "...my_description..."
  name        = "...my_name..."
  team_id     = "...my_team_id..."
}