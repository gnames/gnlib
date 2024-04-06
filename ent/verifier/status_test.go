package verifier_test

import (
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/verifier"
	"github.com/stretchr/testify/assert"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		name string
		in   verifier.TaxonomicStatus
		want string
	}{
		{"Unknown", verifier.UnknownTaxStatus, "N/A"},
		{"Accepted", verifier.AcceptedTaxStatus, "Accepted"},
		{"Synonym", verifier.SynonymTaxStatus, "Synonym"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewStatus(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want verifier.TaxonomicStatus
	}{
		{"Unknown", "N/A", verifier.UnknownTaxStatus},
		{"Accepted", "Accepted", verifier.AcceptedTaxStatus},
		{"Synonym", "Synonym", verifier.SynonymTaxStatus},
		{"Unknown", "Unknown", verifier.UnknownTaxStatus},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := verifier.New(tt.in)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestJSONStatus(t *testing.T) {
	type testData struct {
		Field1 string                     `json:"fieldOne"`
		Field2 []verifier.TaxonomicStatus `json:"fieldTwo"`
	}
	test := testData{
		Field1: "hello",
		Field2: []verifier.TaxonomicStatus{
			verifier.AcceptedTaxStatus,
			verifier.SynonymTaxStatus,
			verifier.UnknownTaxStatus,
		},
	}
	enc := gnfmt.GNjson{}
	res, err := enc.Encode(test)
	assert.Nil(t, err)
	assert.Equal(t, string(res), `{"fieldOne":"hello","fieldTwo":["Accepted","Synonym","N/A"]}`)
	var res2 testData
	err = enc.Decode(res, &res2)
	assert.Nil(t, err)
}
