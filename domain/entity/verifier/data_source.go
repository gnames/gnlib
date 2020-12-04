package verifier

// DataSource provides metadata for an externally collected data-set.
type DataSource struct {
	// ID is a DataSource Id.
	ID int `json:"id"`

	// UUID generated by GlobalNames and associated with the DataSource
	UUID string `json:"uuid,omitempty"`

	// Title is a full title of a DataSource
	Title string `json:"title"`

	// TitleShort is a shortened/abbreviated title of a DataSource.
	TitleShort string `json:"titleShort"`

	// Version of the data-set for a DataSource.
	Version string `json:"version,omitempty"`

	// RevisionDate of a data-set from a data-provider.
	// It follows format of 'YYYY-MM-DD' || 'YYYY-MM' || 'YYYY'
	// This data comes from the information given by the data-provider,
	// while UpdatedAt field is the date of harvesting of the
	// resource.
	RevisionDate string `json:"releaseDate,omitempty"`

	// DOI of a DataSource;
	DOI string `json:"doi,omitempty"`

	// Citation representing a DataSource
	Citation string `json:"citation,omitempty"`

	// Authors associated with the DataSource
	Authors string `json:"authors,omitempty"`

	// Description of the DataSource.
	Description string `json:"description,omitempty"`

	// WebsiteURL is a hompage of a DataSource
	WebsiteURL string `json:"homeURL,omitempty"`

	// IsOutlinkReady is true for data-sources that have enough data and
	// metadata to be recommended for outlinking by third-party applications
	// (be included into preferred data-sources). When false, it does not
	// mean that the original resource is not valuable, it means that
	// its representation at gnames is not complete/resent enough.
	IsOutlinkReady bool `json:"isOutlinkReady"`

	// Curation determines how much of manual or programmatic work is put
	// into assuring the quality of the data.
	Curation CurationLevel `json:"curation"`

	// RecordCount tells how many entries are in a DataSource.
	RecordCount int `json:"recordCount"`

	// UpdatedAt is the last import date (YYYY-MM-DD). In contrast,
	// RevisionDate field indicates when the resource was
	// updated according to its data-provider.
	UpdatedAt string `json:"updatedAt"`
}
