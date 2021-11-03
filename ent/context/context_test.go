package context_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnlib/ent/context"
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
	res := context.CalcContext(hs, 0.7)
	assert.Equal(t, res.Kingdom.Name, "Animalia")
	assert.Equal(t, res.KingdomPC, float32(1.0))
	assert.Equal(t, res.Context.RankStr, "phylum")
	assert.Equal(t, res.Context.Name, "Mollusca")
	assert.Equal(t, res.ContextPC, float32(1.0))

	res = context.CalcContext(hs, 0.5)
	assert.Equal(t, res.Context.RankStr, "class")
	assert.Equal(t, res.Context.Name, "Gastropoda")
	assert.InDelta(t, res.ContextPC, float32(0.55), 0.01)
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

type testHierarchy struct {
	clades []context.Clade
}

func (h testHierarchy) Clades() []context.Clade {
	return h.clades
}

func NewTestHierarchy(idStr, nameStr, rankStr string) testHierarchy {
	ids := strings.Split(idStr, "|")
	names := strings.Split(nameStr, "|")
	ranks := strings.Split(rankStr, "|")
	clades := make([]context.Clade, len(ids))
	for i := range ids {
		clades[i].ID = ids[i]
		clades[i].Name = names[i]
		clades[i].RankStr = ranks[i]
		clades[i].Rank = context.NewRank(ranks[i])
	}

	return testHierarchy{clades: clades}
}
