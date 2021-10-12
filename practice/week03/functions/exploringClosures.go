package main

import "fmt"

func main(){
	innerFn := outerFn()
	fmt.Println(innerFn())
}

func outerFn() func()int {
	value := 0
	fmt.Println("value initilised to ", value, "outer function closes \n")

	return func() int{
		value++
		fmt.Printf("in the internal function now and the value is %d \n", value)
		return value
	}
}