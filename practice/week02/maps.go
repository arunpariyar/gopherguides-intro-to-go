package main

import "fmt"

func main() {
	// creating a map that uses key of string and values of string
	tmnt := map[string]string{}

	tmnt["Leo"] = "Leonardo"
	tmnt["Don"] = "Donatello"
	tmnt["Mike"] = "Micheal Angelo"
	tmnt["Ralph"] = "Raphael"

	fmt.Printf("%v \n", tmnt)

	//length
	fmt.Println(len(tmnt))

	//capacity Maps can hold unlimited values to using cap(tmnt) will result in an error - Invalid argument tmnt (type map[string]string) for cap

	//fmt.Println(cap(tmnt)

	//reteriving values from map
	fmt.Printf("%q \n", tmnt["Leo"])

	key := "Tammy"

	//using okay to check if key exists
	value, ok := tmnt[key]
	if !ok {
		fmt.Printf("Key Not Found: %q \n", key)
	}else{
		fmt.Printf("Key Found : %q = %v \n", key, value)
	}
	
}