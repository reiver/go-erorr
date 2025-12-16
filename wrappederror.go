package erorr

type WrappedError struct {
	callTrace string
	err       error
	msg       string
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
