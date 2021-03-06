// Package cloudwatch provides primitives to interact with Cloudwatch logs
package cloudwatch

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

//CW provides the APIo peration methods for making requests to AWS cloudwatch logs.
type CW struct {
	awsClwClient *cloudwatchlogs.CloudWatchLogs
	debug        *bool
}

// New creates a new instance of the CW client
func New(awsProfile *string, awsRegion *string, d *bool) *CW {
	if *d {
		fmt.Printf("awsProfile: %s, awsRegion: %s\n", *awsProfile, *awsRegion)
	}

	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}

	if awsProfile != nil {
		opts.Profile = *awsProfile
	}

	if awsRegion != nil {
		opts.Config = aws.Config{Region: awsRegion}
	}

	sess := session.Must(session.NewSessionWithOptions(opts))
	return &CW{awsClwClient: cloudwatchlogs.New(sess),
		debug: d}
}
