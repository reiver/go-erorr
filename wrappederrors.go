package erorr

import (
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-calltrace"
)

type WrappedErrors struct {
	callTrace string
	errs      []error
	fields    []field.Field[string]
	msg       string
}

func WrapErrors(errs []error, msg string, fields ...field.Field[string]) WrappedErrors {
	callTrace := calltrace.String()

	return WrappedErrors{
		callTrace:callTrace,
		errs:errs,
		fields:fields,
		msg:msg,
	}
}

func (receiver WrappedErrors) CallTrace() string {
	return receiver.callTrace
}

func (receiver WrappedErrors) Error() string {
	return receiver.msg
}

func (receiver WrappedErrors) Unwrap() []error {
	return receiver.errs
}

var _ error = WrappedErrors{}
