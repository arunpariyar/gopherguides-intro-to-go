package main

import "fmt"

func main() {
	emails := map[string]string{
		"john": "John_doe@gmail.com",
		"mary": "mary_little@gmail.com",
	}

	key := "tammy"

	value, ok := emails[key]
	if !ok {
		fmt.Printf("The key %q doesn't exist \n", key)
	}else{
		fmt.Printf("We found the key %q and the value is %q \n", key, value)
	}
	
	
}

