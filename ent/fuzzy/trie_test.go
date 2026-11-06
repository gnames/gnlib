package fuzzy_test

import (
	"bytes"
	"testing"

	"github.com/gnames/gnlib/ent/fuzzy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var trieNames = []string{
	"Eriopygodes imbecill",
	"Pardosa maest",
	"Pomatomus saltator",
}

func TestTrieMatches(t *testing.T) {
	trie, err := fuzzy.NewTrie(trieNames)
	require.NoError(t, err)

	tests := []struct {
		msg, input string
		dist       int
		res        []string
	}{
		{"exact", "Pardosa maest", 0, []string{"Pardosa maest"}},
		{"exact only", "Pardosa moest", 0, nil},
		{"distance 1", "Pardosa moest", 1, []string{"Pardosa maest"}},
		{"ocr", "Eriopvgodes imbecill", 1, []string{"Eriopygodes imbecill"}},
		{"too far", "Pardosa moost", 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			res := trie.Matches(tt.input, tt.dist)
			assert.ElementsMatch(t, tt.res, res)
		})
	}
}

func TestTrieSaveLoad(t *testing.T) {
	var buf bytes.Buffer
	_, err := fuzzy.NewTrieWrite(trieNames, &buf)
	require.NoError(t, err)

	trie, err := fuzzy.LoadTrie(&buf)
	require.NoError(t, err)
	res := trie.Matches("Pomatomus saltatar", 1)
	assert.Equal(t, []string{"Pomatomus saltator"}, res)
}
