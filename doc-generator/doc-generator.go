package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/doc"
	"github.com/valllabh/ocsf-tool/commands"
	"github.com/valllabh/ocsf-tool/commons"
)

func main() {

	// Declare variables
	var err error
	var docsPath = "./docs/"

	// Remove docsPath if it exists
	err = os.RemoveAll(docsPath)

	// Handle error
	if err != nil {
		fmt.Println("Error in removing directory:", err)
		os.Exit(1)
	}

	// Create directory if it does not exist
	err = commons.EnsureDirExists(docsPath)

	// Handle error
	if err != nil {
		fmt.Println("Error in creating directory:", err)
		os.Exit(1)
	}

	// Generate documentation
	err = doc.GenMarkdownTree(commands.GetRootCmd(), docsPath)

	// Handle error
	if err != nil {
		fmt.Println("Error in generating documentation:", err)
		os.Exit(1)
	}
}
