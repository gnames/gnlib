package verifier

// Input contains options and parameters for the Verify method.
type Input struct {
	// NameStrings contains the list of name-strings to verify.
	NameStrings []string `json:"nameStrings"`

	// DataSources contains data source IDs to limit results to only these
	// sources. The best result is calculated only from this limited set of
	// data. By default only the best result is shown. To see all results, use
	// the WithAllMatches flag. To see all results with the highest score use
	// the WithAllBestResults flag.
	DataSources []int `json:"dataSources"`

	// WithAllMatches when true, returns all results instead of only the best result.
	// The results are sorted by score, not by data source. The top result is
	// the best result.
	WithAllMatches bool `json:"withAllMatches"`

	// WithAllBestResults flag; when true, returns all results with the highest
	// score instead of just one best result. This provides multiple equally
	// good matches when they exist. This flag is ignored when WithAllMatches is
	// true.
	WithAllBestResults bool `json:"withAllBestResults"`

	// Vernaculars contains the list of languages to limit vernacular
	// name results to only these languages. If 'all' is provided as a 'language',
	//  all languages are included. An empty list means no vernacular names will be
	// returned.
	Vernaculars []string `json:"vernaculars"`

	// WithCapitalization when true, capitalizes the first letter of lowercase
	// input name-strings when appropriate.
	WithCapitalization bool `json:"withCapitalization"`

	// WithSpeciesGroup when true, enables species names to be matched by
	// their species group. This includes botanical autonyms and zoological
	// coordinated names in the matching process.
	WithSpeciesGroup bool `json:"withSpeciesGroup"`

	// WithRelaxedFuzzyMatch when true, relaxes the fuzzy matching rules.
	// Normally disabled to reduce false positives and improve verification speed.
	WithRelaxedFuzzyMatch bool `json:"withRelaxedFuzzyMatch"`

	// WithUninomialFuzzyMatch when true, allows uninomial names to undergo
	// fuzzy matching. Normally disabled as it generates too many false positives.
	WithUninomialFuzzyMatch bool `json:"withUninomialFuzzyMatch"`

	// WithStats when true, includes statistical information in the results,
	// such as the most prevalent kingdom and the taxon containing a specified
	// percentage of all names (MainTaxon).
	//
	// For example, a MainTaxon with MainTaxonThreshold of 0.5 corresponds
	// to a taxon containing at least half of all names. The managerial
	// classification of Catalogue of Life is used for MainTaxon calculations.
	WithStats bool `json:"withStats"`

	// MainTaxonThreshold sets the minimum percentage of names a taxon must
	// contain to qualify as the MainTaxon. This field is ignored if
	// WithStats is false.
	//
	// MainTaxon is a taxon containing at least MainTaxonThreshold percentage
	// of all names (genus level and below). The managerial classification
	// of Catalogue of Life is used for MainTaxon calculations.
	MainTaxonThreshold float32 `json:"mainTaxonThreshold"`
}

// Output contains the result returned by the Verify method.
type Output struct {
	// Meta contains the metadata of the request results.
	Meta `json:"metadata"`
	// Names contains the results of name verification.
	Names []Name `json:"names"`
}

