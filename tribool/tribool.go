// package tribool implements a simple nullable tri-way value
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

type Tribool struct {
	Value int
	Valid bool
}

func NewTribool(ints ...int) Tribool {
	if len(ints) == 0 {
		return Tribool{}
	}
	return Tribool{Value: ints[0], Valid: true}
}

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

func (t Tribool) Bool() bool {
	if t.Value > 0 && t.Valid {
		return true
	}
	return false
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
