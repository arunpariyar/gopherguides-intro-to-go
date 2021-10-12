package main

import "fmt"

func main(){
	heroes := []string{"Superman","Batman","Spiderman","Aquaman","Antman"}	
	fmt.Println(heroes)

	fmt.Println("length of heroes slice:", len(heroes))
	fmt.Println("capacity of heroes slice:", cap(heroes))
}