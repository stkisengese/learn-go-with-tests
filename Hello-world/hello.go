package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	engHelloPrefix     = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}
	prefix := engHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(hello("Chris", ""))
}
