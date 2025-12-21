package erorr

import (
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-calltrace"
)

// StampedError is contextual error.
//
// When a StampedError is created includes a call-trace, and a list of fields.
//
// For example:
//
//	err := erorr.Stamp("client is nil",
//		field.String("url", url),
//	)
type StampedError struct {
	callTrace string
	fields    []field.Field[string]
	msg       string
}

// StampError returns a [StampedError].
//
// If you are not sure whetherto use [StampError] or [Stamp] use [Stamp].
func StampError(msg string, fields ...field.Field[string]) StampedError {
	callTrace := calltrace.String()

	return StampedError{
		callTrace:callTrace,
		fields:fields,
		msg:msg,
	}
}

// Stamp returns a [StampedError].
//
// If you are not sure whetherto use [StampError] or [Stamp] use [Stamp].
func Stamp(msg string, fields ...field.Field[string]) error {
	return StampError(msg, fields...)
}

func (receiver StampedError) CallTrace() string {
	return receiver.callTrace
}

func (receiver StampedError) Error() string {
	return receiver.msg
}

func (receiver StampedError) Fields() []field.Field[string] {
	return receiver.fields
}

var _ error = StampedError{}
