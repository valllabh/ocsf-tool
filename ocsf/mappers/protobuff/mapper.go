package protobuff

import (
	"strings"

	"github.com/valllabh/ocsf-schema-processor/ocsf"
)

type Messages map[string]Message

type Mapper struct {
	messages Messages
	Events   []ocsf.Event
	Schema   ocsf.OCSFSchema
}

func (mapper Mapper) Marshal() string {

	pack := []string{}

	mapper.messages = Messages{}

	for _, event := range mapper.Events {
		mapper.messages[event.Name] = Message{
			name:   event.Name,
			fields: mapper.getFieldsFromAttributes(event.Name, event.Attributes),
		}
	}

	for _, m := range mapper.messages {
		pack = append(pack, m.Marshal())
	}

	return strings.Join(pack, "\n\n")
}

func (mapper Mapper) getFieldsFromAttributes(messageDataType string, attributes map[string]ocsf.Attribute) []Field {
	fields := []Field{}

	for k := range attributes {
		attr := attributes[k]
		field := Field{
			name:     k,
			dataType: getDataType(attr),
			optional: attr.Requirement != "required",
			repeated: attr.IsArray,
		}
		field.isReference = field.dataType == "object"

		if field.isReference {
			field.dataType = attr.ObjectType
			attributeIsSelfReferencing := field.dataType == messageDataType
			_, attributeIsAlreadyMapped := mapper.messages[field.dataType]

			if !attributeIsAlreadyMapped && !attributeIsSelfReferencing {
				object, schemaForObjectExists := mapper.Schema.Objects[field.dataType]
				if schemaForObjectExists {
					mapper.messages[field.dataType] = Message{
						name:   object.Name,
						fields: mapper.getFieldsFromAttributes(field.dataType, object.Attributes),
					}
				}
			}

		}

		fields = append(fields, field)
	}

	return fields
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
		t = "float32"
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
