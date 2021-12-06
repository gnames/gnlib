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

func TestFiftyFifty(t *testing.T) {
	tests := []struct {
		msg, paths, ranks, ids string
	}{
		{
			"potentilla",
			"Biota|Plantae|Tracheophyta|Magnoliopsida|Rosales|Rosaceae|Rosoideae|Potentilla|Potentilla erecta",
			"unranked|kingdom|phylum|class|order|family|subfamily|genus|species",
			"5T6MX|P|TP|MG|3Z6|FTK|628NC|6V7H|6VVPW",
		},
		{
			"puma",
			"Biota|Animalia|Chordata|Mammalia|Theria|Eutheria|Carnivora|Feliformia|Felidae|Felinae|Puma|Puma concolor",
			"unranked|kingdom|phylum|class|subclass|infraclass|order|suborder|family|subfamily|genus|species",
			"5T6MX|N|CH2|6224G|6226C|LG|VS|4DL|623RM|JKL|75F9|4QHKG",
		},
		{
			"plantago",
			"Biota|Plantae|Tracheophyta|Magnoliopsida|Lamiales|Plantaginaceae|Digitalidoideae|Plantago|Plantago major",
			"unranked|kingdom|phylum|class|order|family|subfamily|genus|species",
			"5T6MX|P|TP|MG|3F4|6262K|7NLQD|6RHN|4JLPC",
		},
		{
			"bubo",
			"Biota|Animalia|Chordata|Aves|Strigiformes|Strigidae|Striginae|Bubo|Bubo bubo",
			"unranked|kingdom|phylum|class|order|family|subfamily|genus|species",
			"5T6MX|N|CH2|V2|466|GQX|KDK|3DQQ|NKSD",
		},
	}
	hr := make([]context.Hierarch, len(tests))
	for i, v := range tests {
		hr[i] = NewTestHierarchy(v.ids, v.paths, v.ranks)
	}
	res := context.New(hr, 0)
	assert.Equal(t, res.Kingdom.Name, "")
	assert.Equal(t, res.KingdomPercentage, float32(0))
	assert.Equal(t, res.Context.Name, "")
	assert.Equal(t, res.ContextPercentage, float32(0))
}

func testData(t *testing.T) []context.Hierarch {
	var res []context.Hierarch
	var ids, names string
	path := filepath.Join("..", "..", "testdata", "context.txt")
	bytesRead, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	fileContent := string(bytesRead)
	ls := strings.Split(fileContent, "\n")

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

// NewTestHierarchy creates Name that can be used for calculation of
// hierarhcy for the context. It satisfies the Hierarch interface.
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
