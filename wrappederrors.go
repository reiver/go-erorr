package erorr

type WrappedErrors struct {
	callTrace string
	errs      []error
	msg       string
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
