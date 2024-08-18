package main

import "fmt"

const (
	spanish            = "Spanish"
	engHelloPrefix     = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return engHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris", ""))
}
