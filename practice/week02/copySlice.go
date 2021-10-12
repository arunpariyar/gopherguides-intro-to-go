package main

import "fmt"

func main() {
	// creating a exact copy of a slice that does that reference to the original slice itself

	//the original slice
	original := []string{"earth", "water", "fire", "air"}
	//a reference of the original
	ref := original
	//creating duplicate
	dup := make([]string, len(original))
	
	//using copy to copy all items from original 
	copy(dup, original)

	fmt.Println(original)
	fmt.Println(ref)
	fmt.Println(dup)
}