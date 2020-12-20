package verifier

import (
	"errors"
	"strings"
)

// CurationLevel tells if matched result was returned by at least one
// DataSource in the following categories.
type CurationLevel int

const (
	// NotCurated means that all DataSources where the name-string was matched
	// are not curated sufficiently.
	NotCurated CurationLevel = iota

	// AutoCurated means that at least one of the returned DataSources invested
	// significantly in curating their data by scripts.
	AutoCurated

	// Curated means that at least one DataSource is marked as sufficiently
	// curated. It does not mean that the particular match was manually checked
	// though.
	Curated
)

var mapCurationLevel = map[int]string{
	0: "NotCurated",
	1: "AutoCurated",
	2: "Curated",
}

var mapCurationLevelStr = map[string]CurationLevel{
	"NotCurated":  NotCurated,
	"AutoCurated": AutoCurated,
	"Curated":     Curated,
}

func (cl CurationLevel) String() string {
	if match, ok := mapCurationLevel[int(cl)]; ok {
		return match
	}
	return "N/A"
}

// MarshalJSON implements json.Marshaller interface and converts MatchType
// into a string.
func (cl CurationLevel) MarshalJSON() ([]byte, error) {
	return []byte("\"" + cl.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller interface and converts a
// string into MatchType.
func (cl *CurationLevel) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	s := strings.Trim(string(bs), `"`)
	*cl, ok = mapCurationLevelStr[s]
	if !ok {
		err = errors.New("cannot decode as a CurationLevel")
	}
	return err
}
