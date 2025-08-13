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

FireHydrant API: The FireHydrant API is based around REST. It uses Bearer token authentication and returns JSON responses. You can use the FireHydrant API to configure integrations, define incidents, and set up webhooks--anything you can do on the FireHydrant UI.

* [Dig into our API endpoints](https://developers.firehydrant.io/docs/api)
* [View your bot users](https://app.firehydrant.io/organizations/bots)

## Base API endpoint

[https://api.firehydrant.io/v1](https://api.firehydrant.io/v1)

## Current version

v1

## Authentication

All requests to the FireHydrant API require an `Authorization` header with the value set to `Bearer {token}`. FireHydrant supports bot tokens to act on behalf of a computer instead of a user's account. This prevents integrations from breaking when people leave your organization or their token is revoked. See the Bot tokens section (below) for more information on this.

An example of a header to authenticate against FireHydrant would look like:

```
Authorization: Bearer fhb-thisismytoken
```

## Bot tokens

To access the FireHydrant API, you must authenticate with a bot token. (You must have owner permissions on your organization to see bot tokens.) Bot users allow you to interact with the FireHydrant API by using token-based authentication. To create bot tokens, log in to your organization and refer to the **Bot users** [page](https://app.firehydrant.io/organizations/bots).

Bot tokens enable you to create a bot that has no ties to any user. Normally, all actions associated with an API token are associated with the user who created it. Bot tokens attribute all actions to the bot user itself. This way, all data associated with the token actions can be performed against the FireHydrant API without a user.

Every request to the API is authenticated unless specified otherwise.

### Rate Limiting

Currently, requests made with bot tokens are rate limited on a per-account level. If your account has multiple bot token then the rate limit is shared across all of them. As of February 7th, 2023, the rate limit is at least 50 requests per account every 10 seconds, or 300 requests per minute.

Rate limited responses will be served with a `429` status code and a JSON body of:

```json
{"error": "rate limit exceeded"}
```
and headers of:
```
"RateLimit-Limit" -> the maximum number of requests in the rate limit pool
"Retry-After" -> the number of seconds to wait before trying again
```

## How lists are returned

API lists are returned as arrays. A paginated entity in FireHydrant will return two top-level keys in the response object: a data key and a pagination key.

### Paginated requests

The `data` key is returned as an array. Each item in the array includes all of the entity data specified in the API endpoint. (The per-page default for the array is 20 items.)

Pagination is the second key (`pagination`) returned in the overall response body. It includes medtadata around the current page, total count of items, and options to go to the next and previous page. All of the specifications returned in the pagination object are available as URL parameters. So if you want to specify, for example, going to the second page of a response, you can send a request to the same endpoint but pass the URL parameter **page=2**.

For example, you might request **https://api.firehydrant.io/v1/environments/** to retrieve environments data. The JSON returned contains the above-mentioned data section and pagination section. The data section includes various details about an incident, such as the environment name, description, and when it was created.

```
{
  "data": [
    {
      "id": "f8125cf4-b3a7-4f88-b5ab-57a60b9ed89b",
      "name": "Production - GCP",
      "description": "",
      "created_at": "2021-02-17T20:02:10.679Z"
    },
    {
      "id": "a69f1f58-af77-4708-802d-7e73c0bf261c",
      "name": "Staging",
      "description": "",
      "created_at": "2021-04-16T13:41:59.418Z"
    }
  ],
  "pagination": {
    "count": 2,
    "page": 1,
    "items": 2,
    "pages": 1,
    "last": 1,
    "prev": null,
    "next": null
  }
}
```

To request the second page, you'd request the same endpoint with the additional query parameter of `page` in the URL:

```
GET https://api.firehydrant.io/v1/environments?page=2
```

If you need to modify the number of records coming back from FireHydrant, you can use the `per_page` parameter (max is 200):

```
GET https://api.firehydrant.io/v1/environments?per_page=50
```
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Firehydrant Terraform Provider V2](#firehydrant-terraform-provider-v2)
  * [Base API endpoint](#base-api-endpoint)
  * [Current version](#current-version)
  * [Authentication](#authentication)
  * [Bot tokens](#bot-tokens)
  * [How lists are returned](#how-lists-are-returned)
  * [Installation](#installation)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Testing the provider locally](#testing-the-provider-locally)
* [Development](#development)
  * [Contributions](#contributions)
* [terraform-provider-firehydrant-v2](#terraform-provider-firehydrant-v2)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    firehydrant = {
      source  = "firehydrant/firehydrant"
      version = "0.6.0"
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

* [firehydrant_functionality](docs/resources/functionality.md)
### Data Sources

* [firehydrant_functionality](docs/data-sources/functionality.md)
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
