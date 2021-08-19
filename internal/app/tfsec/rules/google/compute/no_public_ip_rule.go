package compute

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

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
		Provider:  provider.GoogleProvider,
		Service:   "compute",
		ShortCode: "no-public-ip",
		Documentation: rule.RuleDocumentation{
			Summary:     "Instances should not have public IP addresses",
			Explanation: `Instances should not be publicly exposed to the internet`,
			Impact:      "Direct exposure of an instance to the public internet",
			Resolution:  "Remove public IP",
			BadExample: []string{`
resource "google_compute_instance" "bad_example" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }
}
`},
			GoodExample: []string{`
resource "google_compute_instance" "good_example" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"
  }
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#access_config",
			},
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"google_compute_instance",
		},
		DefaultSeverity: severity.High,
		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
			if accessConfigBlock := resourceBlock.GetBlock("network_interface").GetBlock("access_config"); !accessConfigBlock.IsNil() {
				set.AddResult().
					WithDescription("Resource '%s' sets access_config", resourceBlock.FullName()).
					WithBlock(accessConfigBlock)
			}
		},
	})
}
