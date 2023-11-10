package main

import (
	"fmt"
	"os"

	"github.com/valllabh/ocsf-tool/commands"
	"github.com/valllabh/ocsf-tool/config"
)

// Main function
func main() {

	// Initialize the config
	config.InitConfig()

	// Execute the root command
	if err := commands.GetRootCmd().Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
