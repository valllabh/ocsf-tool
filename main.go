package main

import (
	"fmt"
	"os"

	"github.com/valllabh/ocsf-tool/commands"
)

// Main function
func main() {
	// Execute the root command
	if err := commands.GetRootCmd().Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
