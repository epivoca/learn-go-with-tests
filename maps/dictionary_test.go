package maps

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "google.com"
	noDefinition := ""
	unknownWordErrMsg := "could not find the word you were looking for"

	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(word)

		assertStrings(t, got, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("dafsfasdf")
		if err == nil {
			t.Fatal("expected to get an errror")
		}

		assertStrings(t, err.Error(), unknownWordErrMsg)
		assertStrings(t, got, noDefinition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
