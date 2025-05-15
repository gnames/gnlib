package gnlib

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// FixUtf8 cleans a string by replacing invalid UTF-8 sequences with U+FFFD
// and normalizing to NFC.
func FixUtf8(s string) string {
	// Estimate capacity: assume most bytes are valid runes (1 rune per 1-4 bytes).
	result := make([]rune, 0, len(s)/2+1)

	// Iterate over the string byte by byte, tracking position.
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size <= 1 {
			// Invalid sequence: append U+FFFD and advance by 1 byte.
			result = append(result, utf8.RuneError)
			i++
		} else {
			// Valid rune: append and advance by rune size.
			result = append(result, r)
			i += size
		}
	}

	return norm.NFC.String(string(result))
}
