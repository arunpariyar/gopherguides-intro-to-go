package week09

import "fmt"

// type drafts []draft

type draft struct {
	title  string
	body   string
	writer string
}

func (d draft) String() string {
	return fmt.Sprintf("%v\n%v\n%v", d.title, d.body, d.writer)
}

func (d draft) IsValid() error {
	switch true {
	case d.title == "":
		return ErrTitleInvalid(d.title)
	case d.body == "":
		return ErrBodyInvalid(d.body)
	case d.writer == "":
		return ErrWriterInvalid(d.writer)
	}
	return nil
}
