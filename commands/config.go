package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valllabh/ocsf-tool/config"
)

// Define the setConfig command
var configCmd = &cobra.Command{
	Use:   `config <extensions|profiles> values...`,
	Short: `Set configuration values for extensions and profiles`,
	Example: `
 ocsf-tool config extensions linux win
 ocsf-tool config profiles cloud linux/linux_users
	`,
	Long: `
 Set configuration values for extensions and profiles
 Possible values for Extensions https://schema.ocsf.io/1.0.0/
 Possible values for Profiles https://schema.ocsf.io/1.0.0/profiles/
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the number of args is less than 2
		if len(args) < 2 {
			println("Invalid number of arguments")
			// print help
			cmd.Help()
			return
		}

		// Extract the variable and values from the args
		variable, values := args[0], args[1:]

		// validate variable
		if variable != "extensions" && variable != "profiles" {
			println("Invalid variable name. Possible values are [extensions, profiles].")
			return
		}

		// set the config value
		viper.Set(variable, values)

		// Write the config file to disk
		config.WriteConfig()
	},
}

func init() {
	// Add the setConfig command to the root command
	rootCmd.AddCommand(configCmd)
}
