# go-erorr

Package **erorr** implements tools to create and manipulate errors, for the Go programming language.

(This package used to be called "fck".)

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/codeberg.org/reiver/go-erorr

[![GoDoc](https://godoc.org/codeberg.org/reiver/go-erorr?status.svg)](https://godoc.org/codeberg.org/reiver/go-erorr)

## Creating Errors

There are two ways to create errors —

With `erorr.Error`:

```
	const err error = erorr.Error("internal-error")
```

**NOTICE THAT THAT ERROR IS A `const` RATHER THAN A `var`.**

**Errors creating using `erorr.Error` can be a Go `const`.**

And with `erorr.Errorf`:

```
	var err error = erorr.Errorf("bad value for id %q", id)
```

## Import

To import package **erorr** use `import` code like the follownig:
```
import "codeberg.org/reiver/go-erorr"
```

## Installation

To install package **erorr** do the following:
```
GOPROXY=direct go get https://codeberg.org/reiver/go-erorr
```

## Author

Package **erorr** was written by [Charles Iliya Krempeaux](http://reiver.link)
