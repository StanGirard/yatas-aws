package autoscaling

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/stangirard/yatas/config"
)

func CheckIfDesiredCapacityMaxCapacityBelow80percent(checkConfig config.CheckConfig, groups []types.AutoScalingGroup, testName string) {
	var check config.Check
	check.InitCheck("Autoscaling maximum capacity is below 80%", "Check if all autoscaling groups have a desired capacity below 80%", testName)
	for _, group := range groups {
		if group.DesiredCapacity != nil && group.MaxSize != nil && float64(*group.DesiredCapacity) > float64(*group.MaxSize)*0.8 {
			Message := "Autoscaling group " + *group.AutoScalingGroupName + " has a desired capacity above 80%"
			result := config.Result{Status: "FAIL", Message: Message, ResourceID: *group.AutoScalingGroupName}
			check.AddResult(result)
		} else {
			Message := "Autoscaling group " + *group.AutoScalingGroupName + " has a desired capacity below 80%"
			result := config.Result{Status: "OK", Message: Message, ResourceID: *group.AutoScalingGroupName}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
