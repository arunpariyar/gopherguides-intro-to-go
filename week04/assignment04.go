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
	//check if audience is null
	if aud == 0 {
		return fmt.Errorf("audience cant be empty")
	}
	//check if artist is null
	if len(artists) == 0 {
		return fmt.Errorf("artist not entered")
	}
	//update the value of venues Audience with audience entered
	v.Audience = aud

	//The Venue should check each Entertainer to see if it implements the Setuper or Teardowner interfaces and call them accordingly.
	for _, artist := range artists {
		if st, ok := artist.(Setuper); ok {
			if err := st.Setup(*v); err != nil {
				return err
			}
		}
		//For each Entertainer call its Perform method passing in the Venue
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

type band struct {
	name string
}

func (b band) Name() string {
	return b.name
}

func (b band) Perform(v Venue) error {
	name := b.Name()
	fmt.Fprintf(v.Log, "%s has performed for %d people.\n", name, v.Audience)
	return nil
}

func (b band) Teardown(v Venue) error {
	name := b.Name()
	fmt.Fprintf(v.Log, "%s has completed teardown.\n", name)
	return nil
}

type comedian struct {
	name string
}

func (c comedian) Name() string {
	return c.name
}

func (c comedian) Setup(v Venue) error {
	name := c.Name()
	fmt.Fprintf(v.Log, "%s has completed setup.\n", name)
	return nil
}

func (c comedian) Perform(v Venue) error {
	name := c.Name()
	fmt.Fprintf(v.Log, "%s has performed for %d people.\n", name, v.Audience)
	return nil
}

func main() {

}
