package main

import "fmt"

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

func main() {
	// It is good to separate your "domain" code from the outside world (side-effects).
 	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
  	// So let's separate these concerns so it's easier to test
	fmt.Println(Hello("world", ""))
}
