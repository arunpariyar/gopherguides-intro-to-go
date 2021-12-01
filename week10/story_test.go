package week10

import "testing"

func Test_Story_IsValid(t *testing.T) {
	table := []struct {
		name string
		s    *story
		exp  error
	}{
		{
			name: "publisher invalid",
			s: &story{
				// publisher: "Daily Planet",
				catagory: "spider-man",
				title:    "Green Goblin Attacks Again",
				body:     "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer:   "Eddie Brock",
			},
			exp: ErrStoryPublisherInvalid(""),
		},
		{
			name: "publisher invalid",
			s: &story{
				publisher: "Daily Planet",
				// catagory: "spider-man",
				title:    "Green Goblin Attacks Again",
				body:     "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer:   "Eddie Brock",
			},
			exp: ErrStoryCatagoryInvalid(""),
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
