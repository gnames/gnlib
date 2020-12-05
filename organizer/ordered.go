package organizer

// Ordered is an object that can provide its index number
// (its position in an order) and unpacking function that allows to
// unload its content.
type Ordered interface {
	// Index returns the position of an object in an order.
	// It is assumed that indices in the original collection do
	// start from 0 and then follow without gaps and duplications.
	Index() int

	// Unpack takes an object of the same type as the type of the content of the
	// ordered object and fills it with the content of the Ordered object.
	Unpack(content interface{}) error
}
