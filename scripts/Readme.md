# OpenAPI to Terraform Provider Pipeline Scripts

This directory contains scripts for processing OpenAPI specifications to generate Terraform providers using Speakeasy.

## Overview

Our pipeline consists of two main phases:
1. **Normalization**: Lightweight schema cleanup and reserved keyword handling
2. **Overlay Generation**: Creates Speakeasy annotations for Terraform provider generation with field mapping and readonly marking

## Workflow

```mermaid
    A[openapi-raw.json] -->|normalize| B[openapi.json]
    B -->|overlay| C[terraform-overlay.yaml]
    C -->|speakeasy| D[Terraform Provider]
```
Note: we point to the openapi3 spec in our Developers repo to ensure an up-to-date generation

## Scripts

### `normalize`
Performs lightweight normalization of an OpenAPI specification to ensure compatibility with Terraform provider generation.

**Usage:**
```bash
go run ./scripts/normalize <input.json> [output.json]
```

**What it does:**
- Replaces Terraform reserved keyword properties with empty objects (only for object/ref types)
- Converts `additionalProperties` to `properties` for better Terraform compatibility
- Normalizes path parameters (e.g., converts integer IDs to strings)
- Converts enums to speakeasy-compatible x-speakeasy-enums structures
- Maintains original schema structures without forcing alignment

### `overlay`
Generates a Speakeasy overlay file with the necessary annotations for Terraform provider generation using field mapping instead of structural changes.

**Usage:**
```bash
go run ./scripts/overlay <input.json> [manual-mappings.yaml]
```

**What it does:**
- Identifies viable Terraform resources (must have at least CREATE and READ operations)
- Adds `x-speakeasy-entity` annotations to entity schemas
- Maps CRUD operations using `x-speakeasy-entity-operation`
- Uses `x-speakeasy-name-override` to map request field names to response field names
- Marks response-only fields with `x-speakeasy-param-readonly: true`
- Handles nested readonly properties within manageable fields
- Only uses `x-speakeasy-ignore` for truly incompatible properties
- Applies manual mappings for edge cases

See speakeasy extensions: https://www.speakeasy.com/docs/speakeasy-reference/extensions

## Key Concepts

1. **During Normalization**:
   - Minimal structural changes to preserve original API design
   - Reserved keyword handling for Terraform compatibility
   - Schema cleanup without forced alignment

2. **During Overlay Generation**:
   - Field mapping using `x-speakeasy-name-override` instead of structural changes
   - Response fields marked as `x-speakeasy-param-readonly: true` when they don't exist in requests
   - Required fields are never marked with `x-speakeasy-ignore`
   - If issues persist, we use `x-speakeasy-param-optional` instead
   - Computed fields use `x-speakeasy-param-readonly: true`

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
- Sufficient field overlap between create and entity schemas (mapped via overlays)

## Field Mapping Strategy

Instead of modifying schemas during normalization, we use Speakeasy annotations to handle field differences:

- **Request fields** that don't match response fields get `x-speakeasy-name-override` to map to the response field name
- **Response fields** that don't exist in requests get `x-speakeasy-param-readonly: true`
- **Computed fields** (like timestamps, IDs) get `x-speakeasy-param-readonly: true`

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

field_mappings:
  - request_schema: create_user
    response_schema: UserEntity
    mappings:
      - request_field: user_name
        response_field: name
      - request_field: user_email
        response_field: email
```

## Example

Given this input structure:
```json
{
  "components": {
    "schemas": {
      "create_user": {
        "properties": {
          "user_name": { "type": "string" },
          "user_email": { "type": "string" },
          "settings": {
            "type": "object",
            "properties": {
              "theme": { "type": "string" },
              "notifications": { "type": "boolean" }
            }
          }
        },
        "required": ["user_name", "user_email"]
      },
      "UserEntity": {
        "properties": {
          "id": { "type": "string" },
          "name": { "type": "string" },
          "email": { "type": "string" },
          "settings": {
            "type": "object", 
            "properties": {
              "theme": { "type": "string" },
              "notifications": { "type": "boolean" },
              "last_login": { "type": "string" }
            }
          },
          "created_at": { "type": "string" }
        }
      }
    }
  }
}
```

The normalization process will:
1. Leave schemas structurally unchanged
2. Handle any reserved keyword properties
3. Perform minimal cleanup operations

The overlay generation will:
1. Mark `UserEntity` with `x-speakeasy-entity`
2. Map CRUD operations
3. Add `x-speakeasy-name-override: "name"` to `create_user.user_name`
4. Add `x-speakeasy-name-override: "email"` to `create_user.user_email`
5. Mark `UserEntity.id` as `x-speakeasy-param-readonly: true`
6. Mark `UserEntity.created_at` as `x-speakeasy-param-readonly: true`
7. Mark `UserEntity.settings.last_login` as `x-speakeasy-param-readonly: true`
8. Keep `settings.theme` and `settings.notifications` manageable in both schemas

## Script Debugging

Run scripts locally. Use `speakeasy run` to attempt to generate the provider.

### Debugging Normalization
- Check for reserved keyword replacements
- Verify minimal structural changes were applied
- Look for enum and parameter normalizations

### Debugging Overlay
- Check for `x-speakeasy-name-override` mappings on request fields
- Verify `x-speakeasy-param-readonly: true` usage on response-only fields
- Look for nested readonly properties within manageable fields
- Ensure required fields weren't incorrectly ignored

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