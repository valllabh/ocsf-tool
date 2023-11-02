package commands

import "github.com/spf13/cobra"

// Define the root command
var rootCmd = &cobra.Command{Use: "ocsf-tool"}

// Initialize the root command
func init() {
	// Add the generate command to the root command
	rootCmd.AddCommand(GenerateProtoCmd)
	rootCmd.AddCommand(SchemaClassListCmd)
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
