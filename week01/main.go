package main

import "fmt"

func main(){
	//printing hello world
	fmt.Println("Hello, World!")
	
	//using printf and correct print verbs to print the respective values
	fmt.Printf("Printing, %s, %d, %t \n", string("Go"), int(42), bool(true))
}