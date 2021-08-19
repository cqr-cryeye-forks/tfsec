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
		ShortCode: "enable-shielded-vm",
		Documentation: rule.RuleDocumentation{
			Summary:     "Instances should have Shielded VM enabled",
			Explanation: `A Shielded VM is a VM with enhanced defences/detection for rootkits/bootkits.`,
			Impact:      "Unable to detect rootkits",
			Resolution:  "Enable Shielded VM",
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

  shielded_instance_config {
    enable_vtpm = false
    enable_integrity_monitoring = false
  }
}
`,
				`
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

  shielded_instance_config {
    enable_vtpm = true
    enable_integrity_monitoring = false
  }
}
`,
				`
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

  shielded_instance_config {
    enable_vtpm = false
    enable_integrity_monitoring = true
  }
}
`},
			GoodExample: []string{`
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

  shielded_instance_config {
    enable_vtpm = true
    enable_integrity_monitoring = true
  }
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#enable_vtpm",
			},
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"google_compute_instance",
		},
		DefaultSeverity: severity.Medium,
		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
			if enableVtpmAttr := resourceBlock.GetBlock("shielded_instance_config").GetAttribute("enable_vtpm"); enableVtpmAttr.IsFalse() {
				set.AddResult().
					WithDescription("Resource '%s' has shielded_instance_config.enable_vtpm set to false", resourceBlock.FullName()).
					WithAttribute(enableVtpmAttr)
			}
			if enableIMAttr := resourceBlock.GetBlock("shielded_instance_config").GetAttribute("enable_integrity_monitoring"); enableIMAttr.IsFalse() {
				set.AddResult().
					WithDescription("Resource '%s' has shielded_instance_config.enable_integrity_monitoring set to false", resourceBlock.FullName()).
					WithAttribute(enableIMAttr)
			}
		},
	})
}
