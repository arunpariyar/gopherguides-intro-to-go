package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var fn greetFn

	fn := sayHello
	greet(fn)

	fn = func() string {
		return strings.Join(os.Args, ", ")
	}
	greet(fn)

	greet(sayHello)

	greet(func() string {
		return "Arun"
	})

}

// greetFn does something
type greetFn func() string

func greet(fn greetFn) {
	s := fn()
	fmt.Printf("Hello, %s\n", s)
}

func sayHello() string {
	return "Hello"
}