package context_test

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnlib/ent/context"
	"github.com/gnames/gnlib/ent/verifier"
	"github.com/stretchr/testify/assert"
)

func TestTestData(t *testing.T) {
	hs := testData(t)
	assert.Equal(t, len(hs), 69)
	for i := range hs {
		clades := hs[i].Clades()
		assert.Greater(t, len(clades), 8)
	}
}

func TestContext(t *testing.T) {
	assert := assert.New(t)
	hs := testData(t)
	res := context.New(hs, 0.7)
	assert.Equal(res.Kingdom.Name, "Animalia")
	assert.Equal(res.KingdomPercentage, float32(1.0))
	assert.Equal(res.Context.RankStr, "phylum")
	assert.Equal(res.Context.Name, "Mollusca")
	assert.Equal(res.ContextPercentage, float32(1.0))

	res = context.New(hs, 0.5)
	assert.Equal(res.Context.RankStr, "class")
	assert.Equal(res.Context.Name, "Gastropoda")
	assert.InDelta(res.ContextPercentage, float32(0.55), 0.01)
}

// TestFishes tests situation where some sequence of ranks varies from
// name to name, and some of the names are are higher than genus.
func TestFishes(t *testing.T) {
	hs := context2(t)
	// there are 9 names
	assert.Equal(t, 9, len(hs))
	res := context.New(hs, 0.5)
	// of of the names is bigger than genus and is removed
	assert.Equal(t, 8, res.NamesNum)
	assert.Equal(t, "Animalia", res.Kingdom.Name)
	assert.Equal(t, float32(1.0), res.KingdomPercentage)
	assert.Equal(t, "Actinopterygii", res.Context.Name)
	assert.Equal(t, float32(0.75), res.ContextPercentage)
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

func context2(t *testing.T) []context.Hierarch {
	var res []context.Hierarch
	path := filepath.Join("..", "..", "testdata", "context2.csv")

	f, err := os.Open(path)
	assert.Nil(t, err)
	defer f.Close()
	r := csv.NewReader(f)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		assert.Nil(t, err)
		n := NewTestHierarchy(row[2], row[0], row[1])
		res = append(res, n)
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
