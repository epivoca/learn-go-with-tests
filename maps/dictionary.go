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
	_, err := d.Search(word)

	// Must abort Add for existing words because it would have a modify them which is not
	// an expected behavior
	//
	// Having a switch like this provides an extra safety net,
	// in case Search returns an error other than ErrNotFound.
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
