package main

import (
	"bytes"
	"fmt"
	"testing"
)

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

func (b band) Setup(v Venue) error {
	name := b.Name()
	fmt.Fprintf(v.Log, "%s has completed setup.\n", name)
	return nil
}

func (b band) Teardown(v Venue) error {
	name := b.Name()
	fmt.Fprintf(v.Log, "%s has completed teardown.\n", name)
	return nil
}

func Test_Entertain(t *testing.T){
	
	res := &bytes.Buffer{}
	v := Venue{
		Audience: 0,
		Log:      res,
	}

	b := band{
		name: "Guns and Roses",
	}

	err := v.Entertain(2, b)
	if err != nil { 
		t.Error(err)
	}

	exp := fmt.Sprintf("Guns and Roses has completed setup.\nGuns and Roses has performed for %d people.\nGuns and Roses has completed teardown.\n",v.Audience)

	if res.String() != exp {
		t.Error("logs don't match")
	}
}