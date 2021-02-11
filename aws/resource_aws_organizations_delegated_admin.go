package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsOrganizationsDelegatedAdminAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsOrganizationsDelegatedAdminAccountCreate,
		Read:   resourceAwsOrganizationsDelegatedAdminAccountRead,
		Delete: resourceAwsOrganizationsDelegatedAdminAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"admin_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAwsAccountId,
			},
			"aws_service_access_principal": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsOrganizationsDelegatedAdminAccountCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).organizationsconn

	adminAccountID := d.Get("admin_account_id").(string)
	awsServiceAccessPrincipal := d.Get("aws_service_access_principal").(string)

	input := &organizations.RegisterDelegatedAdministratorInput{
		AccountId:        aws.String(adminAccountID),
		ServicePrincipal: aws.String(awsServiceAccessPrincipal),
	}

	_, err := conn.RegisterDelegatedAdministrator(input)

	if err != nil {
		return fmt.Errorf("error enabling Organizations Delegated Admin Account (%s): %w", adminAccountID, err)
	}

	d.SetId(adminAccountID)

	if err := waitForOrganizationDelegatedAdminEnable(conn, d.Id()); err != nil {
		return fmt.Errorf("error waiting for policy type (%s) enabling in Organization (%s): %s", awsServiceAccessPrincipal, d.Id(), err)
	}

	return resourceAwsOrganizationsDelegatedAdminAccountRead(d, meta)
}

func resourceAwsOrganizationsDelegatedAdminAccountRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).organizationsconn

	log.Printf("[INFO] Reading Organization: %s", d.Id())

	adminAccount, err := getOrganizationsDelegatedAdminAccount(conn, d)

	if err != nil {
		return fmt.Errorf("error reading Delegated Organization Admin Account (%s): %w", d.Id(), err)
	}

	if adminAccount == nil {
		log.Printf("[WARN] Delegated Organization Admin Account (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("admin_account_id", adminAccount.Id)

	return nil
}

func resourceAwsOrganizationsDelegatedAdminAccountDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).organizationsconn

	awsServiceAccessPrincipal := d.Get("aws_service_access_principal").(string)

	input := &organizations.DeregisterDelegatedAdministratorInput{
		AccountId:        aws.String(d.Id()),
		ServicePrincipal: aws.String(awsServiceAccessPrincipal),
	}

	_, err := conn.DeregisterDelegatedAdministrator(input)

	if err != nil {
		return fmt.Errorf("error disabling Delegated Organization Admin Account (%s): %w", d.Id(), err)
	}

	if err := waitForOrganizationDelegatedAdminDisable(conn, d.Id()); err != nil {
		return fmt.Errorf("error waiting for policy type (%s) enabling in Organization (%s): %s", awsServiceAccessPrincipal, d.Id(), err)
	}
	return nil
}

func getOrganizationsDelegatedAdminAccount(conn *organizations.Organizations, d *schema.ResourceData) (*organizations.DelegatedAdministrator, error) {
	awsServiceAccessPrincipal := d.Get("aws_service_access_principal").(string)

	input := &organizations.ListDelegatedAdministratorsInput{
		ServicePrincipal: aws.String(awsServiceAccessPrincipal),
	}
	var result *organizations.DelegatedAdministrator

	err := conn.ListDelegatedAdministratorsPages(input, func(page *organizations.ListDelegatedAdministratorsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, adminAccount := range page.DelegatedAdministrators {
			if adminAccount == nil {
				continue
			}

			adminAccountId := d.Get("admin_account_id").(string)

			if aws.StringValue(adminAccount.Id) == adminAccountId {
				result = adminAccount
				return false
			}
		}

		return !lastPage
	})

	return result, err
}

func waitForOrganizationDelegatedAdminEnable(conn *organizations.Organizations, policyType string) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			organizationsPolicyTypeStatusDisabled,
			organizations.PolicyTypeStatusPendingEnable,
		},
		Target:  []string{organizations.PolicyTypeStatusEnabled},
		Refresh: getOrganizationDefaultRootPolicyTypeRefreshFunc(conn, policyType),
		Timeout: 5 * time.Minute,
	}

	_, err := stateConf.WaitForState()

	return err
}

func waitForOrganizationDelegatedAdminDisable(conn *organizations.Organizations, policyType string) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			organizationsPolicyTypeStatusDisabled,
			organizations.PolicyTypeStatusPendingEnable,
		},
		Target:  []string{organizations.PolicyTypeStatusEnabled},
		Refresh: getOrganizationDefaultRootPolicyTypeRefreshFunc(conn, policyType),
		Timeout: 5 * time.Minute,
	}

	_, err := stateConf.WaitForState()

	return err
}
