package maps

// Search searches for a word in the dictionary
// and returns its definition
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
