package protobuff_v3

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (ev *EnumValue) Marshal() string {
	content := []string{}

	content = append(content, ToEnumValueName(ev.enum.Name+" "+ev.Name))

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
		return fmt.Sprint(value)
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
