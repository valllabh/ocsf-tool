package protobuff_v3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"github.com/valllabh/ocsf-tool/commons"
	"github.com/valllabh/ocsf-tool/ocsf/schema"
)

var _mapper *mapper

func NewMapper(schema *schema.OCSFSchema) *mapper {
	_mapper = &mapper{
		Schema: schema,
		Preprocessor: Preprocessor{
			MessageName: messageNamePreprocessor,
		},
		Messages:    Messages{},
		Enums:       Enums{},
		RootPackage: NewPackage("ocsf", nil),
		Cache: CacheMap{
			Messages:   *commons.NewCache(),
			Enums:      *commons.NewCache(),
			EnumValues: *commons.NewCache(),
		},
		Fs: afero.NewOsFs(),
	}

	return GetMapper()
}

func GetMapper() *mapper {
	return _mapper
}

func messageNamePreprocessor(name string) string {
	splitName := strings.Split(name, "/")
	return splitName[len(splitName)-1]
}

func (mapper *mapper) Marshal(events []schema.Event) {

	for _, event := range events {

		m := Message{
			Name:     event.Name,
			GroupKey: "Event: " + event.Category,
			Comment:  Comment{},
			Package:  mapper.PackageRef("events", event.Category),
		}

		m.Comment["Event"] = event.Category
		m.Comment["Event UID"] = fmt.Sprintf("%d", event.Uid)
		m.Comment["URL"] = fmt.Sprintf("https://schema.ocsf.io/%s/classes/%s", mapper.Schema.Version, event.Name)

		mapper.populateFieldsFromAttributes(&m, event.Attributes)

		AddMessage(&m)
	}

	mapper.RootPackage.Marshal()
}

func (mapper *mapper) populateFieldsFromAttributes(message *Message, attributes map[string]schema.Attribute) {
	for k := range attributes {
		attr := attributes[k]

		// Build Field
		field := Field{
			Name:     k,
			DataType: getDataType(attr),
			Required: attr.Requirement == "required",
			Repeated: attr.IsArray,
			Map:      attr.IsMap, // Map is not natively supported by OCSF Schema
		}

		// Add Comments
		comments := Comment{
			"Caption": attr.Caption,
		}
		if len(attr.Profile) > 0 {
			comments["Profile"] = attr.Profile
		}
		field.Comment = comments

		// Detect Type
		field.Type = FIELD_TYPE_PRIMITIVE

		if field.DataType == "object" {
			field.Type = FIELD_TYPE_OBJECT
		}

		if len(attr.Enum) > 0 {
			field.Type = FIELD_TYPE_ENUM
		}

		// Processing Based on Type
		switch field.Type {
		case FIELD_TYPE_OBJECT:
			field.DataType = attr.ObjectType
			attributeIsSelfReferencing := field.DataType == message.Name
			_, isObjectMapped := GetMessage(field.DataType)
			if !isObjectMapped && !attributeIsSelfReferencing {
				object, schemaForObjectExists := mapper.getObject(field.DataType)
				if schemaForObjectExists {
					m := &Message{
						Name:     object.Name,
						GroupKey: "Object",
						Package:  mapper.PackageRef("objects"),
					}

					AddMessage(m)

					mapper.populateFieldsFromAttributes(m, object.Attributes)
				}
			}

		case FIELD_TYPE_ENUM:
			enumName := message.Name + " " + field.Name
			field.DataType = enumName
			e, exists := GetEnum(enumName)

			if !exists {
				e = &Enum{
					Name:    enumName,
					Package: message.Package.NewPackage("enums"),
				}
			}
			for aek, aev := range attr.Enum {
				ev, evExists := e.GetValue(aev.Caption)

				if !evExists {
					v, _ := strconv.ParseInt(aek, 10, 64)
					ev = &EnumValue{
						Name:  aev.Caption,
						Value: v,
						Comment: Comment{
							"Type": "OCSF_VALUE",
						},
					}
				}
				e.AddValue(ev)
			}

			AddEnum(e)

		}

		// Add Field to Message
		message.AddField(&field)
	}
}

func (mapper *mapper) getObject(dataType string) (schema.Object, bool) {
	object, exists := mapper.Schema.Objects[dataType]
	return object, exists
}

func (mapper *mapper) PackageRef(pkgs ...string) *Pkg {
	pkgRef := mapper.RootPackage
	for _, pkg := range pkgs {
		pkgRef = pkgRef.NewPackage(pkg)
	}
	return pkgRef
}

func getDataType(attr schema.Attribute) string {
	var t string
	switch attr.Type {
	case "boolean_t":
		t = "bool"
	case "integer_t":
		t = "int32"
	case "long_t":
		t = "int64"
	case "string_t", "bytestring_t", "datetime_t", "email_t", "file_hash_t",
		"file_name_t", "hostname_t", "ip_t", "json_t", "mac_t", "process_name_t",
		"resource_uid_t", "subnet_t", "url_t", "username_t", "uuid_t":
		t = "string"
	case "float_t":
		t = "float"
	case "port_t":
		t = "int32"
	case "timestamp_t":
		t = "int64"
	case "object_t":
		t = "object"
	default:
		t = "unknown"
	}

	return t
}

func AddEnum(enum *Enum) {
	GetMapper().Enums[enum.Name] = enum
}

func GetEnum(name string) (*Enum, bool) {
	value, exists := GetMapper().Enums[name]
	return value, exists
}

func AddMessage(message *Message) {
	GetMapper().Messages[ToMessageName(message.Name)] = message
}

func GetMessage(name string) (*Message, bool) {
	value, exists := GetMapper().Messages[ToMessageName(name)]
	return value, exists
}
