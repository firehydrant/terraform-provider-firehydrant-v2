---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_signals_api_grouping Resource - terraform-provider-firehydrant"
subcategory: ""
description: |-
  SignalsAPIGrouping Resource
---

# firehydrant_signals_api_grouping (Resource)

SignalsAPIGrouping Resource

## Example Usage

```terraform
resource "firehydrant_signals_api_grouping" "my_signals_api_grouping" {
  action_input = {
    fyi = {
      slack_channel_ids = [
        "..."
      ]
    }
    link = false
  }
  reference_alert_time_period = "...my_reference_alert_time_period..."
  strategy = {
    substring = {
      field_name = "...my_field_name..."
      value      = "...my_value..."
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `reference_alert_time_period` (String) How long to group alerts for
- `strategy` (Attributes) The strategy to use for grouping alerts (see [below for nested schema](#nestedatt--strategy))

### Optional

- `action_input` (Attributes) The action to take when grouping alerts (see [below for nested schema](#nestedatt--action_input))

### Read-Only

- `action` (Attributes) (see [below for nested schema](#nestedatt--action))
- `id` (String) The ID of this resource.

<a id="nestedatt--strategy"></a>
### Nested Schema for `strategy`

Optional:

- `substring` (Attributes) The type of strategy to use for grouping alerts (see [below for nested schema](#nestedatt--strategy--substring))

<a id="nestedatt--strategy--substring"></a>
### Nested Schema for `strategy.substring`

Optional:

- `field_name` (String) The field to use for grouping alerts. Not Null
- `value` (String) The value to use for grouping alerts. Not Null



<a id="nestedatt--action_input"></a>
### Nested Schema for `action_input`

Optional:

- `fyi` (Attributes) Send FYI notification (see [below for nested schema](#nestedatt--action_input--fyi))
- `link` (Boolean) Link the alerts and do not notify anyone

<a id="nestedatt--action_input--fyi"></a>
### Nested Schema for `action_input.fyi`

Required:

- `slack_channel_ids` (List of String) The slack channel ids to send the notification to



<a id="nestedatt--action"></a>
### Nested Schema for `action`

Read-Only:

- `fyi` (Attributes) (see [below for nested schema](#nestedatt--action--fyi))
- `link` (Boolean)

<a id="nestedatt--action--fyi"></a>
### Nested Schema for `action.fyi`

Read-Only:

- `slack_channels` (Attributes List) (see [below for nested schema](#nestedatt--action--fyi--slack_channels))

<a id="nestedatt--action--fyi--slack_channels"></a>
### Nested Schema for `action.fyi.slack_channels`

Read-Only:

- `id` (String)
- `name` (String)
- `slack_channel_id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import firehydrant_signals_api_grouping.my_firehydrant_signals_api_grouping ""
```
