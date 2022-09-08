package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/stangirard/yatas/config"
)

func checkIfRDSPrivateEnabled(checkConfig config.CheckConfig, instances []types.DBInstance, testName string) {
	var check config.Check
	check.InitCheck("RDS aren't publicly accessible", "Check if RDS private is enabled", testName)
	for _, instance := range instances {
		if instance.PubliclyAccessible {
			Message := "RDS private is not enabled on " + *instance.DBInstanceIdentifier
			result := config.Result{Status: "FAIL", Message: Message, ResourceID: *instance.DBInstanceArn}
			check.AddResult(result)
		} else {
			Message := "RDS private is enabled on " + *instance.DBInstanceIdentifier
			result := config.Result{Status: "OK", Message: Message, ResourceID: *instance.DBInstanceArn}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
