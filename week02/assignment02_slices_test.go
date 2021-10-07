package main

import "testing"

func TestCopySlice(t *testing.T){
	exp := []string{"Mount Everest", "K2", "Kangchenjunga", "Lhotse", "Makalu", "Cho Oyu", "Dhaulagiri I", "Manaslu", "Nanga Parbat", "Annapurna I"}

	//creating act slice with the length of exp
	act := make([]string, 0, len(exp))

	isSame := copySlice(exp, act)
	
	if isSame != true {
		t.Error("Expected", true, "result", isSame)
	}
}