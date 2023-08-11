// package reconciler describes entities for implementation of a
// Reconciliation Service API v0.2
// https://www.w3.org/community/reports/reconciliation/CG-FINAL-specs-0.2-20230410/
package reconciler

// Input contains fields necessary for the reconciliation process.
// This intput is compatible with W3C Reconciliation API.
type Input struct {
	// Queries contains all requested queries. Queries are identified
	// by an identification string.
	Queries map[string]Query `form:"queries"`
}

// Query is a set of fields used for verification.
type Query struct {
	// Query contains a name-string we try to reconcile.
	Query string `json:"query"`

	// Type allows to constrain reconciliation agains specific type described
	// in the manifest.
	Type string `json:"type"`

	// Limit restricts the number of candidates returned by the query.
	Limit int `json:"limit"`

	// Properties allow to add additional filters to the reconciliation
	// process.
	Properties []PropertyInfo `json:"properties,omitempty"`

	// TypeStrict is a legacy deprecated field that came from FreeBase.
	TypeStrict string `json:"type_strict,omitempty"`
}

// ExtendQuery provides input for getting additional properties associated
// with the name-string ID.
type ExtendQuery struct {
	//IDs contain an entity IDs.
	IDs []string `json:"ids"`

	// Properties contains a slice of properties.
	Properties []Property `json:"properties"`
}

// Property implements the APIs `property`, an attribute of an entity.
type Property struct {
	// ID is a property identifier.
	ID string `json:"id"`

	// Name is human-friendly title of a property.
	Name string `json:"name"`
}

// PropertyInfo combines a property ID with the value of the property.
// PropertyInfo can be used to further filter list of candidates, similar to
// a WHERE cause in SQL.
// This implementation is less flexible than W3C standard and takes only
// one value. We will expand it if necessary.
type PropertyInfo struct {
	// PropertyID is the same as Property.ID
	PropertyID string `json:"pid"`

	// PropertyValue is the value of PropertyValue.Str for this property.
	PropertyValue string `json:"v"`
}

// Ouput is a map where the key is the provided identifier of a query,
// and the ReconciliationResult contains all found ReconciliationCandidates.
type Output map[string]ReconciliationResult

// ReconciliationResult is a slice where results are sorted by their score.
type ReconciliationResult struct {
	// Result contains all candidates for reconciliation.
	Result []ReconciliationCandidate `json:"result"`
}

// ReconciliationCandidate contains the details of a reconciliation item.
type ReconciliationCandidate struct {
	// ID can be used to lookup the entity in a corresponding service.
	ID string `json:"id"`

	// Name contains reconciled name-string.
	Name string `json:"name"`

	// Description provides some metadata about the item.
	Description string `json:"description"`

	// Score is used to estimate chances for a result to be a match.
	// It is calculated from features.
	Score float64 `json:"score"`

	// Features might contain details of reconciliation and be used for
	// the score calculation.
	Features []Feature `json:"features,omitempty"`

	// Types contains types that were assigned to the candidate.
	Types []Type `json:"types"`

	// Match is true if the score is above a threshold and without a
	// reasonable doubt the result is the best match to the query.
	Match bool `json:"match"`
}

// Feature is a matching feature that can be used to determing the matching
// score.
type Feature struct {
	// ID is the name of the feature.
	ID string `json:"id"`

	// Value is a quantitative representation of the feature.
	Value float64 `json:"value"`
}

// PropertyOutput provides information about properties known for a particular
// `entity` type.
type PropertyOutput struct {
	// Type is the type ID for a given entity type.
	Type string `json:"type"`

	// Properties describes corresponding properties of the type.
	Properties []Property `json:"properties"`
}

// ExtendOutput provides data returned by Extend query.
type ExtendOutput struct {
	// Meta describes properties information.
	Meta []Property `json:"meta"`

	// Rows is a map, where the key is an entity ID, and the value is another
	// map where key is the property ID, and the value is a slice of property
	// values (for simplification values are always JSON-encoded strings).
	Rows map[string]map[string][]PropertyValue `json:"rows"`
}

// PropertyValue is simplified compare to API, for now it supports only
// a string value.
type PropertyValue struct {
	Str string `json:"str"`
}
