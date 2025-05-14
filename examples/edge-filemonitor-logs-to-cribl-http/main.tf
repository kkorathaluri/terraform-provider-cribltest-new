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

# Edge Filemonitor Fleet Configuration
resource "cribl-terraform_group" "edge_filemonitor_fleet" {
  config_version = "1.0.1"
  id          = "edge_filemonitor_fleet"
  name        = "edge_filemonitor_fleet"
  description = "Edge Filemonitor Fleet"
  is_fleet    = true
  is_search   = false
  on_prem     = true
  product     = "edge"
  upgrade_version      = "test"
}

# FileMonitor Source Configuration
resource "cribl-terraform_source" "filemonitor_source" {
  id       = "filemonitor_source"
  group_id = cribl-terraform_group.edge_filemonitor_fleet.id
  
  input_file = {
    breaker_rulesets = [
      "json_newline"
    ]
    check_file_mod_time = false
    connections = [
      {
        output   = "cribl_http"
        pipeline = "main"
      }
    ]
    delete_files = true
    depth        = 3
    description  = "FileMonitor Source for Cribl"
    disabled     = false
    filenames = [
      "/opt/cribl/log/cribl.log"
    ]
    force_text                    = true
    hash_len                      = 10
    id                            = "filemonitor_source"
    idle_timeout                  = 60
    include_unidentifiable_binary = true
    interval                      = 60
    max_age_dur                   = "1d"
    metadata = [
      {
        name  = "filemonitor_source"
        value = "filemonitor_source"
      }
    ]
    mode     = "auto"
    path     = "/opt/cribl/log/cribl.log"
    pipeline = "filemonitor_pipeline"
    pq = {
      commit_frequency = 60
      compress         = "none"
      max_buffer_size  = 100
      max_file_size    = "100MB"
      max_size         = "100MB"
      mode             = "smart"
      path             = "/opt/cribl/log/cribl.log"
    }
    pq_enabled             = false
    send_to_routes         = false
    stale_channel_flush_ms = 38000
    status = {
      health = "Red"
      metrics = {
        key = jsonencode("value")
      }
      timestamp          = 8.25
      use_status_from_lb = true
    }
    streamtags = [
      "filemonitor"
    ]
    suppress_missing_path_errors = true
    tail_only                    = true
    type                         = "file"
  }
}

# Cribl HTTP Destination Configuration
resource "cribl-terraform_destination" "cribl_http" {
  id       = "cribl-http-1"
  group_id = cribl-terraform_group.edge_filemonitor_fleet.id

  output_cribl_http = {
    id          = "cribl-http-1"
    type        = "cribl_http"
    description = "Cribl HTTP destination for filemonitor data"
    disabled    = false
    streamtags  = ["filemonitor", "http"]
    compression = "gzip"
    concurrency = 1
    dns_resolve_period_sec = 60
    exclude_self = true
    failed_request_logging_mode = "none"
    flush_period_sec = 2
    load_balance_stats_period_sec = 10
    load_balanced = false
    max_payload_events = 1000
    max_payload_size_kb = 1024
    on_backpressure = "block"
    reject_unauthorized = true
    response_honor_retry_after_header = true
    timeout_sec = 30
    url = "https://cribl-playground.cloud/api/v1/m/default/destination/cribl-http-1"
    use_round_robin_dns = false
  }
}

# Pack Configuration
resource "cribl-terraform_pack" "filemonitor_pack" {
  id       = "filemonitor_pack"
  group_id = cribl-terraform_group.edge_filemonitor_fleet.id
  filename = "EdgeStreamLogs_0.0.1.crbl"  # You'll need to provide the actual pack file
}

# Outputs
output "worker_group_details" {
  value = {
    id   = cribl-terraform_group.edge_filemonitor_fleet.id
    name = cribl-terraform_group.edge_filemonitor_fleet.name
  }
}

output "source_details" {
  value = {
    id   = cribl-terraform_source.filemonitor_source.id
  }
}

output "destination_details" {
  value = {
    id   = cribl-terraform_destination.cribl_http.id
  }
}

output "pack_details" {
  value = {
    id   = cribl-terraform_pack.filemonitor_pack.id
  }
} 