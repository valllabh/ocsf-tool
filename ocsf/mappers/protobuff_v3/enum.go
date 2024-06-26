package protobuff_v3

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
)

// Enum represents a Protobuf Enum
func (e *Enum) AddValue(value *EnumValue) {
	if e.values == nil {
		e.values = EnumValues{}
	}
	value.enum = e
	e.values[value.Name] = value
}

// Enum represents a Protobuf Enum
func (e *Enum) GetValue(name string) (*EnumValue, bool) {
	value, exists := e.values[name]
	return value, exists
}

// Get enum values sorted by name (with UNKNOWN coming first)
func (e *Enum) GetValues() []*EnumValue {

	// Add UNSPECIFIED if not present
	if !e.HasUNSPECIFIED() {
		e.AddValue(&EnumValue{
			Name:  "UNSPECIFIED",
			Value: 0,
			Comment: Comment{
				"Type": "NON_OCSF_VALUE",
			},
		})
	}

	values := maps.Values(e.values)

	// Sort values by name (with UNKNOWN coming first)
	sort.Slice(values, func(i, j int) bool {
		return values[i].Value < values[j].Value
	})

	return values
}

// Marshal returns the Enum as a string
func (e *Enum) Marshal() string {
	content := []string{}
	enumName := ToEnumName(e.Name)

	// Start Enum
	content = append(content, fmt.Sprintf("enum %s {", enumName))

	// Get enum values
	values := e.GetValues()

	// Marshal values and add to content
	for _, v := range values {
		content = append(content, "\t"+v.Marshal())
	}

	// Close Enum
	content = append(content, "}")

	// Return content
	return strings.Join(content, "\n")
}

// GetName returns the name of the Enum
func (e *Enum) GetName() string {
	return ToEnumName(e.Name)
}

// GetReference returns the reference of the Enum
func (e *Enum) GetReference() string {
	return e.GetPackage() + "." + e.GetName()
}

// GetPackage returns the package name of the Enum
func (e *Enum) GetPackage() string {
	return e.Package.GetFullName()
}

// Enum has at least one value ending with HasUNSPECIFIED
func (e *Enum) HasUNSPECIFIED() bool {
	for _, v := range e.values {
		if v.Value == 0 {
			return true
		}
	}
	return false
}

// ToEnumName converts a string to a valid Enum Name
func ToEnumName(input string) string {

	// Return if Cache exists
	value, exists := GetMapper().Cache.Enums.Get(input)

	if exists {
		return fmt.Sprint(value)
	}

	output := input

	// Apply Name Processor
	if GetMapper().Preprocessor.EnumName != nil {
		output = GetMapper().Preprocessor.EnumName(input)
	}

	// Clean Name
	output = cleanName(output)
	output = strcase.ToScreamingSnake(output)

	// Set Cache
	GetMapper().Cache.Enums.Set(input, output)

	return output
}
