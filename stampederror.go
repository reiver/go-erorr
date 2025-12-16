package erorr

type StampedError struct {
	callTrace string
	msg       string
}

func (receiver StampedError) CallTrace() string {
	return receiver.callTrace
}

func (receiver StampedError) Error() string {
	return receiver.msg
}

var _ error = StampedError{}
