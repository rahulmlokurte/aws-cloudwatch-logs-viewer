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

var folder_name string = "replace_folder_to_search"

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ea76cb")).
	PaddingTop(1).
	Italic(true).
	PaddingLeft(4)

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

	input := &cloudwatchlogs.DescribeQueryDefinitionsInput{
		QueryDefinitionNamePrefix: &folder_name,
	}

	result, err := svc.DescribeQueryDefinitions(context.TODO(), input)
	if err != nil {
		fmt.Println("Error listing saved queries:", err)
		return
	}

	folders := make(map[string]string)

	for index, queries := range result.QueryDefinitions {
		if strings.HasPrefix(*queries.Name, folder_name) {
			folderName := *queries.Name
			folders[fmt.Sprintf("%d", index+1)] = folderName
			fmt.Printf("%d. %s\n", index+1, folderName)
		}

	}

	// Prompt user to select a folder
	var selectedFolder string
	fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#eed49f")).PaddingTop(1).Render("Select the query: "))
	fmt.Scanln(&selectedFolder)

	// Retrieve queries within the selected folder
	selectedFolderName, ok := folders[selectedFolder]
	if !ok {
		fmt.Println("Invalid folder index.")
		return
	}

	fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#a6da95")).Padding(1).Italic(true).Render("You have selected " + selectedFolderName))

	for _, q := range result.QueryDefinitions {
		if strings.HasPrefix(*q.Name, selectedFolderName) {
			fmt.Println(style.Render(*q.QueryString))
		}
	}

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
