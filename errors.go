package erorr

type Errors []error

func (receiver Errors) Error() string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	for _, err := range []error(receiver) {
		if nil == err {
			continue
		}

		p = append(p, err.Error()...)
		p = append(p, '\n')
	}

	return string(p)
}

// Before Go added support for:
//
//	Unwrap() []error
//
// This method was called:
//
//	Errors() []error
//
// But was renamed.
func (receiver Errors) Unwrap() []error {
	if len([]error(receiver)) <= 0 {
		return nil
	}

	return append([]error(nil), []error(receiver)...)
}
