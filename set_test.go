package gnlib_test

import (
	"testing"

	"github.com/gnames/gnlib"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	assert := assert.New(t)
	s := make(gnlib.Set[int])
	s.Add(1)
	s.Add(2)
	s.Add(1) // duplicate

	assert.Equal(2, s.Len())
	assert.True(s.Has(1))
	assert.True(s.Has(2))
}

func TestSet_Has(t *testing.T) {
	assert := assert.New(t)
	s := make(gnlib.Set[string])
	s.Add("hello")

	assert.True(s.Has("hello"))
	assert.False(s.Has("world"))
}

func TestSet_Remove(t *testing.T) {
	assert := assert.New(t)
	s := make(gnlib.Set[int])
	s.Add(1)
	s.Add(2)
	s.Del(1)

	assert.False(s.Has(1))
	assert.True(s.Has(2))
	assert.Equal(1, s.Len())
}

func TestSet_Remove_NonExistent(t *testing.T) {
	assert := assert.New(t)
	s := make(gnlib.Set[int])
	s.Add(1)
	s.Del(999) // should not panic

	assert.Equal(1, s.Len())
}

func TestSet_Len(t *testing.T) {
	assert := assert.New(t)
	s := make(gnlib.Set[int])

	assert.Equal(0, s.Len())

	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.Equal(3, s.Len())
}

func TestSet_WithDifferentTypes(t *testing.T) {
	assert := assert.New(t)
	type point struct{ x, y int }
	s := make(gnlib.Set[point])
	s.Add(point{1, 2})
	s.Add(point{3, 4})

	assert.True(s.Has(point{1, 2}))
	assert.Equal(2, s.Len())
}
