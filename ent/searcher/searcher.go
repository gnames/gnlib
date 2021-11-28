package searcher

import (
	"github.com/gnames/gnlib/ent/verifier"
	"github.com/gnames/gnquery/ent/search"
)

type Searcher struct {
	Meta       `json:"metadata"`
	Canonicals []Canonical
}

type Meta struct {
	search.Input
}

type Canonical struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	MatchType        string                 `json:"matchType"`
	BestResult       *verifier.ResultData   `json:"bestResult,omitempty"`
	PreferredResults []*verifier.ResultData `json:"preferredResults,omitempty"`
}
