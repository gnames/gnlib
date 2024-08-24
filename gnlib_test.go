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
