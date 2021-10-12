//the base case in simple terms where functions can be used as arguments is if we have similar tasks to be performed on a variable, then it is possible to chain then as well
package main

import "fmt"

func main(){

	result := 0

	result = addTen(addTwenty(result)) // chaining the functions together // may not be according to effective Go

	fmt.Println(result)

}

func addTen(nbr int) int {
	return nbr + 10
}

func addTwenty(nbr int) int {
	return nbr + 20 
}

func functionApply (fn func(int) int, operand int)int { 
	return fn(operand)
}





