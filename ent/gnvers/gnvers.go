package gnvers

// @Description Version provides information about the version
// @Description of an application.
type Version struct {
	// Version specifies the version of the app, usually in the v0.0.0 format.
	Version string `json:"version" example:"v1.0.2"`

	// Build contains the timestamp or other details
	// indicating when the app was compiled.
	Build string `json:"build" example:"2023-08-03_18:58:38UTC"`
}
