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
	helloPrefix := greetingPrefix(language)

	if _name != "" {
		name = _name
	}

	return helloPrefix + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Portuguese:
		prefix = portugueseHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
