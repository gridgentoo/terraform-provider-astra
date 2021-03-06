---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "astra_private_link_endpoint Resource - terraform-provider-astra"
subcategory: ""
description: |-
  astra_private_link_endpoint completes the creation of a private link endpoint by associating it with your endpoint.
---

# astra_private_link_endpoint (Resource)

`astra_private_link_endpoint` completes the creation of a private link endpoint by associating it with your endpoint.

## Example Usage

```terraform
resource "astra_private_link" "example" {
  allowed_principals = ["arn:aws:iam::445559476293:user/Sebastian"]
  database_id        = "a6bc9c26-e7ce-424f-84c7-0a00afb12588"
  datacenter_id      = "a6bc9c26-e7ce-424f-84c7-0a00afb12588"
}
resource "aws_vpc_endpoint" "example" {
  vpc_id             = "vpc-f939e884"
  service_name       = astra_private_link.example.service_name
  vpc_endpoint_type  = "Interface"
  subnet_ids         = ["subnet-4d376300", "subnet-4d85066c", "subnet-030e8b65"]
  security_group_ids = ["sg-74ae4d41"]
}
resource "astra_private_link_endpoint" "example" {
  database_id   = "a6bc9c26-e7ce-424f-84c7-0a00afb12588"
  datacenter_id = "a6bc9c26-e7ce-424f-84c7-0a00afb12588"
  endpoint_id   = aws_vpc_endpoint.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **database_id** (String) The ID of the Astra database.
- **datacenter_id** (String) Astra datacenter in the region where the private link will be created.
- **endpoint_id** (String) Endpoint created in your cloud provider account example: "vpce-svc-1148ea04af8675309"

### Optional

- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import astra_private_link_endpoint.example a6bc9c26-e7ce-424f-84c7-0a00afb12588
```
