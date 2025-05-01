terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  #server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1/m/default"
}

resource "cribl-terraform_source" "my_source" {
  id = "CriblMetrics"
  input_tcp = {
    auth_type = "manual"
    breaker_rulesets = [
      "..."
    ]
    connections = [
      {
        output   = "test"
        pipeline = "test"
      }
    ]
    description         = "test"
    disabled            = false
    enable_header       = false
    enable_proxy_header = false
    environment         = "test"
    host                = "test"
    id                  = "test"
    max_active_cxn      = 5.5
    metadata = [
      {
        name  = "test"
        value = "test"
      }
    ]
    pipeline = "test"
    port     = 47674.49
    pq = {
      commit_frequency = 1.66
      compress         = "none"
      max_buffer_size  = 51.41
      max_file_size    = 1000000
      max_size         = 1000000
      mode             = "always"
      path             = "test"
    }
    pq_enabled = false
    preprocess = {
      args = [
        "..."
      ]
      command  = "test"
      disabled = false
    }
    send_to_routes         = false
    socket_ending_max_wait = 6.18
    socket_idle_timeout    = 0.36
    socket_max_lifespan    = 5.19
    stale_channel_flush_ms = 8063309.13
    streamtags = [
      "..."
    ]
    type = "tcp"
  }
}

output "source_details" {
  value = {
    id = cribl-terraform_source.my_source.id
  }
} 
