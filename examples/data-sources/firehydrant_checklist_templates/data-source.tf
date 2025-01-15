data "firehydrant_checklist_templates" "my_checklisttemplates" {
  page     = 1
  per_page = 10
  query    = "...my_query..."
}