package ocsf

import (
	"encoding/json"
	"log"
	"os"
)

type EnumAttribute struct {
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

type Enum map[string]EnumAttribute

type Constraints struct {
	AtLeastOne []string `json:"at_least_one"`
	JustOne    []string `json:"just_one"`
}

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

type Object struct {
	Attributes  map[string]Attribute `json:"attributes"`
	Caption     string               `json:"caption"`
	Constraints Constraints          `json:"constraints"`
	Description string               `json:"description"`
	Extends     string               `json:"extends"`
	Name        string               `json:"name"`
}

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

type OCSFSchema struct {
	BaseEvent Event             `json:"base_event"`
	Classes   map[string]Event  `json:"classes"`
	Objects   map[string]Object `json:"objects"`
	Types     map[string]Type   `json:"types"`
	Version   string            `json:"version"`
}

func LoadOCSFSchema() (OCSFSchema, error) {
	// Specify the path to the JSON file
	filePath := "./ocsf/schema.json"

	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Define a variable of the struct type
	var schema OCSFSchema

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(data, &schema); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	return schema, err
}
