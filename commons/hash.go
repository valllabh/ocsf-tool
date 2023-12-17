package commons

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Hash(args ...string) string {
	hasher := md5.New()
	hasher.Write([]byte(strings.Join(args, "")))
	return hex.EncodeToString(hasher.Sum(nil))
}
