terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  #server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1/"
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
}

resource "cribl-terraform_group" "my_group" {
  config_version         = "1.0.1"
  deploying_worker_count = 1
  description            = "test"
  id                        = "edge-test"
  incompatible_worker_count = 1
  is_fleet                  = true
  is_search                 = false
  name                      = "edge-test"
  product                   = "edge"
  tags                 = "test"
  upgrade_version      = "test"
}