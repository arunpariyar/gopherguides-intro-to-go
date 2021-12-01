package week10

import (
	"testing"
)

func Test_Stories_String(t *testing.T) {
	d := draft{
		title:  "Green Goblin Attacks Again",
		body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer: "Eddie Brock",
	}
	act := d.String()
	exp := "Green Goblin Attacks Again\nNew York is being terrorised by Green Goblin yet again and spiderman is no where to be found.\nEddie Brock"
	if act != exp {
		t.Fatalf("expected \n %#v \n got \n %#v", exp, d.String())
	}
}

func Test_Stories_IsValid(t *testing.T) {
	table := []struct {
		name string
		d    draft
		exp  error
	}{
		{
			name: "valid story",
			d: draft{
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			exp: nil,
		},
		{
			name: "title missing",
			d: draft{
				title:  "",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			exp: ErrTitleInvalid(""),
		},
		{
			name: "body missing",
			d: draft{
				title:  "Green Goblin Attacks Again",
				body:   "",
				writer: "Eddie Brock",
			},
			exp: ErrBodyInvalid(""),
		},
		{
			name: "writer missing",
			d: draft{
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "",
			},
			exp: ErrWriterInvalid(""),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.d.IsValid()
			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}

}
