// Copyright 2022 Giuseppe Calabrese.
// This file is distributed under the terms of the ISC license.

package errore

import "fmt"

// A batch of errors.
//
// It is meant to be used in cases in which errors can be deferred and
// then handled in bulk.
type Log []error

var _ error = Log(nil)

// Log adds an error to the log if nonnil.
func (l *Log) Log(err error) bool {
	if err == nil {
		return false
	}

	*l = append(*l, err)
	return true
}

// Error returns a combined error message.
func (l Log) Error() string {
	if l == nil {
		panic("Error() called on nil error log")
	}

	if len(l) == 1 {
		return l[0].Error()
	}

	msg := fmt.Appendf(
		nil, "%s\n... and other (%d) errors:", l[0], len(l)-1)
	for _, e := range l[1:] {
		msg = fmt.Appendf(msg, "\n=> %s", e)
	}
	return string(msg)
}

// Unwrap returns the oldest error.
func (l Log) Unwrap() error {
	return l[0]
}
