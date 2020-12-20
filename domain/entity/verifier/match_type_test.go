package verifier_test

import (
	"testing"

	ver "github.com/gnames/gnlib/domain/entity/verifier"
	"github.com/gnames/gnlib/encode"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testData := []struct {
		mt  ver.MatchTypeValue
		res string
	}{
		{ver.NoMatch, "NoMatch"},
		{ver.Exact, "Exact"},
		{ver.Fuzzy, "Fuzzy"},
		{ver.PartialExact, "PartialExact"},
		{ver.PartialFuzzy, "PartialFuzzy"},
	}
	for _, v := range testData {
		assert.Equal(t, v.mt.String(), v.res)
	}
}

func TestNew(t *testing.T) {
	testData := []struct {
		res ver.MatchTypeValue
		s   string
	}{
		{ver.NoMatch, "NoMatch"},
		{ver.Exact, "Exact"},
		{ver.Fuzzy, "Fuzzy"},
		{ver.PartialExact, "PartialExact"},
		{ver.PartialFuzzy, "PartialFuzzy"},
		{ver.NoMatch, ""},
		{ver.NoMatch, "??хррроо"},
	}
	for _, v := range testData {
		assert.Equal(t, ver.NewMatchType(v.s), v.res)
	}
}

func TestJSON(t *testing.T) {
	type testData struct {
		Field1 string               `json:"fieldOne"`
		Field2 []ver.MatchTypeValue `json:"fieldTwo"`
	}
	test := testData{
		Field1: "hello",
		Field2: []ver.MatchTypeValue{ver.Exact, ver.Fuzzy, ver.NoMatch},
	}
	enc := encode.GNjson{}
	res, err := enc.Encode(test)
	assert.Nil(t, err)
	assert.Equal(t, string(res), `{"fieldOne":"hello","fieldTwo":["Exact","Fuzzy","NoMatch"]}`)
	var res2 testData
	err = enc.Decode(res, &res2)
	assert.Nil(t, err)
	assert.Equal(t, res2, testData{Field1: "hello", Field2: []ver.MatchTypeValue{ver.Exact, ver.Fuzzy, ver.NoMatch}})
}

func TestErrJSON(t *testing.T) {
	enc := encode.GNjson{}
	res, err := enc.Encode("notType")
	assert.Nil(t, err)
	var res2 ver.MatchTypeValue
	err = enc.Decode(res, &res2)
	assert.Contains(t, err.Error(), "cannot decode as a MatchType")

	res, err = enc.Encode(3.1415926)
	assert.Nil(t, err)
	err = enc.Decode(res, &res2)
	assert.Contains(t, err.Error(), "cannot decode as a MatchType")
}
