package gnlib

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

// Pre-compiled regular expressions for tag replacement
var (
	titleRe   = regexp.MustCompile(`<title>(.*?)</title>`)
	warningRe = regexp.MustCompile(`<warn>(.*?)</warn>`)
	emRe      = regexp.MustCompile(`<em>(.*?)</em>`)
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

// MessageBase is a base type that implements the Error interface. It can be
// embedded in other error types to provide the basic functionality for
// formatted, colorized messages. It can also be used independently for
// non-error messages that need colorization.
type MessageBase struct {
	// Msg is the message string. It can be a format string for fmt.Sprintf.
	Msg string
	// Vars is a slice of arguments for the Msg format string.
	Vars []any
}

// UserMessage formats the error message with its variables and applies
// terminal colors based on tags.
//
// It replaces format verbs in the Msg string with values from the Vars slice.
// It also parses the following tags and replaces them with colored output:
//   - <title>...</title> for green text.
//   - <warning>...</warning> for red text.
//   - <em>...</em> for yellow text.
func (mb MessageBase) UserMessage() string {
	return FormatMessage(mb.Msg, mb.Vars)
}

// FormatMessage is a static function that takes a message string with optional
// format variables and returns a colorized string based on tags.
//
// It replaces format verbs in the message with values from the vars slice.
// It also parses the following tags and replaces them with colored output:
//   - <title>...</title> for green text.
//   - <warning>...</warning> for red text.
//   - <em>...</em> for yellow text.
//
// Example:
//
//	msg := FormatMessage("Processing <title>%s</title>", []any{"file.txt"})
func FormatMessage(msg string, vars []any) string {
	if len(vars) > 0 {
		msg = fmt.Sprintf(msg, vars...)
	}

	msg = titleRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.GreenString(titleRe.FindStringSubmatch(match)[1])
	})
	msg = warningRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.RedString(warningRe.FindStringSubmatch(match)[1])
	})
	msg = emRe.ReplaceAllStringFunc(msg, func(match string) string {
		return color.YellowString(emRe.FindStringSubmatch(match)[1])
	})

	return msg
}
