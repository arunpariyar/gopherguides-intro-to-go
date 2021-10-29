package demo 

import "testing"
import "reflect"

func Test_db(t *testing.T){
	
	s := &Store{}

	table := []struct{
		name	string
		store 	*Store
		exp 	data 
	}{
		{name:"empty store", store:s, exp:s.data,},
	}

	for _, tt := range table{
		t.Run(tt.name, func(t *testing.T){
			data := tt.store.db()

			if !reflect.DeepEqual(s.data, data){
				t.Fatalf("%s expected %q got %q",tt.name, tt.exp, data)
			}



		})
	}

}