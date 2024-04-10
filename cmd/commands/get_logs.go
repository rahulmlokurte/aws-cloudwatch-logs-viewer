package commands

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/rahulmlokurte/aclv/pkg/config"
	"github.com/spf13/cobra"
)

func getLogEventsCmd() *cobra.Command {

	getLogEvents := &cobra.Command{
		Use:     "get-logs",
		Short:   "Get the logs given a log group name",
		Aliases: []string{"gl"},
		Run:     getLogs,
	}
	return getLogEvents
}

func getLogs(cmd *cobra.Command, args []string) {
	awsconfig := config.AwsLogin()

	var groupName = args[0]
	var logStream = args[1]
	//var limit int32 = 50
	fmt.Println(groupName)
	svc := cloudwatchlogs.NewFromConfig(awsconfig)
	input := cloudwatchlogs.GetLogEventsInput{
		LogStreamName: &logStream,
		LogGroupName:  &groupName,
	}
	resp, err := svc.GetLogEvents(context.Background(), &input)
	if err != nil {
		fmt.Println("Got error getting log events:")
		fmt.Println(err)
		return
	}

	fmt.Println("Event messages for stream " + *input.LogStreamName)
	gotToken := ""
	nextToken := ""

	for _, event := range resp.Events {
		gotToken = nextToken
		nextToken = *resp.NextForwardToken
		if gotToken == nextToken {
			break
		}
		fmt.Println(" ", *event.Message)
	}
}
