package main

import "fmt"

var helloPrefixTranslation map[string]string = map[string]string {
    "ru": "Привет, ",
    "en": "Hello, ",
    "es": "Hola, ",
}


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if helloPrefix, helloPrefixExist := helloPrefixTranslation[language]; helloPrefixExist {
		return helloPrefix + name + "!"
	}

	return helloPrefixTranslation["en"] + name + "!"
}

func main() {
	// It is good to separate your "domain" code from the outside world (side-effects).
 	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
  	// So let's separate these concerns so it's easier to test
	fmt.Println(Hello("world", ""))
}
