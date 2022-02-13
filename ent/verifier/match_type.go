package verifier

import (
	"errors"
	"strings"
)

// MatchTypeValue describes how a name-string matched a name in gnames database.
type MatchTypeValue int

const (
	// NoMatch means that matching failed.
	NoMatch MatchTypeValue = iota

	// PartialFuzzy is the same as PartialExact, but also the match was not
	// exact. We never do fuzzy matches for uninomials, due to high rate of false
	// positives.
	PartialFuzzy

	// PartialExact used if GNames failed to match full name string. Now the match
	// happened by removing either middle species epithets, or by choppping the
	// 'tail' words of the input name-string canonical form.
	PartialExact

	// Fuzzy means that matches were not exact due to similarity of name-strings,
	// OCR or typing errors. Take these results with more suspition than
	// Exact matches. Fuzzy match is never done on uninomials due to the
	// high rate of false positives.
	Fuzzy

	// Exact means either canonical form, or the whole name-string matched
	// perfectlly.
	Exact

	// Virus names are matched in the database. `Virus` is a wide
	// term and includes a variety of non-cellular terms (virus, prion, plasmid,
	// vector etc.)
	Virus

	// FacetedSearch is a match made by search procedure. It does not happen
	// during verification.
	FacetedSearch
)

var mapMatchType = map[int]string{
	0: "NoMatch",
	1: "PartialFuzzy",
	2: "PartialExact",
	3: "Fuzzy",
	4: "Exact",
	5: "Virus",
	6: "FacetedSearch",
}

var mapMatchTypeStr = map[string]MatchTypeValue{
	"NoMatch":       NoMatch,
	"Virus":         Virus,
	"Exact":         Exact,
	"Fuzzy":         Fuzzy,
	"PartialExact":  PartialExact,
	"PartialFuzzy":  PartialFuzzy,
	"FacetedSearch": FacetedSearch,
}

// NewMatchType takes a string and converts it into a MatchType. If
// the string is unkown, it returns NoMatch type.
func NewMatchType(t string) MatchTypeValue {
	if match, ok := mapMatchTypeStr[t]; ok {
		return match
	}
	return NoMatch
}

// String implements fmt.String interface and returns a string representation
// of a MatchType. The returned string can be converted back to MatchType
// via NewMatchType function.
func (mt MatchTypeValue) String() string {
	if match, ok := mapMatchType[int(mt)]; ok {
		return match
	}
	return "N/A"
}

// MarshalJSON implements json.Marshaller interface and converts MatchType
// into a string.
func (mt MatchTypeValue) MarshalJSON() ([]byte, error) {
	return []byte("\"" + mt.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller interface and converts a
// string into MatchType.
func (mt *MatchTypeValue) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	s := strings.Trim(string(bs), `"`)
	*mt, ok = mapMatchTypeStr[s]
	if !ok {
		err = errors.New("cannot decode as a MatchType")
	}
	return err
}
