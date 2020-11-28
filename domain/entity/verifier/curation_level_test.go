package verifier_test

import (
	"testing"

	ver "github.com/gnames/gnlib/domain/entity/verifier"
	"github.com/gnames/gnlib/encode"
	"github.com/stretchr/testify/assert"
)

func TestCurLevelString(t *testing.T) {
	testData := []struct {
		cl  ver.CurationLevel
		res string
	}{
		{ver.NotCurated, "NotCurated"},
		{ver.AutoCurated, "AutoCurated"},
		{ver.Curated, "Curated"},
	}

	for _, v := range testData {
		assert.Equal(t, v.cl.String(), v.res)
	}
}

func TestCurLevelJSON(t *testing.T) {
	type testData struct {
		Field1 string              `json:"fieldOne"`
		Field2 []ver.CurationLevel `json:"fieldTwo"`
	}
	test := testData{
		Field1: "hello",
		Field2: []ver.CurationLevel{ver.NotCurated, ver.AutoCurated, ver.Curated},
	}
	enc := encode.GNjson{}
	res, err := enc.Encode(test)
	assert.Nil(t, err)
	assert.Equal(t, string(res), `{"fieldOne":"hello","fieldTwo":["NotCurated","AutoCurated","Curated"]}`)
	var res2 testData
	err = enc.Decode(res, &res2)
	assert.Nil(t, err)
	assert.Equal(t, res2, testData{Field1: "hello", Field2: []ver.CurationLevel{ver.NotCurated, ver.AutoCurated, ver.Curated}})
}

func TestCurLevelErrJSON(t *testing.T) {
	enc := encode.GNjson{}
	res, err := enc.Encode("notType")
	assert.Nil(t, err)
	var res2 ver.CurationLevel
	err = enc.Decode(res, &res2)
	assert.Contains(t, err.Error(), "cannot convert 'notType'")

	res, err = enc.Encode(3.1415926)
	assert.Nil(t, err)
	err = enc.Decode(res, &res2)
	assert.Contains(t, err.Error(), "cannot unmarshal number")
}