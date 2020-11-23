package gn

// Version is output from GetVersion method.
type Version struct {
	// Version of gnmatcher.
	Version string `json:"version"`
	// Build timestamp of gnmatcher.
	Build string `json:"build,omitempty"`
}
