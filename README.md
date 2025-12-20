# go-erorr

Package **erorr** provides tools for creating and manipulating **errors**, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/codeberg.org/reiver/go-erorr

[![GoDoc](https://godoc.org/codeberg.org/reiver/go-erorr?status.svg)](https://godoc.org/codeberg.org/reiver/go-erorr)

## History

This package's origin goes back to _July 23rd, 2015 and before_.

It started off as package `manyerrors` — which provided a way to combined multiple errors into a single error.
Later package `fck` was created — which provided a way to create `const` errors (rather than having to use `var`).
Eventually, package `fck` was renamed to package `erorr` — and later the concepts from package `manyerrors` were merged into it.
Later, package `erorr` provides tool for creating error annotations and contextual errors.
And, finally, package `erorr` switched to using `Unwrap() error` and `Unwrap() []error` methods to match the conventions in the built-in Golang `"errors"` package.

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
