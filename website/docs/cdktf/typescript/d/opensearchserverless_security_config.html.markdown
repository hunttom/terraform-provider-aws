---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_security_config"
description: |-
  Terraform data source for managing an AWS OpenSearch Serverless Security Config.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_opensearchserverless_security_config

Terraform data source for managing an AWS OpenSearch Serverless Security Config.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsOpensearchserverlessSecurityConfig } from "./.gen/providers/aws/data-aws-opensearchserverless-security-config";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsOpensearchserverlessSecurityConfig(this, "example", {
      id: "saml/12345678912/example",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `id` - (Required) The unique identifier of the security configuration.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `configVersion` - The version of the security configuration.
* `createdDate` - The date the configuration was created.
* `description` - The description of the security configuration.
* `lastModifiedDate` - The date the configuration was last modified.
* `samlOptions` - SAML options for the security configuration.
* `type` - The type of security configuration.

### saml_options

SAML options for the security configuration.

* `groupAttribute` - Group attribute for this SAML integration.
* `metadata` - The XML IdP metadata file generated from your identity provider.
* `sessionTimeout` - Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
* `userAttribute` - User attribute for this SAML integration.

<!-- cache-key: cdktf-0.20.8 input-8187974261e12b138d3a7e197de7ca26b20f843eafab8ec86b4f5a9e972ddaa6 -->