package main

func main() {

}

func copyArray(main [7]string, dup [7]string ) int {
	//variable to keep track of unmatched content
	noOfunmatchContents := 0

	//looping through exp
	for i, v := range main {
		//copying values to array act in the corresponding index
		dup[i] = v
		//TO CREATE ERROR
		// dup[i] = v + "error"
	}
	
	//loop through act
	for i := range main {
		if main[i] != dup[i] {
			//incrementing when mismatch is found
			noOfunmatchContents++
		}
	}
	return noOfunmatchContents
}

func copySlice( main []string, dup []string) bool {
	//looping through exp
	for _, v := range main  {
		//using append to push the values to act
		dup = append(dup, v)
		
	}
	//TO CREATE ERROR 
	//adding one more value to the slice manually 
	// dup = append(dup, "ERROR")

	//comparing the length of both act and exp
	if len(main) != len(dup){
		return false
	}else{
		return true
	}
}


