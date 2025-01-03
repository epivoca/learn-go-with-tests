package maps

import (
	"testing"
)

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
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "cat"
		definition := "cute animal with four legs"

		err := dictionary.Add(word, definition)
		if !err {
			t.Fatal("should add new word:", err)
		}
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "cat"
		definition := "cute animal with four legs"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
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
