package main

import "fmt"

func main(){
 //functions as variables
 f := func() string{
	 return "Hello"
 } 
 //storing the result of running the function
 result := f()

 fmt.Println(result)
}
