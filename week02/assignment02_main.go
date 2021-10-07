package main

func main() {

}

func copyArray(main [7]string, dup [7]string) int {
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

func copySlice(main []string, dup []string) bool {
	//looping through exp
	for _, v := range main {
		//using append to push the values to act
		dup = append(dup, v)

	}
	//TO CREATE ERROR
	//adding one more value to the slice manually
	// dup = append(dup, "ERROR")

	//comparing the length of both act and exp
	if len(main) != len(dup) {
		return false
	} else {
		return true
	}
}

func copyMap(main map[string]string, dup map[string]string) int {
	//creating a slice of keys with length and capacity of exp length
	keys := make([]string, 0, len(main))

	mismatch := 0

	//copy the contents to the act
	for k, v := range main {
		dup[k] = v
		keys = append(keys, k)
	}

	//to create error
	keys = append(keys, "error")
	for _, k := range keys {
		_, ok := dup[k]
		if !ok {
			mismatch++
		}
	}

	return mismatch

}
