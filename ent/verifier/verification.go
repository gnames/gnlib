package verifier

// Input is options/parameters for the Verify method.
type Input struct {
	// NameStrings is a list of name-strings to verify.
	NameStrings []string `json:"nameStrings"`

	// DataSources field contains DataSources IDs to limit results to only these
	// sources. The best result is calculated only out of this limited set of
	// data. By default only the BestResult is shown. To see all results use
	// WithAllMatches flag.
	DataSources []int `json:"dataSources"`

	// WithAllMatches provides all results, instead of only the BestResult.
	// The results are sorted by score, not by data-source.
	WithAllMatches bool `json:"withAllMatches"`

	// WithVernaculars indicates if corresponding vernacular results will be
	// returned as well.
	WithVernaculars bool `json:"withVernaculars"`

	// WithCapitalization flag; when true, the first rune of low-case
	// input name-strings will be capitalized if appropriate.
	WithCapitalization bool `json:"withCapitalization"`

	// WithStats flag; when true, results will return the most prevalent
	// kingdom for the text, as well as the taxon which contains a given
	// percentage of all names in the text (MainTaxon).
	//
	// For example MainTaxon with the MainTaxonThreshold of 0.5 would correspond
	// to a taxon that contains at least half of all names. We use the managerial
	// classification of Catalogue of Life for the MainTaxon calculation.
	WithStats bool `json:"withStats"`

	// MainTaxonThreshold sets the minimal percentage of names in a taxon
	// to be counted as a MainTaxon of a text. This field is ignored if
	// WithStats is false.
	//
	// MainTaxon is a taxon that contains at least MainTaxonThreshold percentage
	// of all names (genus and below) in the text. We use the managerial
	// classification of Catalogue of Life for the MainTaxon calculation.
	MainTaxonThreshold float32 `json:"mainTaxonThreshold"`
}

// Output is a result returned by Verify method.
type Output struct {
	// Meta is metadata of the request.
	Meta `json:"metadata"`
	// Names from the request.
	Names []Name `json:"names"`
}

// Meta is metadata of the request. It provides intofmation about parameters
// used for the request, and, optionally give information about the kingdom
// that contains most of the names from the request, as well as the lowest
// taxon that contains majority of the names.
type Meta struct {
	// NamesNumber is the number of name-strings in the request.
	NamesNumber int `json:"namesNumber"`

	// WithAllSources indicates if `Results` will include all matched
	// sources.
	WithAllSources bool `json:"withAllSources,omitempty"`

	// WithAllMatches indicates if response provides more then one result
	// per source, if such results were found.
	WithAllMatches bool `json:"withAllMatches,omitempty"`

	// WithStats indicates that the kingdom and a taxon that contain
	// majority of names (MainTaxon) will be calculated.
	WithStats bool `json:"withStats,omitempty"`

	// WithCapitalization is true, if the was a request to capitalize input
	WithCapitalization bool `json:"withCapitalization,omitempty"`

	// DataSources provides IDs of data-sources from the request.
	DataSources []int `json:"dataSources,omitempty"`

	// MainTaxonThreshold provides a minimal percentage names that a taxon should
	// have to be qualified as a MainTaxon.
	MainTaxonThreshold float32 `json:"mainTaxonThreshold,omitempty"`

	// StatsNamesNum is the number of names qualified for MainTaxon/Kingdoms
	// calculation.
	StatsNamesNum int `json:"statsNamesNum,omitempty"`

	// MainTaxon provides the lowest taxon that contains most of the names from
	// the request.
	//
	// Non-matched names, names that are not in the Catalogue of Life, names
	// higher than genus are not part of the calculation.
	MainTaxon string `json:"mainTaxon,omitempty"`

	// MainTaxonPercentage indicates the percentage of names that are placed
	// in the MainTaxon. This number should be higher than
	// MainTaxonThreshold unless MainTaxon is empty.
	MainTaxonPercentage float32 `json:"mainTaxonPercentage,omitempty"`

	// Kingdom provides what kingdom includes the majority of names from the
	// request accorging to the managerial classification of Catalogue of Life.
	//
	// Non-matched names, or names that are not in Catalogue of Life are
	// not part of the calculation.
	Kingdom string `json:"kingdom,omitempty"`

	// KingdomPercentage provides the percentage of names in the most
	// prevalent kingdom.
	//
	// Non-matched names, or names that are not in Catalogue of Life are
	// not part of the calculation.
	KingdomPercentage float32 `json:"kingdomPercentage,omitempty"`

	// Kingdoms provides all kingdoms with matched names and names distribution
	// between the kingdoms.
	Kingdoms []Kingdom `json:"kingdoms,omitempty"`
}

// Kingdom provides statistics of matched names found in a particular kingdom.
type Kingdom struct {
	// KingdomName is the name of a kingdom.
	KingdomName string `json:"kingdomName"`

	// NamesNumber is the number of names found in a kingdom.
	NamesNumber int `json:"namesNumber"`

	// Percentage is a percentage of names found in a kingdom.
	Percentage float32 `json:"percentage"`
}
