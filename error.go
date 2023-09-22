package erorr

type Error string

func (receiver Error) Error() string {
	return string(receiver)
}

var _ error = Error("")
