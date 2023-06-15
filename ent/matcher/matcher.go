/*
	Package matcher provides main data-structures that describe input and

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
	Names []string `json:"names"`

	// WithSpeciesGroup -- when true, species are searched within species group.
	// It means that autonyms in botany and coordination names in zoology.
	WithSpeciesGroup bool `json:"withSpeciesGroup,omitempty"`

	// WithUninomialFuzzyMatch -- when true, the uninomials go through
	// fuzzy matching together with bi- and tri-nomials.
	WithUninomialFuzzyMatch bool `json:"withUninomialFuzzyMatch,omitempty"`

	// DataSources -- is a list of data-sources that are used to search
	// a for a name-string
	DataSources []int `json:"dataSources,omitempty"`
}

type Output struct {
	// Meta is the metadata of the request results.
	Meta `json:"metadata"`
	// Matches contains results of name-matching.
	Matches []Match `json:"matches"`
}

// Meta contains metadata about the matching request.
type Meta struct {
	// NamesNum is the number of name-strings sent for matching.
	NamesNum int `json:"namesNum"`

	// WithSpeciesGroup is set to the value of the `Input`'s option
	// `WithSpeciesGroup`.
	WithSpeciesGroup bool `json:"withSpeciesGroup,omitempty"`

	// DataSources is set to the value of the `Input`'s option
	// 'WithSpeciesGroup'.
	DataSources []int `json:"dataSources,omitempty"`
}

// Match represents match results for one name-string.
type Match struct {
	// ID is a UUIDv5 string generated from `Name` field.
	ID string `json:"id"`

	// Name is a verbatim name-string from `Input.Names`.
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

	// DataSourcesMap is a set of data-sources that have this particular
	// MatchItem.
	DataSourcesMap map[int]struct{} `json:"-"`

	// DataSources is an array of data-sources that have this particular
	// MatchItem
	DataSources []int `json:"dataSources,omitempty"`
}
