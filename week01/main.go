package main

import "fmt"

func main(){
	//printing hello world
	fmt.Println("Hello, World!")
	
	//various variable declartion and initialization
	var language string	// variable declaration without initialization
	num := uint8(42) 	// short declartion with desired data type
	isTrue := true 		// short declaration
	language = "Go" 	// initialization

	//using printf and correct print verbs to print the respective values
	fmt.Printf("Printing, %s \n", language)
	fmt.Printf("Printing, %d \n", num )
	fmt.Printf("Printing, %t \n", isTrue)
}