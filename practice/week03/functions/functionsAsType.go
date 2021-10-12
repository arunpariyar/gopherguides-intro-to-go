package main

import "fmt"

type greeter func() string

func main() {
	greet(func() string{
		return "My type is greeter"
	})
}

func greet(fn greeter) {
	fmt.Println(fn())
}
