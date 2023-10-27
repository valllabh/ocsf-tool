package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-schema-processor/commands"
)

var rootCmd = &cobra.Command{Use: "ocsf-schema"}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.GenerateCmd)
}
