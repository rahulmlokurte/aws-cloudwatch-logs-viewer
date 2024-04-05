package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "aclv",
		Short: "An utility commandline for AWS cloudwatch",
	}

	rootCmd.AddCommand(savedQueriesCmd())
	rootCmd.AddCommand(logGroupsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
