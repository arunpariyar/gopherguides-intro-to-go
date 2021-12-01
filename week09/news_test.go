package week09

import (
	"testing"
)

func Test_String(t *testing.T) {
	n := news{
		id: 1,
		story: story{
			title:  "Green Goblin Attacks Again",
			body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
			writer: "Eddie Brock",
		},
		publisher: "Daily Bugle",
		cat:       "entertainment",
	}

	exp := `id:1
story:

Green Goblin Attacks Again
New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.
Eddie Brockpublisher:Daily Bugle
catagory:entertainment
`
	act := n.String()
	if act != exp {
		t.Fatalf("expected: \n %v \n got \n %v \n", exp, act)
	}
}
