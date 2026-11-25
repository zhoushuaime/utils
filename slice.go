package utils

// FindFirst returns the first slice element equal to target and its index.
// If no element matches, it returns the zero value of T and -1.
func FindFirst[T comparable](in []T, target T) (T, int) {
	for index, element := range in {
		if element == target {
			return element, index
		}
	}

	var zero T
	return zero, -1
}
