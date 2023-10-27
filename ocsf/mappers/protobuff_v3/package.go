package protobuff_v3

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/commons"
	"golang.org/x/exp/maps"
)

func NewPackage(name string, parent *Pkg) *Pkg {
	p := &Pkg{
		Name:     name,
		Parent:   parent,
		Children: Pkgs{},
	}
	NewProto(p)
	return p
}

func (p *Pkg) NewPackage(pkgName string) *Pkg {
	pkg, exists := p.Children[pkgName]
	if exists {
		return pkg
	} else {
		p.Children[pkgName] = NewPackage(pkgName, p)
	}

	return p.Children[pkgName]
}

func (p *Pkg) GetName() string {
	return p.Name
}

func (p *Pkg) GetFullName() string {
	pkgs := p.GetParentHierarchy()
	pkgNames := []string{}

	for _, pkg := range pkgs {
		pkgNames = append(pkgNames, pkg.GetName())
	}

	return strings.Join(pkgNames, ".")
}

func (p *Pkg) GetParentHierarchy() []*Pkg {
	pkgs := []*Pkg{}

	if p.Parent != nil {
		hierarchy := p.Parent.GetParentHierarchy()
		pkgs = append(pkgs, hierarchy...)
	}

	pkgs = append(pkgs, p)

	return pkgs
}

func (p *Pkg) Marshal() {
	dir := p.GetDirPath()
	commons.EnsureDirExists(dir)
	for _, pkg := range p.Children {
		pkg.Marshal()
	}
	p.Proto.Marshal()
}

func (p *Pkg) GetDirName() string {
	return p.Name
}

func (p *Pkg) GetDirPath() string {

	path := p.Path

	if p.Parent != nil {
		path = p.Parent.GetDirPath()
	}

	return path + "/" + p.GetDirName()
}

func (p *Pkg) GetMessages() []*Message {
	msgs := maps.Values(GetMapper().Messages)

	filterFunc := func(m *Message) bool {
		return m.Package.GetName() == p.GetName()
	}

	return commons.Filter(msgs, filterFunc)
}

func (p *Pkg) GetEnums() []*Enum {
	msgs := maps.Values(GetMapper().Enums)

	filterFunc := func(e *Enum) bool {
		return e.Package.GetName() == p.GetName()
	}

	return commons.Filter(msgs, filterFunc)
}

func cleanPackageName(s string) string {
	return strcase.ToSnake(s)
}
