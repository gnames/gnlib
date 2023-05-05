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
	Properties []Property `json:"properties,omitempty"`

	// TypeStrict is a legacy deprecated field that came from FreeBase.
	TypeStrict string `json:"type_strict,omitempty"`
}

// Property can be used to further filter list of candidates, similar to
// a WHERE cause in SQL.
// This implementation is less flexible than W3C standard and takes only
// one value. We will expand it if necessary.
type Property struct {
	// PID is the property name.
	PID string `json:"pid"`

	// Value is used to filter the property.
	Value string `json:"v"`
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

	// Score is used to estimate how well the result matches the
	// original query.
	Score float64 `json:"score"`

	// Features might contain details of reconciliation and be used for
	// the score determination.
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

	// Value is the a quantitative representation of the feature.
	Value float64 `json:"value"`
}
