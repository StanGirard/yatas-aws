package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stangirard/yatas/config"
)

func CheckIfEC2PublicIP(checkConfig config.CheckConfig, instances []types.Instance, testName string) {
	var check config.Check
	check.InitCheck("EC2s don't have a public IP", "Check if all instances have a public IP", testName)
	for _, instance := range instances {
		if instance.PublicIpAddress != nil {
			Message := "EC2 instance " + *instance.InstanceId + " has a public IP" + *instance.PublicIpAddress
			result := config.Result{Status: "FAIL", Message: Message, ResourceID: *instance.InstanceId}
			check.AddResult(result)
		} else {
			Message := "EC2 instance " + *instance.InstanceId + " has no public IP "
			result := config.Result{Status: "OK", Message: Message, ResourceID: *instance.InstanceId}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
