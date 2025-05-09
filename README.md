# cribl-terraform

Terraform Provider for the *cribl-terraform* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=cribl-terraform&utm_campaign=terraform"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
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
<!-- $toc-max-depth=2 -->
* [cribl-terraform](#cribl-terraform)
  * [🏗 **Welcome to your new Terraform Provider!** 🏗](#welcome-to-your-new-terraform-provider)
  * [Installation](#installation)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Testing the provider locally](#testing-the-provider-locally)
* [Development](#development)
  * [Contributions](#contributions)
* [OpenAPI Schema Status Update Scripts](#openapi-schema-status-update-scripts)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation-1)
  * [Scripts](#scripts)
  * [File Structure](#file-structure)
  * [Example](#example)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    cribl-terraform = {
      source  = "speakeasy/cribl-terraform"
      version = "0.11.69"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [cribl-terraform_destination](docs/resources/destination.md)
* [cribl-terraform_group](docs/resources/group.md)
* [cribl-terraform_pack](docs/resources/pack.md)
* [cribl-terraform_source](docs/resources/source.md)
### Data Sources

* [cribl-terraform_pack](docs/data-sources/pack.md)
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

1. Execute `go build` to construct a binary called `terraform-provider-cribl-terraform`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/speakeasy/cribl-terraform" = "<PATH>"
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

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=cribl-terraform&utm_campaign=terraform)

# OpenAPI Schema Status Update Scripts

This repository contains scripts to update OpenAPI schemas by adding a `status` property to Input and Output schemas.

## Prerequisites

- Python 3.x
- pip (Python package manager)

## Installation

1. Create a virtual environment (recommended):
```bash
python3 -m venv venv
source venv/bin/activate  # On Unix/macOS
# or
.\venv\Scripts\activate  # On Windows
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

## Scripts

### update_schemas.py

This script updates the OpenAPI specification by adding a `status` property to Input and Output schemas. It uses the `overlay.yml` file to determine which schemas to update.

Usage:
```bash
python update_schemas.py
```

The script will:
1. Read the schema names from `overlay.yml`
2. Update the corresponding schemas in `openapi.yml`
3. Add a `status` property with a reference to the `TFStatus` schema

### patch_status.py

This script generates the `overlay.yml` file by scanning the OpenAPI spec and creating patches for all Input and Output schemas.

Usage:
```bash
python patch_status.py
```

## File Structure

- `openapi.yml`: The OpenAPI specification file
- `overlay.yml`: Contains the list of schemas to be updated
- `update_schemas.py`: Script to apply the updates
- `patch_status.py`: Script to generate the overlay file
- `requirements.txt`: Python dependencies

## Example

For each Input and Output schema, the script adds:
```yaml
status:
  $ref: "#/components/schemas/TFStatus"
```
