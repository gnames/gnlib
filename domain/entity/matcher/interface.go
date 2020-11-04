package matcher

import "github.com/gnames/gnlib/domain/entity/gn"

// Matcher describes methods required for matching name-strings to names.
type Matcher interface {
	// Versioner interface gets version of a project
	gn.Versioner

	// MatchAry takes a list of strings and matches each of them
	// to known scientific names.
	MatchAry(names []string) []*Match
}

// FuzzyMatcher describes methods needed for fuzzy matching
type FuzzyMatcher interface {
	// MatchStem takes a stemmed scientific name and max edit distance.
	// The search stops if current edit distance becomes bigger than edit
	// distance. The method returns 0 or more stems that did match the
	// input stem within the edit distance constraint.
	MatchStem(stem string, maxEditDistance int) []string
	// StemToCanonicals takes a stem and returns back canonicals
	// that correspond to that stem.
	StemToMatchItems(stem string) []MatchItem
}
