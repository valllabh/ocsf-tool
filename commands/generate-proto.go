package commands

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff_v3"
	"github.com/valllabh/ocsf-schema-processor/ocsf/schema"
)

var GenerateProtoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Generate a Proto file",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func init() {
	GenerateProtoCmd.Flags().StringP("proto-output", "", "./output/proto", "Output directory for the Proto file")
	GenerateProtoCmd.Flags().StringP("proto-root-package", "", "ocsf", "Base package for the Proto file")
	GenerateProtoCmd.Flags().StringP("golang-root-package", "", "github.com/your-project/generated/golang/", "Golang package prefix")
}

func run(cmd *cobra.Command, args []string) {
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
		mapper.RootPackage = protobuff_v3.NewPackage(strcase.ToSnake(ocsfSchema.Version), rootPackage)
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
	println("Done")
}
