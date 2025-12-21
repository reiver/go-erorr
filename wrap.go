package erorr

import (
	"codeberg.org/reiver/go-field"
)

// Wrap returns a wrapped annotated error, which includes a call-trace and fields.
//
// If `err` is nil then Wrap returns nil.
//
// Example usage:
//
//	err := something()
//	if nil != err {
//		return erorr.Wrap(err, "failed to do something",
//			field.String("service", "banana"),
//		)
//	}
func Wrap(err error, msg string, fields ...field.Field[string]) error {
	if nil == err {
		return nil
	}

	switch casted := err.(type) {
	case Errors:
		return WrapErrors(casted, msg, fields...)
	default:
	        return WrapError(err, msg, fields...)
	}
}
