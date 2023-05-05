package reconciler

// Manifest describes metadata of a W3C Reconciliation Service.
type Manifest struct {
	// Versions returns the versions of W3C Reconciliation API descrived at
	// https://www.w3.org/community/reports/reconciliation/CG-FINAL-specs-0.2-20230410 .
	// Versions can have "1.0" and "2.0" elements. For our purposes
	// it should be set to ["2.0"].
	Versions []string `json:"versions"`

	// Name of the reconciliation service. Should be set to "GlobalNames".
	Name string `json:"name"`

	Preview `json:"preview"`
	View    `json:"view"`

	// IdentifierSpace contans the URI prefix of the reconciliation service.
	// For example "https://verifier.globalnames.org/api/v1/name_strings/"
	IdentifierSpace string `json:"identifierSpace"`

	// SchemaSpace provides the URL pointing to the schema of an entity.
	SchemaSpace string `json:"schemaSpace"`

	// DefaultTypes used for a reconciliation queries.
	DefaultTypes []Type `json:"defaultTypes"`
}

// Type describes types available for the reconciliation service.
type Type struct {
	// ID is a unique identifier for the type.
	ID string `json:"id"`

	// Name is a human friendly short description of the type.
	Name string `json:"name"`
}

type Preview struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type View struct {
	URL string `json:"url"`
}
