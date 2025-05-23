---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cribl-terraform_pack Resource - terraform-provider-cribl-terraform"
subcategory: ""
description: |-
  Pack Resource
---

# cribl-terraform_pack (Resource)

Pack Resource

## Example Usage

```terraform
resource "cribl-terraform_pack" "my_pack" {
  filename = "...my_filename..."
  group_id = "...my_group_id..."
  id       = "...my_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) Group Id
- `id` (String) Pack name

### Optional

- `filename` (String) the file to upload. Requires replacement if changed.

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
- `warnings` (String) Parsed as JSON.

<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

- `data_type` (List of String)
- `domain` (List of String)
- `streamtags` (List of String)
- `technology` (List of String)

## Import

Import is supported using the following syntax:

```shell
terraform import cribl-terraform_pack.my_cribl-terraform_pack ""
```
