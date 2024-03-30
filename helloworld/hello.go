package helloworld

const (
	english = "en"
	spanish = "es"
	russian = "ru"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	russianHelloPrefix = "Привет, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
