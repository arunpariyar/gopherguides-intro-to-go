package main

import "fmt"

func main(){
	name := "John"
	DOB := 1990

	calcAge(DOB)
	fmt.Println("The Age is", DOB)

	greet, age, isAdult := info(name, DOB)

	fmt.Println(greet)
	fmt.Println(age)
	fmt.Println(isAdult)
}
// with single return
func calcAge(DOB int) int {
	return 2021 - DOB
}

func info(name string, DOB int) ( string,  int,  bool){
	greet := "Hey my name is " + name
	age := 2021 - DOB
	isAdult := false
	if(age >= 18){
		isAdult = true
	}
	return greet, age, isAdult
}