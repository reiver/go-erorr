package erorr

type WrappedError struct {
	err error
	msg string
}

func (receiver WrappedError) Error() string {
	return receiver.msg
}

func (receiver WrappedError) Unwrap() error {
	return receiver.err
}

var _ error = WrappedError{}
