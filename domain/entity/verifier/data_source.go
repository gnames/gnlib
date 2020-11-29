package verifier

import "time"

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
	// It follows format of 'year-month-day' || 'year-month' || 'year'
	RevisionDate string `json:"releaseDate,omitempty"`

	// DOI of a DataSource;
	DOI string `json:"doi,omitempty"`

	// Citation representing a DataSource
	Citation string `json:"citation,omitempty"`

	// Authors associated with the Datasource
	Authors string `json:"authors,omitempty"`

	// Description of the DataSource.
	Description string `json:"description,omitempty"`

	// WebsiteURL is a hompage of a DataSource
	WebsiteURL string `json:"homeURL,omitempty"`

	// CurationLevel determines how much of manual or programmatic work is put
	// into assuring the quality of the data.
	CurationLevel `json:"curationLevel,omitempty"`

	// RecordCount tells how many entries are in a data source.
	RecordCount int `json:"recordCount,omitempty"`

	// UpdatedAt is the last import time and date.
	UpdatedAt time.Time `json:"updatedAt"`
}
