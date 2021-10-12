package main

import "fmt"

func main() {
	//declaration and initialisation of exp consisting 7 wonders of the world
	exp := [7]string{"Great Wall of China", "Chichén Itzá", "Petra", "Machu Picchu", "Colosseum", "Taj Mahal", "Christ the Redeemer"}

	//creating a new array that has the same length as exp
	act := [len(exp)]string{}

	//variable to keep track of unmatched content
	noOfunmatchContents := 0
	//creating an array to store mismatch content with length of act 
	unmatchContents := [len(act)]string{}

	//looping through exp
	for i, v := range exp {
		//copying values to array act in the corresponding index
		act[i] = v
	}
	
	//loop through act
	for i, v := range act {
		if act[i] != exp[i] {
			//adding unmatched content to unmatchContents array
			unmatchContents[i] = v
			//incrementing when mismatch is found
			noOfunmatchContents++
		}
	}

	fmt.Printf("The number of unmatch contents : %v \n", noOfunmatchContents)

	if noOfunmatchContents != 0 {
		fmt.Println("Here is the list of contents that didn't match: ")
		//printing the contents that didnt match
		for i, v := range unmatchContents {
			fmt.Println(i+1, v)
		}
		fmt.Println("Arrays act and exp are not identical.")
	} else {
		fmt.Println("Arrays act and exp are identical.")
	}
}
