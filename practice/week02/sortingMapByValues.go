package main

import (
	"fmt"
	"sort"
)

func main() {
	words := map[int]string{
		0: "e",
		1: "c",
		2: "a",
		3: "b",
		4: "d",
	}

	//creating an inverted map for words
	sorted := map[string]int{}

		fmt.Println(sorted)

	//creating a slice to store all the keys
	keys := make([]string, 0, len(words))

	//looping through keys
	for k, v := range words {
		//keys being populated with all the values
		keys = append(keys, v)
		//setting the key value pair for sorted 
		sorted[v] = k
	}
	
	sort.Strings(keys)
	
	//looping through keys
	for _, k := range keys {
		//printing the keys of sorted alongside key slice
		fmt.Println(sorted[k], k)
	}
}
