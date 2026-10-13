// package abbr describes entities for the abbreviated names endpoint.
// The endpoint receives names with an abbreviated genus, such as
// `C. elegans`, that share the same genus abbreviation, and returns
// candidate genera that explain most of these names.
package abbr

// Input contains a group of abbreviated names and parameters that
// control how candidate genera are found and ranked.
type Input struct {
	// Names contains abbreviated names as found in a text. All names must
	// have the same genus abbreviation, which is capitalized and ends with
	// a period, for example `C.` or `Cae.`. Otherwise the request is an
	// error.
	Names []Name `json:"names"`

	// MainTaxon is the lowest taxon that contains most of the names of
	// the source, as calculated by GNverifier statistics.
	MainTaxon string `json:"mainTaxon,omitempty"`

	// DominantClass is the most common class of the names of the source,
	// as calculated by GNverifier statistics.
	DominantClass string `json:"dominantClass,omitempty"`

	// Tolerance is the percentage (0-100) of names that may be unmatched
	// for a genus to stay a candidate. The number of allowed unmatched
	// names is rounded down. If nil, the default is used. Zero requires
	// all names to match.
	Tolerance *int `json:"tolerance,omitempty"`

	// MaxGenusSize is the largest genus, in distinct epithets, that is
	// checked for misspellings. If zero, the default is used.
	MaxGenusSize int `json:"maxGenusSize,omitempty"`

	// GeneraLimit is the maximum number of candidate genera in the result.
	// If zero, the default is used.
	GeneraLimit int `json:"generaLimit,omitempty"`
}

// Name is an abbreviated name with an optional client ID and genus hints.
type Name struct {
	// ID is an optional identifier provided by the client. It allows to
	// map results back to positions in a text, because the same name can
	// occur many times.
	ID string `json:"id,omitempty"`

	// Name is the abbreviated name as found in a text, for example
	// `C. elegans`.
	Name string `json:"name"`

	// GenusHints contains full genera found near the name, sorted by
	// closeness, the closest first. Hints never exclude candidates, they
	// only change the ranking.
	GenusHints []string `json:"genusHints,omitempty"`
}

// Output contains candidate genera for a group of abbreviated names.
type Output struct {
	// Meta contains the metadata of the request results.
	Meta `json:"metadata"`

	// Genera contains candidate genera sorted by their total score. If
	// nothing matched, it is empty.
	Genera []Genus `json:"genera"`
}

// Meta contains metadata about the request, including parameters as they
// were applied.
type Meta struct {
	// Names contains the input names and their client IDs.
	Names []Name `json:"names"`

	// Abbreviation is the normalized genus abbreviation of the names.
	Abbreviation string `json:"abbreviation,omitempty"`

	// MainTaxon is the MainTaxon from the input.
	MainTaxon string `json:"mainTaxon,omitempty"`

	// DominantClass is the DominantClass from the input.
	DominantClass string `json:"dominantClass,omitempty"`

	// Tolerance is the percentage of names that may be unmatched for a
	// genus, as applied.
	Tolerance int `json:"tolerance"`

	// MaxGenusSize is the largest genus checked for misspellings, as
	// applied.
	MaxGenusSize int `json:"maxGenusSize"`

	// GeneraLimit is the maximum number of candidate genera, as applied.
	GeneraLimit int `json:"generaLimit"`

	// NamesNumber is the number of submitted names.
	NamesNumber int `json:"namesNumber"`

	// MatchedNamesNumber is the number of names matched by at least one
	// candidate genus.
	MatchedNamesNumber int `json:"matchedNamesNumber"`

	// UnmatchedNamesNumber is the number of names not matched by any
	// candidate genus.
	UnmatchedNamesNumber int `json:"unmatchedNamesNumber"`

	// Error contains an error message, if any. If Error is not empty,
	// no genera are returned.
	Error string `json:"error,omitempty"`
}

// Genus is a candidate genus for a group of abbreviated names.
type Genus struct {
	// Name is the name of the genus.
	Name string `json:"name"`

	// ClassificationPath contains the classification of the genus as
	// a hierarchy of names separated by `|`.
	ClassificationPath string `json:"classificationPath,omitempty"`

	// ClassificationRanks contains ranks of the ClassificationPath,
	// separated by `|`.
	ClassificationRanks string `json:"classificationRanks,omitempty"`

	// EpithetsNumber is the size of the genus, as the number of its
	// distinct epithet stems.
	EpithetsNumber int `json:"epithetsNumber"`

	// OverMaxSize is true if EpithetsNumber is larger than MaxGenusSize.
	// Misspellings are not checked for such genera.
	OverMaxSize bool `json:"overMaxSize,omitempty"`

	// Score is the score of the genus and its parts.
	Score Score `json:"score"`

	// MatchedNames contains input names matched by the genus.
	MatchedNames []MatchedName `json:"matchedNames"`

	// UnmatchedNames contains input names not matched by the genus,
	// within the Tolerance. Genus hints are not included.
	UnmatchedNames []Name `json:"unmatchedNames,omitempty"`
}

// Score contains weighted parts of the score of a candidate genus and
// their total. The parts add up to the total.
type Score struct {
	// MatchedNames is the part of the score that comes from matched names.
	// A name gives less to a large genus, or if its epithet is common.
	MatchedNames float64 `json:"matchedNames"`

	// GenusHints is the part of the score that comes from genus hints.
	GenusHints float64 `json:"genusHints"`

	// UnmatchedNames is the part of the score that comes from unmatched
	// names. It is zero or negative.
	UnmatchedNames float64 `json:"unmatchedNames"`

	// DominantClass is the part of the score that comes from the genus
	// being in the DominantClass.
	DominantClass float64 `json:"dominantClass"`

	// MainTaxon is the part of the score that comes from the genus being
	// in the MainTaxon.
	MainTaxon float64 `json:"mainTaxon"`

	// Total is the total score of the genus.
	Total float64 `json:"total"`
}

// MatchedName is an input name matched by a candidate genus.
type MatchedName struct {
	// ID is the client ID of the input name.
	ID string `json:"id,omitempty"`

	// Name is the input name.
	Name string `json:"name"`

	// MatchedCanonical is the full canonical form the name matched,
	// for example `Panthera tigris altaica`.
	MatchedCanonical string `json:"matchedCanonical"`

	// MatchType tells how the name matched.
	MatchType MatchTypeValue `json:"matchType"`

	// StemMatch is true if an epithet ending matched only by its stem,
	// for example `alba` and `albus`.
	StemMatch bool `json:"stemMatch,omitempty"`
}
