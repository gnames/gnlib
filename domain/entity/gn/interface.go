package gn

// Versioner finds version of a project
type Versioner interface {
	GetVersion() Version
}
