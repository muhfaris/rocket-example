package cmd

import (
	"github.com/muhfaris/rocket-hexagonal/config"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers"
	"github.com/muhfaris/rocket-hexagonal/internal/app"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:     "rest",
	Short:   "Run rest server",
	PreRunE: restPreRunE,
	RunE:    restRunE,
}

func restPreRunE(cmd *cobra.Command, args []string) error {
	app.InitializeRepository()
	return nil
}

func restRunE(cmd *cobra.Command, args []string) error {
	api := routers.Init(config.App.Port)
	err := api.Run()
	if err != nil {
		return err
	}

	return nil
}
