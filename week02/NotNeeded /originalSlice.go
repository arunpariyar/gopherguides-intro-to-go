package main

func main() {

	// //*** for slice and map types assert that the length of act and exp are the same
	
	// //Slice of the 10 highest mountains in the world
	// exp := []string{"Mount Everest", "K2", "Kangchenjunga", "Lhotse", "Makalu", "Cho Oyu", "Dhaulagiri I", "Manaslu", "Nanga Parbat", "Annapurna I"}

	// //creating act slice with the length of exp
	// act := make([]string, 0, len(exp))

	// isSame := copySlice(exp, act)

	// if(isSame){
	// 	fmt.Println("The length of exp and act is the same")
	// }else{
	// 	fmt.Println("The length of exp and act is not the same")
	// }
}

func copySlice( main []string, dup []string) bool {
	//looping through exp
	for _, v := range main  {
		//using append to push the values to act
		dup = append(dup, v)
		//to create error
		// dup = append(dup, "Error")
	}

	//comparing the length of both act and exp
	if len(main) != len(dup){
		return false
	}else{
		return true
	}
}