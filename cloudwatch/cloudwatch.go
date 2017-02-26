package cloudwatch

import (
	"fmt"

	"github.com/lucagrulla/cloudwatch-tail/timeutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	//	"github.com/fatih/color"
)

func cwClient() *cloudwatchlogs.CloudWatchLogs {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	return cloudwatchlogs.New(sess, aws.NewConfig().WithRegion("eu-west-1"))
}

func params(logGroupName string, streamName string, epochStartTime int64) *cloudwatchlogs.FilterLogEventsInput {
	startTimeInt64 := epochStartTime * 1000
	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: &logGroupName,
		Interleaved:  aws.Bool(true),
		StartTime:    &startTimeInt64}

	if streamName != "" {
		params.LogStreamNames = []*string{aws.String(streamName)}
	}
	return params
}

func Tail(startTime *string, follow *bool, logGroupName *string, streamName *string) {
	cwl := cwClient()
	lastTimestamp := timeutil.ParseTime(*startTime).Unix()
	pageHandler := func(res *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) bool {
		for _, event := range res.Events {
			lastTimestamp = *event.Timestamp / 1000
			d := timeutil.FormatTimestamp(lastTimestamp)
			fmt.Printf("%s - %s\n", d, *event.Message)
		}
		return true
	}

	for *follow || (lastTimestamp == timeutil.ParseTime(*startTime).Unix()) {
		logParam := params(*logGroupName, *streamName, lastTimestamp)
		error := cwl.FilterLogEventsPages(logParam, pageHandler)

		if error != nil {
			panic(error)

		}
	}
}

func Ls() {
	cwl := cwClient()
	params := &cloudwatchlogs.DescribeLogGroupsInput{
	//		LogGroupNamePrefix: aws.String("LogGroupName"),
	}

	handler := func(res *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
		for _, logGroup := range res.LogGroups {
			fmt.Println(*logGroup.LogGroupName)
		}
		return true
	}
	err := cwl.DescribeLogGroupsPages(params, handler)
	if err != nil {
		panic(err)
	}
}