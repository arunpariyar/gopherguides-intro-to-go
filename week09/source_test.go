package week09

import (
	"testing"
)

func Test_Source_Build(t *testing.T) {

	s1 := &source{
		name: "Daily Planet",
	}

	d := draft{
		title:  "Green Goblin Attacks Again",
		body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer: "Eddie Brock",
	}
	exp := story{
		publisher:  "Daily Planet",
		catagories: []catagory{"new york", "spider-man"},
		title:      "Green Goblin Attacks Again",
		body:       "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer:     "Eddie Brock",
	}

	act, err := s1.build(d, "new york", "spider-man")

	if err != nil {
		if act.String() != exp.String() {
			t.Fatalf("expected \n %#v \n got %#v", exp, act)
		}
	}

}

func Test_Source_Push(t *testing.T) {
	s := &source{
		name: "Daily Bugle",
	}
	st := story{
		publisher:  "Daily Planet",
		catagories: []catagory{"new york", "spider-man"},
		title:      "Green Goblin Attacks Again",
		body:       "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer:     "Eddie Brock",
	}

	act := s.push(st)
	exp := 1 
	if len(s.stories) != 1 {
		t.Fatalf("expected %v got %v",exp, act)
	}
	
}
