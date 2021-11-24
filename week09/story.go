package week09

import "fmt"

type catagory string

type publisher string

type catagories []catagory

type story struct {
	publisher  string
	catagories catagories
	title      string
	body       string
	writer     string
}

type stories []story

func (s story) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n", s.publisher, s.catagories, s.title, s.body, s.writer)
}

func (s story) IsValid() error {
	switch true {
	case s.publisher == "":
		return ErrStoryPublisherInvalid(s.publisher)
	case len(s.catagories) == 0:
		return ErrStoryCatagoriesInvalid(len(s.catagories))
	}
	return nil
}
