package v3

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (m *Message) AddField(field *Field) {
	if m.fields == nil {
		m.fields = Fields{}
	}
	field.message = m
	m.fields = append(m.fields, field)
}

func (m *Message) Marshal() string {
	content := []string{}
	content = append(content, "// "+m.GroupKey)
	content = append(content, fmt.Sprintf("message %s {", m.proto.ToMessageName(m.Name)))
	for i, f := range m.fields {
		content = append(content, "\t"+f.Marshal(i+1))
	}
	content = append(content, "}")
	return strings.Join(content, "\n")
}

func (p *Proto) ToMessageName(input string) string {

	// Return if Cache exists
	value, exists := p.cache.Messages.Get(input)

	if exists {
		return fmt.Sprint(value)
	}

	output := input

	// Apply Name Processor
	if p.Preprocessor.MessageName != nil {
		output = p.Preprocessor.MessageName(input)
	}

	// Clean Name
	output = cleanName(output)
	output = strcase.ToCamel(output)

	// Set Cache
	p.cache.Messages.Set(input, output)

	return output
}
