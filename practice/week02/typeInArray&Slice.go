package main

import "fmt"

func main(){
	names := [2]string{}
	//results to an error
	//cannot use 1 (type untyped int) as type string in assignment
	names[0] = 1
	fmt.Println(names)
}