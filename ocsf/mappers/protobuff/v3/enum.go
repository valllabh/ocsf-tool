package v3

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
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

	values := maps.Values(e.values)

	sort.Slice(values, func(i, j int) bool {
		return valueSorter(values, i, j)
	})

	i := 0
	for _, v := range values {
		content = append(content, "\t"+v.Marshal(i))
		i++
	}
	content = append(content, "}")
	return strings.Join(content, "\n")
}
func valueSorter(values []*EnumValue, i int, j int) bool {

	valueI := values[i].Name
	valueJ := values[j].Name
	lenUnknown := len("UNKNOWN")
	lenValueI := len(valueI)
	lenValueJ := len(valueJ)

	// Check if the string at index i ends with "UNKNOWN"
	endsWithUNKNOWNI := false
	if lenValueI >= lenUnknown {
		endsWithUNKNOWNI = valueI[lenValueI-lenUnknown:] == "UNKNOWN"
	}

	// Check if the string at index j ends with "UNKNOWN"
	endsWithUNKNOWNJ := false
	if lenValueJ >= lenUnknown {
		endsWithUNKNOWNJ = valueJ[lenValueJ-lenUnknown:] == "UNKNOWN"
	}

	// If only one of them ends with "UNKNOWN," it should come first
	if endsWithUNKNOWNI && !endsWithUNKNOWNJ {
		return true
	} else if !endsWithUNKNOWNI && endsWithUNKNOWNJ {
		return false
	}

	// If both end with "UNKNOWN" or neither do, sort lexicographically
	return valueI > valueJ
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
