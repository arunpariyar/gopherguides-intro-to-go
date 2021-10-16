package main

import (
	"fmt"
	"testing"
)

func Test_Rate(t *testing.T) {

	m := &Movie{
		Name:    "Avatar",
		Length:  16,
		rating:  98.75,
		plays:   4,
		viewers: 400,
	}

	err := m.Rate(m.rating)
	if err != nil {
		t.Error(err)
	}
}

func Test_Play(t *testing.T) {

	m := &Movie{
		Name:    "Avengers: Endgame",
		Length:  181,
		rating:  88.75,
		plays:   3,
		viewers: 300,
		critics: []float32{9, 9, 9, 8.5},
	}

	//storing the value of play and viewer before adding new viewers
	playsBefore := m.plays
	viewersBefore := m.viewers

	//new viewers for the movie that we want to add
	newViewers := 100

	//call the Play function with
	m.Play(newViewers)

	//storing new values of play and viewer after adding new viewers
	playsAfter := m.plays
	viewersAfter := m.viewers
	// calculating the difference between the before and after adding new viewers
	playsDiff := playsAfter - playsBefore
	viewsDiff := viewersAfter - viewersBefore

	if (playsDiff != 1) || (viewsDiff != newViewers) {
		t.Errorf("expected plays 1 got %d expected viewers %d got %d \n", playsDiff, newViewers, viewsDiff)
	}
}

func Test_Viewers(t *testing.T) {

	m := Movie{
		Name:    "Titanic",
		Length:  195,
		rating:  0,
		plays:   2,
		viewers: 200,
	}
	//storing the result of calling Viewers Method
	result := m.Viewers()
	//logging error if difference found
	if result != m.viewers {
		t.Errorf("expected %d got %d \n", m.viewers, result)
	}
}

func Test_Plays(t *testing.T) {
	m := Movie{
		Name:    "Star Wars: The Force Awakens",
		Length:  135,
		rating:  0,
		plays:   100,
		viewers: 100,
	}
	//storing the result of calling the Plays Method
	result := m.Plays()

	//logging error if difference found
	if result != m.plays {
		t.Errorf("expected %d got %d \n", m.plays, result)
	}
}

func Test_Rating(t *testing.T) {
	m := Movie{
		Name:    "Avatar",
		Length:  16,
		rating:  0,
		plays:   4,
		viewers: 400,
	}
	//calcuating rating
	rating := float64(m.rating) / float64(m.plays)
	//storing the result of calling Rating fuction
	result := m.Rating()

	if result != rating {
		t.Errorf("expected %.2f got %.2f \n", rating, result)
	}
}

func Test_String(t *testing.T) {
	m := Movie{
		Name:    "Avatar",
		Length:  16,
		rating:  0,
		plays:   4,
		viewers: 400,
	}

	expString := fmt.Sprintf("%s (%dm) %.1f%%", m.Name, m.Length, m.rating)

	result := m.String()

	if result != expString {
		t.Errorf("expected %q got %q \n.", expString, result)
	}
}

func Test_TPlay(t *testing.T) {

	m1 := &Movie{
		Name:    "Avatar",
		Length:  16,
		rating:  0,
		plays:   4,
		viewers: 400,
		critics: []float32{10, 10, 10, 9.5},
	}

	m2 := &Movie{
		Name:    "Avengers: Endgame",
		Length:  181,
		rating:  0,
		plays:   3,
		viewers: 300,
		critics: []float32{9, 9, 9, 8.5},
		
	}

	m3 := &Movie{
		Name:    "Titanic",
		Length:  195,
		rating:  0,
		plays:   2,
		viewers: 200,
		critics: []float32{8, 8, 8, 7.5},
	}

	m4 := &Movie{
		Name:    "Star Wars: The Force Awakens",
		Length:  135,
		rating:  0,
		plays:   100,
		viewers: 100,
		critics: []float32{7, 7, 7, 6.5},
	}

	t1 := Theatre{
		name: "The Wall",
	}

	mlist := []*Movie{m1, m2, m3, m4}

	err := t1.TPlay(100, mlist...)
	if err != nil {
		t.Error(err)
	}

}

func Test_Critique(t *testing.T){
	m1 := &Movie{
		Name:    "Avatar",
		Length:  16,
		rating:  0,
		plays:   4,
		viewers: 400,
		critics: []float32{10, 10, 10, 9.5},
	}

	m2 := &Movie{
		Name:    "Avengers: Endgame",
		Length:  181,
		rating:  0,
		plays:   3,
		viewers: 300,
		critics: []float32{9, 9, 9, 8.5},
	}

	m3 := &Movie{
		Name:    "Titanic",
		Length:  195,
		rating:  0,
		plays:   2,
		viewers: 200,
		critics: []float32{8, 8, 8, 7.5},
	}

	m4 := &Movie{
		Name:    "Star Wars: The Force Awakens",
		Length:  135,
		rating:  0,
		plays:   1,
		viewers: 100,
		critics: []float32{7, 7, 7, 6.5},
	}

	t1 := Theatre{
		name : "The Wall",
	}

	mlist := []*Movie{m1,m2,m3,m4}

	err :=  t1.Critique(mlist, Critic)
	if err != nil {
		t.Error(err)
	}
}

func Test_Critic(t *testing.T){
	m4 := &Movie{
			Name:    "Star Wars: The Force Awakens",
			Length:  135,
			rating:  0,
			plays:   2,
			viewers: 200,
			critics: []float32{7, 7, 7, 6.5},
		}	
	
	rating := float32(68.75) // hard calculated

	result, err := Critic(m4)
		if err != nil {
			t.Error(err)
		}
	
	if result != rating {
		t.Errorf("Expected %f got %f \n", rating, result)
	}
	
}
