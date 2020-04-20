package main

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEcsClusterWithLogs(t *testing.T) {
	t.Parallel()

	expectedClusterName := fmt.Sprintf("test_cluster_%s", random.UniqueId())

	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "env/dev/")

	terraformOptions := &terraform.Options{
		TerraformDir: tempTestFolder,

		Vars: map[string]interface{}{
			"cluster_name": expectedClusterName,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	assert.Equal(t, 1, 1)
}

func TestTerraformEcsClusterWithoutLogs(t *testing.T) {
	t.Parallel()

	expectedClusterName := fmt.Sprintf("test_cluster_%s", random.UniqueId())
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "env/dev/")

	terraformOptions := &terraform.Options{
		TerraformDir: tempTestFolder,

		Vars: map[string]interface{}{
			"cluster_name":                expectedClusterName,
			"create_cloudwatch_log_group": false,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	assert.Equal(t, 1, 1)
}
