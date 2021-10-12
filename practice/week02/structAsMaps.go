package main

import "fmt"

func main() {
	type staff struct {
		ID   int
		Name string
	}

	s1 := staff{ 1, "arun"}

	fmt.Println(s1)

	m := map[staff]string{
		s1: "programmer",
	}

	fmt.Println(m)

	fmt.Printf("Staff %v, ID number %d, is a %v \n.", s1.Name, s1.ID, m[s1])

}