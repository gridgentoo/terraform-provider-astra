---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "astra Provider"
subcategory: ""
description: |-
  
---

# astra Provider



## Example Usage

```terraform
variable "token" {}

provider "astra" {
  // This can also be set via ASTRA_API_TOKEN environment variable.
  token = var.token
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **token** (String) Authentication token for Astra API.