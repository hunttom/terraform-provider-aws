---
subcategory: "Web Services Budgets"
layout: "aws"
page_title: "AWS: aws_budgets_budget"
description: |-
  Provides a budgets budget resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_budgets_budget

Provides a budgets budget resource. Budgets use the cost visualization provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new BudgetsBudget(this, "ec2", {
      budgetType: "COST",
      costFilter: [
        {
          name: "Service",
          values: ["Amazon Elastic Compute Cloud - Compute"],
        },
      ],
      limitAmount: "1200",
      limitUnit: "USD",
      name: "budget-ec2-monthly",
      notification: [
        {
          comparisonOperator: "GREATER_THAN",
          notificationType: "FORECASTED",
          subscriberEmailAddresses: ["test@example.com"],
          threshold: 100,
          thresholdType: "PERCENTAGE",
        },
      ],
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      timePeriodEnd: "2087-06-15_00:00",
      timePeriodStart: "2017-07-01_00:00",
      timeUnit: "MONTHLY",
    });
  }
}

```

Create a budget for *$100*.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "cost", {
      budgetType: "COST",
      limitAmount: "100",
      limitUnit: "USD",
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a budget with planned budget limits.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  budgetType: any;
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "cost", {
      plannedLimit: [
        {
          amount: "100",
          startTime: "2017-07-01_00:00",
          unit: "USD",
        },
        {
          amount: "200",
          startTime: "2017-08-01_00:00",
          unit: "USD",
        },
      ],
      budgetType: config.budgetType,
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a budget for s3 with a limit of *3 GB* of storage.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "s3", {
      budgetType: "USAGE",
      limitAmount: "3",
      limitUnit: "GB",
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a Savings Plan Utilization Budget

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "savings_plan_utilization", {
      budgetType: "SAVINGS_PLANS_UTILIZATION",
      costTypes: {
        includeCredit: false,
        includeDiscount: false,
        includeOtherSubscription: false,
        includeRecurring: false,
        includeRefund: false,
        includeSubscription: true,
        includeSupport: false,
        includeTax: false,
        includeUpfront: false,
        useBlended: false,
      },
      limitAmount: "100.0",
      limitUnit: "PERCENTAGE",
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a RI Utilization Budget

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "ri_utilization", {
      budgetType: "RI_UTILIZATION",
      costFilter: [
        {
          name: "Service",
          values: ["Amazon Relational Database Service"],
        },
      ],
      costTypes: {
        includeCredit: false,
        includeDiscount: false,
        includeOtherSubscription: false,
        includeRecurring: false,
        includeRefund: false,
        includeSubscription: true,
        includeSupport: false,
        includeTax: false,
        includeUpfront: false,
        useBlended: false,
      },
      limitAmount: "100.0",
      limitUnit: "PERCENTAGE",
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a cost filter using resource tags

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  budgetType: any;
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "cost", {
      costFilter: [
        {
          name: "TagKeyValue",
          values: [
            "aws:createdBy$Terraform",
            "user:business-unit$human_resources",
          ],
        },
      ],
      budgetType: config.budgetType,
      timeUnit: config.timeUnit,
    });
  }
}

```

