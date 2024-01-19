package gnlib

// Map applies a function to each element of a slice and returns a new slice
// in the same order.
func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements of s for which
// filter function returns true.
func FilterFunc[T any](s []T, f func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
