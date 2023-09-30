package v3

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff/commons"
	"golang.org/x/exp/maps"
)

func NewProto() *Proto {
	proto := Proto{}

	proto.cache.Messages = *commons.NewCache()
	proto.cache.Enums = *commons.NewCache()
	proto.cache.EnumValues = *commons.NewCache()

	return &proto
}

func (p *Proto) AddMessage(message *Message) {
	if p.messages == nil {
		p.messages = Messages{}
	}
	message.proto = p
	p.messages[message.Name] = message
}

func (p *Proto) AddEnum(enum *Enum) {
	if p.enums == nil {
		p.enums = Enums{}
	}
	enum.proto = p
	p.enums[enum.Name] = enum
}

func (p *Proto) EnumExists(name string) bool {
	_, exists := p.enums[name]
	return exists
}

func (p *Proto) MessageExists(name string) bool {
	_, exists := p.messages[name]
	return exists
}

func (p *Proto) Marshal() string {
	content := []string{}
	content = append(content, fmt.Sprintf("syntax = \"%s\";", "proto3"))

	messages := maps.Values(p.messages)

	sort.Slice(messages, func(i, j int) bool {

		group := messages[i].GroupKey < messages[j].GroupKey

		if messages[i].GroupKey == messages[j].GroupKey {
			return messages[i].Name < messages[j].Name
		}

		return group
	})

	for _, m := range messages {
		content = append(content, m.Marshal())
	}
	for _, e := range p.enums {
		content = append(content, e.Marshal())
	}
	return strings.Join(content, "\n\n")
}

func cleanName(name string) string {

	// Remove leading and trailing spaces
	value := strings.TrimSpace(name)

	// Define a regular expression to match non-alphanumeric characters
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Replace all non-alphanumeric characters with spaces
	value = regex.ReplaceAllString(value, " ")

	return value
}
