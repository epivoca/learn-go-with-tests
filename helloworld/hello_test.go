package helloworld

import "testing"

func TestHello(t *testing.T) {
	// Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.
	// A benefit of this approach is you can set up shared code that can be used in the other tests.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty strings are supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hola, World!' when an empty name is supplied", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Chris", "es")
		want := "Hola, Chris!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call (lines 11 and 16)
	// rather than inside our test helper (line 26).
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
