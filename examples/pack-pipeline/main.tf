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

resource "cribl-terraform_pack_pipeline" "my_packpipeline" {
  conf = {
    async_func_timeout = 60
    description        = "Pack pipeline for testing"
    functions = [
      {
        id = "eval"
        filter = "channel == \"ProcessMetrics\" && cid.startsWith(\"service\")"
        conf = {
          add = jsonencode([
            {
              disabled = false
              name = "service_name"
              value = "cid.split(\":\")[1]"
            },
            {
              disabled = false
              name = "process"
              value = "source.match(/\\d+/)? source.match(/\\d+/)[0]: 0"
            }
          ])
        }
        group_id = "test-group"
      },
    ]
    groups = {
      key = {
        description = "Test group"
        disabled    = true
        name        = "test-group"
      }
    }
    output = "test-output"
    streamtags = [
      "test-streamtag"
    ]
  }
  group_id = "default"
  id       = "test-pack-pipeline"
  pack     = "test_ui"
}
