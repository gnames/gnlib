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
	// The results are sorted by score, not by data-source. The top result is
	// the the best result.
	WithAllMatches bool `json:"withAllMatches"`

	// Vernaculars field provides the list of languages to limit vernacular
	// names result to only these languages. If 'all' is provided instead, all
	// languages are included. Empty list means that no vernacular names will be
	// returned.
	Vernaculars []string `json:"vernaculars"`

	// WithCapitalization flag; when true, the first rune of low-case
	// input name-strings will be capitalized if appropriate.
	WithCapitalization bool `json:"withCapitalization"`

	// WithSpeciesGroup flag; when true, species names also get matched by
	// their species group. It means that the request will take in account
	// botanical autonyms and zoological coordinated names.
	WithSpeciesGroup bool `json:"withSpeciesGroup"`

	// WithRelaxedFuzzyMatch flag; when true, the fuzzy matching rules are
	// relaxed. Normally it is switched off to decrease the number of false
	// positives and make verification faster.
	WithRelaxedFuzzyMatch bool `json:"withRelaxedFuzzyMatch"`

	// WithUninomialFuzzyMatch flag; when true, uninomial names are not
	// restricted from fuzzy matching. Normally it creates too many false
	// positives and is switched off.
	WithUninomialFuzzyMatch bool `json:"withUninomialFuzzyMatch"`

	// WithStats flag; when true, results will return the most prevalent
	// kingdom for the text, as well as the taxon which contains a given
	// percentage of all names in the text (MainTaxon).
	//
	// For example MainTaxon with the MainTaxonThreshold of 0.5 would correspond
	// to a taxon that contains at least half of all names. We use the
	// managerial classification of Catalogue of Life for the MainTaxon
	// calculation.
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
	// Meta is the metadata of the request results.
	Meta `json:"metadata"`
	// Names are results of name-verification.
	Names []Name `json:"names"`
}

// Meta is metadata of the request. It provides information about parameters
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

	// WithCapitalization is true if there was a request to capitalize input
	WithCapitalization bool `json:"withCapitalization,omitempty"`

	// WithSpeciesGroup is true if Input included `WithSpeciesGroup` option.
	WithSpeciesGroup bool `json:"withSpeciesGroup,omitempty"`

	// WithRelaxedFuzzyMatch is true if Input included `WithRelaxedFuzzyMatch`
	// option. It means that the fuzzy matching rules are relaxed.
	// Normally it is switched off to decrease the number of false positives.
	WithRelaxedFuzzyMatch bool `json:"withRelaxedFuzzyMatch,omitempty"`

	// WithUninomialFuzzyMatch is true when it when uninomial names go
	// through fuzzy matching. Normally it is switched off to decrease the
	// number of false positives.
	WithUninomialFuzzyMatch bool `json:"withUninomialFuzzyMatch,omitempty"`

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

// NameStringInput is used to get information about a particular name-string.
type NameStringInput struct {
	// ID is the UUID v5 generated from a name-string
	ID string

	// DataSources is a slice of DataSourceIDs which should be used in the
	// output. If the slice is empty, all DataSources are used.
	DataSources []int

	// WithAllMatches controls if only the BestMatch, or all possible matches
	// are returned.
	WithAllMatches bool
}

// NameStringOutput contains data corresponding to the provided name-string ID.
type NameStringOutput struct {
	// NameStringMeta contains metadata from the input.
	NameStringMeta `json:"meta"`

	// Name is the found name data.
	*Name `json:"name"`
}

// NameStringMeta contains metadata from the provided input.
type NameStringMeta struct {
	// ID is the UUID v5 generated for a particular name-string.
	ID string `json:"id"`

	// DataSources is a slice of DataSource IDs. If it is not empty,
	// the output results will be constrained to these IDs.
	DataSources []int `json:"dataSources,omitempty"`

	// WithAllMatches indicates if all matches should be returned, or
	// only the best matches.
	WithAllMatches bool `json:"withAllMatches"`
}
