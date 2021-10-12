package main

import "fmt"

func main(){
	var phonebook map[string]int

	//if a maps is not initilized and we use it we get a assignment to entry in nil map panic 
	phonebook["arun"] = 123456
	fmt.Println(phonebook["arun"])

}