package maps

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "google.com"
	noDefinition := ""

	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search(word)

		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("dafsfasdf")
		if err == nil {
			t.Fatal("expected to get an errror")
		}

		assertStrings(t, got, noDefinition)
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	newWord := "cat"
	newWordDefinition := "cute animal with four legs"

	t.Run("add word", func(t *testing.T) {
		dictionary.Add(newWord, newWordDefinition)
		assertDefinition(t, dictionary, newWord, newWordDefinition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
