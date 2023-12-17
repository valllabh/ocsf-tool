package commands

import (
	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tool/ocsf/schema"
)

// Define SchemaClassList Command
var SchemaClassListCmd = &cobra.Command{
	Use:   "schema-class-list",
	Short: "List all classes in the OCSF schema",
	Args:  cobra.MinimumNArgs(0),
	Run:   runSchemaClassListCmd,
}

// Define the run function for the SchemaClassListCmd command
func runSchemaClassListCmd(cmd *cobra.Command, args []string) {

	ocsfSchema := schema.LoadOCSFSchema()
	classes := ocsfSchema.Classes

	// Group classes by Category
	classesByCategory := make(map[string][]string)
	for _, class := range classes {
		classesByCategory[class.Category] = append(classesByCategory[class.Category], class.Name)
	}

	// Print classes by Category
	for category, classes := range classesByCategory {
		println("\n" + category + ":")
		for _, class := range classes {
			println("    " + class)
		}
	}
}

func init() {
	// Add flags to the SchemaClassListCmd command
	rootCmd.AddCommand(SchemaClassListCmd)
}
