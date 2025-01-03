package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("provided word already exists")
)

// Search searches for a word in the dictionary and returns its definition
// or ErrNotFound if no definition was found
func (d Dictionary) Search(word string) (string, error) {
	definition, definitionExists := d[word]
	if !definitionExists {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a new word and a definition to the dictionary and returns ErrWordExists
// if a word already exists
func (d Dictionary) Add(word, definition string) error {
	// Must abort Add for existing words because it would have a modifying behaviour which is not
	// an expected behavior
	if _, err := d.Search(word); err == nil {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}
