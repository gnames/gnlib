/* Package matcher provides main data-structures that describe input and
output of gnmatcher functionality.
*/
package matcher

import (
	vlib "github.com/gnames/gnlib/ent/verifier"
)

// Input for the matcher includes a slice of name-strings several options
// sent to the matcher.
type Input struct {
	// Names is a slice of name-strings.
	Names []string

	// WithSpeciesGroup -- when true, species are searched within species group.
	// It means that autonyms in botany and coordination names in zoology.
	WithSpeciesGroup bool
}

// Output is output of MatchAry method.
type Output struct {
	// ID is UUIDv5 generated from verbatim input name-string.
	ID string `json:"id"`

	// Name is verbatim input name-string.
	Name string `json:"input"`

	// MatchType describe what kind of match happened.
	MatchType vlib.MatchTypeValue `json:"matchType"`

	// MatchItems provide all matched data. It will be empty if no matches
	// occured.
	MatchItems []MatchItem `json:"matchItems,omitempty"`
}

// MatchItem describes one matched string and its properties.
type MatchItem struct {
	// ID is a UUIDv5 generated out of MatchStr.
	ID string `json:"id"`

	// InputStr is the string used for matching. Usually it is the canonical
	// form of the input name. However, if matching was partial, it is
	// the string that was created by partial matching algorithm.
	InputStr string `json:"inputString"`

	// MatchStr is the string that matched a particular input. More often than
	// not it is a canonical form of a name. However for viruses it
	// can be matched string from the database.
	MatchStr string `json:"matchString"`

	// MatchType describe what kind of match happened.
	MatchType vlib.MatchTypeValue `json:"matchType"`

	// EditDistance is a Levenshtein edit distance between
	// InputStr and MatchStr.
	EditDistance int `json:"editDistance"`

	// EditDistanceStem is a Levenshtein edit distance between stemmed
	// InputStr and stemmed MatchStr.
	EditDistanceStem int `json:"editDistanceStem"`
}
