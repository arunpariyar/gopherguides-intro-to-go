package main

import "fmt"

func main(){
	greeting := "Hey I am from the sayHello function"
	//sayHello with argumentß
	sayHello(greeting)
	
	james :="Hi! I am james"
	laura := "Hi I am Laura"
	
	multipleHello(james, laura)
}

// function with single argument
func sayHello(greeting string){
	fmt.Println(greeting)
}

//function with multiple argument
func multipleHello(greet1 string , greet2 string){
	fmt.Println(greet1)
	fmt.Println(greet2)
}

