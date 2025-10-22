package gnlib_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/gnames/gnlib"
	"github.com/stretchr/testify/assert"
)

// TestErrorBasic verifies the fundamental behavior of creating a custom error
// type that embeds gnlib.MessageBase. It ensures that the custom error satisfies
// the gnlib.Error interface, can be wrapped by standard errors, and can be
// correctly inspected using errors.As.
func TestErrorBasic(t *testing.T) {
	type errTest struct {
		error
		gnlib.MessageBase
	}
	err := errTest{
		error:       errors.New("test error"),
		MessageBase: gnlib.MessageBase{Msg: "hello"},
	}
	assert.Equal(t, "test error", err.Error())
	assert.Equal(t, "hello", err.UserMessage())
	err2 := fmt.Errorf("higher error: %w", err)
	var target gnlib.Error
	assert.True(t, errors.As(err2, &target))
	assert.Equal(t, "hello", target.UserMessage())
}

// TestError_Features is a table-driven test that covers the variable
// substitution and colorization features of the UserMessage method.
func TestError_Features(t *testing.T) {
	// Force color output for testing purposes
	origNoColor := color.NoColor
	color.NoColor = false
	defer func() { color.NoColor = origNoColor }()
	// tests is a collection of test cases for UserMessage.
	tests := []struct {
		// name is a descriptive name for the test case.
		name string
		// msg is the format string for the error.
		msg string
		// vars are the arguments for the format string.
		vars []any
		// expected is the final string expected from UserMessage.
		expected string
	}{
		{
			name:     "no vars, no tags",
			msg:      "Simple message.",
			vars:     nil,
			expected: "Simple message.",
		},
		{
			name:     "string var",
			msg:      "Hello, %s!",
			vars:     []any{"world"},
			expected: "Hello, world!",
		},
		{
			name:     "int var",
			msg:      "Count: %d",
			vars:     []any{42},
			expected: "Count: 42",
		},
		{
			name:     "multiple vars",
			msg:      "%s is %d years old.",
			vars:     []any{"John", 30},
			expected: "John is 30 years old.",
		},
		{
			name:     "title tag",
			msg:      "<title>This is a title</title>",
			vars:     nil,
			expected: "\x1b[32mThis is a title\x1b[0m",
		},
		{
			name:     "warning tag",
			msg:      "<warning>This is a warning</warning>",
			vars:     nil,
			expected: "\x1b[31mThis is a warning\x1b[0m",
		},
		{
			name:     "em tag",
			msg:      "This is <em>important</em>.",
			vars:     nil,
			expected: "This is \x1b[33mimportant\x1b[0m.",
		},
		{
			name:     "multiple tags",
			msg:      "<title>Title</title> and <warning>Warning</warning>",
			vars:     nil,
			expected: "\x1b[32mTitle\x1b[0m and \x1b[31mWarning\x1b[0m",
		},
		{
			name:     "tags and vars",
			msg:      "<title>Processing %s</title>",
			vars:     []any{"file.txt"},
			expected: "\x1b[32mProcessing file.txt\x1b[0m",
		},
		{
			name:     "unknown tags",
			msg:      "<title>Title</title> and <unknown>something</unknown>",
			vars:     nil,
			expected: "\x1b[32mTitle\x1b[0m and <unknown>something</unknown>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := gnlib.MessageBase{Msg: tt.msg, Vars: tt.vars}
			assert.Equal(t, tt.expected, base.UserMessage(), tt.name)
		})
	}
}

// TestFormatMessage verifies the static FormatMessage function that can be
// used independently of the MessageBase type.
func TestFormatMessage(t *testing.T) {
	// Force color output for testing purposes
	origNoColor := color.NoColor
	color.NoColor = false
	defer func() { color.NoColor = origNoColor }()

	tests := []struct {
		name     string
		msg      string
		vars     []any
		expected string
	}{
		{
			name:     "simple message",
			msg:      "Hello, world!",
			vars:     nil,
			expected: "Hello, world!",
		},
		{
			name:     "with variables",
			msg:      "Processing %d files",
			vars:     []any{42},
			expected: "Processing 42 files",
		},
		{
			name:     "with tags",
			msg:      "<title>Success!</title> All <em>%d</em> tests passed.",
			vars:     []any{10},
			expected: "\x1b[32mSuccess!\x1b[0m All \x1b[33m10\x1b[0m tests passed.",
		},
		{
			name:     "complex message",
			msg:      "<warning>Error:</warning> Failed to process <em>%s</em> in <title>%s</title>",
			vars:     []any{"data.csv", "/home/user"},
			expected: "\x1b[31mError:\x1b[0m Failed to process \x1b[33mdata.csv\x1b[0m in \x1b[32m/home/user\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gnlib.FormatMessage(tt.msg, tt.vars)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestPrintUserMessage verifies that the PrintUserMessage function correctly
// prints the user message to stdout for errors that implement the gnlib.Error
// interface, and prints nothing for regular errors.
func TestPrintUserMessage(t *testing.T) {
	// Force color output for testing purposes
	origNoColor := color.NoColor
	color.NoColor = false
	defer func() { color.NoColor = origNoColor }()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name: "gnlib.Error prints user message",
			err: func() error {
				type errTest struct {
					error
					gnlib.MessageBase
				}
				return errTest{
					error:       errors.New("internal error"),
					MessageBase: gnlib.MessageBase{Msg: "Test <title>message</title>"},
				}
			}(),
			expected: "Test \x1b[32mmessage\x1b[0m\n",
		},
		{
			name: "wrapped gnlib.Error prints user message",
			err: func() error {
				type errTest struct {
					error
					gnlib.MessageBase
				}
				err := errTest{
					error:       errors.New("internal error"),
					MessageBase: gnlib.MessageBase{Msg: "Wrapped <warning>error</warning>"},
				}
				return fmt.Errorf("outer: %w", err)
			}(),
			expected: "Wrapped \x1b[31merror\x1b[0m\n",
		},
		{
			name:     "regular error prints nothing",
			err:      errors.New("regular error"),
			expected: "",
		},
		{
			name:     "nil error prints nothing",
			err:      nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			gnlib.PrintUserMessage(tt.err)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)

			assert.Equal(t, tt.expected, buf.String())
		})
	}
}
