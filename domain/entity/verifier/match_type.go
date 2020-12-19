package verifier

import (
	"encoding/json"
	"fmt"
)

// MatchTypeValue describes how a name-string matched a name in gnames database.
type MatchTypeValue int

const (
	// NoMatch means that verification failed.
	NoMatch MatchTypeValue = iota

	// Exact means either canonical form, or the whole name-string matched
	// perfectlly.
	Exact

	// Fuzzy means that matches were not exact due to similarity of name-strings,
	// OCR or typing errors. Take these results with more suspition than
	// Exact matches. Fuzzy match is never done on uninomials due to the
	// high rate of false positives.
	Fuzzy

	// PartialExact used if GNames failed to match full name string. Now the match
	// happened by removing either middle species epithets, or by choppping the
	// 'tail' words of the input name-string canonical form.
	PartialExact

	// PartialFuzzy is the same as PartialExact, but also the match was not
	// exact. We never do fuzzy matches for uninomials, due to high rate of false
	// positives.
	PartialFuzzy
)

var mapMatchType = map[int]string{
	0: "NoMatch",
	1: "Exact",
	2: "Fuzzy",
	3: "PartialExact",
	4: "PartialFuzzy",
}

var mapMatchTypeStr = map[string]MatchTypeValue{
	"NoMatch":      NoMatch,
	"Exact":        Exact,
	"Fuzzy":        Fuzzy,
	"PartialExact": PartialExact,
	"PartialFuzzy": PartialFuzzy,
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
	return json.Marshal(mt.String())
}

// UnmarshalJSON implements json.Unmarshaller interface and converts a
// string into MatchType.
func (mt *MatchTypeValue) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if m, ok := mapMatchTypeStr[s]; ok {
		*mt = m
		return nil
	}
	return fmt.Errorf("cannot convert '%s' to MatchType", s)
}
