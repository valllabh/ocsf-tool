package protobuff_v3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/valllabh/ocsf-tool/commons"
)

func NewProto(p *Pkg) {
	proto := &Proto{
		Pkg: p,
	}

	p.Proto = proto
}

func (p *Proto) Marshal() {
	messages := p.Pkg.GetMessages()
	enums := p.Pkg.GetEnums()
	protoPath := p.Pkg.Proto.GetProtoFilePath()

	if len(messages) == 0 && len(enums) == 0 {
		return
	}

	content := []string{}

	// Proto Header
	content = append(content, fmt.Sprintf("syntax = \"%s\";", "proto3"))
	content = append(content, fmt.Sprintf("package %s;", p.Pkg.GetFullName()))

	// Go Package

	goPackage := p.Pkg.GetFullName()
	goPackage = strings.ReplaceAll(goPackage, ".", "/")

	if GetMapper().Preprocessor.GolangPackageName != nil {
		goPackage = GetMapper().Preprocessor.GolangPackageName(goPackage)
	}
	content = append(content, fmt.Sprintf("option go_package = \"%s\";", goPackage))

	// Proto Body >>>

	// Preparing Messages for Appending
	messageContent := []string{}
	imports := Imports{}

	for _, m := range messages {
		// Appending Message
		messageContent = append(messageContent, m.Marshal())

		// Merge imports map from all messages
		for _, i := range m.GetImports() {
			imports[i.Name] = i
		}
	}

	// Appending imports
	for _, i := range imports {
		if protoPath != i.Name {
			content = append(content, i.Marshal())
		}
	}

	// Appending Messages
	content = append(content, messageContent...)

	// Appending Enums
	for _, e := range enums {
		content = append(content, e.Marshal())
	}

	output := strings.Join(content, "\n\n")

	commons.CreateFile(p.GetProtoFilePath(), []byte(output))
}

func (p *Proto) GetProtoFilePath() string {
	return p.Pkg.GetDirPath() + "/" + p.Pkg.Name + ".proto"
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
