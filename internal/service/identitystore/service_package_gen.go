// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package identitystore

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceGroup,
			TypeName: "aws_identitystore_group",
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_identitystore_user",
		},
		{
			Factory:  DataSourceUsers,
			TypeName: "aws_identitystore_users",
			Name:     "Users",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceGroup,
			TypeName: "aws_identitystore_group",
		},
		{
			Factory:  ResourceGroupMembership,
			TypeName: "aws_identitystore_group_membership",
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_identitystore_user",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IdentityStore
}

var ServicePackage = &servicePackage{}
