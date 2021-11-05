package context

// An interface that allows to produce a normalized verion of a hierarchy
// as a slice of clades, ordered accorting from more general to more specific
// clades.
type Hierarch interface {
	// Clades method produces a slice of clades that represent a path in a
	// hierarchy.
	Clades() []Clade
}
