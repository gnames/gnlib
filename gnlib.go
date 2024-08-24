package gnlib

import (
	"strconv"
	"strings"
)

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

func CmpVersion(a, b string) int {
	if a == b {
		return 0
	}

	asl := verToAri(a)
	bsl := verToAri(b)
	asl, bsl = mkSameLen(asl, bsl)
	for i := range asl {
		if asl[i] > bsl[i] {
			return 1
		}
		if asl[i] < bsl[i] {
			return -1
		}
	}
	return 0
}

func mkSameLen(a, b []int) ([]int, []int) {
	if len(a) == len(b) {
		return a, b
	}
	if len(a) < len(b) {
		newSlice := make([]int, len(b))
		copy(newSlice, a)
		a = newSlice
	} else {
		newSlice := make([]int, len(a))
		copy(newSlice, b)
		b = newSlice
	}

	return a, b
}

func verToAri(ver string) []int {
	ver = strings.TrimPrefix(ver, "v")
	es := strings.Split(ver, ".")
	return Map(es, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
}
