package nomcode_test

import (
	"testing"

	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, inp string
		out      nomcode.Code
	}{
		{"bad", "something", nomcode.Unknown},
		{"zoo1", "zoo", nomcode.Zoological},
		{"zoo2", "Zoological", nomcode.Zoological},
		{"zoo3", "ICZN", nomcode.Zoological},
		{"bot1", "bot", nomcode.Botanical},
		{"bot2", "botanical", nomcode.Botanical},
		{"bot2", "icn", nomcode.Botanical},
		{"cult1", "CULT", nomcode.Cultivars},
		{"cult2", "CultiVar", nomcode.Cultivars},
		{"cult3", "icncp", nomcode.Cultivars},
		{"bact1", "bact", nomcode.Bacterial},
		{"bact2", "bacterial", nomcode.Bacterial},
		{"bact3", "ICNP", nomcode.Bacterial},
	}

	for _, v := range tests {
		res := nomcode.New(v.inp)
		assert.Equal(v.out, res, v.msg)
	}
}

func TestAbbr(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, out string
		inp      nomcode.Code
	}{
		{"zoo", "ICZN", nomcode.Zoological},
		{"bot", "ICN", nomcode.Botanical},
		{"bact", "ICNP", nomcode.Bacterial},
		{"cult", "ICNCP", nomcode.Cultivars},
	}

	for _, v := range tests {
		res := v.inp.Abbr()
		assert.Equal(v.out, res, v.msg)
	}
}

func TestID(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, out string
		inp      nomcode.Code
	}{
		{"zoo", "ZOOLOGICAL", nomcode.Zoological},
		{"bot", "BOTANICAL", nomcode.Botanical},
		{"bact", "BACTERIAL", nomcode.Bacterial},
		{"cult", "CULTIVARS", nomcode.Cultivars},
	}

	for _, v := range tests {
		res := v.inp.ID()
		assert.Equal(v.out, res, v.msg)
	}
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, out string
		inp      nomcode.Code
	}{
		{"zoo", "zoological", nomcode.Zoological},
		{"bot", "botanical", nomcode.Botanical},
		{"bact", "bacterial", nomcode.Bacterial},
		{"cult", "cultivars", nomcode.Cultivars},
	}

	for _, v := range tests {
		res := v.inp.String()
		assert.Equal(v.out, res, v.msg)
	}
}
