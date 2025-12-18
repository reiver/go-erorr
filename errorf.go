package erorr

import (
	"fmt"
	"strings"

	"github.com/reiver/go-calltrace"
)

// Errorf returns a formatted error.
//
// If an one or more errors are passed in the parameters of Errorf, then they are wrapped.
// If there is one error, then a [WrappedError] is returned.
// If there is more-than one error, then a [WrappedErrors] is returned.
// If there are no errors, then a [StampedError] is returned.
func Errorf(format string, a ...any) error {

	callTrace := calltrace.String()

	format = strings.ReplaceAll(format, "%w", "%s")

	msg := fmt.Sprintf(format, a...)

	errs := extractErrors(a...)

	switch len(errs) {
	case 0:
		return StampedError{
			callTrace:callTrace,
			msg:msg,
		}
	case 1:
		err := errs[0]

		return createWrappedError(msg, err)
	default:
		return createWrappedErrors(msg, errs...)
	}
}

func extractErrors(a ...any) []error {
	var errs []error

	for _, datum := range a {
		err, casted := datum.(error)
		if !casted {
			continue
		}

		errs = append(errs, err)
	}

	return errs
}
