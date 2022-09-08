package s3

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stangirard/yatas/config"
)

func TestCheckIfBucketInOneZone(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		buckets     BucketAndNotInRegion
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Check if S3 buckets are in one zone",
			args: args{
				checkConfig: config.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan config.Check, 1),
				},
				buckets: BucketAndNotInRegion{
					Buckets: []types.Bucket{
						{
							Name: aws.String("test"),
						},
					},
					NotInRegion: []types.Bucket{
						{
							Name: aws.String("toto"),
						},
					},
				},
				testName: "AWS_S3_001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfBucketInOneZone(tt.args.checkConfig, tt.args.buckets, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfBucketInOneZone() = %v", check)
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfBucketInOneZoneFail(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		buckets     BucketAndNotInRegion
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Check if S3 buckets are in one zone",
			args: args{
				checkConfig: config.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan config.Check, 1),
				},
				buckets: BucketAndNotInRegion{
					Buckets: []types.Bucket{
						{
							Name: aws.String("test"),
						},
					},
					NotInRegion: []types.Bucket{
						{
							Name: aws.String("test"),
						},
					},
				},
				testName: "AWS_S3_001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfBucketInOneZone(tt.args.checkConfig, tt.args.buckets, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfBucketInOneZone() = %v", check)
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
