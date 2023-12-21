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

// function to check if an element exists in a slice
// Example: Contains([]string{"a", "b", "c"}, "b") -> true
func Contains[T comparable](arr []T, element T) bool {

	// check if arr is not nil and empty
	if len(arr) == 0 {
		return false
	}

	// iterate over the array and check if the element exists
	for _, a := range arr {
		if a == element {
			return true
		}
	}

	// return false if the element does not exist
	return false
}
