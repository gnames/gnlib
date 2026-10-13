package abbr_test

import (
	"encoding/json"
	"testing"

	"github.com/gnames/gnlib/ent/abbr"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		mt  abbr.MatchTypeValue
		res string
	}{
		{abbr.NoMatch, "NoMatch"},
		{abbr.SpeciesLevel, "SpeciesLevel"},
		{abbr.EpithetMoved, "EpithetMoved"},
		{abbr.Full, "Full"},
		{abbr.MatchTypeValue(100), "N/A"},
	}
	for _, v := range tests {
		assert.Equal(v.res, v.mt.String())
	}
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		s   string
		res abbr.MatchTypeValue
	}{
		{"Full", abbr.Full},
		{"EpithetMoved", abbr.EpithetMoved},
		{"SpeciesLevel", abbr.SpeciesLevel},
		{"", abbr.NoMatch},
		{"Exact", abbr.NoMatch},
	}
	for _, v := range tests {
		assert.Equal(v.res, abbr.NewMatchType(v.s))
	}
}

func TestJSON(t *testing.T) {
	assert := assert.New(t)
	mn := abbr.MatchedName{
		Name:             "P. altaica",
		MatchedCanonical: "Panthera tigris altaica",
		MatchType:        abbr.EpithetMoved,
	}
	bs, err := json.Marshal(mn)
	assert.Nil(err)
	assert.Contains(string(bs), `"matchType":"EpithetMoved"`)

	var res abbr.MatchedName
	err = json.Unmarshal(bs, &res)
	assert.Nil(err)
	assert.Equal(mn, res)

	err = json.Unmarshal([]byte(`{"matchType":"Exact"}`), &res)
	assert.NotNil(err)
}
