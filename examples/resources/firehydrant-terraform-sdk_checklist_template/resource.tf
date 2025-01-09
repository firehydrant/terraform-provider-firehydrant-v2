resource "firehydrant-terraform-sdk_checklist_template" "my_checklisttemplate" {
  checks = [
    {
      description = "...my_description..."
      name        = "...my_name..."
    }
  ]
  description = "...my_description..."
  name        = "...my_name..."
  team_id     = "...my_team_id..."
}