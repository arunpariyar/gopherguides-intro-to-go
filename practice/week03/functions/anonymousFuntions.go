package main

import "fmt"

func main() {
	//anonymousFunction with out arguments
	func() {
		fmt.Println("I am Anonymous")
	}()
	//anonymous Function with arguments
	func(x int) {
		fmt.Println("The Meaning of Life is", x)
	}(42)

	//anonymous Function with return
	value := func(x string) string {
		return "The Meaning of Life is " + x
	}("42")

	fmt.Println(value)

	sayHello()

}

func sayHello(func() string {
  return "Hello"
})

