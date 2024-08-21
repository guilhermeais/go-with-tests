package main

import "fmt"

const (
	Portuguese = "Portuguese"
	Spanish    = "Spanish"

	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Olá, "
	spanishHelloPrefix    = "Hola, "
)

func Hello(_name, language string) string {
	name := "World"
	helloPrefix := englishHelloPrefix

	if _name != "" {
		name = _name
	}
	switch language {
	case Portuguese:
		helloPrefix = portugueseHelloPrefix
	case Spanish:
		helloPrefix = spanishHelloPrefix
	}

	return helloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("", ""))
}
