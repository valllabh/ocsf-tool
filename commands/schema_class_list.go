package commands

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tool/commons"
	"github.com/valllabh/ocsf-tool/ocsf/schema"
)

// Define SchemaClassList Command
var SchemaClassListCmd = &cobra.Command{
	Use:   "schema-class-list",
	Short: "List all classes in the OCSF schema",
	Args:  cobra.MinimumNArgs(0),
	Run:   runSchemaClassListCmd,
}

func init() {
	// Add flags to the SchemaClassListCmd command
	rootCmd.AddCommand(SchemaClassListCmd)

	// Add flag to output json file path
	SchemaClassListCmd.Flags().StringP("output", "", "", "Output JSON file path")

}

// Define the run function for the SchemaClassListCmd command
func runSchemaClassListCmd(cmd *cobra.Command, args []string) {

	ocsfSchema := schema.LoadOCSFSchema()
	classes := ocsfSchema.Classes

	// get output file path
	outputFilePath := cmd.Flag("output").Value.String()

	// Group classes by Category
	classesByCategory := make(map[string][]string)
	for _, class := range classes {
		classesByCategory[class.Category] = append(classesByCategory[class.Category], class.Name)
	}

	if outputFilePath != "" {
		contents, _ := json.Marshal(classesByCategory)
		// Write classes by Category to output file
		outputFilePathError := commons.CreateFile(outputFilePath, contents)

		if outputFilePathError != nil {
			println("Error writing to output file: " + outputFilePath)
		} else {
			println("Classes by Category written to output file: " + outputFilePath)
		}
		return
	}

	// Print classes by Category
	for category, classes := range classesByCategory {
		println("\n" + category + ":")
		for _, class := range classes {
			println("    " + class)
		}
	}
}
