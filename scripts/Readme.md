# OpenAPI to Terraform Provider Pipeline Scripts

This directory contains scripts for processing OpenAPI specifications to generate Terraform providers using Speakeasy.

## Overview

Our pipeline consists of two main phases:
1. **Normalization**: Aligns request/response schemas and convert enums to ensure compatibility
2. **Overlay Generation**: Creates Speakeasy annotations for Terraform provider generation

## Workflow

```mermaid
    A[openapi-raw.json] -->|normalize| B[openapi.json]
    B -->|overlay| C[terraform-overlay.yaml]
    C -->|speakeasy| D[Terraform Provider]
```
Note: we point to the openapi3 spec in our Developers repo to ensure an up-to-date generation

## Scripts

### `normalize`
Normalizes an OpenAPI specification to ensure compatibility with Terraform provider generation.

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
- Converts enums to speakeasy-compatible x-speaky-enums structures

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
- Handles property inconsistencies:
  - Required fields: Uses `x-speakeasy-param-optional` to preserve functionality
  - Optional fields: Uses `x-speakeasy-ignore` for incompatible properties
- Applies manual mappings for edge cases

See speakeasy extensions: https://www.speakeasy.com/docs/speakeasy-reference/extensions
## Key Concepts

### Required Field Handling

Our approach ensures that required fields are never ignored in Terraform:

1. **During Normalization**:
   - Required fields missing from entity schemas are automatically added
   - Structural differences are reconciled (e.g., copying nested properties)
   - Added fields are marked as `nullable: true` since they might not be returned by the API

2. **During Overlay Generation**:
   - Required fields are never marked with `x-speakeasy-ignore`
   - If issues persist, we use `x-speakeasy-param-optional` instead
   - Write-only fields can be marked with `x-speakeasy-param-readonly`

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
- Sufficient overlap between create and entity schemas

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
    
  - schema: SomeEntity
    property: problematic_field
    action: ignore_property
```

## Example

Given this input structure:
```json
{
  "components": {
    "schemas": {
      "create_user": {
        "properties": {
          "name": { "type": "string" },
          "email": { "type": "string" },
          "password": { "type": "string" }
        },
        "required": ["name", "email", "password"]
      },
      "UserEntity": {
        "properties": {
          "id": { "type": "string" },
          "name": { "type": "string" },
          "email": { "type": "string" }
          // Note: password is missing (write-only)
        }
      }
    }
  }
}
```

The normalization process will:
1. Add `password` field to `UserEntity` (as nullable)
2. Ensure all structures align

The overlay generation will:
1. Mark `UserEntity` with `x-speakeasy-entity`
2. Map CRUD operations
3. Mark `password` as `x-speakeasy-param-readonly` in the entity

## Debugging

Run scripts locally. Use `speakeasy run` to attempt to generate the provider.
If issues persist, compare normalization script output and generated overlay.
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