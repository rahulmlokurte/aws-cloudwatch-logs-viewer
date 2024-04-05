package commands

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/rahulmlokurte/aclv/pkg/config"
	"github.com/spf13/cobra"
)

func logGroupsCmd() *cobra.Command {
	logGroups := &cobra.Command{
		Use:     "log-groups",
		Short:   "List Logged Groups",
		Aliases: []string{"lg"},
		Run:     listLogGroups,
	}
	return logGroups
}

func listLogGroups(cmd *cobra.Command, args []string) {
	awsConfig := config.AwsLogin()
	queryCloudWatchLogGroups(awsConfig)
}

func queryCloudWatchLogGroups(cfg aws.Config) {
	svc := cloudwatchlogs.NewFromConfig(cfg)
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	result, err := svc.DescribeLogGroups(context.TODO(), input)
	if err != nil {
		fmt.Println("Error listing log groups:", err)
		return
	}
	fmt.Println("Log Groups")
	for _, lg := range result.LogGroups {
		fmt.Println(*lg.LogGroupName)
	}
}
