terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  #server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1"
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
  server_url = "https://app.cribl-playground.cloud"
}

resource "cribl-terraform_pack" "my_pack" {
  id           = "test-pack-id-1"
  group_id     = "default"
  filename = "/Users/kishore/Downloads/cribl-paloalto-xsiam-1.0.0.crbl"
}

# Output the pack details to see the read-only attributes
output "pack_details" {
  value = cribl-terraform_pack.my_pack
} 

/*
# Example of using a data source to fetch an existing pack
data "cribl-terraform_pack" "existing_pack" {
}

# Output the data source details
output "existing_pack_details" {
  value = data.cribl-terraform_pack.existing_pack
}
*/

/*
# Example of importing an existing pack
# First, define a resource block for the pack you want to import
resource "cribl-terraform_pack" "imported_pack" {
  # These values will be overwritten by the import
  id = "HelloPacks"
  description = "This will be replaced after import"
  version = "0.0.0"
  disabled = false
  display_name = "This will be replaced after import"
}

# To import the pack, run the following command:
# terraform import cribl-terraform_pack.imported_pack <pack-id>
# 
# For import only, run:
# terraform import cribl-terraform_pack.imported_pack HelloPacks
# After importing, you can use 'terraform state show' to see the actual values:
# terraform state show cribl-terraform_pack.imported_pack

# Output the imported pack details
output "imported_pack_details" {
  value = cribl-terraform_pack.imported_pack
}

import {
  id = jsonencode({
    id: "HelloPacks"
  })
  to = cribl-terraform_pack.imported_pack
}

*/