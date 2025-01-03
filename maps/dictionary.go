package maps

// Dictionary store definitions to words.
type Dictionary map[string]string

const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search fins a word in the dictionary and returns its definition
func (d Dictionary) Search(word string) (string, error) {
	definition, definitionExists := d[word]
	if !definitionExists {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a new word and a definition to the dictionary
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

// Update changes definition for an existing word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}
