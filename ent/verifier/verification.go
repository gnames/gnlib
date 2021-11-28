package verifier

// Input is options/parameters for the Verify method.
type Input struct {
	// NameStrings is a list of name-strings to verify.
	NameStrings []string `json:"nameStrings"`

	// PreferredSources contain DataSources IDs whos matches will be returned
	// becides the best result. See PreferredResults field in Verirication.
	// If Preferred sources are []int{0}, then all matched Sources are used.
	PreferredSources []int `json:"preferredSources"`

	// ContextThreshold sets the minimal percentage of names in a clade
	// to be counted as a context of a text.
	//
	// Context is a clade that contains at least ContextThreshold percentage
	// of all names in the text. We use the managerial classification of
	// Catalogue of Life for the context calculation.
	ContextThreshold float32

	// WithAllMatches indicates that all matches per data-source are returned,
	// sorted by score (instead of the best match per source).
	WithAllMatches bool `json:"withAllMatches"`

	// WithVernaculars indicates if corresponding vernacular results will be
	// returned as well.
	WithVernaculars bool `json:"withVernaculars"`

	// WithCapitalization flag; when true, the first rune of low-case
	// input name-strings will be capitalized if appropriate.
	WithCapitalization bool `json:"withCapitalization"`

	// WithContext flag; when true, results will return the most prevalent
	// kingdom for the text, as well as the clade which contains a given
	// percentage of all names in the text.
	//
	// For examplle context with threshold 0.5 would correspond to a clade that
	// contains at least half of all names. We use the managerial classification
	// of Catalogue of Life for the context calculation.
	WithContext bool
}

// Verification is a result returned by Verify method.
type Verification struct {
	// Meta is metadata of the request.
	Meta `json:"metadata"`
	// Names from the request.
	Names []*Name `json:"names"`
}

// Meta is metadata of the request. It provides intofmation about parameters
// used for the request, and, optionally give information about the kingdom
// that contains most of the names from the request, as well as the lowest
// clade that contains majority of the names.
type Meta struct {
	// NamesNum is the number of name-strings in the request.
	NamesNum int `json:"namesNum"`

	// Number of names qualified for context/kingdoms calculation
	ContextNamesNum int `json:"contextNamesNum,omitempty"`

	// WithAllSources indicates if preferred results include all matched
	// sources.
	WithAllSources bool `json:"allSources"`

	// WithAllMatches indicates if response provides more then one result
	// per source, if such results were found.
	WithAllMatches bool `json:"allMatches"`

	// WithContext indicates that the kingdom and convergence clade that contain
	// majority of names will be calculated.
	WithContext bool `json:"withContext"`

	// InputCapitalized is true, if the was a request to capitalize input
	InputCapitalized bool `json:"inputCapitalized,omitempty"`

	// ContextThreshold provides a minimal percentage names that a clade should
	// have to be qualified as a Context clade.
	ContextThreshold float32 `json:"contextThreshold,omitempty"`

	// PreferredSources provides IDs of data-sources from the request.
	PreferredSources []int `json:"preferredSources,omitempty"`

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

	// Kingdoms provides distribution of names over the kingdoms
	Kingdoms []Kingdom

	// Context provides the lowest clade that contains most of names from
	// the request.
	//
	// Non-matched names, or names that are not in Catalogue of Life are
	// not part of the calculation.
	Context string `json:"context,omitempty"`

	// ContextPercentage indicates the percentage of names that are placed
	// in the "context" clade. This number should be higher than
	// ContexThreshold unless Context is empty.
	ContextPercentage float32 `json:"contextPercentage,omitempty"`
}

type Kingdom struct {
	Name       string
	NamesNum   int
	Percentage float32
}
