package organizations_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func init() {
	acctest.RegisterServiceErrorCheckFunc(organizations.EndpointsID, testAccErrorCheckSkip)
}

func testAccErrorCheckSkip(t *testing.T) resource.ErrorCheckFunc {
	return acctest.ErrorCheckSkipMessagesContaining(t,
		"MASTER_ACCOUNT_NOT_GOVCLOUD_ENABLED",
	)
}

func TestAccOrganizations_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Organization": {
			"basic":                        testAccOrganization_basic,
			"AwsServiceAccessPrincipals":   testAccOrganization_serviceAccessPrincipals,
			"EnabledPolicyTypes":           testAccOrganization_EnabledPolicyTypes,
			"FeatureSet_Basic":             testAccOrganization_FeatureSet,
			"FeatureSet_Update":            testAccOrganization_FeatureSetUpdate,
			"FeatureSet_ForcesNew":         testAccOrganization_FeatureSetForcesNew,
			"DataSource":                   testAccOrganizationDataSource_basic,
			"ChildAccountsDataSource":      testAccOrganizationalUnitChildAccountsDataSource_basic,
			"DescendantAccountsDataSource": testAccOrganizationalUnitDescendantAccountsDataSource_basic,
		},
		"Account": {
			"basic":           testAccAccount_basic,
			"CloseOnDeletion": testAccAccount_CloseOnDeletion,
			"ParentId":        testAccAccount_ParentID,
			"Tags":            testAccAccount_Tags,
			"GovCloud":        testAccAccount_govCloud,
		},
		"OrganizationalUnit": {
			"basic":      testAccOrganizationalUnit_basic,
			"disappears": testAccOrganizationalUnit_disappears,
			"Name":       testAccOrganizationalUnit_Name,
			"Tags":       testAccOrganizationalUnit_Tags,
		},
		"OrganizationalUnits": {
			"DataSource": testAccOrganizationalUnitsDataSource_basic,
		},
		"Policy": {
			"basic":                  testAccPolicy_basic,
			"concurrent":             testAccPolicy_concurrent,
			"Description":            testAccPolicy_description,
			"Tags":                   testAccPolicy_tags,
			"SkipDestroy":            testAccPolicy_skipDestroy,
			"disappears":             testAccPolicy_disappears,
			"Type_AI_OPT_OUT":        testAccPolicy_type_AI_OPT_OUT,
			"Type_Backup":            testAccPolicy_type_Backup,
			"Type_SCP":               testAccPolicy_type_SCP,
			"Type_Tag":               testAccPolicy_type_Tag,
			"ImportAwsManagedPolicy": testAccPolicy_importManagedPolicy,
		},
		"PolicyAttachment": {
			"Account":            testAccPolicyAttachment_Account,
			"OrganizationalUnit": testAccPolicyAttachment_OrganizationalUnit,
			"Root":               testAccPolicyAttachment_Root,
			"SkipDestroy":        testAccPolicyAttachment_skipDestroy,
			"disappears":         testAccPolicyAttachment_disappears,
		},
		"PolicyDataSource": {
			"UnattachedPolicy": testAccPolicyDataSource_UnattachedPolicy,
		},
		"DelegatedAdministrator": {
			"basic":      testAccDelegatedAdministrator_basic,
			"disappears": testAccDelegatedAdministrator_disappears,
		},
		"ResourceTags": {
			"basic": testAccResourceTagsDataSource_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
