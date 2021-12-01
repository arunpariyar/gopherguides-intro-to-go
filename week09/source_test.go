package week09

import (
	"testing"
)

func Test_Source_Build_Success(t *testing.T) {

	s1 := &source{
		Name: "Daily Planet",
	}

	d := draft{
		title:  "Green Goblin Attacks Again",
		body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer: "Eddie Brock",
	}
	exp := story{
		publisher: "Daily Planet",
		catagory:  "spider-man",
		title:     "Green Goblin Attacks Again",
		body:      "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer:    "Eddie Brock",
	}

	act, err := s1.build(d, "new york")

	if err != nil {
		if act.String() != exp.String() {
			t.Fatalf("expected \n %#v \n got %#v", exp, act)
		}
	}
}

func Test_Source_Build_Fail(t *testing.T) {

	table := []struct {
		name string
		s    *source
		d    draft
		c    catagory
		exp  error
	}{
		{
			name: "title invalid",
			s: &source{
				Name: "Daily Bugle",
			},
			d: draft{
				title:  "",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			c:   "spider-man",
			exp: ErrTitleInvalid(""),
		},
		{
			name: "body invalid",
			s: &source{
				Name: "Daily Bugle",
			},
			d: draft{
				title:  "Green Goblin Back!",
				writer: "Eddie Brock",
			},
			c:   "spider-man",
			exp: ErrBodyInvalid(""),
		},
		{
			name: "body invalid",
			s: &source{
				Name: "Daily Bugle",
			},
			d: draft{
				title: "Green Goblin Back!",
				body:  "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
			},
			c:   "spider-man",
			exp: ErrWriterInvalid(""),
		},
		{
			name: "catagory invalid",
			s: &source{
				Name: "Daily Bugle",
			},
			d: draft{
				title:  "Green Goblin Back!",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			c:   "",
			exp: ErrStoryCatagoryInvalid(""),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, act := tt.s.build(tt.d, tt.c)

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}

		})
	}

}

func Test_Source_Push_Success(t *testing.T) {
	s := &source{
		Name: "Daily Bugle",
	}
	st := story{
		publisher: "Daily Planet",
		catagory:  "new york",
		title:     "Green Goblin Attacks Again",
		body:      "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
		writer:    "Eddie Brock",
	}

	act := s.push(st)
	exp := 1
	if len(s.stories) != 1 {
		t.Fatalf("expected %v got %v", exp, act)
	}

}

func Test_Source_Push_Fail(t *testing.T) {

	table := []struct {
		name string
		s    *source
		st   story
		exp  error
	}{
		{
			s: &source{
				Name: "Daily Bugle",
			},
			st: story{
				// publisher: "Daily Planet",
				catagory: "new york",
				title:    "Green Goblin Attacks Again",
				body:     "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer:   "Eddie Brock",
			},
			exp: ErrStoryPublisherInvalid(""),
		},
		{
			s: &source{
				Name: "Daily Bugle",
			},
			st: story{
				publisher: "Daily Planet",
				// catagory: "new york",
				title:  "Green Goblin Attacks Again",
				body:   "New York is being terrorised by Green Goblin yet again and spiderman is no where to be found.",
				writer: "Eddie Brock",
			},
			exp: ErrStoryCatagoryInvalid(""),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.s.push(tt.st)

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}

}
