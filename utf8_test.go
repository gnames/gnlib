package gnlib_test

import (
	"testing"

	"github.com/gnames/gnlib"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/unicode/norm"
)

func TestFixUtf8(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid ascii",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "valid utf8",
			input:    "こんにちは世界", // Hello world in Japanese
			expected: "こんにちは世界",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "invalid utf8 sequence",
			input:    "hello\xffworld", // \xff is an invalid byte
			expected: "hello\uFFFDworld",
		},
		{
			name:     "multiple invalid utf8 sequences",
			input:    "he\x80llo\xff\xfeworld\xfa",
			expected: "he\uFFFDllo\uFFFD\uFFFDworld\uFFFD",
		},
		{
			name:     "mixed valid and invalid",
			input:    "valid\xe2\x28\xa1invalid",  // \xe2\x28\xa1 is an incomplete sequence start
			expected: "valid\uFFFD(\uFFFDinvalid", // The ( is what \x28 becomes after \xe2 is replaced
		},
		{
			name:     "string needing nfc normalization",
			input:    "e\u0301lite", // e + combining acute accent
			expected: "\u00E9lite",  // é (precomposed)
		},
		{
			name:     "invalid sequence then nfc",
			input:    "invalid\xffthen_e\u0301lite",
			expected: "invalid\uFFFDthen_\u00E9lite",
		},
		{
			name:     "long string with invalid char",
			input:    "This is a long string with an inv\xc3alid character.",
			expected: "This is a long string with an inv\uFFFDalid character.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gnlib.FixUtf8(tt.input)
			assert.Equal(norm.NFC.String(tt.expected), got, "Input: %q", tt.input)
		})
	}
}
