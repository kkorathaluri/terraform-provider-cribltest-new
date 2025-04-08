resource "cribl-terraform_group" "my_group" {
  cloud = {
    provider = "azure"
    region   = "...my_region..."
  }
  config_version         = "...my_config_version..."
  deploying_worker_count = 4.87
  description            = "...my_description..."
  estimated_ingest_rate  = 4.46
  git = {
    commit        = "...my_commit..."
    local_changes = 8.64
    log = [
      {
        author_email = "...my_author_email..."
        author_name  = "...my_author_name..."
        date         = "...my_date..."
        hash         = "...my_hash..."
        message      = "...my_message..."
        short        = "...my_short..."
      }
    ]
  }
  id                        = "...my_id..."
  incompatible_worker_count = 0.87
  inherits                  = "...my_inherits..."
  is_fleet                  = false
  is_search                 = true
  name                      = "...my_name..."
  on_prem                   = false
  product                   = "stream"
  provisioned               = true
  streamtags = [
    "..."
  ]
  tags                 = "...my_tags..."
  upgrade_version      = "...my_upgrade_version..."
  worker_count         = 7.37
  worker_remote_access = false
}