package erorr

import (
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-calltrace"
)

type WrappedError struct {
	callTrace string
	err       error
	fields    []field.Field[string]
	msg       string
}

func WrapError(err error, msg string, fields ...field.Field[string]) WrappedError {
	callTrace := calltrace.String()

	return WrappedError{
		callTrace:callTrace,
		err:err,
		fields:fields,
		msg:msg,
	}
}

func (receiver WrappedError) CallTrace() string {
	return receiver.callTrace
}

func (receiver WrappedError) Error() string {
	return receiver.msg
}

func (receiver WrappedError) Fields() []field.Field[string] {
	return receiver.fields
}

func (receiver WrappedError) Unwrap() error {
	return receiver.err
}

var _ error = WrappedError{}
