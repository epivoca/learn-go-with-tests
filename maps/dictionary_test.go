package maps

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "google.com"

	// dictionary := map[string]string{word: definition}
	dictionary := Dictionary{word: definition}

	got := dictionary.Search(word)
	want := definition

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
