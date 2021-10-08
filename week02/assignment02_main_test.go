package main

import "testing"

func TestCopySlice(t *testing.T){
	exp := []string{"Mount Everest", "K2", "Kangchenjunga", "Lhotse", "Makalu", "Cho Oyu", "Dhaulagiri I", "Manaslu", "Nanga Parbat", "Annapurna I"}

	//creating act slice // note since we have to compare the length of slice with each other at the end didnt set the length here.
	act := make([]string, 0)

	isSame := copySlice(exp, act)
	
	if isSame != true {
		t.Error("Length of slices is different: Expected", true, "result", isSame)
	}
}

func TestCopyArray(t *testing.T){
	exp := [7]string{"Great Wall of China", "Chichén Itzá", "Petra", "Machu Picchu", "Colosseum", "Taj Mahal", "Christ the Redeemer"}

	//creating a new array that has the same length as exp
	act := [len(exp)]string{}

	mismatch := copyArray(exp, act)
		if mismatch > 0 {
			t.Error("Content mismatch found : Expected", 0, "Mismatch", mismatch)
		}
}

func TestCopyMap(t *testing.T){
	exp := map[string]string{
		"india":"Taj Mahal",
		"china":"The Great Wall",
		"jordan":"Petra",
		"mexico":"Chichen Itza",
		"peru":"Machu Picchu",
		"italy":"The Colosseum",
		"brazil":"Christ The Redeemer",
	}

	act := map[string]string{}

	mismatch := copyMap(exp, act)

	if mismatch > 0 {
		t.Error("Content mismatch found : Expected", 0, "Mismatch", mismatch)
	}

}
