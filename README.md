# go-erorr

Package **erorr** implements tools to create and manipulate errors, for the Go programming language.

(This package used to be called "fck".)

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-erorr

[![GoDoc](https://godoc.org/github.com/reiver/go-erorr?status.svg)](https://godoc.org/github.com/reiver/go-erorr)

## Creating Errors

There are two ways to create errors —

With `erorr.Error`:
```
	const err error = erorr.Error("internal-error")
```
And with `erorr.Errorf`:
```
	var err error = erorr.Errorf("bad value for id %q", id)
```

**One thing to notice is that `erorr.Error` errors can be a Go `const`.**
