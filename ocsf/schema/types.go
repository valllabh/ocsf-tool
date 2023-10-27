package schema

// EnumAttribute represents an attribute of an enum.
type EnumAttribute struct {
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

// Enum represents a map of string keys to EnumAttribute values.
type Enum map[string]EnumAttribute

// Constraints represents the constraints for an object.
type Constraints struct {
	AtLeastOne []string `json:"at_least_one"`
	JustOne    []string `json:"just_one"`
}

// Attribute represents an attribute of an object or event.
type Attribute struct {
	Caption     string   `json:"caption"`
	Description string   `json:"description"`
	Enum        Enum     `json:"enum"`
	Group       string   `json:"group"`
	Requirement string   `json:"requirement"`
	Sibling     string   `json:"sibling"`
	Type        string   `json:"type"`
	TypeName    string   `json:"type_name"`
	Default     any      `json:"default"`
	IsArray     bool     `json:"is_array"`
	ObjectName  string   `json:"object_name"`
	ObjectType  string   `json:"object_type"`
	Attributes  []string `json:"attributes"`
	Profile     string   `json:"profile"`
}

// Event represents an event in the schema.
type Event struct {
	Attributes   map[string]Attribute `json:"attributes"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Uid          int                  `json:"uid"`
	Category     string               `json:"category"`
	Caption      string               `json:"caption"`
	Profiles     []string             `json:"profiles"`
	CategoryName interface{}          `json:"category_name"`
}

// Object represents an object in the schema.
type Object struct {
	Attributes  map[string]Attribute `json:"attributes"`
	Caption     string               `json:"caption"`
	Constraints Constraints          `json:"constraints"`
	Description string               `json:"description"`
	Extends     string               `json:"extends"`
	Name        string               `json:"name"`
}

// Type represents a type in the schema.
type Type struct {
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Values      []bool `json:"values"`
	Type        string `json:"type"`
	TypeName    string `json:"type_name"`
	Regex       string `json:"regex"`
	Observable  int    `json:"observable"`
	MaxLen      int    `json:"max_len"`
	Range       []int  `json:"range"`
}

// OCSFSchema represents the entire schema.
type OCSFSchema struct {
	BaseEvent Event             `json:"base_event"`
	Classes   map[string]Event  `json:"classes"`
	Objects   map[string]Object `json:"objects"`
	Types     map[string]Type   `json:"types"`
	Version   string            `json:"version"`
}