// Meta contains metadata about the request, including information about parameters
// used for the request. Optionally provides information about the kingdom
// containing most names from the request, as well as the lowest
// taxon containing the majority of the names.
type Meta struct {
	// NamesNumber contains the total number of name-strings in the request.
	NamesNumber int `json:"namesNumber"`

	// Vernaculars contains the languages that were requested for finding
	// vernacular names.
	Vernaculars []string `json:"vernaculars,omitempty"`

	// WithAllSources indicates whether results will include all matched
	// data sources.
	WithAllSources bool `json:"withAllSources,omitempty"`

	// WithAllMatches indicates whether the response provides more than one result
	// per source when such results are found.
	WithAllMatches bool `json:"withAllMatches,omitempty"`

	// WithAllBestResults indicates that the response will return all
	// results with the highest score. This flag should not be shown
	// when WithAllMatches is true, as all best results will be
	// already included in all matches.
	WithAllBestResults bool `json:"withAllBestResults,omitempty"`

	// WithStats indicates that statistical information including the kingdom
	// and taxon containing the majority of names (MainTaxon) will be calculated.
	WithStats bool `json:"withStats,omitempty"`

	// WithCapitalization is true if there was a request to capitalize input
	// name-strings.
	WithCapitalization bool `json:"withCapitalization,omitempty"`

	// WithSpeciesGroup is true if the input included the WithSpeciesGroup option.
	WithSpeciesGroup bool `json:"withSpeciesGroup,omitempty"`

	// WithRelaxedFuzzyMatch is true if the input included the WithRelaxedFuzzyMatch
	// option, meaning the fuzzy matching rules are relaxed.
	// Normally disabled to reduce false positives.
	WithRelaxedFuzzyMatch bool `json:"withRelaxedFuzzyMatch,omitempty"`

	// WithUninomialFuzzyMatch is true when uninomial names undergo
	// fuzzy matching. Normally disabled to reduce false positives.
	WithUninomialFuzzyMatch bool `json:"withUninomialFuzzyMatch,omitempty"`

	// DataSources contains the IDs of data sources from the request.
	DataSources []int `json:"dataSources,omitempty"`

	// MainTaxonThreshold contains the minimum percentage of names a taxon
	// must have to qualify as the MainTaxon.
	MainTaxonThreshold float32 `json:"mainTaxonThreshold,omitempty"`

	// StatsNamesNum contains the number of names that qualified for MainTaxon
	// and Kingdom calculations.
	StatsNamesNum int `json:"statsNamesNum,omitempty"`

	// MainTaxon contains the lowest taxon that includes most of the names from
	// the request.
	//
	// Non-matched names, names not in the Catalogue of Life, and names
	// higher than genus level are not included in the calculation.
	MainTaxon string `json:"mainTaxon,omitempty"`

	// MainTaxonPercentage contains the percentage of names placed
	// in the MainTaxon. This value should be higher than
	// MainTaxonThreshold unless MainTaxon is empty.
	MainTaxonPercentage float32 `json:"mainTaxonPercentage,omitempty"`

	// Kingdom contains the kingdom that includes the majority of names from the
	// request according to the managerial classification of Catalogue of Life.
	//
	// Non-matched names and names not in the Catalogue of Life are
	// not included in the calculation.
	Kingdom string `json:"kingdom,omitempty"`

	// KingdomPercentage contains the percentage of names in the most
	// prevalent kingdom.
	//
	// Non-matched names and names not in the Catalogue of Life are
	// not included in the calculation.
	KingdomPercentage float32 `json:"kingdomPercentage,omitempty"`

	// Kingdoms contains all kingdoms with matched names and their
	// distribution statistics.
	Kingdoms []Kingdom `json:"kingdoms,omitempty"`
}

// Kingdom contains statistics for matched names found in a particular kingdom.
type Kingdom struct {
	// KingdomName contains the name of the kingdom.
	KingdomName string `json:"kingdomName"`

	// NamesNumber contains the number of names found in this kingdom.
	NamesNumber int `json:"namesNumber"`

	// Percentage contains the percentage of names found in this kingdom.
	Percentage float32 `json:"percentage"`
}

// NameStringInput contains parameters for retrieving information about a particular name-string.
type NameStringInput struct {
	// ID contains the UUID v5 generated from a name-string.
	ID string

	// DataSources contains data source IDs to be used in the
	// output. If empty, all data sources are used.
	DataSources []int

	// WithAllMatches controls whether only the best match or all possible matches
	// are returned.
	WithAllMatches bool
}

// NameStringOutput contains the data corresponding to the provided name-string ID.
type NameStringOutput struct {
	// NameStringMeta contains the metadata from the input.
	NameStringMeta `json:"meta"`

	// Name contains the found name data.
	*Name `json:"name"`
}

// NameStringMeta contains the metadata from the provided input.
type NameStringMeta struct {
	// ID contains the UUID v5 generated for a particular name-string.
	ID string `json:"id"`

	// DataSources contains data source IDs. If not empty,
	// the output results will be limited to these IDs.
	DataSources []int `json:"dataSources,omitempty"`

	// WithAllMatches indicates whether all matches should be returned
	// or only the best matches.
	WithAllMatches bool `json:"withAllMatches"`
}
