package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris!"

	if got != want {
		// For tests %q is very useful as it wraps your values in double quotes.
		t.Errorf("got %q, want %q", got, want)
	}
}
