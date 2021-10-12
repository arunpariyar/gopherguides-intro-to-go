package main

func main() {

}

func copyArray(main [7]string, dup [7]string) int {
	//variable to keep track of unmatched content
	noOfMismatch := 0

	//looping through main
	for i, v := range main {
		//copying values to array act in the corresponding index
		dup[i] = v
		
		//TO CREATE ERROR
		// dup[i] = v + "error"
	}

	//loop through main
	for i := range main {
		if main[i] != dup[i] {
			//incrementing when mismatch is found
			noOfMismatch++
		}
	}
	return noOfMismatch
}

func copySlice(main []string, dup []string) bool {
	
	//using append while unpacking the main slice using the .... notation
	dup = append(dup, main...)

	
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

	//TO CREATE ERROR
	// keys = append(keys, "error")

	//loop through the keys
	for _, k := range keys {
		//checking if the keys exist in dup
		_, ok := dup[k]
		if !ok {
			//incrementing mismatch to keep track of mismatches
			mismatch++
		}
	}

	return mismatch
}
