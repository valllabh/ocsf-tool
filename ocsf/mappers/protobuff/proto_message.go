package protobuff

import (
	"fmt"
	"strings"
)

type Field struct {
	name        string
	dataType    string
	optional    bool
	repeated    bool
	isReference bool
}

type Message struct {
	name   string
	fields []Field
}

func (m Message) Marshal() string {
	content := []string{}
	content = append(content, fmt.Sprintf("message %s {", toValidMessageName(m.name)))
	for i, f := range m.fields {
		content = append(content, "\t"+f.Marshal(i+1))
	}
	content = append(content, "}")
	return strings.Join(content, "\n")
}

func (f Field) Marshal(index int) string {
	content := []string{}
	if f.optional {
		content = append(content, "optional")
	}
	if f.repeated {
		content = append(content, "repeated")
	}
	if f.isReference {
		content = append(content, toValidMessageName(f.dataType))
	} else {
		content = append(content, f.dataType)
	}
	content = append(content, f.name)
	content = append(content, fmt.Sprintf("= %d;", index))
	return strings.Join(content, " ")
}

var cache = map[string]string{}

func toValidMessageName(input string) string {

	cachedValue, cacheExists := cache[input]

	if cacheExists {
		return cachedValue
	}

	// Remove spaces, hyphens, and underscores and split into words
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	// Convert words to title case (uppercase first letter of each word)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}

	// Join the words without spaces, hyphens, or underscores
	className := strings.Join(words, "")

	cache[input] = className

	return className
}
