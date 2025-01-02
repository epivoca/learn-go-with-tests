package maps

import "errors"

type Dictionary map[string]string

// Search searches for a word in the dictionary
// and returns its definition
func (d Dictionary) Search(word string) (string, error) {
	definition, definitionExists := d[word]
	if !definitionExists {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
