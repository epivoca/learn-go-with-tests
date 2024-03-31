package iteration

// strings.Repeat(character, repeatCount) is standard lib solution for that task
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
