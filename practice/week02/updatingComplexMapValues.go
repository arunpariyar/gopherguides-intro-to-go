package main

import "fmt"

func main(){
	type turtle struct {
		Name string
		Weapon string
	}

	tmnt := map[string]turtle{}
	// Creating Dynamically assign both key and value
	tmnt["Leo"] = turtle{"Leonardo"," Twin Katana"}

	// Create from an instance of struct
	// Creating the struct 
	r := turtle{Name:"Ralph", Weapon: "Ninja Sai"}
	//assigning key to struct value
	tmnt[r.Name] = r 

	r = turtle{Name:"Donnatello",Weapon: "Staff"}
	tmnt[r.Name] = r

	r = turtle{Name:"Michealangelo", Weapon: ""}
	tmnt[r.Name] = r 

	//update Michealangelo Weapon
	key := "Michealangelo"
	t, ok := tmnt[key]
	if !ok {
		fmt.Errorf("Couldn't find %q", key)
	} else {
		t.Weapon = "NunChuks"
	}

	tmnt[t.Name] = t


	fmt.Println(tmnt)
}