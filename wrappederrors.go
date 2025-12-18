package erorr

import (
	"github.com/reiver/go-calltrace"
)

type WrappedErrors struct {
	callTrace string
	errs      []error
	msg       string
}

func createWrappedErrors(msg string, errs ...error) WrappedErrors {
	callTrace := calltrace.String()

	return WrappedErrors{
		callTrace:callTrace,
		errs:errs,
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
