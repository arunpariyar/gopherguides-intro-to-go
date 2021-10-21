package main

import (
	"fmt"
	"io"
)

type Venue struct {
	Audience int
	Log      io.Writer
}

type Entertainer interface {
	Name() string
	Perform(v Venue) error
}

type Setuper interface {
	Setup(v Venue) error
}

type Teardowner interface {
	Teardown(v Venue) error
}

func (v *Venue) Entertain(aud int, artists ...Entertainer) error {

	if aud == 0 {
		return fmt.Errorf("audience cant be empty")
	}
	if len(artists) == 0 {
		return fmt.Errorf("artist not entered")
	}
	v.Audience = aud

	for _, artist := range artists {
		if st, ok := artist.(Setuper); ok {
			if err := st.Setup(*v); err != nil {
				return err
			}
		}
		
		if err := artist.Perform(*v); err != nil {
			return err
		}

		if td, ok := artist.(Teardowner); ok {
			if err := td.Teardown(*v); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {

}
