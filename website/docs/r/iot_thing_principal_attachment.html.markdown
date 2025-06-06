---
subcategory: "IoT Core"
layout: "aws"
page_title: "AWS: aws_iot_thing_principal_attachment"
description: |-
    Provides AWS IoT Thing Principal attachment.
---

# Resource: aws_iot_thing_principal_attachment

Attaches Principal to AWS IoT Thing.

## Example Usage

```terraform
resource "aws_iot_thing" "example" {
  name = "example"
}

resource "aws_iot_certificate" "cert" {
  csr    = file("csr.pem")
  active = true
}

resource "aws_iot_thing_principal_attachment" "att" {
  principal = aws_iot_certificate.cert.arn
  thing     = aws_iot_thing.example.name
}
```

## Argument Reference

This resource supports the following arguments:

* `principal` - (Required) The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
* `thing` - (Required) The name of the thing.

## Attribute Reference

This resource exports no additional attributes.
