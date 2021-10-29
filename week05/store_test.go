package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_db(t *testing.T) {

	s := &Store{}

	table := []struct {
		name  string
		store *Store
		exp   data
	}{
		{name: "empty store", store: s, exp: s.data},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.store.db()

			if !reflect.DeepEqual(s.data, data) {
				t.Fatalf("%s expected %q got %q", tt.name, tt.exp, data)
			}
		})
	}
}

func Test_All(t *testing.T) {
	tn1 := "users"
	tn2 := "orders"

	s1 := &Store{}
	db1 := s1.db()

	err1 := ErrTableNotFound{
		table: "users",
	}

	errMsg1 := fmt.Errorf("%s", err1.Error())

	s2 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}
	db2 := s2.db()

	table := []struct {
		name  string
		store *Store
		tn    string
		mods  Models
		err   error
	}{
		{name: "empty store", store: s1, tn: tn1, mods: db1[tn1], err: errMsg1},
		{name: "store with orders", store: s2, tn: tn2, mods: db2[tn2], err: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.store.All(tt.tn)

			if err != nil {
				//need to ask naveen about this
				if err == tt.err {
					t.Fatal(err)
				}
			}

			if !reflect.DeepEqual(tt.mods, mods) {
				t.Fatalf("%s expected %q got %q", tt.name, tt.mods, mods)
			}
		})
	}
}

func Test_Len(t *testing.T){
	tn1 := "users"
	tn2 := "orders"

	s1 := &Store{}

	s2 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}

	err1 := ErrTableNotFound{
		table: "users",
	}

	errMsg1 := fmt.Errorf("%s", err1.Error())

	table := []struct{
		name 	string
		store 	*Store
		tn 		string
		len		int
		err 	error
	}{
		{name: "empty store", store:s1, tn: tn1,len:0, err:errMsg1},
		{name: "store with orders", store:s2, tn:tn2,len:2, err:nil},
	}

	for _, tt := range table{
		t.Run(tt.name, func(t *testing.T){
			len, err := tt.store.Len(tt.tn)
			if err != nil {
				//need to ask naveen about this
			if err == tt.err {
					t.Fatal(err)
				}
			}
			if len != tt.len {
				t.Fatalf("%s expected length %q got %q", tt.name, tt.len, len)
			}

		})
	}
}

func Test_Insert(t *testing.T){


	tn := "orders"

	s1 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}

	m1 := Model {"id": 3, "item": "mobile"}
	m2 := Model {"id": 4, "item": "smartwatch"}

	addOne := []Model{m1}
	addTwo := []Model{m1,m2}

	expMods1 := []Model{
		{"id": 1, "item": "desktop"}, 
		{"id": 2, "item": "laptop"},
		{"id": 3, "item": "mobile"},}
		
	expMods2 := []Model{
		{"id": 1, "item": "desktop"},
		 {"id": 2, "item": "laptop"},
		 {"id": 3, "item": "mobile"},
		 {"id": 4, "item": "smartwatch"}}

	table := []struct{
		name 	string
		store 	*Store
		tn		string
		mods 	[]Model
		exp 	[]Model
	}{
		{name:"insert one item", store: s1, tn:tn, mods:addOne,exp:expMods1},
		{name:"insert two item", store: s1, tn:tn, mods:addTwo,exp:expMods2},
	}

	for _, tt := range table{
		t.Run(tt.name, func(t *testing.T){
			tt.store.Insert(tt.tn, tt.mods...)
			db := tt.store.db()
			mods := db[tt.tn]
			
			if reflect.DeepEqual(tt.exp, mods){
				t.Fatalf("%s expected models %v got %v", tt.name, tt.exp, mods)
			}

		})
	}
}


