resource "firehydrant_task_list" "my_tasklist" {
  description = "...my_description..."
  name        = "...my_name..."
  task_list_items = [
    {
      description = "...my_description..."
      summary     = "...my_summary..."
    }
  ]
}