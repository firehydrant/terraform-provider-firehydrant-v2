---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_current_users Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  CurrentUsers DataSource
---

# firehydrant_current_users (Data Source)

CurrentUsers DataSource

## Example Usage

```terraform
data "firehydrant_current_users" "my_currentusers" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `account_id` (Number)
- `email` (String)
- `id` (String) The ID of this resource.
- `name` (String)
- `organization_id` (String)
- `organization_name` (String)
- `role` (String)
- `source` (String)
