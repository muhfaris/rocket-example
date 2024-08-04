package cmd

import (
	"fmt"
	"os"

	"github.com/muhfaris/rocket-example/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "rocket-example",
	PreRunE: preRunE,
	RunE:    mainRunE,
}

func preRunE(cmd *cobra.Command, args []string) error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	return nil
}

func mainRunE(cmd *cobra.Command, args []string) error {
	return nil
}

// Execute is root function
func Execute() {
	rootCmd.AddCommand(restCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
