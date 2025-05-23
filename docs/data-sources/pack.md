---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cribl-terraform_pack Data Source - terraform-provider-cribl-terraform"
subcategory: ""
description: |-
  Pack DataSource
---

# cribl-terraform_pack (Data Source)

Pack DataSource

## Example Usage

```terraform
data "cribl-terraform_pack" "my_pack" {
  group_id = "...my_group_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) Group Id

### Read-Only

- `items` (Attributes List) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `author` (String)
- `description` (String)
- `display_name` (String)
- `exports` (List of String)
- `id` (String)
- `is_disabled` (Boolean)
- `min_log_stream_version` (String)
- `settings` (Map of String)
- `source` (String)
- `spec` (String)
- `tags` (Attributes) (see [below for nested schema](#nestedatt--items--tags))
- `version` (String)

<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

- `data_type` (List of String)
- `domain` (List of String)
- `streamtags` (List of String)
- `technology` (List of String)
