package fuzzy

import (
	"io"

	"github.com/dvirsky/levenshtein"
)

// Trie finds strings that are close to a given string. It is a wrapper of
// levenshtein automata, gnmatcher and gndb build it from stemmed canonical
// forms.
type Trie struct {
	tree *levenshtein.MinTree
}

// NewTrie creates a Trie. The names must be unique and sorted by bytes.
func NewTrie(names []string) (*Trie, error) {
	tree, err := levenshtein.NewMinTree(names)
	if err != nil {
		return nil, err
	}
	return &Trie{tree: tree}, nil
}

// NewTrieWrite creates a Trie and saves it to w, so LoadTrie can restore it
// later. The names must be unique and sorted by bytes.
func NewTrieWrite(names []string, w io.Writer) (*Trie, error) {
	tree, err := levenshtein.NewMinTreeWrite(names, w)
	if err != nil {
		return nil, err
	}
	return &Trie{tree: tree}, nil
}

// LoadTrie restores a Trie saved by NewTrieWrite.
func LoadTrie(r io.Reader) (*Trie, error) {
	tree, err := levenshtein.LoadMinTree(r)
	if err != nil {
		return nil, err
	}
	return &Trie{tree: tree}, nil
}

// Matches returns names that are within maxEditDist of the input. The
// distance is computed without the checks of EditDistance, so callers have
// to run EditDistance on the results.
func (t *Trie) Matches(input string, maxEditDist int) []string {
	return t.tree.FuzzyMatches(input, maxEditDist)
}
