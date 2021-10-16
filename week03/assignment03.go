package main

import (
	"fmt"
)

type Movie struct {
	Name    string
	Length  int
	rating  float32
	plays   int
	viewers int
	critics []float32
}

type Theatre struct {
	name string
}

type Critiquefn func(*Movie) (float32, error)

func (m *Movie) Rate(rating float32) error {
	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}
	if rating == 0 {
		return fmt.Errorf("rating can't be a zero value")
	}
	m.rating = rating
	return nil
}

func (m *Movie) Play(viewers int) {
	m.viewers += viewers
	m.plays++
}

func (m Movie) Viewers() int {
	return m.viewers
}

func (m Movie) Plays() int {
	return m.plays
}

func (m Movie) Rating() float64 {
	rating := float64(m.rating) / float64(m.plays)
	return rating
}

func (m Movie) String() string {
	return fmt.Sprintf("%s (%dm) %.1f%%", m.Name, m.Length, m.rating)
}

func (t *Theatre) Play(viewers int, movies ...*Movie) error {
	if (viewers) == 0 {
		return fmt.Errorf("viewers cant be zero")
	}
	if len(movies) == 0 {
		return fmt.Errorf("no movies to play")
	}
	for _, m := range movies {
		m.Play(viewers)
	}
	return nil
}

func (t Theatre) Critique(m []*Movie, cfn Critiquefn) error {
	if len(m) == 0 {
		return fmt.Errorf("no movies to play")
	}
	for _, m := range m {
		m.Play(1)
		rating, err := cfn(m)

		if err != nil {
			return err
		}

		err = m.Rate(rating)

		if err != nil {
			return err
		}
	}
	return nil
}

func Critic(m *Movie) (float32, error) {
	if m.plays == 1 {
		return 0, fmt.Errorf("this is the films first critique play value expected more than 1 got %d", m.plays)
	}

	var rating float32

	for _, v := range m.critics {
		rating += v
	}

	rating = rating / float32(len(m.critics)) * 10
	return float32(rating), nil
}

func main() {

}
