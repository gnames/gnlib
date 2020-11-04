package format

// Outputter interface us a uniform way to create an output of a datum
type Outputter interface {
	// FormattedOutput takes a record and returns a string representation of
	// the record accourding to supplied format.
	Output(record interface{}, f Format) string
}
