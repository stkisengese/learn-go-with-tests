package main

import "fmt"

const engHelloPrefix = "Hello, "

func hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return engHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris", ""))
}
