---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_severity_matrix_condition Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  SeverityMatrixCondition DataSource
---

# firehydrant_severity_matrix_condition (Data Source)

SeverityMatrixCondition DataSource

## Example Usage

```terraform
data "firehydrant_severity_matrix_condition" "my_severitymatrix_condition" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String)
- `position` (Number) Position is used to determine ordering of conditions in API responses and dropdowns. The condition with the lowest position (typically 0) will be considered the Default Condition
