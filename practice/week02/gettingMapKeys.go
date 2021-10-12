package main

import (
	"fmt"
	"sort"
)

func main(){
	tmnt := map[string]string{}

	tmnt["Leo"] = "Leonardo"
	tmnt["Don"] = "Donatello"
	tmnt["Mike"] = "Micheal Angelo"
	tmnt["Ralph"] = "Raphael"

	//print the map
	// for k, v := range tmnt{
	// 	fmt.Println(k,v)
	// }

	//***retrevings the keys
	//creating a slice of string with length of zero and capcity equal to the length of the tmnt map.
	keys := make([]string, 0, len(tmnt))

	//using for loop to go through all of tmnt
	for k := range tmnt{
		//appending the key to the keys slice
		keys = append(keys, k)
	}

	fmt.Printf("Keys collected %v \n",keys)

	//sorting
	sort.Strings(keys)
	fmt.Printf("Keys after sorting %v \n",keys)

	
	for _, k := range keys {
		fmt.Println(k, tmnt[k])
	}

	





}