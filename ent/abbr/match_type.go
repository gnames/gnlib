package abbr

import (
	"errors"
	"strings"
)

// MatchTypeValue describes how an abbreviated name matched a name of a
// candidate genus.
type MatchTypeValue int

const (
	// NoMatch means that matching failed.
	NoMatch MatchTypeValue = iota

	// SpeciesLevel means that the input and the matched name have the
	// same species epithet, but their last epithets differ. The match
	// can be coarser (`P. tigris altaica` matched `Panthera tigris`),
	// finer (`P. tigris` matched `Panthera tigris altaica`), or another
	// infraspecies of the same species (`P. tigris altaica` matched
	// `Panthera tigris amurensis`). It is not the name from the text,
	// but it still shows the genus.
	SpeciesLevel

	// EpithetMoved means that the input and the matched name have the
	// same last epithet, but other epithets differ. The epithet moved to
	// another rank or species, for example `P. tigris altaica` matched
	// `Panthera altaica`, or `P. altaica` matched
	// `Panthera tigris altaica`.
	EpithetMoved

	// Full means that the input and the matched name have the same
	// epithets, for example `P. tigris altaica` matched
	// `Panthera tigris altaica`. Nominotypical repeats are ignored, so
	// `B. bubo` matches `Bubo bubo bubo` fully.
	Full
)

var mapMatchType = map[MatchTypeValue]string{
	NoMatch:      "NoMatch",
	SpeciesLevel: "SpeciesLevel",
	EpithetMoved: "EpithetMoved",
	Full:         "Full",
}

var mapMatchTypeStr = map[string]MatchTypeValue{
	"NoMatch":      NoMatch,
	"SpeciesLevel": SpeciesLevel,
	"EpithetMoved": EpithetMoved,
	"Full":         Full,
}

// NewMatchType takes a string and converts it into a MatchType. If
// the string is unknown, it returns NoMatch type.
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
	if match, ok := mapMatchType[mt]; ok {
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
