package commands

import (
	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tool/config"
)

// Define the root command
var rootCmd = &cobra.Command{
	Use: "ocsf-tool",

	PostRun: func(cmd *cobra.Command, args []string) {
		// Write the config file to disk
		config.WriteConfig()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		println("--")
	},
}

// Initialize the root command
func init() {

}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
