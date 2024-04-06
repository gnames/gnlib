package verifier

import (
	"errors"
	"strings"
)

type TaxonomicStatus int

const (
	UnknownTaxStatus TaxonomicStatus = iota
	AcceptedTaxStatus
	SynonymTaxStatus
)

var txStatusMap = map[string]TaxonomicStatus{
	"N/A":      UnknownTaxStatus,
	"Accepted": AcceptedTaxStatus,
	"Synonym":  SynonymTaxStatus,
}

var txStatusStringMap = func() map[TaxonomicStatus]string {
	res := make(map[TaxonomicStatus]string)
	for k, v := range txStatusMap {
		res[v] = k
	}
	return res
}()

func New(txStatus string) TaxonomicStatus {
	if res, ok := txStatusMap[txStatus]; ok {
		return res
	}
	return UnknownTaxStatus
}

func (ts TaxonomicStatus) String() string {
	return txStatusStringMap[ts]
}

// MarshalJSON implements json.Marshaller interface and converts MatchType
// into a string.
func (ts TaxonomicStatus) MarshalJSON() ([]byte, error) {
	return []byte("\"" + ts.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller interface and converts a
// string into MatchType.
func (ts *TaxonomicStatus) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	s := strings.Trim(string(bs), `"`)
	*ts, ok = txStatusMap[s]
	if !ok {
		err = errors.New("cannot decode as a TaxonomicStatus")
	}
	return err
}
