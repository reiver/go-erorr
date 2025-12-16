package erorr_test

import (
	"fmt"

	"codeberg.org/reiver/go-erorr"
)

func ExampleErrorf_oneError() {

	const ErrFailure = erorr.Error("failure!")

	// ...

	service := "AuthService"

	var err error = ErrFailure
	err = erorr.Errorf("Uh oh, something bad happened to %s: %w", service, err)

	fmt.Printf("error-message: %s\n", err)
	fmt.Printf("errors-type: %T\n", err)

	// Output:
	// error-message: Uh oh, something bad happened to AuthService: failure!
	// errors-type: erorr.WrappedError
}

func ExampleErrorf_twoErrors() {

	const ErrFailure = erorr.Error("failure!")
	const ErrProblem = erorr.Error("problem!")

	// ...

	service := "AuthService"

	err := erorr.Errorf("Uh oh, some bad things happened to %s: %w and %w", service, ErrFailure, ErrProblem)

	fmt.Printf("error-message: %s\n", err)
	fmt.Printf("errors-type: %T\n", err)

	// Output:
	// error-message: Uh oh, some bad things happened to AuthService: failure! and problem!
	// errors-type: erorr.WrappedErrors
}
