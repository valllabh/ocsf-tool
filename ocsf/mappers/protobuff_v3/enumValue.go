package protobuff_v3

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (ev *EnumValue) Marshal() string {
	content := []string{}

	// NOTE: some OCSF Values have non-Zero UNKNOKWN values, while protos want a zero value unknown.
	// <code>class_uid * 100 + activity_id</code>. -> These are OCSF values, and we still emit a UNSPECIFRIED=0 enum
	// value
	baseName := ev.enum.Name + " " + ev.Name
	content = append(content, ToEnumValueName(baseName))

	content = append(content, fmt.Sprintf("= %d;", ev.Value))

	if len(ev.Comment) > 0 {
		content = append(content, "//")
		for k, v := range ev.Comment {
			content = append(content, fmt.Sprintf("%s: %s;", k, v))
		}
	}

	return strings.Join(content, " ")
}

func ToEnumValueName(input string) string {

	// Return if Cache exists
	value, exists := GetMapper().Cache.EnumValues.Get(input)

	if exists {
		return value.(string)
	}

	output := input

	// Apply Name Processor
	if GetMapper().Preprocessor.EnumName != nil {
		output = GetMapper().Preprocessor.EnumValueName(input)
	}

	// Clean Name
	output = cleanName(output)
	output = strcase.ToScreamingSnake(output)

	// Set Cache
	GetMapper().Cache.EnumValues.Set(input, output)

	return output
}
