package context_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnlib/ent/context"
	"github.com/gnames/gnlib/ent/verifier"
	"github.com/stretchr/testify/assert"
)

func TestTestData(t *testing.T) {
	test := testData(t)
	assert.Equal(t, len(test), 69)
	for i := range test {
		clades := test[i].Clades()
		assert.Greater(t, len(clades), 8)
	}
}

func TestContext(t *testing.T) {
	hs := testData(t)
	res := context.New(hs, 0.7)
	assert.Equal(t, res.Kingdom.Name, "Animalia")
	assert.Equal(t, res.KingdomPercentage, float32(1.0))
	assert.Equal(t, res.Context.RankStr, "phylum")
	assert.Equal(t, res.Context.Name, "Mollusca")
	assert.Equal(t, res.ContextPercentage, float32(1.0))

	res = context.New(hs, 0.5)
	assert.Equal(t, res.Context.RankStr, "class")
	assert.Equal(t, res.Context.Name, "Gastropoda")
	assert.InDelta(t, res.ContextPercentage, float32(0.55), 0.01)
}

func testData(t *testing.T) []context.Hierarch {
	var res []context.Hierarch
	var ids, names string
	path := filepath.Join("..", "..", "testdata", "context.txt")
	bytesRead, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	file_content := string(bytesRead)
	ls := strings.Split(file_content, "\n")

	for _, v := range ls {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\"")
		if ids == "" {
			ids = v
		} else if names == "" {
			names = v
		} else {
			h := NewTestHierarchy(ids, names, v)
			res = append(res, h)
			ids, names = "", ""
		}
	}
	return res
}

func NewTestHierarchy(idStr, nameStr, rankStr string) *verifier.Name {
	rd := verifier.ResultData{
		DataSourceID:        1,
		ClassificationIDs:   idStr,
		ClassificationPath:  nameStr,
		ClassificationRanks: rankStr,
	}
	name := verifier.Name{
		BestResult: &rd,
	}

	return &name
}
