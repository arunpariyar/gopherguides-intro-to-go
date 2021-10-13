package main

import "fmt"

type Movie struct {
	Name string
	Length int 
	rating float64
	plays int
	viewers int 
}

type Critiquefn func( *Movie)(float32, error)

type Theatre struct {

}

func (m *Movie) Rate(rating float32) error {

	//Calling Rate should track this rating. If the number of plays is 0 return the following error: fmt.Errorf("can't review a movie without watching it first")

	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}else{ 
		return nil
	}

}

func (m *Movie) Play(viewers int){
	//Calling Play should increase both the number of viewers, as well as the number of plays, for the movie.
	m.viewers += viewers
	m.plays += viewers
}

func (m Movie) Viewers() int {
	//returns the number (int) of people who have viewed the movie.
	return m.viewers
}

func (m Movie) Rating() float64 {
	//Rating takes no arguments and returns the rating (float64) of the movie. This can be calculated by the total ratings for the movie divided by the number of times the movie has been played.
	rating := m.rating / float64(m.plays)

	return rating
}

func (m Movie) String() string {
  // String should return a string that that includes the name, length, and rating of the film. Ex. Wizard of Oz (102m) 99.0%	
  return fmt.Sprintf("%s (%dm) %.1f", m.Name, m.Length, m.rating)
}


/* 
type Movie struct {
	Name string
	Length int 
	rating float64
	plays int
	viewers int 
*/

func main() {
	 m := Movie{
		 Name : "Wizard of Oz",
		 Length : 102,
		 rating : 0,
		 plays : 10000,
		 viewers: 10000,
	 }

	fmt.Println(m.String())
	
}