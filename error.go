package erorr

// Error is a simple error type.
//
// Error is implemented in a way such that you can create `const` errors.
//
// For example:
//
//	const ErrNilReceiver = erorr.Error("nil receiver")
//
// Use Error for your core errors.
type Error string

func (receiver Error) Error() string {
	return string(receiver)
}

var _ error = Error("")
