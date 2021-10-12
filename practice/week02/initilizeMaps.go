package main

import "fmt"

func main(){
	// the recommendation is to initilize during time of variable decleration
	highestGrossingFilms := map[string]int{
		"avatar": 2_847_246_203,
		"avengers endgame":2_797_501_328,
		"titanic":2_187_425_379,
	}
	fmt.Println(highestGrossingFilms)
	fmt.Printf("Avater is the highest films making %d in the cinemas \n", highestGrossingFilms["avatar"])

	//deleting keys from a map
	delete(highestGrossingFilms, "titanic")
	fmt.Println(highestGrossingFilms)
}