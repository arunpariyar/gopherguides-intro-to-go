package main

import "fmt"

func main(){
	//here i am saying please make a slice of string with the length of 5 and capacity of 15 
	heros := make([]string,5,15)

	fmt.Println("Length of heroes slice", len(heros))
	fmt.Println("Capacity of heroes slice",cap(heros))

	fmt.Printf("%v \n", heros)


}