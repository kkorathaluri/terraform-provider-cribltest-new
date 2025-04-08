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


resource "cribl-terraform_destination" "my_destination" {
  id = "test_4"
  output_cribl_http = {
    compression            = "gzip"
    concurrency            = 11.56
    description            = "...my_description..."
    dns_resolve_period_sec = 5488.51
    environment            = "...my_environment..."
    exclude_fields = [
      "..."
    ]
    exclude_self = true
    extra_http_headers = [
      {
        name  = "...my_name..."
        value = "...my_value..."
      }
    ]
    failed_request_logging_mode   = "none"
    flush_period_sec              = 2.35
    id                            = "test_4"
    load_balance_stats_period_sec = 12.74
    load_balanced                 = false
    max_payload_events            = 6.83
    max_payload_size_kb           = 2827.75
    on_backpressure               = "block"
    pipeline                      = "...my_pipeline..."
    pq_compress                   = "none"
    pq_controls = {
      # ...
    }
    pq_max_file_size                  = 10000000
    pq_max_size                       = 10000000
    pq_mode                           = "always"
    pq_on_backpressure                = "drop"
    pq_path                           = "...my_pq_path..."
    reject_unauthorized               = true
    response_honor_retry_after_header = false
    response_retry_settings = [
      {
        backoff_rate    = 18.82
        http_status     = 531.39
        initial_backoff = 117330.79
        max_backoff     = 52596.23
      }
    ]
    safe_headers = [
      "..."
    ]
    streamtags = [
      "..."
    ]
    system_fields = [
      "..."
    ]
    timeout_retry_settings = {
      backoff_rate    = 19.69
      initial_backoff = 173747.34
      max_backoff     = 163675.23
      timeout_retry   = true
    }
    timeout_sec = 8157421429180492
    tls = {
      ca_path             = "...my_ca_path..."
      cert_path           = "...my_cert_path..."
      certificate_name    = "...my_certificate_name..."
      disabled            = false
      max_version         = "TLSv1.3"
      min_version         = "TLSv1.3"
      passphrase          = "...my_passphrase..."
      priv_key_path       = "...my_priv_key_path..."
      reject_unauthorized = false
      servername          = "...my_servername..."
    }
    token_ttl_minutes = 38.63
    type              = "cribl_http"
    url               = "https://cribl-playground.cloud/api/v1/m/default/destination/test_1"
    urls = [
      {
        url    = "https://cribl-playground.cloud/api/v1/m/default/destination/test_1"
        weight = 0.48
      }
    ]
    use_round_robin_dns = false
  }
}

output "destination_details" {
  value = {
    id = cribl-terraform_destination.my_destination.output_cribl_http.id
    type = cribl-terraform_destination.my_destination.output_cribl_http.type
    environment = cribl-terraform_destination.my_destination.output_cribl_http.environment
    pipeline = cribl-terraform_destination.my_destination.output_cribl_http.pipeline
  }
}