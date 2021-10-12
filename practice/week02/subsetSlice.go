package main

import "fmt"

func main() {
	avengers := []string{"Ironman", "Thor", "Captain America", "Hawk", "Scarlet Witch", "Hulk", "Black Widow", "Vision"}

	fmt.Println(avengers)
	fmt.Println(avengers[:3])  // up to the third index
	fmt.Println(avengers[0:3]) //same as the above
	fmt.Println(avengers[2:5])

	fmt.Println(avengers[5:])              // from the fith index to last
	fmt.Println(avengers[5:len(avengers)]) // same as the above
	// Very important to note that mutating the subset mutates the original slice
	leader := avengers[:1]

	fmt.Println(leader)
	leader[0] = "Aquaman" // this will change Ironman to Aquaman in the actual avengers slice

	fmt.Println(avengers)
}
