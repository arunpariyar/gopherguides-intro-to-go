package main

type Movie struct { 
	Name string
	Length int
	rate float32
	play int
	viewers int 
	plays int
	rating float64
}

// Defining a new function type
type CritiqueFn func (*Movie) (float32, error)

//creating methog Rate

func (m *Movie) Rate(float32) error {
	//Calling Rate should track this rating. If the number of plays is 0 return the following error: fmt.Errorf("can't review a movie without watching it first")	
}




func main(){

}	