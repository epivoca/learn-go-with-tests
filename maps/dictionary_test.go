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
		assertError(t, err, nil)
		assertStrings(t, got, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("dafsfasdf")
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

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "cat"
		definition := "cute animal with four legs"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update definition for existing word", func(t *testing.T) {
		word := "dog"
		definition := "friendliest pet"
		dictionary := Dictionary{word: definition}

		newDefinition := "frinedly version of a wolf"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update definition for new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "dog"
		definition := "friendliest pet"

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assertError(t, err, nil)
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
