package week09

import (
	"testing"
)
func Test_Stories_String(t *testing.T){
	s1 := story{
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			}
	act := s1.String()	
	exp := "Green Goblin Attacks Again\nNew York is being terrorised by Green Goblin yet again and spiderman is no where to be found.\nEddie Brock"
	if act != exp{ 
		t.Fatalf("expected \n %#v \n got \n %#v", exp, s1.String())
	}
}

func Test_Stories_IsValid(t *testing.T) {
	table := []struct {
		name string
		s    story
		exp  error
	}{
		{
			name: "valid story",
			s: story{
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			exp: nil,
		},
		{
			name: "title missing",
			s: story{
				title:  "",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			exp: ErrTitleInvalid(""),
		},
		{
			name: "body missing",
			s: story{
				title:  "Green Goblin Attacks Again",
				body:   "",
				writer: "Eddie Brock",
			},
			exp: ErrBodyInvalid(""),
		},
		{
			name: "writer missing",
			s: story{
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "",
			},
			exp: ErrWriterInvalid(""),
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
