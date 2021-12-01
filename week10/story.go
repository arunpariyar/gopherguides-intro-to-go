package week10

import "fmt"

type catagory string

// type publisher string

// type catagories []catagory

type story struct {
	publisher string
	catagory  catagory
	title     string
	body      string
	writer    string
}

type stories []story

func (s story) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v", s.publisher, s.catagory, s.title, s.body, s.writer)
}

func (s story) IsValid() error {
	switch true {
	case s.publisher == "":
		return ErrStoryPublisherInvalid(s.publisher)
	case s.catagory == "":
		return ErrStoryCatagoryInvalid(s.catagory)
	}
	return nil
}
