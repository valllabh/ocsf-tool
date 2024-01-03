package protobuff_v3

import (
	"github.com/spf13/afero"
	"github.com/valllabh/ocsf-tool/commons"
	"github.com/valllabh/ocsf-tool/ocsf/schema"
)

type Comment map[string]string

type Messages map[string]*Message
type Fields []*Field
type Enums map[string]*Enum
type EnumValues map[string]*EnumValue

type FieldType int16

const (
	FIELD_TYPE_OBJECT    FieldType = 100
	FIELD_TYPE_PRIMITIVE FieldType = 110
	FIELD_TYPE_ENUM      FieldType = 120
)

type Pkg struct {
	Name     string
	Children Pkgs
	Parent   *Pkg
	Proto    *Proto
	Path     string
}

type Pkgs map[string]*Pkg

type Proto struct {
	Pkg *Pkg
	// messages Messages
	// enums    Enums
}

type Import struct {
	Name string
}

type Imports map[string]*Import

type Message struct {
	Name     string
	fields   Fields
	GroupKey string
	Comment  Comment
	Package  *Pkg
}

type Field struct {
	Name     string
	DataType string
	Required bool
	Repeated bool
	Map      bool
	Type     FieldType
	message  *Message
	Comment  Comment
}

type Enum struct {
	Name    string
	values  EnumValues
	Package *Pkg
}

type EnumValue struct {
	Name    string  `json:"name"`
	Value   int64   `json:"value"`
	Comment Comment `json:",omitempty"`
	enum    *Enum
}

type ProtoFile []string
type Preprocessor struct {
	MessageName       func(string) string
	EnumName          func(string) string
	EnumValueName     func(string) string
	GolangPackageName func(string) string
	JavaPackageName   func(string) string
}

type CacheMap struct {
	Messages   commons.Cache
	Enums      commons.Cache
	EnumValues commons.Cache
}

type mapper struct {
	Schema       *schema.OCSFSchema
	Preprocessor Preprocessor
	Messages     Messages
	Enums        Enums
	Cache        CacheMap
	RootPackage  *Pkg
	OutputPath   string
	Fs           afero.Fs
}
