package erorr

type StampedError struct {
	callTrace string
	msg       string
}

func createStampedError(msg string) StampedError {
	callTrace := calltrace.String()

	return StampedError{
		callTrace:callTrace,
		msg:msg,
	}
}

func (receiver StampedError) CallTrace() string {
	return receiver.callTrace
}

func (receiver StampedError) Error() string {
	return receiver.msg
}

var _ error = StampedError{}
