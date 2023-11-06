package commands

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-tool/ocsf/mappers/protobuff_v3"
	"github.com/valllabh/ocsf-tool/ocsf/schema"
)

// Define the GenerateProtoCmd command
var GenerateProtoCmd = &cobra.Command{
	Use:     "generate-proto [ocsf_class_name]...",
	Short:   "Generate a Proto file",
	Example: "ocsf-tool generate-proto file_activity process_activity",
	Long:    "Generate a Proto file for the specified OCSF classes.\nUse the `ocsf-tool schema-class-list` command to see a list of all OCSF classes.",
	Args:    cobra.MinimumNArgs(1),
	Run:     runGenerateProtoCmd,
}

// Initialize the GenerateProtoCmd command
func init() {
	// Add flags to the GenerateProtoCmd command

	// Specifies the output directory for the Proto file.
	// The default value is `./output/proto`.
	GenerateProtoCmd.Flags().StringP("proto-output", "", "./output/proto", "Output directory for the Proto file")

	// Specifies the base package for the Proto file.
	// The default value is `ocsf`.
	GenerateProtoCmd.Flags().StringP("proto-root-package", "", "ocsf", "Base package for the Proto file")

	// Specifies the Golang package prefix.
	// The default value is `github.com/your-project/generated/golang/`.
	GenerateProtoCmd.Flags().StringP("golang-root-package", "", "github.com/your-project/generated/golang/", "Golang package prefix")
}

// Define the run function for the GenerateProtoCmd command
func runGenerateProtoCmd(cmd *cobra.Command, args []string) {
	var errors error
	protoOutput, _ := cmd.Flags().GetString("proto-output")
	protoRootPackage, _ := cmd.Flags().GetString("proto-root-package")
	golangPackageName, _ := cmd.Flags().GetString("golang-root-package")

	ocsfSchema, _ := schema.LoadOCSFSchema()
	events := []schema.Event{}
	mapper := protobuff_v3.NewMapper(ocsfSchema)

	// Validate Base Package and Output Path
	if len(protoRootPackage) > 0 && len(protoOutput) > 0 {
		rootPackage := protobuff_v3.NewPackage(protoRootPackage, nil)
		rootPackage.Path = protoOutput
		mapper.RootPackage = protobuff_v3.NewPackage(toVersionPackage(strcase.ToSnake(ocsfSchema.Version)), rootPackage)
	}

	if len(golangPackageName) > 0 {
		mapper.Preprocessor.GolangPackageName = func(name string) string {
			return golangPackageName + name
		}
	}

	// Validate and Load OCSF Event
	for _, eventName := range args {
		event, exists := ocsfSchema.Classes[eventName]
		if exists {
			events = append(events, event)
		} else {
			// Error if invalid OCSF Class is passed
			errors = multierror.Append(errors, fmt.Errorf("[%s] is not a valid OCSF Class", eventName))
		}
	}

	// Throw Error
	if errors != nil {
		fmt.Println(errors.Error())
		return
	}

	// Produce Output
	mapper.Marshal(events)

	// Print a console message indicating where the output is generated
	fmt.Printf("Proto files are generated in %s\n", protoOutput)
}

// Convert version to package name
func toVersionPackage(version string) string {
	return fmt.Sprintf("v%s", version)
}
