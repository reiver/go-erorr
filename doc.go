/*
# Core Errors

To create your core errors, use [Error].
For example:

	const (
		ErrNilReceiver = erorr.Error("nil receiver")
		ErrNotFound    = erorr.Error("not found")
		ErrUndefined   = erorr.Error("undefined")
	)

Note that you can create errors with [Error] using `const` (rather than `var`).

# Error Annotations

Use [Wrap] to annotate your core errors and other errors.

	if nil != err {
		return erorr.Wrap(err, "API request failed.",
			field.String("request-uri", requestURI),
			field.String("service", "monitor"),
		)
	}


Note that error annotations of this type automatically contains a call-trace.
*/
package erorr

