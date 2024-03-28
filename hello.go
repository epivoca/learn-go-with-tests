package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	// It is good to separate your "domain" code from the outside world (side-effects).
 	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
  	// So let's separate these concerns so it's easier to test
	fmt.Println(Hello("world"))
}
