# terraform-provider-firehydrant-v2

Developer-friendly & type-safe Terraform SDK specifically catered to leverage *firehydrant-terraform-sdk* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=firehydrant-terraform-sdk&utm_campaign=terraform"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new Terraform Provider!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Add resources and datasources to your SDK by [annotating your OAS](https://www.speakeasy.com/docs/customize-terraform/terraform-extensions#map-api-entities-to-terraform-resources)
- [ ] ♻️ Refine your terraform provider quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your terraform provider to hashicorp registry by [configuring automatic publishing](https://www.speakeasy.com/docs/terraform-publishing)
- [ ] ✨ When ready to productionize, delete this section from the README

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
    firehydrant-terraform-sdk = {
      source  = "speakeasy/firehydrant-terraform-sdk"
      version = "0.1.1"
    }
  }
}

provider "firehydrant-terraform-sdk" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [firehydrant-terraform-sdk_checklist_template](docs/resources/checklist_template.md)
* [firehydrant-terraform-sdk_environment](docs/resources/environment.md)
* [firehydrant-terraform-sdk_functionality](docs/resources/functionality.md)
* [firehydrant-terraform-sdk_incident_role](docs/resources/incident_role.md)
* [firehydrant-terraform-sdk_incident_type](docs/resources/incident_type.md)
* [firehydrant-terraform-sdk_priority](docs/resources/priority.md)
* [firehydrant-terraform-sdk_runbook](docs/resources/runbook.md)
* [firehydrant-terraform-sdk_service](docs/resources/service.md)
* [firehydrant-terraform-sdk_service_dependency](docs/resources/service_dependency.md)
* [firehydrant-terraform-sdk_severity](docs/resources/severity.md)
* [firehydrant-terraform-sdk_status_update_template](docs/resources/status_update_template.md)
* [firehydrant-terraform-sdk_task_list](docs/resources/task_list.md)
* [firehydrant-terraform-sdk_team](docs/resources/team.md)
* [firehydrant-terraform-sdk_webhook](docs/resources/webhook.md)
### Data Sources

* [firehydrant-terraform-sdk_checklist_template](docs/data-sources/checklist_template.md)
* [firehydrant-terraform-sdk_checklist_templates](docs/data-sources/checklist_templates.md)
* [firehydrant-terraform-sdk_environment](docs/data-sources/environment.md)
* [firehydrant-terraform-sdk_environments](docs/data-sources/environments.md)
* [firehydrant-terraform-sdk_functionalities](docs/data-sources/functionalities.md)
* [firehydrant-terraform-sdk_functionality](docs/data-sources/functionality.md)
* [firehydrant-terraform-sdk_incident_role](docs/data-sources/incident_role.md)
* [firehydrant-terraform-sdk_incident_type](docs/data-sources/incident_type.md)
* [firehydrant-terraform-sdk_incident_types](docs/data-sources/incident_types.md)
* [firehydrant-terraform-sdk_priorities](docs/data-sources/priorities.md)
* [firehydrant-terraform-sdk_priority](docs/data-sources/priority.md)
* [firehydrant-terraform-sdk_runbook](docs/data-sources/runbook.md)
* [firehydrant-terraform-sdk_runbooks](docs/data-sources/runbooks.md)
* [firehydrant-terraform-sdk_service](docs/data-sources/service.md)
* [firehydrant-terraform-sdk_service_dependency](docs/data-sources/service_dependency.md)
* [firehydrant-terraform-sdk_services](docs/data-sources/services.md)
* [firehydrant-terraform-sdk_severities](docs/data-sources/severities.md)
* [firehydrant-terraform-sdk_severity](docs/data-sources/severity.md)
* [firehydrant-terraform-sdk_status_update_template](docs/data-sources/status_update_template.md)
* [firehydrant-terraform-sdk_status_update_templates](docs/data-sources/status_update_templates.md)
* [firehydrant-terraform-sdk_task_list](docs/data-sources/task_list.md)
* [firehydrant-terraform-sdk_task_lists](docs/data-sources/task_lists.md)
* [firehydrant-terraform-sdk_team](docs/data-sources/team.md)
* [firehydrant-terraform-sdk_teams](docs/data-sources/teams.md)
* [firehydrant-terraform-sdk_users](docs/data-sources/users.md)
* [firehydrant-terraform-sdk_webhook](docs/data-sources/webhook.md)
* [firehydrant-terraform-sdk_webhooks](docs/data-sources/webhooks.md)
* [firehydrant-terraform-sdk_webhook_target](docs/data-sources/webhook_target.md)
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

1. Execute `go build` to construct a binary called `terraform-provider-firehydrant-terraform-sdk`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/speakeasy/firehydrant-terraform-sdk" = "<PATH>"
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
