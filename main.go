package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

const (
	folder_name = "query_folder_name"
)

var rootCmd = &cobra.Command{
	Use:   "aclv",
	Short: "A CLI app to view AWS CloudWatch Logs",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(logGroupsCmd)
	rootCmd.AddCommand(savedQueriesCmd)
}

var logGroupsCmd = &cobra.Command{
	Use:   "log-groups",
	Short: "List CloudWatch log groups",
	Run: func(cmd *cobra.Command, args []string) {
		listLogGroups()
	},
}

var savedQueriesCmd = &cobra.Command{
	Use:   "saved-queries",
	Short: "List CloudWatch saved queries",
	Run: func(cmd *cobra.Command, args []string) {
		listSavedQueries()
	},
}

func listLogGroups() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		return
	}
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

func listSavedQueries() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		return
	}

	svc := cloudwatchlogs.NewFromConfig(cfg)

	input := &cloudwatchlogs.DescribeQueryDefinitionsInput{}

	result, err := svc.DescribeQueryDefinitions(context.TODO(), input)
	if err != nil {
		fmt.Println("Error listing saved queries:", err)
		return
	}

	var output strings.Builder
	output.WriteString(strings.TrimSpace(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00")).Render("Saved Queries in " + folder_name + " folder:\n")))
	for _, queries := range result.QueryDefinitions {
		if strings.HasPrefix(*queries.Name, folder_name) {
			styledName := lipgloss.NewStyle().Foreground(lipgloss.Color("#36C5F0")).PaddingLeft(4).Italic(true).Render(*queries.Name)
			output.WriteString(styledName + "\n")
		}

	}
	fmt.Print(output.String())

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
