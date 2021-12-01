package week09

import "testing"

func Test_Subscriber_IsValid(t *testing.T){
	table := []struct{
		name string
		s 	subscriber
		exp error 
	}{
		{
			name: "success",
			s : subscriber{
				name : "The Daily Planet",
				cat : catagories{"sports", "international"},
			},
			exp : nil,

		},
		{
			name: "name invalid",
			s : subscriber{
				// name : "The Daily Planet",
				cat : catagories{"sports", "international"},
			},
			exp : ErrSubscriberInvalidName(""),
		},
		{
			name: "catagories invalid",
			s : subscriber{
				name : "The Daily Planet",
				// cat : catagories{"sports", "international"},
			},
			exp : ErrSubscriberCatgoriesInvalid(0),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.s.IsValid()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}