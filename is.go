package erorr

import (
	"errors"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}
