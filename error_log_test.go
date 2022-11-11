// Copyright 2022 Giuseppe Calabrese.
// This file is distributed under the terms of the ISC license.

package errore

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleLog() {
	var errs Log

	tryDoSomething := func(i int) error {
		return fmt.Errorf("example error #%d", i)
	}

	for i := 0; i < 5; i++ {
		errs.Log(
			tryDoSomething(i),
		)
	}

	fmt.Println(errs)
	// Output:
	// example error #0
	// ... and other (4) errors:
	// => example error #1
	// => example error #2
	// => example error #3
	// => example error #4
}

func TestLog_Log(t *testing.T) {
	var errs Log

	if errs.Log(nil) {
		t.Error(".Log(nil) should return false")
	}

	if !errs.Log(errors.New("a")) || !errs.Log(errors.New("b")) {
		t.Error(".Log(nonnil) should return true")
	}

	if len(errs) != 2 {
		t.Error(".Log() should modify its receiver")
	}
}

func TestLog_Error_nil(t *testing.T) {
	var errs Log
	defer func() {
		if v := recover(); v == nil {
			t.Error(".Error() should fail on nil receiver")
		}
	}()
	_ = errs.Error()
}

func TestLog_Error(t *testing.T) {
	var errs Log
	errs.Log(errors.New("a"))
	errs.Log(errors.New("b"))

	// Test formatting.
	want := `a
... and other (1) errors:
=> b`
	if got := errs.Error(); got != want {
		t.Errorf("got message %#v; want %#v", got, want)
	}
}

func TestLog_Unwrap_nil(t *testing.T) {
	var errs Log
	defer func() {
		if v := recover(); v == nil {
			t.Error(".Unwrap() should fail on nil receiver")
		}
	}()
	errs.Unwrap()
}

func TestLog_Unwrap(t *testing.T) {
	var errs Log
	err := errors.New("a")
	errs.Log(err)
	errs.Log(errors.New("b"))

	if errs.Unwrap() != err {
		t.Error(".Unwrap() should return the first error")
	}
}
