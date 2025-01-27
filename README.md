# Firehydrant Terraform Provider V2

Developer-friendly & type-safe Terraform SDK specifically catered to leverage *firehydrant-terraform-sdk* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=firehydrant-terraform-sdk&utm_campaign=terraform"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [Installation](#installation)
* [Available Resources and Data Sources](#available-resources-and-data-sources)
* [Testing the provider locally](#testing-the-provider-locally)
<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    firehydrant = {
      source  = "firehydrant/firehydrant"
      version = "0.1.5"
    }
  }
}

provider "firehydrant" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [firehydrant_checklist_template](docs/resources/checklist_template.md)
* [firehydrant_environment](docs/resources/environment.md)
* [firehydrant_functionality](docs/resources/functionality.md)
* [firehydrant_incident_role](docs/resources/incident_role.md)
* [firehydrant_incident_type](docs/resources/incident_type.md)
* [firehydrant_priority](docs/resources/priority.md)
* [firehydrant_runbook](docs/resources/runbook.md)
* [firehydrant_service](docs/resources/service.md)
* [firehydrant_service_dependency](docs/resources/service_dependency.md)
* [firehydrant_severity](docs/resources/severity.md)
* [firehydrant_status_update_template](docs/resources/status_update_template.md)
* [firehydrant_task_list](docs/resources/task_list.md)
* [firehydrant_team](docs/resources/team.md)
* [firehydrant_webhook](docs/resources/webhook.md)
### Data Sources

* [firehydrant_checklist_template](docs/data-sources/checklist_template.md)
* [firehydrant_checklist_templates](docs/data-sources/checklist_templates.md)
* [firehydrant_environment](docs/data-sources/environment.md)
* [firehydrant_environments](docs/data-sources/environments.md)
* [firehydrant_functionalities](docs/data-sources/functionalities.md)
* [firehydrant_functionality](docs/data-sources/functionality.md)
* [firehydrant_incident_role](docs/data-sources/incident_role.md)
* [firehydrant_incident_type](docs/data-sources/incident_type.md)
* [firehydrant_incident_types](docs/data-sources/incident_types.md)
* [firehydrant_priorities](docs/data-sources/priorities.md)
* [firehydrant_priority](docs/data-sources/priority.md)
* [firehydrant_runbook](docs/data-sources/runbook.md)
* [firehydrant_runbooks](docs/data-sources/runbooks.md)
* [firehydrant_service](docs/data-sources/service.md)
* [firehydrant_service_dependency](docs/data-sources/service_dependency.md)
* [firehydrant_services](docs/data-sources/services.md)
* [firehydrant_severities](docs/data-sources/severities.md)
* [firehydrant_severity](docs/data-sources/severity.md)
* [firehydrant_status_update_template](docs/data-sources/status_update_template.md)
* [firehydrant_status_update_templates](docs/data-sources/status_update_templates.md)
* [firehydrant_task_list](docs/data-sources/task_list.md)
* [firehydrant_task_lists](docs/data-sources/task_lists.md)
* [firehydrant_team](docs/data-sources/team.md)
* [firehydrant_teams](docs/data-sources/teams.md)
* [firehydrant_users](docs/data-sources/users.md)
* [firehydrant_webhook](docs/data-sources/webhook.md)
* [firehydrant_webhooks](docs/data-sources/webhooks.md)
* [firehydrant_webhook_target](docs/data-sources/webhook_target.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-firehydrant`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/firehydrant/firehydrant" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Contributions

While we value open-source contributions to this terraform provider, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=firehydrant-terraform-sdk&utm_campaign=terraform)
# terraform-provider-firehydrant-v2
