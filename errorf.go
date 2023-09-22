package erorr

import (
	"fmt"
)

func Errorf(format string, a ...interface{}) error {

	var s string = fmt.Sprintf(format, a...)

	return Error(s)
}
