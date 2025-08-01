---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_alerts Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  Alerts DataSource
---

# firehydrant_alerts (Data Source)

Alerts DataSource

## Example Usage

```terraform
data "firehydrant_alerts" "my_alerts" {
  alert_id = "...my_alert_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alert_id` (String)

### Read-Only

- `child_alerts` (Attributes List) (see [below for nested schema](#nestedatt--child_alerts))
- `conversations` (Attributes List) (see [below for nested schema](#nestedatt--conversations))
- `description` (String)
- `duration_iso8601` (String)
- `duration_ms` (Number)
- `ends_at` (String)
- `environments` (Attributes List) (see [below for nested schema](#nestedatt--environments))
- `events` (Attributes List) (see [below for nested schema](#nestedatt--events))
- `id` (String) The ID of this resource.
- `incidents` (Attributes List) (see [below for nested schema](#nestedatt--incidents))
- `integration_name` (String)
- `is_expired` (Boolean)
- `is_noise` (Boolean)
- `labels` (Attributes) Arbitrary key:value pairs of labels. (see [below for nested schema](#nestedatt--labels))
- `parent_alerts` (Attributes List) (see [below for nested schema](#nestedatt--parent_alerts))
- `position` (Number)
- `priority` (String)
- `remote_id` (String)
- `remote_url` (String)
- `services` (Attributes List) (see [below for nested schema](#nestedatt--services))
- `signal_id` (String)
- `signal_rule` (Attributes) Signals_API_Rule model (see [below for nested schema](#nestedatt--signal_rule))
- `signal_target` (Attributes) (see [below for nested schema](#nestedatt--signal_target))
- `source_icon` (String)
- `starts_at` (String)
- `status` (String)
- `summary` (String)
- `tags` (List of String)
- `team_id` (String)
- `team_name` (String)

<a id="nestedatt--child_alerts"></a>
### Nested Schema for `child_alerts`

Read-Only:

- `description` (String)
- `ends_at` (String)
- `id` (String)
- `labels` (Attributes) Arbitrary key:value pairs of labels. (see [below for nested schema](#nestedatt--child_alerts--labels))
- `remote_id` (String)
- `remote_url` (String)
- `signal_id` (String)
- `signal_rule` (Attributes) Signals_API_Rule model (see [below for nested schema](#nestedatt--child_alerts--signal_rule))
- `starts_at` (String)
- `status` (String)
- `summary` (String)
- `tags` (List of String)

<a id="nestedatt--child_alerts--labels"></a>
### Nested Schema for `child_alerts.labels`


<a id="nestedatt--child_alerts--signal_rule"></a>
### Nested Schema for `child_alerts.signal_rule`

Read-Only:

- `create_incident_condition_when` (String)
- `created_at` (String)
- `created_by` (Attributes) (see [below for nested schema](#nestedatt--child_alerts--signal_rule--created_by))
- `deduplication_expiry` (String) Duration for deduplicating similar alerts (ISO8601 duration format e.g., 'PT30M', 'PT2H', 'P1D')
- `expression` (String)
- `id` (String)
- `incident_type` (Attributes) (see [below for nested schema](#nestedatt--child_alerts--signal_rule--incident_type))
- `name` (String)
- `notification_priority_override` (String)
- `target` (Attributes) (see [below for nested schema](#nestedatt--child_alerts--signal_rule--target))
- `team_id` (String)
- `updated_at` (String)

<a id="nestedatt--child_alerts--signal_rule--created_by"></a>
### Nested Schema for `child_alerts.signal_rule.created_by`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)
- `source` (String)


<a id="nestedatt--child_alerts--signal_rule--incident_type"></a>
### Nested Schema for `child_alerts.signal_rule.incident_type`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--child_alerts--signal_rule--target"></a>
### Nested Schema for `child_alerts.signal_rule.target`

Read-Only:

- `id` (String)
- `is_pageable` (Boolean)
- `name` (String)
- `team_id` (String)
- `type` (String)




<a id="nestedatt--conversations"></a>
### Nested Schema for `conversations`

Read-Only:

- `channel` (Attributes) (see [below for nested schema](#nestedatt--conversations--channel))
- `comments_url` (String)
- `field` (String)
- `id` (String)
- `resource_class` (String)
- `resource_id` (String)

<a id="nestedatt--conversations--channel"></a>
### Nested Schema for `conversations.channel`

Read-Only:

- `name` (String)



<a id="nestedatt--environments"></a>
### Nested Schema for `environments`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--events"></a>
### Nested Schema for `events`

Read-Only:

- `created_at` (String)
- `data` (Attributes) The event's payload (see [below for nested schema](#nestedatt--events--data))
- `id` (String)
- `type` (String)

<a id="nestedatt--events--data"></a>
### Nested Schema for `events.data`



<a id="nestedatt--incidents"></a>
### Nested Schema for `incidents`

Read-Only:

- `id` (String)
- `name` (String)
- `number` (Number)


<a id="nestedatt--labels"></a>
### Nested Schema for `labels`


<a id="nestedatt--parent_alerts"></a>
### Nested Schema for `parent_alerts`

Read-Only:

- `description` (String)
- `ends_at` (String)
- `id` (String)
- `labels` (Attributes) Arbitrary key:value pairs of labels. (see [below for nested schema](#nestedatt--parent_alerts--labels))
- `remote_id` (String)
- `remote_url` (String)
- `signal_id` (String)
- `signal_rule` (Attributes) Signals_API_Rule model (see [below for nested schema](#nestedatt--parent_alerts--signal_rule))
- `starts_at` (String)
- `status` (String)
- `summary` (String)
- `tags` (List of String)

<a id="nestedatt--parent_alerts--labels"></a>
### Nested Schema for `parent_alerts.labels`


<a id="nestedatt--parent_alerts--signal_rule"></a>
### Nested Schema for `parent_alerts.signal_rule`

Read-Only:

- `create_incident_condition_when` (String)
- `created_at` (String)
- `created_by` (Attributes) (see [below for nested schema](#nestedatt--parent_alerts--signal_rule--created_by))
- `deduplication_expiry` (String) Duration for deduplicating similar alerts (ISO8601 duration format e.g., 'PT30M', 'PT2H', 'P1D')
- `expression` (String)
- `id` (String)
- `incident_type` (Attributes) (see [below for nested schema](#nestedatt--parent_alerts--signal_rule--incident_type))
- `name` (String)
- `notification_priority_override` (String)
- `target` (Attributes) (see [below for nested schema](#nestedatt--parent_alerts--signal_rule--target))
- `team_id` (String)
- `updated_at` (String)

<a id="nestedatt--parent_alerts--signal_rule--created_by"></a>
### Nested Schema for `parent_alerts.signal_rule.created_by`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)
- `source` (String)


<a id="nestedatt--parent_alerts--signal_rule--incident_type"></a>
### Nested Schema for `parent_alerts.signal_rule.incident_type`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--parent_alerts--signal_rule--target"></a>
### Nested Schema for `parent_alerts.signal_rule.target`

Read-Only:

- `id` (String)
- `is_pageable` (Boolean)
- `name` (String)
- `team_id` (String)
- `type` (String)




<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--signal_rule"></a>
### Nested Schema for `signal_rule`

Read-Only:

- `create_incident_condition_when` (String)
- `created_at` (String)
- `created_by` (Attributes) (see [below for nested schema](#nestedatt--signal_rule--created_by))
- `deduplication_expiry` (String) Duration for deduplicating similar alerts (ISO8601 duration format e.g., 'PT30M', 'PT2H', 'P1D')
- `expression` (String)
- `id` (String)
- `incident_type` (Attributes) (see [below for nested schema](#nestedatt--signal_rule--incident_type))
- `name` (String)
- `notification_priority_override` (String)
- `target` (Attributes) (see [below for nested schema](#nestedatt--signal_rule--target))
- `team_id` (String)
- `updated_at` (String)

<a id="nestedatt--signal_rule--created_by"></a>
### Nested Schema for `signal_rule.created_by`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)
- `source` (String)


<a id="nestedatt--signal_rule--incident_type"></a>
### Nested Schema for `signal_rule.incident_type`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--signal_rule--target"></a>
### Nested Schema for `signal_rule.target`

Read-Only:

- `id` (String)
- `is_pageable` (Boolean)
- `name` (String)
- `team_id` (String)
- `type` (String)



<a id="nestedatt--signal_target"></a>
### Nested Schema for `signal_target`

Read-Only:

- `id` (String)
- `is_pageable` (Boolean)
- `name` (String)
- `team_id` (String)
- `type` (String)
