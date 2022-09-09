package cloudfront

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/stangirard/yatas/plugins/commons"
)

func TestCheckIfStandardLogginEnabled(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		d           []SummaryToConfig
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfStandardLogginEnabled",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				d: []SummaryToConfig{
					{
						summary: types.DistributionSummary{
							Id: aws.String("test"),
						},
						config: types.DistributionConfig{
							Logging: &types.LoggingConfig{
								Enabled: aws.Bool(true),
							},
						},
					},
				},
				testName: "TestCheckIfStandardLogginEnabled",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfStandardLogginEnabled(tt.args.checkConfig, tt.args.d, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfStandardLogginEnabled() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfStandardLogginEnabledFail(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		d           []SummaryToConfig
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfStandardLogginEnabled",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				d: []SummaryToConfig{
					{
						summary: types.DistributionSummary{
							Id: aws.String("test"),
						},
						config: types.DistributionConfig{
							Logging: &types.LoggingConfig{},
						},
					},
				},
				testName: "TestCheckIfStandardLogginEnabled",
			},
		},
		{
			name: "TestCheckIfStandardLogginEnabled",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				d: []SummaryToConfig{
					{
						summary: types.DistributionSummary{
							Id: aws.String("test"),
						},
						config: types.DistributionConfig{
							Logging: &types.LoggingConfig{
								Enabled: aws.Bool(false),
							},
						},
					},
				},
				testName: "TestCheckIfStandardLogginEnabled",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfStandardLogginEnabled(tt.args.checkConfig, tt.args.d, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfStandardLogginEnabled() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
