package commons

func Filter[T any](arr []T, filter func(T) bool) []T {
	var filtered []T
	for _, element := range arr {
		if filter(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}
