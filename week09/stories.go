package week09

import "fmt"

type story struct {
	title  string
	body   string
	writer string
}

func (s story) String() string {
	return fmt.Sprintf("%v\n%v\n%v", s.title, s.body, s.writer)
}

func (s story) IsValid() error {
	switch true {
	case s.title == "":
		return ErrTitleInvalid(s.title)
	case s.body == "":
		return ErrBodyInvalid(s.body)
	case s.writer == "":
		return ErrWriterInvalid(s.writer)
	}
	return nil
}
