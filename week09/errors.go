package week09

import "fmt"

type ErrTitleInvalid string

func (e ErrTitleInvalid) Error() string {
	return fmt.Sprintf("draft must have a title, got %#v", string(e))
}

type ErrBodyInvalid string

func (e ErrBodyInvalid) Error() string {
	return fmt.Sprintf("draft must have a body, got %#v", string(e))
}

type ErrWriterInvalid string

func (e ErrWriterInvalid) Error() string {
	return fmt.Sprintf("draft must have a writer, got %#v", string(e))
}

type ErrStoryPublisherInvalid string

func (e ErrStoryPublisherInvalid) Error() string {
	return fmt.Sprintf("story must hava a publisher got %#v", string(e))
}

type ErrStoryCatagoriesInvalid int

func (e ErrStoryCatagoriesInvalid) Error() string {
	return "story catagories is empty"
}
