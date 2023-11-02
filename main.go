package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tools/commands"
)

// Define the root command
var rootCmd = &cobra.Command{Use: "ocsf-schema"}

// Main function
func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

// Initialize the root command
func init() {
	// Add the generate command to the root command
	rootCmd.AddCommand(commands.GenerateCmd)
}
