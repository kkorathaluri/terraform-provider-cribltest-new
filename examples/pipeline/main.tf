terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
}

resource "cribl-terraform_pipeline" "my_pipeline" {
  id = "test-pipeline-1"
  group_id = "default"
  conf = {
    description = "Test pipeline for Terraform provider"
    functions = [
      {
        id = "eval"
        description = "Set organizationId test"
        disabled = false
        filter = "true"
        final = false
        group_id = "default"
        conf = {
          fields = {
            organizationId = "tenantId || __forwardedAttrs.__metadata.env.TENANT_ID"
          }
        }
      }
    ]
    output = "events"
    streamtags = ["terraform", "test"]
  }
}

# Output the pipeline details
output "pipeline_details" {
  value = cribl-terraform_pipeline.my_pipeline
}

/*
# Example of using a data source to fetch an existing pipeline
data "cribl-terraform_pipeline" "existing_pipeline" {
  id = "existing-pipeline-id"
}

# Output the data source details
output "existing_pipeline_details" {
  value = data.cribl-terraform_pipeline.existing_pipeline
}
*/

/*
# Example of importing an existing pipeline
resource "cribl-terraform_pipeline" "imported_pipeline" {
  # These values will be overwritten by the import
  id = "existing-pipeline-id"
  group_id = "default"
  conf = {
    description = "This will be replaced after import"
  }
}

# To import the pipeline, run:
# terraform import cribl-terraform_pipeline.imported_pipeline <pipeline-id>

# Output the imported pipeline details
output "imported_pipeline_details" {
  value = cribl-terraform_pipeline.imported_pipeline
}
*/ 