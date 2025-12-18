package erorr

import (
	"github.com/reiver/go-calltrace"
)

type WrappedError struct {
	callTrace string
	err       error
	msg       string
}

func createWrappedError(msg string, err error) WrappedError {
	callTrace := calltrace.String()

	return WrappedError{
		callTrace:callTrace,
		err:err,
		msg:msg,
	}
}

func (receiver WrappedError) CallTrace() string {
	return receiver.callTrace
}

func (receiver WrappedError) Error() string {
	return receiver.msg
}

func (receiver WrappedError) Unwrap() error {
	return receiver.err
}

var _ error = WrappedError{}
