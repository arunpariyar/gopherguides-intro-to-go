package main

import "fmt"

func main() {
	trinity := []string{}
	otherForms := []string{"Krishna", "Ram", "Buddha"}
	allGods := []string{}

	//append one item to trinity
	trinity = append(trinity, "Brahma")
	fmt.Println(trinity)

	//append multiple item to trinity
	trinity = append(trinity, "Vishnu", "Mahesh")
	fmt.Println(trinity)

	//append a slice with another slice
	allGods = append(allGods, trinity...)

	//append a slice with another slice
	allGods = append(allGods, otherForms...)

	fmt.Println(allGods)
}
