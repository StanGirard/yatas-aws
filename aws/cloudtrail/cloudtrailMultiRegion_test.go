package cloudtrail

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/padok-team/yatas/plugins/commons"
)

func TestCheckIfCloudtrailsMultiRegion(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		cloudtrails []types.Trail
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfCloudtrailsMultiRegion",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				cloudtrails: []types.Trail{
					{
						Name:                       aws.String("test"),
						TrailARN:                   aws.String("test"),
						IncludeGlobalServiceEvents: aws.Bool(false),
						IsMultiRegionTrail:         aws.Bool(true),
					},
				},
				testName: "TestCheckIfCloudtrailsMultiRegion",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfCloudtrailsMultiRegion(tt.args.checkConfig, tt.args.cloudtrails, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfCloudtrailsMultiRegion() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfCloudtrailsMultiRegionFail(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		cloudtrails []types.Trail
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfCloudtrailsMultiRegion",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				cloudtrails: []types.Trail{
					{
						Name:                       aws.String("test"),
						TrailARN:                   aws.String("test"),
						IncludeGlobalServiceEvents: aws.Bool(false),
						IsMultiRegionTrail:         aws.Bool(false),
					},
				},
				testName: "TestCheckIfCloudtrailsMultiRegion",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfCloudtrailsMultiRegion(tt.args.checkConfig, tt.args.cloudtrails, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfCloudtrailsMultiRegion() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
