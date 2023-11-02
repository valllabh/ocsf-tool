package protobuff_v3

import "github.com/valllabh/ocsf-tool/commons"

var cache CacheMap

func InitCache() {
	cache = CacheMap{
		Messages:   *commons.NewCache(),
		Enums:      *commons.NewCache(),
		EnumValues: *commons.NewCache(),
	}
}

func Cache() *CacheMap {
	return &cache
}
