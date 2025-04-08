terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1/m/default"
}

data "cribl-terraform_source" "my_source" {
  id = "CriblMetrics"
}

output "source_details" {
  value = {
    count_total = data.cribl-terraform_source.my_source.count_total
  }
} 
