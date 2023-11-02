package commands

import (
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate",
	Run: func(cmd *cobra.Command, args []string) {
		println("Use subcommands to generate files")
	},
}

func init() {
	GenerateCmd.AddCommand(GenerateProtoCmd)
}
