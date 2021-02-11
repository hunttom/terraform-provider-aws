package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccAwsOrganizationsDelegatedAdminAccount_basic(t *testing.T) {
	resourceName := "aws_organizations_delegated_admin_.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccOrganizationsAccountPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsOrganizationsDelegatedAdminAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsOrganizationsDelegatedAdminAccountConfigSelf(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsOrganizationsDelegatedAdminAccountExists(resourceName),
					testAccCheckResourceAttrAccountID(resourceName, "admin_account_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAwsOrganizationsDelegatedAdminAccountDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).organizationsconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_organizations_delegated_admin" {
			continue
		}

		accountInfo, err := getOrganizationsDelegatedAdminAccount(conn, rs)
		adminAccount := accountInfo.Get("admin_account_id")

		if isAWSErr(err, organizations.ConstraintViolationExceptionReasonCannotRemoveDelegatedAdministratorFromOrg, "unable to remove delegated administrator") {
			continue
		}

		if err != nil {
			f
			return err
		}

		if adminAccount == nil {
			continue
		}

		return fmt.Errorf("expected Delegated Organization Admin Account (%s) to be removed", rs.Primary.ID)
	}

	return nil
}

func testAccCheckAwsOrganizationsDelegatedAdminAccountExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		conn := testAccProvider.Meta().(*AWSClient).organizationsconn

		adminAccount, err := getOrganizationsDelegatedAdminAccount(conn, rs.Primary)

		if err != nil {
			return err
		}

		if adminAccount == nil {
			accountID := rs.Get("admin_account_id")
			return fmt.Errorf("Delegated Organization Admin Account (%s) not found", accountID)
		}

		return nil
	}
}

func testAccCheckAwsOrganizationsDelegatedAdminAccountConfigSelf() string {
	return `
data "aws_caller_identity" "current" {}

resource "aws_organizations_delegated_admin" "test" {
  admin_account_id = data.aws_caller_identity.current.account_id
  aws_service_access_principal = "config.amazonaws.com"
}
`
}
