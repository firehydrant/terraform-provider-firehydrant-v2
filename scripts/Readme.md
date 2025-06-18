# OpenAPI to Terraform Provider Pipeline Scripts

This directory contains scripts for processing OpenAPI specifications to generate Terraform providers using Speakeasy.

## Overview

Our pipeline consists of two main phases:
1. **Normalization**: Deep structure alignment between request/response schemas with reference resolution
2. **Overlay Generation**: Creates Speakeasy annotations for Terraform provider generation with granular readonly marking

## Workflow

```mermaid
    A[openapi-raw.json] -->|normalize| B[openapi.json]
    B -->|overlay| C[terraform-overlay.yaml]
    C -->|speakeasy| D[Terraform Provider]
```
Note: we point to the openapi3 spec in our Developers repo to ensure an up-to-date generation

## Scripts

### `normalize`
Normalizes an OpenAPI specification to ensure compatibility with Terraform provider generation through deep structure analysis.

**Usage:**
```bash
go run ./scripts/normalize <input.json> [output.json]
```

**What it does:**
- Identifies entity schemas and their corresponding create/update schemas
- Adds missing required fields from request schemas to entity schemas
- Aligns property structures between request and response schemas
- Converts `additionalProperties` to `properties` for better Terraform compatibility
- Normalizes path parameters (e.g., converts integer IDs to strings)
- Converts enums to speakeasy-compatible x-speakeasy-enums structures

### `overlay`
Generates a Speakeasy overlay file with the necessary annotations for Terraform provider generation.

**Usage:**
```bash
go run ./scripts/overlay <input.json> [manual-mappings.yaml]
```

**What it does:**
- Identifies viable Terraform resources (must have at least CREATE and READ operations)
- Adds `x-speakeasy-entity` annotations to entity schemas
- Maps CRUD operations using `x-speakeasy-entity-operation`
- Detects properties normalized during the normalize phase
- Marks response-only fields with `x-speakeasy-param-readonly` instead of ignoring
- Handles nested readonly properties within manageable fields
- Only uses `x-speakeasy-ignore` for truly incompatible properties
- Applies manual mappings for edge cases

See speakeasy extensions: https://www.speakeasy.com/docs/speakeasy-reference/extensions
## Key Concepts

1. **During Normalization**:
   - Deep structure alignment ensures compatibility
   - Required fields missing from entity schemas are automatically added
   - Added fields are marked as `nullable: true` since they might not be returned by the API

2. **During Overlay Generation**:
   - Required fields are never marked with `x-speakeasy-ignore`
   - If issues persist, we use `x-speakeasy-param-optional` instead
   - Computed fields use `x-speakeasy-param-readonly`

### Entity Identification

Entities are identified by:
- Schema name ending with `Entity`
- Contains an `id` or `slug` property
- Has associated create/update schemas (e.g., `UserEntity` → `create_user`, `update_user`)

### Resource Viability

A resource is considered viable for Terraform if it has:
- At least CREATE and READ operations
- A valid create schema
- Identifiable primary ID parameter
- Sufficient overlap between create and entity schemas after normalization

## Manual Mappings

The `manual-mappings.yaml` file allows for handling edge cases:

```yaml
operations:
  - path: /api/v1/special/resource
    method: post
    action: entity
    value: SpecialResourceEntity
  
  - path: /api/v1/problematic
    method: get
    action: ignore
  
  - path: /api/v1/resource/{resource_id}
    method: get
    action: match
    value: "resource_id:id"
    
properties:
  - schema: SomeEntity
    property: problematic_field
    action: ignore_property
```

## Example

Given this input structure with nested objects and references:
```json
{
  "components": {
    "schemas": {
      "create_user": {
        "properties": {
          "name": { "type": "string" },
          "email": { "type": "string" },
          "settings": {
            "type": "object",
            "properties": {
              "theme": { "type": "string" },
              "notifications": { "type": "boolean" }
            }
          }
        },
        "required": ["name", "email"]
      },
      "UserEntity": {
        "properties": {
          "id": { "type": "string" },
          "name": { "type": "string" },
          "email": { "type": "string" },
          "settings": { "$ref": "#/components/schemas/UserSettings" },
          "created_at": { "type": "string" }
        }
      },
      "UserSettings": {
        "properties": {
          "theme": { "type": "string" },
          "notifications": { "type": "boolean" },
          "last_login": { "type": "string" }  // computed field
        }
      }
    }
  }
}
```

The normalization process will:
1. Detect that `settings` structures are semantically equivalent
2. Align the response structure to match the request structure
3. Preserve computed fields like `last_login`

The overlay generation will:
1. Mark `UserEntity` with `x-speakeasy-entity`
2. Map CRUD operations
3. Mark `created_at` as `x-speakeasy-param-readonly`
4. Mark `settings.last_login` as `x-speakeasy-param-readonly`
5. Keep `settings.theme` and `settings.notifications` manageable

## Script Debugging

Run scripts locally. Use `speakeasy run` to attempt to generate the provider.

### Debugging Normalization
- Look for fields that were added from request schemas
- Verify that complex structures were properly aligned

### Debugging Overlay
- Check for `x-speakeasy-param-readonly` vs `x-speakeasy-ignore` usage
- Verify that normalized properties weren't incorrectly ignored
- Look for nested readonly properties within manageable fields

To view the combined normalized spec + overlay run `speakeasy overlay apply -s {specPath} -o {overlayPath} > {outputPath}` This output will always be in yaml, so name accordingly

Lint the spec with `speakeasy openapi`, select lint and follow prompts

To focus on a sub-section of the spec:
Collect the operationIds that encompass your desired scope, e.g. create_incident,get_incident,delete_incident
Run `speakeasy openapi` and select transform, continue past prompts until you see `transform remove-unused` or `schema to transform` prompt. Enter the operationIds, follow prompts. A smaller version of the spec will be created for testing/debugging.

## Integration

This pipeline is typically run as part of a GitHub Action:

```yaml
- name: Process OpenAPI spec
  run: |
    go run ./scripts/normalize openapi-raw.json openapi.json
    go run ./scripts/overlay openapi.json ./scripts/overlay/manual-mappings.yaml
    mkdir -p .speakeasy
    mv terraform-overlay.yaml .speakeasy/
```

The generated overlay is then used by Speakeasy to create the Terraform provider.

## Provider Local Testing

After `speakeasy run` has successfully completed. 
Add dev.tfrc to directory root (this sometimes needs to be named .terraformrc)
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/firehydrant/firehydrant" = "/Users/jonahstewart/Developer/terraform-provider-firehydrant-v2"
  }
  
  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
This will tell terraform to use your local build instead of downloading the latest published version

Add a test.tf terraform config to the root directory for your chosen resource testing, Can be modeled after the corresponding entity file in docs>resources, which will be generated with `speakeasy run`.

add terraform.tfvars to directory root with:
``` 
api_key = {staging apikey}
```

Run `go build -o terraform-provider-firehydrant` so you have a local binary to use
after that you can terraform init, plan, apply, destroy etc to test against staging