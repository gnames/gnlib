// package tribool implements a simple nullable three-valued construct.
// It can be safely serialized/deserialized by JSON libraries.
package tribool

import (
	"bytes"
	"errors"
)

var (
	null  = []byte{'n', 'u', 'l', 'l'}
	yes   = []byte{'"', 'y', 'e', 's', '"'}
	no    = []byte{'"', 'n', 'o', '"'}
	maybe = []byte{'"', 'm', 'a', 'y', 'b', 'e', '"'}
)

// Tribool is a nullable three-value construct.
type Tribool struct {
	Value int
	Valid bool
}

// NewTribool creates a Tribool out of an integer. Positive integer
// creates 'yes' value, 0 creates 'maybe' values, negative integer generates
// 'no value. If no number is given, the result represents null, if more than
// one number is given, other numbers are ignored.
func NewTribool(ints ...int) Tribool {
	if len(ints) == 0 {
		return Tribool{}
	}

	val := ints[0]
	if val < 0 {
		val = -1
	} else if val > 0 {
		val = 1
	}
	return Tribool{Value: val, Valid: true}
}

// Implements fmt.String interface
func (t Tribool) String() string {
	if !t.Valid {
		return ""
	}
	if t.Value < 0 {
		return "no"
	}
	if t.Value > 0 {
		return "yes"
	}
	return "maybe"
}

// Degrades data to a boolean
func (t Tribool) Bool() bool {
	if t.Value > 0 && t.Valid {
		return true
	}
	return false
}

// Returns integer representation of Tribool. For "yes" it returns 1,
// for "no" and "null" it returns -1, and for "maybe it returns 0.
func (t Tribool) Int() int {
	if !t.Valid {
		return -1
	}
	return t.Value
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int is null.
func (t Tribool) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return null, nil
	}
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller.
func (t *Tribool) UnmarshalJSON(bs []byte) error {
	t.Valid = true

	if bytes.Equal(bs, null) {
		t.Valid = false
	} else if bytes.Equal(bs, yes) {
		t.Value = 1
	} else if bytes.Equal(bs, no) {
		t.Value = -1
	} else if bytes.Equal(bs, maybe) {
		t.Value = 0
	} else {
		return errors.New("cannot decode to tribool")
	}

	return nil
}
