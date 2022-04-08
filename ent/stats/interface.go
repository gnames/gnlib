package stats

// An interface that allows to produce a normalized verion of a hierarchy
// as a slice of taxons, ordered accorting from more general to more specific
// taxons.
type Hierarchy interface {
	// Taxons method produces a slice of taxons that represent a path in a
	// hierarchy.
	Taxons() []Taxon
}
