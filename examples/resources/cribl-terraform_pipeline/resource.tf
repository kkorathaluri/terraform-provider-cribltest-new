resource "cribl-terraform_pipeline" "my_pipeline" {
  conf = {
    async_func_timeout = 9527
    description        = "...my_description..."
    functions = [
      {
        conf = {
          key = jsonencode("value")
        }
        description = "...my_description..."
        disabled    = true
        filter      = "...my_filter..."
        final       = true
        group_id    = "...my_group_id..."
        id          = "...my_id..."
      }
    ]
    groups = {
      key = {
        description = "...my_description..."
        disabled    = false
        name        = "...my_name..."
      }
    }
    output = "...my_output..."
    streamtags = [
      "..."
    ]
  }
  group_id = "...my_group_id..."
  id       = "...my_id..."
}