Create a cost filter using resource tags, obtaining the tag value from a Terraform variable

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
interface MyConfig {
  budgetType: any;
  timeUnit: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new BudgetsBudget(this, "cost", {
      costFilter: [
        {
          name: "TagKeyValue",
          values: ["TagKey$${var.TagValue}"],
        },
      ],
      budgetType: config.budgetType,
      timeUnit: config.timeUnit,
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `budgetType` - (Required) Whether this budget tracks monetary cost or usage.
* `timeUnit` - (Required) The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.

The following arguments are optional:

* `accountId` - (Optional) The ID of the target account for budget. Will use current user's account_id by default if omitted.
* `autoAdjustData` - (Optional) Object containing [AutoAdjustData](#auto-adjust-data) which determines the budget amount for an auto-adjusting budget.
* `costFilter` - (Optional) A list of [CostFilter](#cost-filter) name/values pair to apply to budget.
* `costTypes` - (Optional) Object containing [CostTypes](#cost-types) The types of cost included in a budget, such as tax and subscriptions.
* `limitAmount` - (Optional) The amount of cost or usage being measured for a budget.
* `limitUnit` - (Optional) The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
* `name` - (Optional) The name of a budget. Unique within accounts.
* `namePrefix` - (Optional) The prefix of the name of a budget. Unique within accounts.
* `notification` - (Optional) Object containing [Budget Notifications](#budget-notification). Can be used multiple times to define more than one budget notification.
* `plannedLimit` - (Optional) Object containing [Planned Budget Limits](#planned-budget-limits). Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
* `tags` - (Optional) Map of tags assigned to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `timePeriodEnd` - (Optional) The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
* `timePeriodStart` - (Optional) The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.

For more detailed documentation about each argument, refer to the [AWS official
documentation](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-budget.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the budget.
* `id` - id of resource.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

### Auto Adjust Data

The parameters that determine the budget amount for an auto-adjusting budget.

* `autoAdjustType` (Required) - The string that defines whether your budget auto-adjusts based on historical or forecasted data. Valid values: `FORECAST`,`HISTORICAL`
* `historicalOptions` (Optional) - Configuration block of [Historical Options](#historical-options). Required for `autoAdjustType` of `HISTORICAL` Configuration block that defines the historical data that your auto-adjusting budget is based on.
* `lastAutoAdjustTime` (Optional) - The last time that your budget was auto-adjusted.

### Historical Options

* `budgetAdjustmentPeriod` (Required) - The number of budget periods included in the moving-average calculation that determines your auto-adjusted budget amount.
* `lookbackAvailablePeriods` (Optional) - The integer that describes how many budget periods in your BudgetAdjustmentPeriod are included in the calculation of your current budget limit. If the first budget period in your BudgetAdjustmentPeriod has no cost data, then that budget period isn’t included in the average that determines your budget limit. You can’t set your own LookBackAvailablePeriods. The value is automatically calculated from the `budgetAdjustmentPeriod` and your historical cost data.

### Cost Types

Valid keys for `costTypes` parameter.

* `includeCredit` - A boolean value whether to include credits in the cost budget. Defaults to `true`
* `includeDiscount` - Whether a budget includes discounts. Defaults to `true`
* `includeOtherSubscription` - A boolean value whether to include other subscription costs in the cost budget. Defaults to `true`
* `includeRecurring` - A boolean value whether to include recurring costs in the cost budget. Defaults to `true`
* `includeRefund` - A boolean value whether to include refunds in the cost budget. Defaults to `true`
* `includeSubscription` - A boolean value whether to include subscriptions in the cost budget. Defaults to `true`
* `includeSupport` - A boolean value whether to include support costs in the cost budget. Defaults to `true`
* `includeTax` - A boolean value whether to include tax in the cost budget. Defaults to `true`
* `includeUpfront` - A boolean value whether to include upfront costs in the cost budget. Defaults to `true`
* `useAmortized` - Whether a budget uses the amortized rate. Defaults to `false`
* `useBlended` - A boolean value whether to use blended costs in the cost budget. Defaults to `false`

Refer to [AWS CostTypes documentation](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CostTypes.html) for further detail.

### Cost Filter

Valid keys for `costFilter` parameter.

* name - (Required) The name of the cost filter. Valid values are `AZ`, `BillingEntity`, `CostCategory`, `InstanceType`, `InvoicingEntity`, `LegalEntityName`, `LinkedAccount`, `Operation`, `PurchaseType`, `Region`, `Service`, `TagKeyValue`, `UsageType`, and `UsageTypeGroup`.
* values - (Required) The list of values used for filtering.

Refer to [AWS CostFilter documentation](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-create-filters.html) for further detail.

### Budget Notification

Valid keys for `notification` parameter.

* `comparisonOperator` - (Required) Comparison operator to use to evaluate the condition. Can be `LESS_THAN`, `EQUAL_TO` or `GREATER_THAN`.
* `threshold` - (Required) Threshold when the notification should be sent.
* `thresholdType` - (Required) What kind of threshold is defined. Can be `PERCENTAGE` OR `ABSOLUTE_VALUE`.
* `notificationType` - (Required) What kind of budget value to notify on. Can be `ACTUAL` or `FORECASTED`
* `subscriberEmailAddresses` - (Optional) E-Mail addresses to notify. Either this or `subscriberSnsTopicArns` is required.
* `subscriberSnsTopicArns` - (Optional) SNS topics to notify. Either this or `subscriberEmailAddresses` is required.

### Planned Budget Limits

Valid keys for `plannedLimit` parameter.

* `startTime` - (Required) The start time of the budget limit. Format: `2017-01-01_12:00`. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
* `amount` - (Required) The amount of cost or usage being measured for a budget.
* `unit` - (Required) The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import budgets using `AccountID:BudgetName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BudgetsBudget } from "./.gen/providers/aws/budgets-budget";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    BudgetsBudget.generateConfigForImport(
      this,
      "myBudget",
      "123456789012:myBudget"
    );
  }
}

```

Using `terraform import`, import budgets using `AccountID:BudgetName`. For example:

```console
% terraform import aws_budgets_budget.myBudget 123456789012:myBudget
```

<!-- cache-key: cdktf-0.20.8 input-dd9b63f9a398499a48eeeb8ff1356cc5ea01c7583a6ad256695d6226fb2a2535 -->