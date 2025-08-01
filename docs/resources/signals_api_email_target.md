---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_signals_api_email_target Resource - terraform-provider-firehydrant"
subcategory: ""
description: |-
  SignalsAPIEmailTarget Resource
---

# firehydrant_signals_api_email_target (Resource)

SignalsAPIEmailTarget Resource

## Example Usage

```terraform
resource "firehydrant_signals_api_email_target" "my_signals_api_emailtarget" {
  allowed_senders = [
    "..."
  ]
  description            = "...my_description..."
  level_cel              = "...my_level_cel..."
  name                   = "...my_name..."
  rule_matching_strategy = "...my_rule_matching_strategy..."
  rules = [
    "..."
  ]
  slug       = "...my_slug..."
  status_cel = "...my_status_cel..."
  target_input = {
    id   = "...my_id..."
    type = "...my_type..."
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The email target's name.

### Optional

- `allowed_senders` (List of String) A list of email addresses that are allowed to send events to the target. Must be exact match.
- `description` (String) A detailed description of the email target.
- `level_cel` (String) The CEL expression that defines the level of an incoming email that is sent to the target.
- `rule_matching_strategy` (String) Whether or not all rules must match, or if only one rule must match.
- `rules` (List of String) A list of CEL expressions that should be evaluated and matched to determine if the target should be notified.
- `slug` (String) The email address that will be listening to events.
- `status_cel` (String) The CEL expression that defines the status of an incoming email that is sent to the target.
- `target_input` (Attributes) The target that the email target will notify. This object must contain a `type`
field that specifies the type of target and an `id` field that specifies the ID of
the target. The `type` field must be one of "escalation_policy", "on_call_schedule",
"team", "user", or "slack_channel". (see [below for nested schema](#nestedatt--target_input))

### Read-Only

- `created_at` (String)
- `created_by` (Attributes) (see [below for nested schema](#nestedatt--created_by))
- `email` (String)
- `id` (String) The ID of this resource.
- `target` (Attributes) (see [below for nested schema](#nestedatt--target))
- `team_id` (String) The team ID that the email target belongs to, if applicable
- `updated_at` (String)

<a id="nestedatt--target_input"></a>
### Nested Schema for `target_input`

Required:

- `id` (String) The ID of the target that the inbound email will notify when matched.
- `type` (String) The type of target that the inbound email will notify when matched.


<a id="nestedatt--created_by"></a>
### Nested Schema for `created_by`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)
- `source` (String)


<a id="nestedatt--target"></a>
### Nested Schema for `target`

Read-Only:

- `id` (String)
- `is_pageable` (Boolean)
- `name` (String)
- `team_id` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import firehydrant_signals_api_email_target.my_firehydrant_signals_api_email_target ""
```
