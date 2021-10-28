package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Entertain1(t *testing.T) {

	res := &bytes.Buffer{}

	v := Venue{
		Audience: 0,
		Log:      res,
	}

	b := band{
		name: "Guns and Roses",
	}

	var _ Teardowner = b

	err := v.Entertain(2, b)
	if err != nil {
		t.Fatal(err)
	}

	exp := fmt.Sprintf("Guns and Roses has performed for %d people.\nGuns and Roses has completed teardown.\n", v.Audience)

	act := res.String()
	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func Test_Entertain2(t *testing.T) {

	res := &bytes.Buffer{}

	v := Venue{
		Audience: 0,
		Log:      res,
	}

	b := band{
		name: "Guns and Roses",
	}

	c := comedian{
		name: "Trevor Noah",
	}

	var _ Setuper = c
	var _ Teardowner = b

	err := v.Entertain(2, b, c)
	if err != nil {
		t.Error(err)
	}

	exp := fmt.Sprintf("Guns and Roses has performed for %d people.\nGuns and Roses has completed teardown.\nTrevor Noah has completed setup.\nTrevor Noah has performed for %d people.\n", v.Audience, v.Audience)

	act := res.String()
	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}

}
