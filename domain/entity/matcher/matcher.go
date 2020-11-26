/* Package matcher provides main data-structures that describe input and
output of gnmatcher functionality.
*/
package matcher

import (
	vlib "github.com/gnames/gnlib/domain/entity/verifier"
)

// Match is output of MatchAry method.
type Match struct {
	// ID is UUIDv5 generated from verbatim input name-string.
	ID string `json:"id"`
	// Name is verbatim input name-string.
	Name string `json:"input"`
	// VirusMatch is true if matching
	VirusMatch bool `json:"isVirus,omitempty"`
	// MatchType describe what kind of match happened.
	MatchType vlib.MatchType `json:"matchType"`
	// MatchItems provide all matched data. It will be empty if no matches
	// occured.
	MatchItems []MatchItem `json:"matchItems,omitempty"`
}

// MatchItem describes one matched string and its properties.
type MatchItem struct {
	// ID is a UUIDv5 generated out of MatchStr.
	ID string `json:"id"`
	// MatchStr is the string that matched a particular input. More often than
	// not it is a canonical form of a name. However for viruses it
	// can be matched string from the database.
	MatchStr string `json:"matchString"`
	// EditDistance is a Levenshtein edit distance between normalized
	// input and MatchStr.
	EditDistance int `json:"editDistance"`
	// EditDistanceStem is a Levenshtein edit distance between stemmed input and
	// stemmed MatchStr.
	EditDistanceStem int `json:"editDistanceStem"`
}
