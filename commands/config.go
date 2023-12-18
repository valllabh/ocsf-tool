package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valllabh/ocsf-tool/commons"
	"github.com/valllabh/ocsf-tool/config"
)

type ConfigVariable struct {
	variable string
	help     string
	example  string
	type_    string
}

var CONFIGURATIONS map[string]ConfigVariable = map[string]ConfigVariable{
	"extensions.selected": ConfigVariable{
		variable: "extensions.selected",
		help:     "OCSF Extensions",
		example:  "linux win",
		type_:    "[]string",
	},
	"profiles.selected": ConfigVariable{
		variable: "profiles.selected",
		help:     "OCSF Profiles",
		example:  "cloud linux/linux_users",
		type_:    "[]string",
	},
}

// Define the setConfig command
var configCmd = &cobra.Command{
	Use:   `config <extensions|profiles> values...`,
	Short: `Set configuration values for extensions and profiles`,
	Example: `
 ocsf-tool config extensions linux win
 ocsf-tool config profiles cloud linux/linux_users
 ocfs-tool config schema.loading.strategy repository
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

		configMeta, configError := getConfig(variable)

		// validate variable
		if configError != nil {
			println("Invalid config variable:", variable)
			// print help
			cmd.Help()
			return
		}

		// switch on the type of the config variable and set the value
		switch configMeta.type_ {
		case "[]string":
			viper.Set(variable, values)
		case "string":
			viper.Set(variable, values[0])
		}

		// Write the config file to disk
		config.WriteConfig()
	},
}

// getConfig returns ConfigVariable of the given config variable or error if the config variable is invalid
func getConfig(variable string) (ConfigVariable, error) {
	// validate variable
	if !isValidConfigVariable(variable) {
		return ConfigVariable{}, errors.New("Invalid config variable: " + variable)
	}

	return CONFIGURATIONS[variable], nil
}

// getValidConfigVariables returns a list of valid config variables
func getValidConfigVariables() []string {
	return commons.GetMapKeys(CONFIGURATIONS)
}

// isValidConfigVariable checks if the given variable is a valid config variable
func isValidConfigVariable(variable string) bool {
	return commons.Contains(getValidConfigVariables(), variable)
}

func init() {
	// Add the setConfig command to the root command
	rootCmd.AddCommand(configCmd)
}
