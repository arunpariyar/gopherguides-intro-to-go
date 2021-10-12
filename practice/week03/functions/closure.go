package main

import "fmt"

func main() {

	name := "Secret data"

	IcanAcceesName := func(){
		fmt.Printf("I can access %q \n", name)
	}
	runFunction(IcanAcceesName)
}

func runFunction(fn func()){
	fn()
}

