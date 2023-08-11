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

	Extend `json:"extend"`

	// IdentifierSpace contans the URI prefix of the reconciliation service.
	// For example "https://verifier.globalnames.org/api/v1/name_strings/"
	IdentifierSpace string `json:"identifierSpace"`

	// SchemaSpace provides the URL pointing to the schema of an entity.
	SchemaSpace string `json:"schemaSpace"`

	// DefaultTypes used for a reconciliation queries.
	DefaultTypes []Type `json:"defaultTypes"`

	// BatchSize sets maximum amount of queris in one batch.
	BatchSize int
}

// Choice provides flexibility to PropertySetting, allowing several different
// variants.
type Choice struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Entity describes an entity that can be used during reconciliation.
type Entity struct {
	// ID of an entity.
	ID string `json:"id"`

	// Name is a human-friendly title for an entity.
	Name string `json:"string"`

	// Description explains what an entity is.
	Description string `json:"description"`

	// Type or types of an entity.
	Type []Type `json:"type"`
}

// Type describes types available for the reconciliation service.
type Type struct {
	// ID is a unique identifier for the type.
	ID string `json:"id"`

	// Name is a human friendly short description of the type.
	Name string `json:"name"`
}

// Preview sets options to provide a widget with more details about a
// reconciliation candidate.
type Preview struct {
	// Height is the vertical size of a widget in pixels.
	Height int `json:"height"`

	// Width is a horisontal size of a widget in pixels.
	Width int `json:"width"`

	// URL provides a template in a form of `https://host/path/{{id}}`
	// where '{{id}}' will be substituted with an Entity ID
	URL string `json:"url"`
}

// View provides options for outlink where it shows details about a
// reconciliation candidate on a remote web page.
type View struct {
	// URL provides a template in a form of `http://host/path{{id}}`.
	// This URL is an outlink to an entity with the given ID.
	URL string `json:"url"`
}

// Extend provides information about optional additional information connected
// to the reconciliation candidate.
type Extend struct {
	// ProposeProperties contains metadata for Extend service.
	ProposeProperties `json:"propose_properties,omitempty"`

	// PropertySettings (optional) describes existing properties.
	PropertySettings []PropertySetting `json:"property_settings"`
}

// ProposeProperties provides metadata for Extend service.
type ProposeProperties struct {
	// ServiceURL contains Extend service URL without a path.
	ServiceURL string `json:"service_url,omitempty"`

	// ServicePath is the path part of the Extend service.
	ServicePath string `json:"service_path"`
}

// PropertySetting provides metadata of optional property settings
// for defining properties in Extend queries.
type PropertySetting struct {
	// Name of the property settings.
	Name string `json:"name"`

	// Label of the property settings.
	Label string `json:"label"`

	// Type of the property settings.
	Type string `json:"type"`

	// Default value of the property settings.
	Default string `json:"default"`

	// HelpText explains the property setting.
	HelpText string `json:"help_text,omitempty"`

	// Choices for the setting (optional)
	Choices []Choice `json:"choices,omitempty"`
}
