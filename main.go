package main

import (
	"fmt"
	"os"

	"github.com/valllabh/ocsf-tool/commands"
)

// Main function
func main() {
	rootCmd := commands.GetRootCmd()
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
