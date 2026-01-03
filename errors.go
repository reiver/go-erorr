package erorr

type Errors []error

func (receiver Errors) Error() string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	var appended bool
	for _, err := range []error(receiver) {
		if nil == err {
			continue
		}

		if appended {
			p = append(p, ", "...)
		}
		p = append(p, "⸨"...)
		p = append(p, err.Error()...)
		p = append(p, "⸩"...)
		appended = true
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
