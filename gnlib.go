package gnlib

import (
	"context"
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

// SliceMap takes a slice and returns back a lookup map which allows to find
// index for each element of the slice. If the value happens several times
// in the slice, the index corresponds to the first matching element.
func SliceMap[T comparable](s []T) map[T]int {
	res := map[T]int{}
	for i, v := range s {
		if _, ok := res[v]; !ok {
			res[v] = i
		}
	}
	return res
}

// ChunkChannel reads from an input channel and sends slices (chunks) of
// up to `chunkSize` items to an output channel. If the input channel is
// closed, any remaining items are sent as a final slice. The output channel
// is closed after all data is processed.
//
// Parameters:
//   - `ctx`: Context for cancellation.
//   - `input`: Input channel.
//   - `chunkSize`: Max items per chunk.
//
// Returns:
//   - Output channel with slices of items.
//
// Example:
//
//	input := make(chan int)
//	go func() { for i := 1; i <= 10; i++ { input <- i }; close(input) }()
//	chunked := ChunkChannel(context.Background(), input, 3)
//	for chunk := range chunked { fmt.Println(chunk) }
//	// Output: [1 2 3] [4 5 6] [7 8 9] [10]
func ChunkChannel[T any](ctx context.Context, input <-chan T, chunkSize int) <-chan []T {
	output := make(chan []T)
	go func() {
		defer close(output) // Close output channel when done
		var chunk []T       // Buffer to collect items
		for {
			select {
			case <-ctx.Done(): // Handle cancellation
				return
			case val, ok := <-input:
				if !ok { // Input channel closed
					if len(chunk) > 0 {
						output <- chunk // Send remaining items
					}
					return
				}
				chunk = append(chunk, val)
				if len(chunk) == chunkSize {
					output <- chunk // Send full chunk
					chunk = nil     // Reset chunk
				}
			}
		}
	}()
	return output
}

// CmpVersion compares two semantic versions (eg v0.1.3 vs v0.2.0) as a and b.
// It returns 0 if the versions are equal, 1 if a is greater than b, and -1
// if a is less than b. The version strings are expected to be in a format
// that can be split into integer components for comparison,
// such as "1.2.3" or "1.0.0".
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
