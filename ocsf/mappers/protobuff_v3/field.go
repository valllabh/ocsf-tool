package protobuff_v3

import (
	"fmt"
	"strings"
)

func (f *Field) Marshal(index int) string {
	content := []string{}

	// Required option is removed from proto3
	// TODO: confirm alternative for optional/required
	// if f.Required {
	// 	content = append(content, "required")
	// }
	if f.Repeated {
		content = append(content, "repeated")
	}

	switch f.Type {
	case FIELD_TYPE_OBJECT:
		m, _ := GetMessage(f.DataType)
		content = append(content, m.GetReference())
	case FIELD_TYPE_PRIMITIVE:
		content = append(content, f.DataType)
	case FIELD_TYPE_ENUM:
		e, _ := GetEnum(f.message.Name + " " + f.Name)
		content = append(content, e.GetReference())
	}

	content = append(content, f.Name)
	content = append(content, fmt.Sprintf("= %d;", index))

	if len(f.Comment) > 0 {
		content = append(content, "//")
		for k, v := range f.Comment {
			content = append(content, fmt.Sprintf("%s: %s;", k, v))
		}
	}

	return strings.Join(content, " ")
}
