package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	engHelloPrefix     = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("Chris", ""))
}
