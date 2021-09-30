package main

import "fmt"

func main(){
	//printing hello world
	fmt.Println("Hello, World!")
	
	//using printf and correct print verbs to print the respective values
	fmt.Printf("Printing, %s \n", string("Go"))
	fmt.Printf("Printing, %d \n", int(42))
	fmt.Printf("Printing, %t \n", bool(true))
}