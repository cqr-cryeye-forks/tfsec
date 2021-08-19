package documentdb

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

// generator-locked
import (
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/provider"
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/aquasecurity/tfsec/pkg/severity"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		Provider:  provider.AWSProvider,
		Service:   "documentdb",
		ShortCode: "enable-storage-encryption",
		Documentation: rule.RuleDocumentation{
			Summary:     "DocumentDB storage must be encrypted",
			Explanation: `Encryption of the underlying storage used by DocumentDB ensures that if their is compromise of the disks, the data is still protected.`,
			Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
			Resolution:  "Enable storage encryption",
			BadExample: []string{`
resource "aws_docdb_cluster" "bad_example" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  storage_encrypted = false
}
`},
			GoodExample: []string{`
resource "aws_docdb_cluster" "good_example" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  storage_encrypted = true
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#storage_encrypted",
			},
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_docdb_cluster",
		},
		DefaultSeverity: severity.High,
		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
			if storageEncryptedAttr := resourceBlock.GetAttribute("storage_encrypted"); storageEncryptedAttr.IsNil() { // alert on use of default value
				set.AddResult().
					WithDescription("Resource '%s' uses default value for storage_encrypted", resourceBlock.FullName())
			} else if storageEncryptedAttr.IsFalse() {
				set.AddResult().
					WithDescription("Resource '%s' does not have storage_encrypted set to true", resourceBlock.FullName()).
					WithAttribute(storageEncryptedAttr)
			}
		},
	})
}
