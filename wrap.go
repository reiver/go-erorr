package erorr

import (
	"codeberg.org/reiver/go-field"
)

func Wrap(err error, msg string, fields ...field.Field[string]) error {
	switch casted := err.(type) {
	case Errors:
		return WrapErrors(casted, msg, fields...)
	default:
	        return WrapError(err, msg, fields...)
	}
}
