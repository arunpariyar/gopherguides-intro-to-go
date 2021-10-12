package main

import "fmt"

func main(){
	fmt.Println(MeaningOfLife())
}

func MeaningOfLife() (rc int) {
	defer func() { rc = 0 }()
	return 41
}