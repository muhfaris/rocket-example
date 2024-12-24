package cmd

import (
	"fmt"
	"os"

	"github.com/muhfaris/rocket-hexagonal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "rocket-hexagonal",
	RunE: mainRunE,
}

func readConfig() {
	err := config.LoadConfig()
	if err != nil {
		return
	}
}

func mainRunE(cmd *cobra.Command, args []string) error {
	return nil
}

// Execute is root function
func Execute() {
	cobra.OnInitialize(readConfig)
	rootCmd.AddCommand(restCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
