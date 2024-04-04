package commands

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"strings"
)

var (
	startsWith string
)

func savedQueriesCmd() *cobra.Command {
	savedQueries := &cobra.Command{
		Use:     "saved-queries",
		Short:   "List saved queries",
		Aliases: []string{"sq"},
		Run:     listSavedQueries,
	}
	savedQueries.Flags().StringVarP(&startsWith, "startsWith", "s", "", "List saved queries starting with this string")
	return savedQueries
}

func listSavedQueries(cmd *cobra.Command, args []string) {
	folderName, _ := cmd.Flags().GetString("startsWith")
	awsConfig := awsLogin()
	result := queryCloudWatchNames(awsConfig, folderName)
	queryCloudWatchNamesList(result)
}

func queryCloudWatchNamesList(queryDefinitionsOutput *cloudwatchlogs.DescribeQueryDefinitionsOutput) {
	folders := make(map[string]string)
	for index, queries := range queryDefinitionsOutput.QueryDefinitions {
		queryDefinitionName := *queries.Name
		folders[fmt.Sprintf("%d", index)] = queryDefinitionName
		fmt.Printf("%d. %s\n", index, queryDefinitionName)
	}
	promptUser(folders, queryDefinitionsOutput)

}

func promptUser(folders map[string]string, output *cloudwatchlogs.DescribeQueryDefinitionsOutput) {
	var selectedQueryDefinitionName string
	fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#eed49f")).PaddingTop(1).Italic(true).Blink(true).Render("Select the query number: "))
	_, err := fmt.Scanln(&selectedQueryDefinitionName)
	if err != nil {
		return
	}
	// Retrieve queries within the selected folder
	selectedFolderName, ok := folders[selectedQueryDefinitionName]
	if !ok {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).PaddingTop(1).Italic(true).Strikethrough(true).Blink(true).Render("Invalid folder index."))
		return
	}

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a6da95")).Padding(1).Italic(true).Render("You have selected " + selectedFolderName))
	for _, q := range output.QueryDefinitions {
		if strings.HasPrefix(*q.Name, selectedFolderName) {
			fmt.Println(*q.QueryString)
		}
	}

}

func queryCloudWatchNames(cfg aws.Config, queryDefinitionName string) *cloudwatchlogs.DescribeQueryDefinitionsOutput {
	svc := cloudwatchlogs.NewFromConfig(cfg)
	var input *cloudwatchlogs.DescribeQueryDefinitionsInput
	if len(queryDefinitionName) > 0 {
		input = &cloudwatchlogs.DescribeQueryDefinitionsInput{QueryDefinitionNamePrefix: &queryDefinitionName}
	} else {
		input = &cloudwatchlogs.DescribeQueryDefinitionsInput{}
	}
	result, err := svc.DescribeQueryDefinitions(context.TODO(), input)
	if err != nil {
		fmt.Println("Error listing saved queries:", err)
		return nil
	}
	return result
}

func awsLogin() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		return aws.Config{}
	}
	return cfg
}
