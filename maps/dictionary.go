package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// Search searches for a word in the dictionary
// and returns its definition
func (d Dictionary) Search(word string) (string, error) {
	definition, definitionExists := d[word]
	if !definitionExists {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a new word and a definition to the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
