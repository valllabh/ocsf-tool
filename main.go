package main

import (
	"github.com/valllabh/ocsf-schema-processor/ocsf"
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff_v3"
)

func main() {

	// Loads to provided OCSF schema version in schema.json
	ocsfSchema, _ := ocsf.LoadOCSFSchema()

	mapToProtoFile(ocsfSchema)
}

func mapToProtoFile(ocsfSchema ocsf.OCSFSchema) {

	protobuff_v3.InitMapper(ocsfSchema, "./output/proto")

	events := []ocsf.Event{ocsfSchema.Classes["file_activity"]}

	protobuff_v3.Mapper().Marshal(events)

}
