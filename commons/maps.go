package commons

// func GetMapKeys is to get keys of map storing any type of values
// Example: GetMapKeys(map[string]int{"a": 1, "b": 2}) -> []string{"a", "b"}
func GetMapKeys[T any](m map[string]T) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
