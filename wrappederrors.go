package erorr

type WrappedErrors struct {
	errs []error
	msg  string
}

func (receiver WrappedErrors) Error() string {
	return receiver.msg
}

func (receiver WrappedErrors) Unwrap() []error {
	return receiver.errs
}

var _ error = WrappedErrors{}
