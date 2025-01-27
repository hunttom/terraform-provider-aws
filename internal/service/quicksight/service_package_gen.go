// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newFolderMembershipResource,
			TypeName: "aws_quicksight_folder_membership",
			Name:     "Folder Membership",
		},
		{
			Factory:  newIAMPolicyAssignmentResource,
			TypeName: "aws_quicksight_iam_policy_assignment",
			Name:     "IAM Policy Assignment",
		},
		{
			Factory:  newIngestionResource,
			TypeName: "aws_quicksight_ingestion",
			Name:     "Ingestion",
		},
		{
			Factory:  newNamespaceResource,
			TypeName: "aws_quicksight_namespace",
			Name:     "Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newRefreshScheduleResource,
			TypeName: "aws_quicksight_refresh_schedule",
			Name:     "Refresh Schedule",
		},
		{
			Factory:  newTemplateAliasResource,
			TypeName: "aws_quicksight_template_alias",
			Name:     "Template Alias",
		},
		{
			Factory:  newVPCConnectionResource,
			TypeName: "aws_quicksight_vpc_connection",
			Name:     "VPC Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
		},
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  dataSourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccountSubscription,
			TypeName: "aws_quicksight_account_subscription",
			Name:     "Account Subscription",
		},
		{
			Factory:  resourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDashboard,
			TypeName: "aws_quicksight_dashboard",
			Name:     "Dashboard",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDataSource,
			TypeName: "aws_quicksight_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFolder,
			TypeName: "aws_quicksight_folder",
			Name:     "Folder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  resourceGroupMembership,
			TypeName: "aws_quicksight_group_membership",
			Name:     "Group Membership",
		},
		{
			Factory:  resourceTemplate,
			TypeName: "aws_quicksight_template",
			Name:     "Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.QuickSight
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*quicksight.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*quicksight.Options){
		quicksight.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	return quicksight.NewFromConfig(cfg, optFns...), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
