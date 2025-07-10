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
- Extracts inline request schemas: Moves inline `requestBody` schemas to `components/schemas` with `$ref` references
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
   - Inline request parameter extraction into request schemas
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
1. Extract any inline request schemas to `components/schemas` with proper `$ref` references
2. Leave existing schemas structurally unchanged
3. Handle any reserved keyword properties
4. Perform minimal cleanup operations

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
- Verify that extracted schemas appear in `components/schemas`
- Confirm inline request bodies now use `$ref` references
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

## Release Process

Our automated release process ensures consistent versioning and reliable publishing of the Terraform provider.

### Overview

The release process consists of three main GitHub Actions workflows:

1. **SDK Generation** (`sdk_generation.yaml`) - Daily automated provider generation
2. **Tag Release** (`tag_release.yaml`) - Creates release tags when provider code changes
3. **Release Terraform Provider** (`release.yaml`) - Publishes the provider to registries

### Automated Release Flow

```mermaid
flowchart TD
    A [Daily: SDK Generation Workflow] --> [Speakeasy Bot Opens PR]
    B --> Merge PR to main
    C --> {PR Contains /internal Changes?}
    C -->|Yes| [Tag Release Workflow Triggered]
    C -->|No| [No release]
    D --> [Extract Version from PR Title]
    E --> [Create Git Tag]
    F --> [Release Workflow Triggered]
    G --> [GoReleaser Publishes Provider]
```

### Workflow Details

#### 1. SDK Generation (`sdk_generation.yaml`)
**Triggers:** Daily at midnight, manual dispatch
**Purpose:** Generate updated provider from latest OpenAPI spec

**Process:**
- Fetches latest OpenAPI spec from developers repo
- Runs normalization and overlay scripts
- Generates provider via Speakeasy
- Speakeasy bot opens PR with changes

#### 2. Tag Release (`tag_release.yaml`)
**Triggers:** Push to main with `/internal/**` changes, manual dispatch
**Purpose:** Create release tags with proper versioning

**Version Sources (in priority order):**
1. **Manual dispatch** with `set_version` input
2. **Speakeasy PR title** - Extracts version from pattern: `"chore: 🐝 Update SDK - Generate X.X.X"`
3. **Workflow fails** if no version found (no auto-increment to prevent wrong versions)

**Process:**
- Detects version from PR title or manual input
- Creates annotated Git tag (e.g., `v0.2.7`)
- Pushes tag to repository

#### 3. Release Terraform Provider (`release.yaml`)
**Triggers:** Git tag creation (`v*` pattern), manual dispatch
**Purpose:** Publish provider to Terraform Registry

**Process:**
- Triggered automatically when tag is created
- Runs GoReleaser to build and publish provider
- Creates GitHub release with assets
- Publishes to Terraform Registry

### Manual Release Process

For manual releases or hotfixes:

1. **Direct tag creation:**
   ```bash
   git tag v0.2.8
   git push origin v0.2.8
   ```

2. **Via workflow dispatch:**
   - Go to Actions → "Tag Release" workflow
   - Click "Run workflow"
   - Set `set_version` to desired version (e.g., `0.2.8`)
   - Run workflow

### Version Management

- **Current version scheme:** `v0.2.X` (semantic versioning)
- **Speakeasy determines versions** based on API changes
- **No auto-increment** to prevent accidental releases
- **Version extraction** from PR titles ensures accuracy

### Troubleshooting

#### Tag Mismatch Errors
If you see errors like "git tag v0.2.6 was not made against commit...", fix with:
```bash
git tag -d v0.2.6
git push origin :refs/tags/v0.2.6
git tag v0.2.6
git push origin v0.2.6
```

#### Missing Version in PR
If the Tag Release workflow fails due to missing version:
- Check PR title follows format: `"chore: 🐝 Update SDK - Generate X.X.X"`
- Use manual dispatch with `set_version` as fallback

#### Release Workflow Fails
- Verify tag exists and points to correct commit
- Check GitHub tokens and secrets are configured
- Ensure GoReleaser configuration is valid

### Testing Release Process

To test the release process safely:

1. **Create test branch:**
   ```bash
   git checkout -b test-release
   ```

2. **Modify tag workflow** to trigger on test branch
3. **Use dry-run mode** to validate without creating actual tags
4. **Test version extraction** with sample commit messages

### Prerequisites

Ensure these secrets are configured in GitHub:
- `GITHUB_TOKEN` (automatic)
- `terraform_gpg_private_key` (for signing)
- `terraform_gpg_passphrase` (for signing)
- `FH_OPS_SSH_KEY` (for accessing developers repo)
- `SPEAKEASY_API_KEY` (for SDK generation)

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

Run `go build -o terraform-provider-firehydrant-v2` so you have a local binary to use
after that you can terraform init, plan, apply, destroy etc to test against staging