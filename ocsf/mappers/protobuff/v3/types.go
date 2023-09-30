package v3

import (
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff/commons"
)

type Messages map[string]*Message
type Fields []*Field
type Enums map[string]*Enum
type EnumValues []*EnumValue

type FieldType int16

const (
	FIELD_TYPE_OBJECT    FieldType = 100
	FIELD_TYPE_PRIMITIVE FieldType = 110
	FIELD_TYPE_ENUM      FieldType = 120
)

type Preprocessor struct {
	MessageName   func(string) string
	EnumName      func(string) string
	EnumValueName func(string) string
}

type CacheMap struct {
	Messages   commons.Cache
	Enums      commons.Cache
	EnumValues commons.Cache
}

type Proto struct {
	Preprocessor Preprocessor
	messages     Messages
	enums        Enums
	cache        CacheMap
}

type Message struct {
	Name     string
	fields   Fields
	proto    *Proto
	GroupKey string
}

type Field struct {
	Name     string
	DataType string
	Required bool
	Repeated bool
	Type     FieldType
	message  *Message
	Comment  commons.Comment
}

type Enum struct {
	Name   string
	values EnumValues
	proto  *Proto
}

type EnumValue struct {
	Name    string
	Value   int64
	Comment commons.Comment
	enum    *Enum
}
