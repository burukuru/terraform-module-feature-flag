package main

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"

	awsSDK "github.com/aws/aws-sdk-go/aws"
)

func TestTerraformEcsClusterWithLogs(t *testing.T) {
	t.Parallel()

	expectedClusterName := fmt.Sprintf("test_cluster_%s", random.UniqueId())
	region := "us-east-1"

	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "env/dev/")

	terraformOptions := &terraform.Options{
		TerraformDir: tempTestFolder,

		Vars: map[string]interface{}{
			"cluster_name": expectedClusterName,
			"region":       region,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	cluster := aws.GetEcsCluster(t, region, expectedClusterName)
	service := aws.GetEcsService(t, region, expectedClusterName, "test_service")

	assert.Equal(t, expectedClusterName, awsSDK.StringValue(cluster.ClusterName))
	assert.Equal(t, "test_service", awsSDK.StringValue(service.ServiceName))
}

func TestTerraformEcsClusterWithoutLogs(t *testing.T) {
	t.Parallel()

	expectedClusterName := fmt.Sprintf("test_cluster_%s", random.UniqueId())
	region := "us-east-1"
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "env/dev/")

	terraformOptions := &terraform.Options{
		TerraformDir: tempTestFolder,

		Vars: map[string]interface{}{
			"region":                      region,
			"cluster_name":                expectedClusterName,
			"create_cloudwatch_log_group": false,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	cluster := aws.GetEcsCluster(t, region, expectedClusterName)
	service := aws.GetEcsService(t, region, expectedClusterName, "test_service")

	assert.Equal(t, expectedClusterName, awsSDK.StringValue(cluster.ClusterName))
	assert.Equal(t, "test_service", awsSDK.StringValue(service.ServiceName))
}
