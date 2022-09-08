package cloudtrail

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/stangirard/yatas/config"
)

func TestCheckIfCloudtrailsEncrypted(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		cloudtrails []types.Trail
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfCloudtrailsEncrypted",
			args: args{
				checkConfig: config.CheckConfig{Queue: make(chan config.Check, 1), Wg: &sync.WaitGroup{}},
				cloudtrails: []types.Trail{
					{
						Name:     aws.String("test"),
						KmsKeyId: aws.String("test"),
						TrailARN: aws.String("test"),
					},
				},
				testName: "TestCheckIfCloudtrailsEncrypted",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfCloudtrailsEncrypted(tt.args.checkConfig, tt.args.cloudtrails, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfCloudtrailsEncrypted() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfCloudtrailsEncryptedFail(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		cloudtrails []types.Trail
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfCloudtrailsEncrypted",
			args: args{
				checkConfig: config.CheckConfig{Queue: make(chan config.Check, 1), Wg: &sync.WaitGroup{}},
				cloudtrails: []types.Trail{
					{
						Name:     aws.String("test"),
						TrailARN: aws.String("test"),
					},
				},
				testName: "TestCheckIfCloudtrailsEncrypted",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfCloudtrailsEncrypted(tt.args.checkConfig, tt.args.cloudtrails, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfCloudtrailsEncrypted() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
