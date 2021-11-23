# go-fck

Package **fck** implements tools to create and manipulate errors.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fck

[![GoDoc](https://godoc.org/github.com/reiver/go-fck?status.svg)](https://godoc.org/github.com/reiver/go-fck)

## Creating Errors

There are two ways to create errors —

With `fck.Error`:
```
	const err error = fck.Error("internal-error")
```
And with `fck.Errorf`:
```
	var err error = fck.Errorf("bad value for id %q", id)
```

**One thing to notice is that `fck.Error` errors can be a Go `const`.**
