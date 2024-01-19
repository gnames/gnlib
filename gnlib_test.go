package gnlib_test

import (
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
