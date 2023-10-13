package protobuff_v3

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
	if len(m.Comment) > 0 {
		for k, v := range m.Comment {
			content = append(content, fmt.Sprintf("// %s: %s", k, v))
		}
	}

	content = append(content, fmt.Sprintf("message %s {", ToMessageName(m.Name)))
	for i, f := range m.fields {
		content = append(content, "\t"+f.Marshal(i+1))
	}
	content = append(content, "}")
	return strings.Join(content, "\n")
}

func ToMessageName(input string) string {

	// Return if Cache exists
	value, exists := Mapper().Cache.Messages.Get(input)

	if exists {
		return fmt.Sprint(value)
	}

	output := input

	// Apply Name Processor
	if Mapper().Preprocessor.MessageName != nil {
		output = Mapper().Preprocessor.MessageName(input)
	}

	// Clean Name
	output = cleanName(output)
	output = strcase.ToCamel(output)

	// Set Cache
	Mapper().Cache.Messages.Set(input, output)

	return output
}

func (m *Message) GetName() string {
	return ToMessageName(m.Name)
}

func (m *Message) GetReference() string {
	return m.GetPackage() + "." + m.GetName()
}

func (m *Message) GetPackage() string {
	return m.Package.GetFullName()
}

func (m *Message) GetImports() Imports {

	imports := Imports{}

	for _, f := range m.fields {
		p := ""
		switch f.Type {
		case FIELD_TYPE_OBJECT:
			m, _ := GetMessage(f.DataType)
			p = m.Package.Proto.GetProtoFilePath()
		case FIELD_TYPE_ENUM:
			e, _ := GetEnum(f.message.Name + " " + f.Name)
			p = e.Package.Proto.GetProtoFilePath()
		}

		if len(p) != 0 {
			imports[p] = &Import{
				Name: p,
			}
		}
	}

	return imports
}
