package tribool_test

import (
	"testing"

	"github.com/gnames/gnlib/encode"
	tbl "github.com/gnames/gnlib/tribool"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data := []struct {
		opts  []int
		valid bool
		str   string
		b     bool
	}{
		{[]int{}, false, "", false},
		{[]int{1}, true, "yes", true},
		{[]int{0}, true, "maybe", false},
		{[]int{-1}, true, "no", false},
		{[]int{-1, 44}, true, "no", false},
	}
	for i := range data {
		tb := tbl.NewTribool(data[i].opts...)
		assert.Equal(t, tb.Valid, data[i].valid)
		assert.Equal(t, tb.String(), data[i].str)
		assert.Equal(t, tb.Bool(), data[i].b)
	}
}

func TestJSON(t *testing.T) {
	enc := encode.GNjson{}
	type dataStruct struct {
		Field1 string      `json:"f1"`
		Tb     tbl.Tribool `json:"tb"`
		Field2 []int       `json:"f2"`
	}
	data := []struct {
		dataStruct
		res string
	}{
		{dataStruct{"null", tbl.NewTribool(), []int{1, 2}},
			`{"f1":"null","tb":null,"f2":[1,2]}`},
		{dataStruct{"yes", tbl.NewTribool(10), []int{}},
			`{"f1":"yes","tb":"yes","f2":[]}`},
		{dataStruct{"maybe", tbl.NewTribool(0), []int{5}},
			`{"f1":"maybe","tb":"maybe","f2":[5]}`},
		{dataStruct{"no", tbl.NewTribool(-3), []int{3}},
			`{"f1":"no","tb":"no","f2":[3]}`},
	}
	for i := range data {
		var v dataStruct
		res, err := enc.Encode(data[i])
		assert.Nil(t, err)
		assert.Equal(t, string(res), data[i].res)
		err = enc.Decode(res, &v)
		assert.Nil(t, err)
		assert.Equal(t, v.Tb.Valid, data[i].Tb.Valid)
		assert.Equal(t, v.Tb.String(), data[i].Tb.String())
	}
}
