package week09

import (
	"testing"
)

func Test_News_Service_Search(t *testing.T) {

	ns := &NewsService{}
	exp := 2
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
	n2 := news{
		id: 2,
		story: story{
			title:  "Avengers Defeat Thanos",
			body:   "With a click of a finger and a line that will be for years to come Iron Man and the avengers defeat thanos",
			writer: "peter parkar",
		},
		publisher: "Daily Bugle",
		cat:       "world",
	}

	n3 := news{
		id: 3,
		story: story{
			title:  "Superman dead",
			body:   "The red cape flutters no more",
			writer: "louis lane",
		},
		publisher: "The Daily Planet",
		cat:       "superman",
	}
	//creating a map news and assigning it to history
	m := make(map[int]news)
	m[1] = n
	m[2] = n2
	m[3] = n3
	ns.History = m

	act := ns.Search(1, 5, 3)

	if len(act) != exp {
		t.Fatalf("expected %v got %v", len(act), exp)
	}
}

func Test_News_Service_Stop(t *testing.T) {
	ns := &NewsService{
		Stopped: false,
	}

	ns.Stop()

	if ns.Stopped != true {
		t.Fatalf("expected true got %v", ns.Stopped)
	}
}
