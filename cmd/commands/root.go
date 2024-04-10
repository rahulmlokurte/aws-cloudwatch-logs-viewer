package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "aclv",
		Short: "An utility commandline for AWS cloudwatch",
	}

	rootCmd.AddCommand(savedQueriesCmd())
	rootCmd.AddCommand(logGroupsCmd())
	rootCmd.AddCommand(getLogEventsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
