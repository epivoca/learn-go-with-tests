package maps

type Dictionary map[string]string

// Search searches for a word in the dictionary
// and returns its definition
func (d Dictionary) Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
