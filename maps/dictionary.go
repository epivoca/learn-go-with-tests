package maps

type Dictionary map[string]string

// Search searches for a word in the dictionary
// and returns its definition
func (d Dictionary) Search(word string) string {
	return d[word]
}
