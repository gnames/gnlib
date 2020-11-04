package verifier

type Verifier interface {
	// Verify takes names-strings and options and returns verification result.
	Verify(VerifyParams) []*Verification

	// DataSources takes data-source id and opts and returns the data-source
	// metadata.  If no id is provided, it returns metadata for all data-sources.
	DataSources(DataSourcesOpts) []*DataSource
}
