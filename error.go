package gnlib

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

func PrintUserMessage(err error) {
	var target Error
	printable := errors.As(err, &target)
	if printable {
		fmt.Println(target.UserMessage())
	}
}

// Error is an interface for errors that can produce a formatted, user-friendly
// message. It is intended for errors that need to be displayed to the end-user
// on STDOUT.
type Error interface {
	error
	// UserMessage returns a formatted and colorized string that is safe to
	// display to an end-user.
	UserMessage() string
}

// ErrorBase is a base type that implements the Error interface. It can be
// embedded in other error types to provide the basic functionality.
type ErrorBase struct {
	// Msg is the message string. It can be a format string for fmt.Sprintf.
	Msg string
	// Vars is a slice of arguments for the Msg format string.
	Vars []any
}

// NewError is a constructor for ErrorBase. It takes a message string (which
// can be a format string) and a slice of arguments for the format string.
func NewError(msg string, vars []any) ErrorBase {
	res := ErrorBase{
		Msg:  msg,
		Vars: vars,
	}
	return res
}

// UserMessage formats the error message with its variables and applies
// terminal colors based on tags.
//
// It replaces format verbs in the Msg string with values from the Vars slice.
// It also parses the following tags and replaces them with colored output:
//   - <title>...</title> for green text.
//   - <warning>...</warning> for red text.
//   - <em>...</em> for yellow text.
func (gnerr ErrorBase) UserMessage() string {
	msg := gnerr.Msg
	if len(gnerr.Vars) > 0 {
		msg = fmt.Sprintf(msg, gnerr.Vars...)
	}

	res := gnerr.colorize(msg)
	return res
}

// colorize parses a string for color tags and replaces them with
// ANSI-colored text.
func (gnerr ErrorBase) colorize(msg string) string {
	tags := map[string]func(format string, a ...any) string{
		"title":   color.GreenString,
		"warning": color.RedString,
		"em":      color.YellowString,
	}

	// Process each tag type separately since Go regex doesn't support backreferences
	for tagName, colorFunc := range tags {
		// Create a regex for this specific tag
		pattern := fmt.Sprintf(`<%s>(.*?)</%s>`, tagName, tagName)
		re := regexp.MustCompile(pattern)

		msg = re.ReplaceAllStringFunc(msg, func(match string) string {
			submatches := re.FindStringSubmatch(match)
			if len(submatches) < 2 {
				return match
			}
			content := submatches[1]
			return colorFunc(content)
		})
	}

	return msg
}
