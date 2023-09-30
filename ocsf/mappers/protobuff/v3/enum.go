package v3

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (e *Enum) AddValue(value *EnumValue) {
	if e.values == nil {
		e.values = EnumValues{}
	}
	value.enum = e
	e.values[value.Name] = value
}

func (e *Enum) GetValue(name string) (*EnumValue, bool) {
	value, exists := e.values[name]
	return value, exists
}

func (e *Enum) Marshal() string {
	content := []string{}
	content = append(content, fmt.Sprintf("enum %s {", e.proto.ToEnumName(e.Name)))
	i := 0
	for _, v := range e.values {
		content = append(content, "\t"+v.Marshal(i))
		i++
	}
	content = append(content, "}")
	return strings.Join(content, "\n")
}

func (p *Proto) ToEnumName(input string) string {

	// Return if Cache exists
	value, exists := p.cache.Enums.Get(input)

	if exists {
		return fmt.Sprint(value)
	}

	output := input

	// Apply Name Processor
	if p.Preprocessor.EnumName != nil {
		output = p.Preprocessor.EnumName(input)
	}

	// Clean Name
	output = cleanName(output)
	output = strcase.ToScreamingSnake(output)

	// Set Cache
	p.cache.Enums.Set(input, output)

	return output
}
