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

## Example

Here are some examples.

### Core Errors

Create your **core errors** with source-code similar to the following:

```golang
const (
	ErrNilReceiver = erorr.Error("nil receiver")
	ErrNotFound    = erorr.Error("not found")
	ErrUndefined   = erorr.Error("undefined")
)
```

Notice that a `const` was used (rather than a `var`).

### Annotated Errors

Create an **annotated error** with source-code similar to the following:

```golang
if nil != err {
	return erorr.Wrap(err, "API request failed.",
		field.String("request-uri", requestURI),
		field.String("service", "monitor"),
	)
}
```

### Contextual Errors

Create a **contextual error** with source-code similar to the following:

```golang
if nil != err {
	return erorr.Wrap("API request failed.",
		field.String("request-uri", requestURI),
		field.String("service", "monitor"),
	)
}
```

Note that `erorr.Stamp()` is similar to `erorr.Wrap()`, with the difference being that, `erorr.Stamp()` does not wrap another error.

### Errorf

The function `erorr.Errorf()` is also available.
For example:

```golang
if !isValidID(id) {
	return erorr.Errorf("bad value for id %q", id)
}
```

Or also, for example:

```golang
if nil != err {
	return erorr.Errorf("Uh oh, something bad happened to %s: %w", service, err)
}
```

Note that a different because this package's `erorr.Errorf()` function and the `fmt.Errorf()` in the built-in Go `"fmt"` package is that the error returned from this package's `erorr.Errorf()` function will include a call-trace and other annotation and contextual information.

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
