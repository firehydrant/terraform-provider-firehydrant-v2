---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_ticketing_projects_project_list_items Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  TicketingProjectsProjectListItems DataSource
---

# firehydrant_ticketing_projects_project_list_items (Data Source)

TicketingProjectsProjectListItems DataSource

## Example Usage

```terraform
data "firehydrant_ticketing_projects_project_list_items" "my_ticketing_projects_projectlistitems" {
  configured_projects   = false
  connection_ids        = "...my_connection_ids..."
  page                  = 5
  per_page              = 7
  providers             = "...my_providers..."
  query                 = "...my_query..."
  supports_ticket_types = "...my_supports_ticket_types..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `configured_projects` (Boolean)
- `connection_ids` (String)
- `page` (Number)
- `per_page` (Number)
- `providers` (String)
- `query` (String)
- `supports_ticket_types` (String)

### Read-Only

- `attribute` (String)
- `connection_id` (String)
- `connection_slug` (String)
- `connection_type` (String)
- `details` (Attributes) A config object containing details about the project config. Can be one of: Ticketing::JiraCloud::ProjectConfig, Ticketing::JiraOnprem::ProjectConfig, or Ticketing::Shortcut::ProjectConfig (see [below for nested schema](#nestedatt--details))
- `external_field` (String)
- `id` (String) The ID of this resource.
- `logic` (Map of String) An unstructured object of key/value pairs describing the logic for applying the rule.
- `name` (String)
- `presentation` (String)
- `strategy` (String)
- `ticketing_project_id` (String)
- `ticketing_project_name` (String)
- `type` (String)
- `updated_at` (String)
- `user_data` (Map of String)
- `value` (String)

<a id="nestedatt--details"></a>
### Nested Schema for `details`
