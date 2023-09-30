package protobuff

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/valllabh/ocsf-schema-processor/ocsf"
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff/commons"
	v3 "github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff/v3"
)

type mapper struct {
	proto  *v3.Proto
	Schema ocsf.OCSFSchema
}

func NewMapper(schema ocsf.OCSFSchema) mapper {

	proto := v3.NewProto()
	proto.Preprocessor = v3.Preprocessor{
		MessageName: messageNamePreprocessor,
	}

	return mapper{
		proto:  proto,
		Schema: schema,
	}
}

func messageNamePreprocessor(name string) string {
	splitName := strings.Split(name, "/")
	return splitName[len(splitName)-1]
}

func (mapper *mapper) Marshal(events []ocsf.Event) string {

	for _, event := range events {

		m := v3.Message{
			Name:     event.Name,
			GroupKey: "Event: " + event.Category,
			Comment:  commons.Comment{},
		}

		m.Comment["Event"] = event.Category
		m.Comment["Event UID"] = fmt.Sprintf("%d", event.Uid)
		m.Comment["URL"] = fmt.Sprintf("https://schema.ocsf.io/%s/classes/%s", mapper.Schema.Version, event.Name)

		mapper.populateFieldsFromAttributes(&m, event.Attributes)

		mapper.proto.AddMessage(&m)

	}

	return mapper.proto.Marshal()
}

func (mapper *mapper) populateFieldsFromAttributes(message *v3.Message, attributes map[string]ocsf.Attribute) {
	for k := range attributes {
		attr := attributes[k]

		// Build Field
		field := v3.Field{
			Name:     k,
			DataType: getDataType(attr),
			Required: attr.Requirement == "required",
			Repeated: attr.IsArray,
		}

		// Add Comments
		comments := commons.Comment{
			"Caption": attr.Caption,
		}
		if len(attr.Profile) > 0 {
			comments["Profile"] = attr.Profile
		}
		field.Comment = comments

		// Detect Type
		field.Type = v3.FIELD_TYPE_PRIMITIVE

		if field.DataType == "object" {
			field.Type = v3.FIELD_TYPE_OBJECT
		}

		if len(attr.Enum) > 0 {
			field.Type = v3.FIELD_TYPE_ENUM
		}

		// Processing Based on Type
		switch field.Type {
		case v3.FIELD_TYPE_OBJECT:
			field.DataType = attr.ObjectType
			attributeIsSelfReferencing := field.DataType == message.Name
			isObjectMapped := mapper.proto.MessageExists(field.DataType)

			if !isObjectMapped && !attributeIsSelfReferencing {
				object, schemaForObjectExists := mapper.getObject(field.DataType)
				if schemaForObjectExists {
					m := v3.Message{
						Name:     object.Name,
						GroupKey: "Object",
					}
					mapper.populateFieldsFromAttributes(&m, object.Attributes)

					mapper.proto.AddMessage(&m)
				}
			}
		case v3.FIELD_TYPE_ENUM:
			enumName := message.Name + " " + field.Name
			field.DataType = enumName
			e, exists := mapper.proto.GetEnum(enumName)

			if !exists {
				e = &v3.Enum{
					Name: enumName,
				}
			}
			for aek, aev := range attr.Enum {
				ev, evExists := e.GetValue(aev.Caption)

				if !evExists {
					v, _ := strconv.ParseInt(aek, 10, 64)
					ev = &v3.EnumValue{
						Name:  aev.Caption,
						Value: v,
					}
				}
				e.AddValue(ev)
			}

			mapper.proto.AddEnum(e)

		}

		// Add Field to Message
		message.AddField(&field)
	}
}

func (mapper *mapper) getObject(dataType string) (ocsf.Object, bool) {
	object, exists := mapper.Schema.Objects[dataType]
	return object, exists
}

func getDataType(attr ocsf.Attribute) string {
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
