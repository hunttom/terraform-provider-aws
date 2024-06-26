---
subcategory: "DRS (Elastic Disaster Recovery)"
layout: "aws"
page_title: "AWS: drs_replication_configuration_template"
description: |-
  Provides an Elastic Disaster Recovery replication configuration template resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_drs_replication_configuration_template

Provides an Elastic Disaster Recovery replication configuration template resource.

~> **NOTE:** This resource is provided on a best-effort basis and may not function as intended. Due to challenges with DRS permissions, it has not been fully tested. We are collaborating with AWS to enhance its functionality and [welcome your feedback](https://github.com/hashicorp/terraform-provider-aws/issues/new/choose).

## Example Usage

### Basic configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DrsReplicationConfigurationTemplate } from "./.gen/providers/aws/drs-replication-configuration-template";
interface MyConfig {
  ebsEncryption: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new DrsReplicationConfigurationTemplate(this, "example", {
      associateDefaultSecurityGroup: false,
      bandwidthThrottling: 12,
      createPublicIp: false,
      dataPlaneRouting: "PRIVATE_IP",
      defaultLargeStagingDiskType: "GP2",
      ebs_ecryption: "DEFAULT",
      ebsEncryptionKeyArn:
        "arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab",
      pitPolicy: [
        {
          enabled: true,
          interval: 1,
          retentionDuration: 1,
          units: "DAY",
        },
      ],
      replicationServerInstanceType: "t3.small",
      replicationServersSecurityGroupsIds: Token.asList(
        Fn.lookupNested(awsSecurityGroupExample, ["*", "id"])
      ),
      stagingAreaSubnetId: Token.asString(awsSubnetExample.id),
      useDedicatedReplicationServer: false,
      ebsEncryption: config.ebsEncryption,
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `associateDefaultSecurityGroup` - (Required) Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.
* `bandwidthThrottling` - (Required) Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.
* `createPublicIp` - (Required) Whether to create a Public IP for the Recovery Instance by default.
* `dataPlaneRouting` - (Required) Data plane routing mechanism that will be used for replication. Valid values are `PUBLIC_IP` and `PRIVATE_IP`.
* `defaultLargeStagingDiskType` - (Required) Staging Disk EBS volume type to be used during replication. Valid values are `GP2`, `GP3`, `ST1`, or `AUTO`.
* `ebsEncryption` - (Required) Type of EBS encryption to be used during replication. Valid values are `DEFAULT` and `CUSTOM`.
* `ebsEncryptionKeyArn` - (Required) ARN of the EBS encryption key to be used during replication.
* `pitPolicy` - (Required) Configuration block for Point in time (PIT) policy to manage snapshots taken during replication. [See below](#pit_policy).
* `replicationServerInstanceType` - (Required) Instance type to be used for the replication server.
* `replicationServersSecurityGroupsIds` - (Required) Security group IDs that will be used by the replication server.
* `stagingAreaSubnetId` - (Required) Subnet to be used by the replication staging area.
* `stagingAreaTags` - (Required) Set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
* `useDedicatedReplicationServer` - (Required) Whether to use a dedicated Replication Server in the replication staging area.

The following arguments are optional:

* `autoReplicateNewDisks` - (Optional) Whether to allow the AWS replication agent to automatically replicate newly added disks.
* `tags` - (Optional) Set of tags to be associated with the Replication Configuration Template resource.

### `pitPolicy`

* `enabled` - (Optional) Whether this rule is enabled or not.
* `interval` - (Required) How often, in the chosen units, a snapshot should be taken.
* `retentionDuration` - (Required) Duration to retain a snapshot for, in the chosen `units`.
* `ruleId` - (Optional) ID of the rule. Valid values are integers.
* `units` - (Required) Units used to measure the `interval` and `retentionDuration`. Valid values are `MINUTE`, `HOUR`, and `DAY`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Replication configuration template ARN.
* `id` - Replication configuration template ID.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `20m`)
- `update` - (Default `20m`)
- `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DRS Replication Configuration Template using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DrsReplicationConfigurationTemplate } from "./.gen/providers/aws/drs-replication-configuration-template";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DrsReplicationConfigurationTemplate.generateConfigForImport(
      this,
      "example",
      "templateid"
    );
  }
}

```

Using `terraform import`, import DRS Replication Configuration Template using the `id`. For example:

```console
% terraform import aws_drs_replication_configuration_template.example templateid
```

<!-- cache-key: cdktf-0.20.1 input-14df6a78db1c9ae29f6e63bd7a2f593dd19bd956dce0d2743d9229bc0501cdd6 -->