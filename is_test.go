package erorr

import (
	"testing"

	"io"
)

func TestIs(t *testing.T) {
	const CoreError = Error("core error")


	tests := []struct{
		Error    error
		Target   error
		Expected bool
	}{
		{
			Error:  nil,
			Target: nil,
			Expected: true,
		},
		{
			Error:  error(nil),
			Target: nil,
			Expected: true,
		},
		{
			Error:  CoreError,
			Target: nil,
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: nil,
			Expected: false,
		},



		{
			Error:  nil,
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  error(nil),
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  CoreError,
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  nil,
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  io.EOF,
			Target: CoreError,
			Expected: false,
		},



		{
			Error:  nil,
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: io.EOF,
			Expected: true,
		},


		//==================================================
		//==================================================
		//==================================================


		{
			Error:  nil,
			Target: WrappedError{err:nil},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:nil},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:nil},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:nil},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:error(nil)},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:error(nil)},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:error(nil)},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:error(nil)},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:CoreError},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:CoreError},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:CoreError},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:CoreError},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:io.EOF},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:io.EOF},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:io.EOF},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:io.EOF},
			Expected: false,
		},



		//==================================================



		{
			Error:  WrappedError{err:nil},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:error(nil)},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:CoreError},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:io.EOF},
			Target: nil,
			Expected: false,
		},



		{
			Error:  WrappedError{err:nil},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:error(nil)},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:CoreError},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:io.EOF},
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  WrappedError{err:nil},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:error(nil)},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:CoreError},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:io.EOF},
			Target: CoreError,
			Expected: false,
		},



		{
			Error:  WrappedError{err:nil},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:error(nil)},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:CoreError},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:io.EOF},
			Target: io.EOF,
			Expected: true,
		},



		//==================================================



		{
			Error:  nil,
			Target: WrappedError{err:WrappedError{err:nil}},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:WrappedError{err:nil}},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:WrappedError{err:nil}},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:WrappedError{err:nil}},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:WrappedError{err:error(nil)}},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:WrappedError{err:error(nil)}},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:WrappedError{err:error(nil)}},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:WrappedError{err:error(nil)}},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:WrappedError{err:CoreError}},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:WrappedError{err:CoreError}},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:WrappedError{err:CoreError}},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:WrappedError{err:CoreError}},
			Expected: false,
		},



		{
			Error:  nil,
			Target: WrappedError{err:WrappedError{err:io.EOF}},
			Expected: false,
		},
		{
			Error:  error(nil),
			Target: WrappedError{err:WrappedError{err:io.EOF}},
			Expected: false,
		},
		{
			Error:  CoreError,
			Target: WrappedError{err:WrappedError{err:io.EOF}},
			Expected: false,
		},
		{
			Error:  io.EOF,
			Target: WrappedError{err:WrappedError{err:io.EOF}},
			Expected: false,
		},



		//==================================================



		{
			Error:  WrappedError{err:WrappedError{err:nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:CoreError}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:io.EOF}},
			Target: nil,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:CoreError}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:io.EOF}},
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:nil}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:error(nil)}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:io.EOF}},
			Target: CoreError,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:nil}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:error(nil)}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:CoreError}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:io.EOF}},
			Target: io.EOF,
			Expected: true,
		},



		//==================================================



		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:CoreError}}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:io.EOF}}},
			Target: nil,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:CoreError}}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:io.EOF}}},
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:io.EOF}}},
			Target: CoreError,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:nil}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:error(nil)}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:CoreError}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedError{err:WrappedError{err:io.EOF}}},
			Target: io.EOF,
			Expected: true,
		},



		//==================================================
		//==================================================
		//==================================================



		{
			Error:  WrappedErrors{errs:nil},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error(nil)},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF}},
			Target: nil,
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{nil, nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, CoreError}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, io.EOF}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{error(nil), nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), CoreError}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), io.EOF}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{CoreError, nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, CoreError}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, io.EOF}},
			Target: nil,
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, nil}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, error(nil)}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, CoreError}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, io.EOF}},
			Target: nil,
			Expected: false,
		},



		//==================================================



		{
			Error:  WrappedErrors{errs:nil},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error(nil)},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF}},
			Target: error(nil),
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{nil, nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, CoreError}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, io.EOF}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{error(nil), nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), CoreError}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), io.EOF}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{CoreError, nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, CoreError}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, io.EOF}},
			Target: error(nil),
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, nil}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, error(nil)}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, CoreError}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, io.EOF}},
			Target: error(nil),
			Expected: false,
		},



		//==================================================



		{
			Error:  WrappedErrors{errs:nil},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error(nil)},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil)}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{nil, nil}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, error(nil)}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, io.EOF}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{error(nil), nil}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), error(nil)}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), io.EOF}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{CoreError, nil}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, error(nil)}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, io.EOF}},
			Target: CoreError,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, nil}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, error(nil)}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, CoreError}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, io.EOF}},
			Target: CoreError,
			Expected: false,
		},




		//==================================================



		{
			Error:  WrappedErrors{errs:nil},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error(nil)},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil)}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{nil, nil}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, error(nil)}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, CoreError}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{nil, io.EOF}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{error(nil), nil}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), error(nil)}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), CoreError}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{error(nil), io.EOF}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{CoreError, nil}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, error(nil)}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, CoreError}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedErrors{errs:[]error{CoreError, io.EOF}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, nil}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, error(nil)}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, CoreError}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedErrors{errs:[]error{io.EOF, io.EOF}},
			Target: io.EOF,
			Expected: true,
		},



		//==================================================
		//==================================================
		//==================================================



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, CoreError}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, io.EOF}}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), CoreError}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), io.EOF}}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, CoreError}}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, io.EOF}}},
			Target: nil,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, CoreError}}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, io.EOF}}},
			Target: nil,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, CoreError}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, io.EOF}}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), CoreError}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), io.EOF}}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, CoreError}}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, io.EOF}}},
			Target: error(nil),
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, CoreError}}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, io.EOF}}},
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{nil, io.EOF}}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{error(nil), io.EOF}}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, nil}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, error(nil)}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{CoreError, io.EOF}}},
			Target: CoreError,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{io.EOF, io.EOF}}},
			Target: CoreError,
			Expected: false,
		},



		//==================================================
		//==================================================
		//==================================================



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, CoreError}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, io.EOF}}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, CoreError}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, io.EOF}}},
			Target: nil,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, CoreError}}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, io.EOF}}},
			Target: nil,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, nil}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, error(nil)}}},
			Target: nil,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, CoreError}}},
			Target: nil,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, io.EOF}}},
			Target: nil,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, CoreError}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, io.EOF}}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, CoreError}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, io.EOF}}},
			Target: error(nil),
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, CoreError}}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, io.EOF}}},
			Target: error(nil),
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, nil}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, error(nil)}}},
			Target: error(nil),
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, CoreError}}},
			Target: error(nil),
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, io.EOF}}},
			Target: error(nil),
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, io.EOF}}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, io.EOF}}},
			Target: CoreError,
			Expected: false,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, nil}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, error(nil)}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, io.EOF}}},
			Target: CoreError,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, nil}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, error(nil)}}},
			Target: CoreError,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, CoreError}}},
			Target: CoreError,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, io.EOF}}},
			Target: CoreError,
			Expected: false,
		},



		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, nil}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, error(nil)}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, CoreError}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:nil}, io.EOF}}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, nil}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, error(nil)}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, CoreError}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:error(nil)}, io.EOF}}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, nil}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, error(nil)}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, CoreError}}},
			Target: io.EOF,
			Expected: false,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:CoreError}, io.EOF}}},
			Target: io.EOF,
			Expected: true,
		},
		
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, nil}}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, error(nil)}}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, CoreError}}},
			Target: io.EOF,
			Expected: true,
		},
		{
			Error:  WrappedError{err:WrappedErrors{errs:[]error{WrappedError{err:io.EOF}, io.EOF}}},
			Target: io.EOF,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := Is(test.Error, test.Target)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("ERROR:  (%T) %#v", test.Error, test.Error)
			t.Logf("TARGET: (%T) %#v", test.Target, test.Target)
			continue
		}
	}
}
