package commands

import (
	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tool/config"
)

// Define the root command
var rootCmd = &cobra.Command{
	Use: "ocsf-tool",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		if err := config.InitConfig(configFile); err != nil {
			return err
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		println("--")
	},
}

// Initialize the root command
func init() {
	rootCmd.PersistentFlags().String("config", "config.yaml", "config file path")
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
