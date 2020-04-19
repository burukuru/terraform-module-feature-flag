package main

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEcsClusterWithLogs(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../env/dev/",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	assert.Equal(t, 1, 1)
}

func TestTerraformEcsClusterWithoutLogs(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../env/dev/",

		Vars: map[string]interface{} {
			"create_cloudwatch_log_group": false,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	assert.Equal(t, 1, 1)
}
