package gnlib_test

import (
	"context"
	"strings"
	"testing"

	"github.com/gnames/gnlib"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)
	test := []string{"a", "b", "c"}
	res := gnlib.Map(test, func(s string) string {
		return strings.ToUpper(s)
	})
	assert.Equal([]string{"A", "B", "C"}, res)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	test := []string{"a", "b", "c"}
	res := gnlib.FilterFunc(test, func(s string) bool {
		return s != "b"
	})
	assert.Equal([]string{"a", "c"}, res)
}

func TestSliceMap(t *testing.T) {
	assert := assert.New(t)
	sm := gnlib.SliceMap([]int{1, 2, 3})
	assert.Equal(1, sm[2])
	sm2 := gnlib.SliceMap([]string{"one", "two", "three"})
	assert.Equal(1, sm2["two"])
}

func TestIsVersion(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, input string
		isVer      bool
	}{
		{"not version", "sdfs aas", false},
		{"too short", "v1.2", false},
		{"version", "v23.443.102", true},
		{"patch", "v1.55.2.1", true},
		{"version2", "v1.2.3", true},
		{"version3", "v0.0.0", true},
		{"typo", "v.1.5.3", false},
		{"typo2", "v.1.5.3.", false},
	}

	for _, v := range tests {
		res := gnlib.IsVersion(v.input)
		assert.Equal(v.isVer, res, v.msg)
	}
}

func TestCmpVersion(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		ver1, ver2 string
		res        int
	}{
		{"v1.2.3", "v1.2.3", 0},
		{"v2.0.0", "v1.55.3", 1},
		{"v1.3.1", "v1.4.2", -1},
		{"v1.6.0", "v1.6.0.3", -1},
		{"v1.6.0", "v1.6.0.3b", 0},
	}

	for _, v := range tests {
		res := gnlib.CmpVersion(v.ver1, v.ver2)
		assert.Equal(v.res, res)
	}
}

// TestChunkChannel tests the ChunkChannel function with various input
// scenarios.
func TestChunkChannel(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		chunkSize int
		expected  [][]int
	}{
		{
			name:      "normal",
			input:     []int{1, 2, 3, 4, 5, 6},
			chunkSize: 2,
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "partial",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: 2,
			expected:  [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:      "empty",
			input:     []int{},
			chunkSize: 2,
			expected:  nil,
		},
		{
			name:      "chunk size 1",
			input:     []int{1, 2, 3},
			chunkSize: 1,
			expected:  [][]int{{1}, {2}, {3}},
		},
		{
			name:      "chunk size larger",
			input:     []int{1, 2, 3},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3}},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			// Create input channel
			input := make(chan int)

			// Start goroutine to send input values and close the channel
			go func() {
				for _, v := range v.input {
					input <- v
				}
				close(input)
			}()

			// Call ChunkChannel to get the output channel
			output := gnlib.ChunkChannel(context.Background(), input, v.chunkSize)

			// Collect all chunks from the output channel
			var result [][]int
			for chunk := range output {
				result = append(result, chunk)
			}

			// Assert that the result matches the expected output
			assert.Equal(t, v.expected, result)
		})
	}
}
