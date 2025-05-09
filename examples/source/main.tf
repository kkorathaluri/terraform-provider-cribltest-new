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
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
}

resource "cribl-terraform_source" "my_source_1" {
  id = "test"
  group_id = "default"
  input_tcp = {
    auth_type = "manual"
    breaker_rulesets = [
      "..."
    ]
    connections = [
      {
        output   = "test1"
        pipeline = "HelloPacks"
      }
    ]
    description         = "test"
    disabled            = false
    enable_header       = false
    enable_proxy_header = false
    host                = "0.0.0.0"
    id                  = "test"
    max_active_cxn      = 0
    metadata = [
      {
        name  = "test1"
        value = "test1"
      }
    ]
    pipeline = "test"
    port     = 20001
   
    send_to_routes         = false
    socket_ending_max_wait = 6
    socket_idle_timeout    = 30
    socket_max_lifespan    = 5
    stale_channel_flush_ms = 30000
    streamtags = [
      "test"
    ]
    type = "tcp"
  }
}

output "source_details" {
  value = {
    id = cribl-terraform_source.my_source_1.id
  }
} 


resource "cribl-terraform_source" "my_source_2" {
  id = "test_2"
  group_id = "default"
  input_cribl_http = {
    activity_log_sample_rate = 100
    auth_tokens = [
      "token1",
      "token2"
    ]
    capture_headers = false
    connections = [
      {
        output   = "default"
        pipeline = "main"
      }
    ]
    description             = "HTTP Source for Cribl"
    disabled                = false
    enable_health_check     = true
    enable_proxy_header     = false
    host                    = "0.0.0.0"
    id                      = "cribl_http_source"
    keep_alive_timeout      = 300
    max_active_req          = 1000
    max_requests_per_socket = 4
    metadata = [
      {
        name  = "source_type"
        value = "http"
      },
      {
        name  = "environment"
        value = "production"
      }
    ]
    pipeline = "main"
    port     = 10081
    pq = {
      commit_frequency = 5
      compress         = "gzip"
      max_buffer_size  = 48
      max_file_size    = "1 MB"
      max_size         = "100 MB"
      mode             = "smart"
      path             = "$CRIBL_HOME/state/buffers"
    }
    pq_enabled      = true
    request_timeout = 30
    send_to_routes  = true
    socket_timeout  = 10
    streamtags = [
      "http",
      "cribl",
      "source"
    ]
    tls = {
      ca_path             = "/opt/cribl/certs/ca.pem"
      cert_path           = "/opt/cribl/certs/cert.pem"
      certificate_name    = "cribl_cert"
      disabled            = true
      max_version         = "TLSv1.3"
      min_version         = "TLSv1.2"
      passphrase          = "cribl_passphrase"
      priv_key_path       = "/opt/cribl/certs/key.pem"
      reject_unauthorized = true
      request_cert        = false
    }
    type = "cribl_http"
  }
}