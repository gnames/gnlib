package verifier

import (
	"encoding/json"
	"fmt"
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

func (c CurationLevel) String() string {
	if match, ok := mapCurationLevel[int(c)]; ok {
		return match
	}
	return "N/A"
}

// MarshalJSON implements json.Marshaller interface and converts MatchType
// into a string.
func (c *CurationLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON implements json.Unmarshaller interface and converts a
// string into MatchType.
func (c *CurationLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if v, ok := mapCurationLevelStr[s]; ok {
		*c = v
		return nil
	}
	return fmt.Errorf("cannot convert '%s' to CurationLevel", s)
}
