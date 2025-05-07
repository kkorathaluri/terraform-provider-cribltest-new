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
  cloud = {
    provider = "aws"
    region   = "us-west-2"
  }
  config_version         = "1.0.1"
  deploying_worker_count = 1
  description            = "test"
  estimated_ingest_rate  = 1024
  git = {
    commit        = "...my_commit..."
    local_changes = 8.64
    log = [
      {
        author_email = "test@test.com"
        author_name  = "test"
        date         = "2021-01-01"
        hash         = "test"
        message      = "test"
        short        = "test"
      }
    ]
  }
  id                        = "test"
  incompatible_worker_count = 1
  is_fleet                  = false
  is_search                 = false
  name                      = "test"
  on_prem                   = false
  product                   = "stream"
  provisioned               = true
  streamtags = [
    "..."
  ]
  tags                 = "test"
  upgrade_version      = "test"
  worker_count         = 1
  worker_remote_access = false
}