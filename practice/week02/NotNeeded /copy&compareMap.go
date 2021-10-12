package main

import (
	"fmt"
)

// For map when iterating through act to assert it has the same contents of exp, assert that the key being requested from exp exists. Hint: Use the “magic ok”.

func main(){
	exp := map[string]string{
		"india":"Taj Mahal",
		"china":"The Great Wall",
		"jordan":"Petra",
		"mexico":"Chichen Itza",
		"peru":"Machu Picchu",
		"italy":"The Colosseum",
		"brazil":"Christ The Redeemer",
	}
	act := map[string]string{}

	//creating a slice of keys with length and capacity of exp length
	keys := make([]string,0,len(exp))

	//copy the contents to the act
	for k, v := range exp{
		act[k] = v
		keys = append(keys, k)
	}
	fmt.Println(keys)
	
	//to create error
	// keys = append(keys, "error")

	fmt.Println(keys)

	unmatched := 0

	//assert that all content are present 
	for _, k := range keys{
		_ ,ok := act[k]
		if !ok {
			unmatched++
		}
	}

	if unmatched > 0 {
		fmt.Println("Key mismatch found")
	}else{
		fmt.Println("All keys present")
	}
}