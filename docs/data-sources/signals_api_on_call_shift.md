---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_signals_api_on_call_shift Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  SignalsAPIOnCallShift DataSource
---

# firehydrant_signals_api_on_call_shift (Data Source)

SignalsAPIOnCallShift DataSource

## Example Usage

```terraform
data "firehydrant_signals_api_on_call_shift" "my_signals_api_oncallshift" {
  id          = "...my_id..."
  schedule_id = "...my_schedule_id..."
  team_id     = "...my_team_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `schedule_id` (String)
- `team_id` (String)

### Read-Only

- `color` (String)
- `coverage_request` (String)
- `end_time` (String)
- `id` (String) The ID of this resource.
- `on_call_rotation` (Attributes) (see [below for nested schema](#nestedatt--on_call_rotation))
- `on_call_schedule` (Attributes) (see [below for nested schema](#nestedatt--on_call_schedule))
- `start_time` (String)
- `team` (Attributes) (see [below for nested schema](#nestedatt--team))
- `time_zone` (String)
- `user` (Attributes) (see [below for nested schema](#nestedatt--user))

<a id="nestedatt--on_call_rotation"></a>
### Nested Schema for `on_call_rotation`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--on_call_schedule"></a>
### Nested Schema for `on_call_schedule`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--team"></a>
### Nested Schema for `team`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--user"></a>
### Nested Schema for `user`

Read-Only:

- `id` (String)
- `name` (String)
