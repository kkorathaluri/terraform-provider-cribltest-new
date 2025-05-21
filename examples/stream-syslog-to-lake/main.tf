terraform {
  required_providers {
    criblio = {
      source = "criblio/criblio"
    }
  }
}

provider "criblio" {
  organization_id = "beautiful-nguyen-y8y4azd"
  workspace_id = "main"
  server_url = "https://app.cribl-playground.cloud"
}

# Worker Group Configuration
resource "criblio_group" "syslog_worker_group" {
  cloud = {
    provider = "aws"
    region   = "us-west-2"
  }
  config_version         = "4.12.0"
  deploying_worker_count = 1
  description            = "Worker group for handling syslog data"
  estimated_ingest_rate  = 1024
  id                     = "syslog-workers"
  is_fleet               = false
  is_search              = false
  name                   = "syslog-workers"
  on_prem                = false
  product                = "stream"
  provisioned            = true
  worker_count           = 1
  worker_remote_access   = false
}

# Syslog Source Configuration
resource "criblio_source" "syslog_source" {
  id       = "syslog-input"
  group_id = criblio_group.syslog_worker_group.id
  
  input_syslog = {
    input_syslog_syslog1 = {
      allow_non_standard_app_name = false
      connections = [
        {
          output   = "cribl-lake-2"
          pipeline = "palo_alto_traffic"
        }
      ]
      description           = "Syslog input source"
      disabled              = false
      enable_load_balancing = true
      enable_proxy_header   = false
      host                  = "0.0.0.0"
      id                    = "syslog-input"
      infer_framing         = false
      keep_fields_list      = []
      max_active_cxn        = 0
      max_buffer_size       = 8
      metadata              = []
      octet_counting        = false
      pipeline              = "palo_alto_traffic"
      pq = {
        commit_frequency = 5
        compress         = "none"
        max_buffer_size  = 50
        max_file_size    = "1 MB"
        max_size         = "100 MB"
        mode             = "smart"
        path             = "$CRIBL_HOME/state/buffers"
      }
      pq_enabled             = true
      send_to_routes         = false
      single_msg_udp_packets = false
      socket_ending_max_wait = 6
      socket_idle_timeout    = 30
      socket_max_lifespan    = 5
      streamtags             = ["syslog", "network"]
      strictly_infer_octet_counting = true
      tcp_port               = 20005
      timestamp_timezone     = "UTC"
      tls = {
        ca_path             = ""
        cert_path           = ""
        certificate_name    = ""
        common_name_regex   = "{}"
        disabled            = true
        max_version         = "TLSv1.3"
        min_version         = "TLSv1.2"
        passphrase          = ""
        priv_key_path       = ""
        reject_unauthorized = true
        request_cert        = false
      }
      type                   = "syslog"
      udp_port               = 20005
      udp_socket_rx_buf_size = 262144
    }
  }
}

# Cribl Lake Destination Configuration
resource "criblio_destination" "cribl_lake" {
  id       = "cribl-lake-2"
  group_id = criblio_group.syslog_worker_group.id

  output_cribl_lake = {
    id          = "cribl-lake-2"
    type        = "cribl_lake"
    description = "Cribl Lake destination for syslog data"
    disabled    = false
    streamtags  = ["syslog", "lake"]
    dest_path   = "default_logs"
    format      = "json"
    compress    = "gzip"
    add_id_to_stage_path = true
    aws_authentication_method = "auto"
    base_file_name = "CriblOut"
    file_name_suffix = "'.gz'"
    max_file_size_mb = 32
    max_open_files = 100
    write_high_water_mark = 64
    on_backpressure = "block"
    deadletter_enabled = false
    on_disk_full_backpressure = "block"
    max_file_open_time_sec = 300
    max_file_idle_time_sec = 30
    verify_permissions = true
    max_closing_files_to_backpressure = 100
    max_concurrent_file_parts = 1
    empty_dir_cleanup_sec = 300
    max_retry_num = 20
  }
}

# Pack Configuration
resource "criblio_pack" "syslog_pack" {
  id       = "syslog-processing"
  group_id = criblio_group.syslog_worker_group.id
  filename = "cribl-palo-alto-networks-source-1.0.0.crbl" 
}

# Outputs
output "worker_group_details" {
  value = {
    id   = criblio_group.syslog_worker_group.id
    name = criblio_group.syslog_worker_group.name
  }
}

output "source_details" {
  value = {
    id   = criblio_source.syslog_source.id
  }
}

output "destination_details" {
  value = {
    id   = criblio_destination.cribl_lake.id
  }
}

output "pack_details" {
  value = {
    id   = criblio_pack.syslog_pack.id
  }
} 