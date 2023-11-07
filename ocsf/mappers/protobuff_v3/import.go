package protobuff_v3

import (
	"fmt"
	"strings"
)

func (i *Import) Marshal() string {
	path := i.Name
	path = strings.ReplaceAll(path, GetMapper().OutputPath+"/", "")
	return fmt.Sprintf("import \"%s\";", path)
}
