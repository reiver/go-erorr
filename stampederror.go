package erorr

import (
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-calltrace"
)

type StampedError struct {
	callTrace string
	fields    []field.Field[string]
	msg       string
}

func StampError(msg string, fields ...field.Field[string]) StampedError {
	callTrace := calltrace.String()

	return StampedError{
		callTrace:callTrace,
		fields:fields,
		msg:msg,
	}
}

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
