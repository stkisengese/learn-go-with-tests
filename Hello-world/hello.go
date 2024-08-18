package main

import "fmt"

const engHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
        return "Hello, World"
    }
	return engHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris"))
}
