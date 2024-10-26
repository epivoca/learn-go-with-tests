package maps

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "what is test, actually? in terms of programming, I guess, it's just a way to make sure, that your code runs properly"
	dictionary := map[string]string{word: definition}

	got := Search(dictionary, word)
	want := definition

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, dictionary)
	}
}
